// repositories/entgo/secretrepository.go

package entgo

import (
	"context"
	"github.com/google/uuid"
	"tribemedia.io/m/ent"
	"tribemedia.io/m/ent/secret"
	"tribemedia.io/m/repositories/interfaces"
)

type SecretRepositoryEntgo struct {
	client *ent.Client
}

func NewSecretRepositoryEntgo(client *ent.Client) interfaces.SecretRepository {
	return &SecretRepositoryEntgo{client: client}
}

func (r *SecretRepositoryEntgo) CreateSecret(ctx context.Context, s *ent.Secret) (*ent.Secret, error) {
	return r.client.Secret.Create().
		SetName(s.Name).
		SetValue(s.Value).
		SetProjectID(s.ProjectID).
		Save(ctx)
}

func (r *SecretRepositoryEntgo) GetSecretByID(ctx context.Context, id uuid.UUID) (*ent.Secret, error) {
	return r.client.Secret.Get(ctx, id)
}

func (r *SecretRepositoryEntgo) ListSecretsByProjectID(ctx context.Context, projectID uuid.UUID) ([]*ent.Secret, error) {
	return r.client.Secret.Query().
		Where(secret.ProjectID(projectID)).
		All(ctx)
}

func (r *SecretRepositoryEntgo) UpdateSecret(ctx context.Context, s *ent.Secret) (*ent.Secret, error) {
	return r.client.Secret.UpdateOneID(s.ID).
		SetName(s.Name).
		SetValue(s.Value).
		Save(ctx)
}

func (r *SecretRepositoryEntgo) DeleteSecret(ctx context.Context, id uuid.UUID) error {
	return r.client.Secret.DeleteOneID(id).Exec(ctx)
}
