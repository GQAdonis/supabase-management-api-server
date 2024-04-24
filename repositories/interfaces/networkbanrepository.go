// repositories/interfaces/networkbanrepository.go

package interfaces

import (
	"context"
	"github.com/google/uuid"
	"tribemedia.io/m/ent"
)

type NetworkBanRepository interface {
	CreateNetworkBan(ctx context.Context, networkBan *ent.NetworkBan) (*ent.NetworkBan, error)
	GetNetworkBanByID(ctx context.Context, id uuid.UUID) (*ent.NetworkBan, error)
	ListNetworkBansByProjectID(ctx context.Context, projectID uuid.UUID) ([]*ent.NetworkBan, error)
	UpdateNetworkBan(ctx context.Context, networkBan *ent.NetworkBan) (*ent.NetworkBan, error)
	DeleteNetworkBan(ctx context.Context, id uuid.UUID) error
}
