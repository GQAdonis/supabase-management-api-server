// repositories/interfaces/customhostnamerepository.go

package interfaces

import (
	"context"
	"github.com/google/uuid"
	"tribemedia.io/m/ent"
)

type CustomHostnameRepository interface {
	CreateCustomHostname(ctx context.Context, customHostname *ent.CustomHostname) (*ent.CustomHostname, error)
	GetCustomHostnameByID(ctx context.Context, id uuid.UUID) (*ent.CustomHostname, error)
	ListCustomHostnamesByProjectID(ctx context.Context, projectID uuid.UUID) ([]*ent.CustomHostname, error)
	UpdateCustomHostname(ctx context.Context, customHostname *ent.CustomHostname) (*ent.CustomHostname, error)
	DeleteCustomHostname(ctx context.Context, id uuid.UUID) error
}
