package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/pwanstax/backend-personal-webapp/graph/generated"
	"github.com/pwanstax/backend-personal-webapp/graph/model"
)

// CreateOrder is the resolver for the createOrder field.
func (r *mutationResolver) CreateOrder(ctx context.Context, input model.OrderInput) (*model.Order, error) {
	panic(fmt.Errorf("not implemented: CreateOrder - createOrder"))
}

// UpdateOrder is the resolver for the updateOrder field.
func (r *mutationResolver) UpdateOrder(ctx context.Context, orderID int, input model.OrderInput) (*model.Order, error) {
	panic(fmt.Errorf("not implemented: UpdateOrder - updateOrder"))
}

// DeleteOrder is the resolver for the deleteOrder field.
func (r *mutationResolver) DeleteOrder(ctx context.Context, orderID int) (bool, error) {
	panic(fmt.Errorf("not implemented: DeleteOrder - deleteOrder"))
}

// Orders is the resolver for the orders field.
func (r *queryResolver) Orders(ctx context.Context) ([]*model.Order, error) {
	panic(fmt.Errorf("not implemented: Orders - orders"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
