// repositories/interfaces/projectrepository.go

package interfaces

import (
	"context"
	"github.com/google/uuid"
	"tribemedia.io/m/ent"
)

type ProjectRepository interface {
	CreateProject(ctx context.Context, project *ent.Project) (*ent.Project, error)
	GetProjectByID(ctx context.Context, id uuid.UUID) (*ent.Project, error)
	ListProjects(ctx context.Context) ([]*ent.Project, error)
	UpdateProject(ctx context.Context, project *ent.Project) (*ent.Project, error)
	DeleteProject(ctx context.Context, id uuid.UUID) error
}
