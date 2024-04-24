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
	"time"
	"tribemedia.io/m/ent"
	"tribemedia.io/m/handlers"
	"tribemedia.io/m/repositories/entgo"
	"tribemedia.io/m/services"
	s "tribemedia.io/m/supabase"
)

func Open(databaseUrl string) *ent.Client {
	db, err := sql.Open("pgx", databaseUrl)
	if err != nil {
		log.Fatal(err)
	}

	// Create an ent.Driver from `db`.
	drv := entsql.OpenDB(dialect.Postgres, db)
	return ent.NewClient(ent.Driver(drv))
}

func main() {

	pflag.String("dsn", "", "database connection string")
	pflag.Parse()

	// Bind the command line flags to Viper
	err := viper.BindPFlags(pflag.CommandLine)
	if err != nil {
		log.Fatal(err)
	}

	configFile := viper.GetString("config")
	if configFile != "" {
		viper.SetConfigFile(configFile)
	} else {
		viper.SetConfigName("config") // Name of the config file (without extension)
		viper.AddConfigPath(".")      // Look for config in the current directory
	}

	// Attempt to read the configuration file
	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Error reading config file, %s", err)
	}

	// Automatically read configuration options from environment variables
	viper.AutomaticEnv()

	//dsn := "postgresql://postgres:postgres@127.0.0.1:4322/postgres"

	// Use the configuration values
	dsn := viper.GetString("dsn")
	supabaseUrl := viper.GetString("supabaseUrl")
	supabaseAnonKey := viper.GetString("supabaseAnonKey")
	supabaseServiceRoleKey := viper.GetString("supabaseServiceRoleKey")
	projectRef := viper.GetString("projectRef")

	fmt.Println("DSN:", dsn)
	fmt.Println("Supabase URL:", supabaseUrl)
	fmt.Println("Supabase Anon Key:", supabaseAnonKey)
	fmt.Println("Supabase Service Role Key:", supabaseServiceRoleKey)
	fmt.Println("Project Ref:", projectRef)

	client := Open(dsn)
	defer func(client *ent.Client) {
		err := client.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(client)

	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	supabaseClient, err := supabase.NewClient(
		supabaseUrl, supabaseServiceRoleKey, nil)
	if err != nil {
		fmt.Println("cannot initalize client", err)
	}

	goTrueClient := gotrue.New(
		projectRef,
		supabaseAnonKey,
	)

	authService := services.NewAuthService(goTrueClient)

	//secretRepo := entgo.NewSecretRepositoryEntgo(client)
	//functionRepo := entgo.NewFunctionRepositoryEntgo(client)
	customHostnameRepo := entgo.NewCustomHostnameRepositoryEntgo(client)
	//networkBanRepo := entgo.NewNetworkBanRepositoryEntgo(client)
	//pgsodiumConfigRepo := entgo.NewPgsodiumConfigRepositoryEntgo(client)
	projectRepo := entgo.NewProjectRepositoryEntgo(client, supabaseClient)
	apiHandler := handlers.NewApiHandler(
		projectRepo,
		customHostnameRepo,
	)

	// Create generated server.
	srv, err := s.NewServer(apiHandler, handlers.NewSecurityHandler(authService))
	if err != nil {
		log.Fatal(err)
	}

	// Create a new router
	router := http.NewServeMux()

	// Register the server routes
	router.Handle("/", srv)

	// Register the Swagger UI endpoint
	opts := middleware.SwaggerUIOpts{SpecURL: "/openapi.json"}
	sh := middleware.SwaggerUI(opts, nil)
	router.Handle("/docs", sh)

	// Register the openapi.json endpoint
	router.HandleFunc("/openapi.json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./openapi.json")
	})

	// Start the server
	if err := http.ListenAndServe(":8180", router); err != nil {
		log.Fatal(err)
	}
}

func globalLoggingHook(next ent.Mutator) ent.Mutator {
	return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
		start := time.Now()
		defer func() {
			log.Printf("Op=%s\tType=%s\tTime=%s\tConcreteType=%T\n", m.Op(), m.Type(), time.Since(start), m)
		}()
		return next.Mutate(ctx, m)
	})
}
