// repositories/entgo/pgsodiumconfigrepository.go

package entgo

import (
	"context"
	"github.com/google/uuid"
	"tribemedia.io/m/ent"
	"tribemedia.io/m/ent/pgsodiumconfig"
	"tribemedia.io/m/repositories/interfaces"
)

type PgsodiumConfigRepositoryEntgo struct {
	client *ent.Client
}

func NewPgsodiumConfigRepositoryEntgo(client *ent.Client) interfaces.PgsodiumConfigRepository {
	return &PgsodiumConfigRepositoryEntgo{client: client}
}

func (r *PgsodiumConfigRepositoryEntgo) CreatePgsodiumConfig(ctx context.Context, pc *ent.PgsodiumConfig) (*ent.PgsodiumConfig, error) {
	return r.client.PgsodiumConfig.Create().
		SetID(pc.ID).
		SetProjectID(pc.ProjectID).
		SetEnabled(pc.Enabled).
		Save(ctx)
}

func (r *PgsodiumConfigRepositoryEntgo) GetPgsodiumConfigByID(ctx context.Context, id uuid.UUID) (*ent.PgsodiumConfig, error) {
	return r.client.PgsodiumConfig.Get(ctx, id)
}

func (r *PgsodiumConfigRepositoryEntgo) ListPgsodiumConfigsByProjectID(ctx context.Context, projectID uuid.UUID) ([]*ent.PgsodiumConfig, error) {
	return r.client.PgsodiumConfig.Query().
		Where(pgsodiumconfig.ProjectID(projectID)).
		All(ctx)
}

func (r *PgsodiumConfigRepositoryEntgo) UpdatePgsodiumConfig(ctx context.Context, pc *ent.PgsodiumConfig) (*ent.PgsodiumConfig, error) {
	update := r.client.PgsodiumConfig.UpdateOneID(pc.ID)
	update.SetEnabled(pc.Enabled)
	return update.Save(ctx)
}

func (r *PgsodiumConfigRepositoryEntgo) DeletePgsodiumConfig(ctx context.Context, id uuid.UUID) error {
	return r.client.PgsodiumConfig.DeleteOneID(id).Exec(ctx)
}
