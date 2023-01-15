package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"time"

	"github.com/OSBC-LLC/togo-subgraph-main/ent"
	"github.com/OSBC-LLC/togo-subgraph-main/ent/image"
	"github.com/OSBC-LLC/togo-subgraph-main/ent/profile"
	"github.com/OSBC-LLC/togo-subgraph-main/ent/user"
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

// CreateDefaultUser is the resolver for the createDefaultUser field.
func (r *mutationResolver) CreateDefaultUser(ctx context.Context, session model.Session, details model.NewUser) (*ent.User, error) {
	t := time.Now()
	defaultImage, err := r.client.Image.Query().Where(image.URL("default")).Only(ctx)
	if err != nil {
		return nil, err
	}
	defaultProfile, err := r.client.Profile.Query().Where(profile.Name("standard")).Only(ctx)
	if err != nil {
		return nil, err
	}
	return r.client.User.Create().
		SetID(uuid.New()).
		SetFirstName(details.FirstName).
		SetLastName(details.LastName).
		SetCreatedAt(t).
		SetUpdatedAt(t).
		SetImage(defaultImage).
		SetProfile(defaultProfile).
		Save(ctx)
}

// GetAppData is the resolver for the getAppData field.
func (r *queryResolver) GetAppData(ctx context.Context, session model.Session) (*ent.User, error) {
	return r.client.User.Query().
		Where(user.FirstName("Will")).
		Only(ctx)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
