package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"time"

	"github.com/OSBC-LLC/togo-subgraph-main/ent"
	"github.com/OSBC-LLC/togo-subgraph-main/graph/generated"
	"github.com/OSBC-LLC/togo-subgraph-main/graph/model"
	"github.com/google/uuid"
)

// CreateProfile is the resolver for the createProfile field.
func (r *mutationResolver) CreateProfile(ctx context.Context, session model.Session, details model.NewProfile) (*ent.Profile, error) {
	t := time.Now()
	return r.client.Profile.Create().
		SetID(uuid.New()).
		SetName(details.Name).
		SetDescription(details.Description).
		SetCreatedAt(t).
		SetUpdatedAt(t).
		Save(ctx)
}

// CreateImage is the resolver for the createImage field.
func (r *mutationResolver) CreateImage(ctx context.Context, session model.Session, details model.NewImage) (*ent.Image, error) {
	t := time.Now()
	return r.client.Image.Create().
		SetID(uuid.New()).
		SetURL(details.URL).
		SetWidth(details.Width).
		SetHeight(details.Height).
		SetType(details.Type).
		SetCreatedAt(t).
		SetUpdatedAt(t).
		Save(ctx)
}

// GetAppData is the resolver for the getAppData field.
func (r *queryResolver) GetAppData(ctx context.Context, session model.Session) (*ent.User, error) {
	panic(fmt.Errorf("not implemented: GetAppData - getAppData"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
