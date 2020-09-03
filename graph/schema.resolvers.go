package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/bb3104/gqlgen_gorm_sample/db_model"
	"github.com/bb3104/gqlgen_gorm_sample/graph/generated"
	"github.com/bb3104/gqlgen_gorm_sample/graph/model"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*db_model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Users(ctx context.Context) ([]*db_model.User, error) {
	users := []*db_model.User{}

	err := r.DB.Preload("Items").Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *queryResolver) User(ctx context.Context, input *model.FetchUser) (*db_model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
