// repositories/entgo/customhostnamerepository.go

package entgo

import (
	"context"
	"github.com/google/uuid"
	"tribemedia.io/m/ent"
	"tribemedia.io/m/ent/customhostname"
	"tribemedia.io/m/repositories/interfaces"
)

type CustomHostnameRepositoryEntgo struct {
	client *ent.Client
}

func NewCustomHostnameRepositoryEntgo(client *ent.Client) interfaces.CustomHostnameRepository {
	return &CustomHostnameRepositoryEntgo{client: client}
}

func (r *CustomHostnameRepositoryEntgo) CreateCustomHostname(ctx context.Context, ch *ent.CustomHostname) (*ent.CustomHostname, error) {
	return r.client.CustomHostname.Create().
		SetID(ch.ID).
		SetProjectID(ch.ProjectID).
		SetHostname(ch.Hostname).
		SetSslStatus(ch.SslStatus).
		Save(ctx)
}

func (r *CustomHostnameRepositoryEntgo) GetCustomHostnameByID(ctx context.Context, id uuid.UUID) (*ent.CustomHostname, error) {
	return r.client.CustomHostname.Get(ctx, id)
}

func (r *CustomHostnameRepositoryEntgo) ListCustomHostnamesByProjectID(ctx context.Context, projectID uuid.UUID) ([]*ent.CustomHostname, error) {
	return r.client.CustomHostname.Query().
		Where(customhostname.ProjectID(projectID)).
		All(ctx)
}

func (r *CustomHostnameRepositoryEntgo) UpdateCustomHostname(ctx context.Context, ch *ent.CustomHostname) (*ent.CustomHostname, error) {
	update := r.client.CustomHostname.UpdateOneID(ch.ID)
	if ch.Hostname != "" {
		update.SetHostname(ch.Hostname)
	}
	if ch.SslStatus != "" {
		update.SetSslStatus(ch.SslStatus)
	}
	return update.Save(ctx)
}

func (r *CustomHostnameRepositoryEntgo) DeleteCustomHostname(ctx context.Context, id uuid.UUID) error {
	return r.client.CustomHostname.DeleteOneID(id).Exec(ctx)
}
