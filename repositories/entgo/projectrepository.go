package entgo

import (
	"context"
	"github.com/google/uuid"
	"github.com/supabase-community/supabase-go"
	"tribemedia.io/m/ent"
	"tribemedia.io/m/repositories/interfaces"
)

type ProjectRepositoryEntgo struct {
	client         *ent.Client
	supabaseClient *supabase.Client
}

func NewProjectRepositoryEntgo(client *ent.Client,
	supabaseClient *supabase.Client) interfaces.ProjectRepository {
	return &ProjectRepositoryEntgo{
		client:         client,
		supabaseClient: supabaseClient,
	}
}

func (r *ProjectRepositoryEntgo) CreateProject(ctx context.Context, p *ent.Project) (*ent.Project, error) {
	return r.client.Project.Create().
		SetName(p.Name).
		Save(ctx)
}

func (r *ProjectRepositoryEntgo) GetProjectByID(ctx context.Context, id uuid.UUID) (*ent.Project, error) {
	return r.client.Project.Get(ctx, id)
}

func (r *ProjectRepositoryEntgo) ListProjects(ctx context.Context) ([]*ent.Project, error) {
	return r.client.Project.Query().All(ctx)
}

func (r *ProjectRepositoryEntgo) UpdateProject(ctx context.Context, p *ent.Project) (*ent.Project, error) {
	return r.client.Project.UpdateOne(p).
		Save(ctx)
}

func (r *ProjectRepositoryEntgo) DeleteProject(ctx context.Context, id uuid.UUID) error {
	return r.client.Project.DeleteOneID(id).Exec(ctx)
}
