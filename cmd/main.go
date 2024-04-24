package main

import (
	"context"
	"database/sql"
	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"fmt"
	"github.com/go-openapi/runtime/middleware"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"github.com/supabase-community/gotrue-go"
	"github.com/supabase-community/supabase-go"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"tribemedia.io/m/ent"
	"tribemedia.io/m/handlers"
	"tribemedia.io/m/repositories/entgo"
	"tribemedia.io/m/services"
	s "tribemedia.io/m/supabase"
)

// Open initializes the Ent client with a connection to the database.
func Open(databaseUrl string) *ent.Client {
	var db *sql.DB
	var err error
	// Implementing a retry mechanism for database connection
	for i := 0; i < 5; i++ {
		db, err = sql.Open("pgx", databaseUrl)
		if err == nil {
			break
		}
		log.Printf("Failed to connect to database. Retry in 5 seconds...")
		time.Sleep(5 * time.Second)
	}
	if err != nil {
		log.Fatalf("Failed to connect to database after retries: %v", err)
	}

	// Create an ent.Driver from `db`.
	drv := entsql.OpenDB(dialect.Postgres, db)
	return ent.NewClient(ent.Driver(drv))
}

func main() {
	// Define command-line flags
	pflag.String("dsn", "", "database connection string")
	pflag.Int("port", 8180, "Port to run the server on")
	pflag.Parse()

	// Bind the command line flags to Viper
	if err := viper.BindPFlags(pflag.CommandLine); err != nil {
		log.Fatal(err)
	}

	// Configuration file setup
	configFile := viper.GetString("config")
	if configFile != "" {
		viper.SetConfigFile(configFile)
	} else {
		viper.SetConfigName("config")
		viper.AddConfigPath(".")
	}

	// Attempt to read the configuration file
	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Error reading config file, %s", err)
	}

	// Automatically read configuration options from environment variables
	viper.AutomaticEnv()

	// Use the configuration values
	dsn := viper.GetString("dsn")
	supabaseUrl := viper.GetString("supabaseUrl")
	supabaseAnonKey := viper.GetString("supabaseAnonKey")
	supabaseServiceRoleKey := viper.GetString("supabaseServiceRoleKey")
	projectRef := viper.GetString("projectRef")
	port := viper.GetInt("port") // Using the port number from Viper

	// Database client initialization with retry logic
	client := Open(dsn)
	defer client.Close()

	// Auto migration
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	// Supabase client setup
	supabaseClient, err := supabase.NewClient(supabaseUrl, supabaseServiceRoleKey, nil)
	if err != nil {
		log.Fatalf("cannot initialize Supabase client: %v", err)
	}

	// GoTrue client setup for Supabase authentication
	goTrueClient := gotrue.New(projectRef, supabaseAnonKey)
	authService := services.NewAuthService(goTrueClient)

	// Repository and handler initialization
	customHostnameRepo := entgo.NewCustomHostnameRepositoryEntgo(client)
	projectRepo := entgo.NewProjectRepositoryEntgo(client, supabaseClient)
	apiHandler := handlers.NewApiHandler(projectRepo, customHostnameRepo)

	// Generated server setup
	srv, err := s.NewServer(apiHandler, handlers.NewSecurityHandler(authService))
	if err != nil {
		log.Fatal(err)
	}

	// HTTP server and router setup
	router := http.NewServeMux()
	router.Handle("/", srv)
	opts := middleware.SwaggerUIOpts{SpecURL: "/openapi.json"}
	sh := middleware.SwaggerUI(opts, nil)
	router.Handle("/docs", sh)
	router.HandleFunc("/openapi.json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./openapi.json")
	})

	// Graceful shutdown setup
	httpSrv := &http.Server{Addr: fmt.Sprintf(":%d", port), Handler: router}
	go func() {
		log.Printf("Server is running at http://localhost:%d", port)
		if err := httpSrv.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Signal handling for graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := httpSrv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}
