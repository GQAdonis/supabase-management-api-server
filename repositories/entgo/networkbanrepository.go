// repositories/entgo/networkbanrepository.go

package entgo

import (
	"context"
	"github.com/google/uuid"
	"tribemedia.io/m/ent"
	"tribemedia.io/m/ent/networkban"
	"tribemedia.io/m/repositories/interfaces"
)

type NetworkBanRepositoryEntgo struct {
	client *ent.Client
}

func NewNetworkBanRepositoryEntgo(client *ent.Client) interfaces.NetworkBanRepository {
	return &NetworkBanRepositoryEntgo{client: client}
}

func (r *NetworkBanRepositoryEntgo) CreateNetworkBan(ctx context.Context, nb *ent.NetworkBan) (*ent.NetworkBan, error) {
	return r.client.NetworkBan.Create().
		SetID(nb.ID).
		SetProjectID(nb.ProjectID).
		SetIPAddress(nb.IPAddress).
		SetReason(nb.Reason).
		Save(ctx)
}

func (r *NetworkBanRepositoryEntgo) GetNetworkBanByID(ctx context.Context, id uuid.UUID) (*ent.NetworkBan, error) {
	return r.client.NetworkBan.Get(ctx, id)
}

func (r *NetworkBanRepositoryEntgo) ListNetworkBansByProjectID(ctx context.Context, projectID uuid.UUID) ([]*ent.NetworkBan, error) {
	return r.client.NetworkBan.Query().
		Where(networkban.ProjectID(projectID)).
		All(ctx)
}

func (r *NetworkBanRepositoryEntgo) UpdateNetworkBan(ctx context.Context, nb *ent.NetworkBan) (*ent.NetworkBan, error) {
	update := r.client.NetworkBan.UpdateOneID(nb.ID)
	if nb.IPAddress != "" {
		update.SetIPAddress(nb.IPAddress)
	}
	if nb.Reason != "" {
		update.SetReason(nb.Reason)
	}
	return update.Save(ctx)
}

func (r *NetworkBanRepositoryEntgo) DeleteNetworkBan(ctx context.Context, id uuid.UUID) error {
	return r.client.NetworkBan.DeleteOneID(id).Exec(ctx)
}
