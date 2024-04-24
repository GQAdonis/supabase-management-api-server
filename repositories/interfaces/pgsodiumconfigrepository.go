// repositories/interfaces/pgsodiumconfigrepository.go

package interfaces

import (
	"context"
	"github.com/google/uuid"
	"tribemedia.io/m/ent"
)

type PgsodiumConfigRepository interface {
	CreatePgsodiumConfig(ctx context.Context, pgsodiumConfig *ent.PgsodiumConfig) (*ent.PgsodiumConfig, error)
	GetPgsodiumConfigByID(ctx context.Context, id uuid.UUID) (*ent.PgsodiumConfig, error)
	ListPgsodiumConfigsByProjectID(ctx context.Context, projectID uuid.UUID) ([]*ent.PgsodiumConfig, error)
	UpdatePgsodiumConfig(ctx context.Context, pgsodiumConfig *ent.PgsodiumConfig) (*ent.PgsodiumConfig, error)
	DeletePgsodiumConfig(ctx context.Context, id uuid.UUID) error
}
