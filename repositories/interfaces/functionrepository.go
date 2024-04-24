// repositories/interfaces/functionrepository.go

package interfaces

import (
	"context"
	"github.com/google/uuid"
	"tribemedia.io/m/ent"
)

type FunctionRepository interface {
	CreateFunction(ctx context.Context, function *ent.Function) (*ent.Function, error)
	GetFunctionByID(ctx context.Context, id uuid.UUID) (*ent.Function, error)
	ListFunctionsByProjectID(ctx context.Context, projectID uuid.UUID) ([]*ent.Function, error)
	UpdateFunction(ctx context.Context, function *ent.Function) (*ent.Function, error)
	DeleteFunction(ctx context.Context, id uuid.UUID) error
}
