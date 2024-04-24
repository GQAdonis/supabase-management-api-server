// repositories/entgo/functionrepository.go

package entgo

import (
	"context"
	"github.com/google/uuid"
	"tribemedia.io/m/ent"
	"tribemedia.io/m/ent/function"
	"tribemedia.io/m/repositories/interfaces"
)

type FunctionRepositoryEntgo struct {
	client *ent.Client
}

func NewFunctionRepositoryEntgo(client *ent.Client) interfaces.FunctionRepository {
	return &FunctionRepositoryEntgo{client: client}
}

func (r *FunctionRepositoryEntgo) CreateFunction(ctx context.Context, f *ent.Function) (*ent.Function, error) {
	return r.client.Function.Create().
		SetID(f.ID).
		SetProjectID(f.ProjectID).
		SetName(f.Name).
		SetRuntime(f.Runtime).
		SetSourceCode(f.SourceCode).
		Save(ctx)
}

func (r *FunctionRepositoryEntgo) GetFunctionByID(ctx context.Context, id uuid.UUID) (*ent.Function, error) {
	return r.client.Function.Get(ctx, id)
}

func (r *FunctionRepositoryEntgo) ListFunctionsByProjectID(ctx context.Context, projectID uuid.UUID) ([]*ent.Function, error) {
	return r.client.Function.Query().
		Where(function.ProjectID(projectID)).
		All(ctx)
}

func (r *FunctionRepositoryEntgo) UpdateFunction(ctx context.Context, f *ent.Function) (*ent.Function, error) {
	update := r.client.Function.UpdateOneID(f.ID)
	if f.Name != "" {
		update.SetName(f.Name)
	}
	if f.Runtime != "" {
		update.SetRuntime(f.Runtime)
	}
	if f.SourceCode != "" {
		update.SetSourceCode(f.SourceCode)
	}
	return update.Save(ctx)
}

func (r *FunctionRepositoryEntgo) DeleteFunction(ctx context.Context, id uuid.UUID) error {
	return r.client.Function.DeleteOneID(id).Exec(ctx)
}
