// repositories/interfaces/secretrepository.go

package interfaces

import (
	"context"
	"github.com/google/uuid"
	"tribemedia.io/m/ent"
)

type SecretRepository interface {
	CreateSecret(ctx context.Context, secret *ent.Secret) (*ent.Secret, error)
	GetSecretByID(ctx context.Context, id uuid.UUID) (*ent.Secret, error)
	ListSecretsByProjectID(ctx context.Context, projectID uuid.UUID) ([]*ent.Secret, error)
	UpdateSecret(ctx context.Context, secret *ent.Secret) (*ent.Secret, error)
	DeleteSecret(ctx context.Context, id uuid.UUID) error
}
