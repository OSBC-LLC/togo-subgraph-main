package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/OSBC-LLC/togo-subgraph-main/ent"
	"github.com/OSBC-LLC/togo-subgraph-main/graph/generated"
	"github.com/google/uuid"
)

// Node is the resolver for the node field.
func (r *queryResolver) Node(ctx context.Context, id uuid.UUID) (ent.Noder, error) {
	return r.client.Noder(ctx, id)
}

// Nodes is the resolver for the nodes field.
func (r *queryResolver) Nodes(ctx context.Context, ids []uuid.UUID) ([]ent.Noder, error) {
	return r.client.Noders(ctx, ids)
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
