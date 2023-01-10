package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/OSBC-LLC/togo-subgraph-main/ent"
	"github.com/OSBC-LLC/togo-subgraph-main/graph/generated"
	"github.com/google/uuid"
)

// DogImgID is the resolver for the dogImgID field.
func (r *dogResolver) DogImgID(ctx context.Context, obj *ent.Dog) (string, error) {
	panic(fmt.Errorf("not implemented: DogImgID - dogImgID"))
}

// BreedID is the resolver for the breedID field.
func (r *dogProfileBreedResolver) BreedID(ctx context.Context, obj *ent.DogProfileBreed) (string, error) {
	panic(fmt.Errorf("not implemented: BreedID - breedID"))
}

// DogID is the resolver for the dogID field.
func (r *dogProfileBreedResolver) DogID(ctx context.Context, obj *ent.DogProfileBreed) (string, error) {
	panic(fmt.Errorf("not implemented: DogID - dogID"))
}

// OwnerID is the resolver for the ownerID field.
func (r *dogProfileOwnerResolver) OwnerID(ctx context.Context, obj *ent.DogProfileOwner) (string, error) {
	panic(fmt.Errorf("not implemented: OwnerID - ownerID"))
}

// DogID is the resolver for the dogID field.
func (r *dogProfileOwnerResolver) DogID(ctx context.Context, obj *ent.DogProfileOwner) (string, error) {
	panic(fmt.Errorf("not implemented: DogID - dogID"))
}

// Node is the resolver for the node field.
func (r *queryResolver) Node(ctx context.Context, id uuid.UUID) (ent.Noder, error) {
	panic(fmt.Errorf("not implemented"))
}

// Nodes is the resolver for the nodes field.
func (r *queryResolver) Nodes(ctx context.Context, ids []uuid.UUID) ([]ent.Noder, error) {
	panic(fmt.Errorf("not implemented"))
}

// UserImageID is the resolver for the userImageID field.
func (r *userResolver) UserImageID(ctx context.Context, obj *ent.User) (string, error) {
	panic(fmt.Errorf("not implemented: UserImageID - userImageID"))
}

// ProfileID is the resolver for the profileID field.
func (r *userResolver) ProfileID(ctx context.Context, obj *ent.User) (string, error) {
	panic(fmt.Errorf("not implemented: ProfileID - profileID"))
}

// BreedID is the resolver for the breedID field.
func (r *dogProfileBreedWhereInputResolver) BreedID(ctx context.Context, obj *ent.DogProfileBreedWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: BreedID - breedID"))
}

// BreedIDNeq is the resolver for the breedIDNEQ field.
func (r *dogProfileBreedWhereInputResolver) BreedIDNeq(ctx context.Context, obj *ent.DogProfileBreedWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: BreedIDNeq - breedIDNEQ"))
}

// BreedIDIn is the resolver for the breedIDIn field.
func (r *dogProfileBreedWhereInputResolver) BreedIDIn(ctx context.Context, obj *ent.DogProfileBreedWhereInput, data []string) error {
	panic(fmt.Errorf("not implemented: BreedIDIn - breedIDIn"))
}

// BreedIDNotIn is the resolver for the breedIDNotIn field.
func (r *dogProfileBreedWhereInputResolver) BreedIDNotIn(ctx context.Context, obj *ent.DogProfileBreedWhereInput, data []string) error {
	panic(fmt.Errorf("not implemented: BreedIDNotIn - breedIDNotIn"))
}

// BreedIDGt is the resolver for the breedIDGT field.
func (r *dogProfileBreedWhereInputResolver) BreedIDGt(ctx context.Context, obj *ent.DogProfileBreedWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: BreedIDGt - breedIDGT"))
}

// BreedIDGte is the resolver for the breedIDGTE field.
func (r *dogProfileBreedWhereInputResolver) BreedIDGte(ctx context.Context, obj *ent.DogProfileBreedWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: BreedIDGte - breedIDGTE"))
}

// BreedIDLt is the resolver for the breedIDLT field.
func (r *dogProfileBreedWhereInputResolver) BreedIDLt(ctx context.Context, obj *ent.DogProfileBreedWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: BreedIDLt - breedIDLT"))
}

// BreedIDLte is the resolver for the breedIDLTE field.
func (r *dogProfileBreedWhereInputResolver) BreedIDLte(ctx context.Context, obj *ent.DogProfileBreedWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: BreedIDLte - breedIDLTE"))
}

// DogID is the resolver for the dogID field.
func (r *dogProfileBreedWhereInputResolver) DogID(ctx context.Context, obj *ent.DogProfileBreedWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: DogID - dogID"))
}

// DogIDNeq is the resolver for the dogIDNEQ field.
func (r *dogProfileBreedWhereInputResolver) DogIDNeq(ctx context.Context, obj *ent.DogProfileBreedWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: DogIDNeq - dogIDNEQ"))
}

// DogIDIn is the resolver for the dogIDIn field.
func (r *dogProfileBreedWhereInputResolver) DogIDIn(ctx context.Context, obj *ent.DogProfileBreedWhereInput, data []string) error {
	panic(fmt.Errorf("not implemented: DogIDIn - dogIDIn"))
}

// DogIDNotIn is the resolver for the dogIDNotIn field.
func (r *dogProfileBreedWhereInputResolver) DogIDNotIn(ctx context.Context, obj *ent.DogProfileBreedWhereInput, data []string) error {
	panic(fmt.Errorf("not implemented: DogIDNotIn - dogIDNotIn"))
}

// DogIDGt is the resolver for the dogIDGT field.
func (r *dogProfileBreedWhereInputResolver) DogIDGt(ctx context.Context, obj *ent.DogProfileBreedWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: DogIDGt - dogIDGT"))
}

// DogIDGte is the resolver for the dogIDGTE field.
func (r *dogProfileBreedWhereInputResolver) DogIDGte(ctx context.Context, obj *ent.DogProfileBreedWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: DogIDGte - dogIDGTE"))
}

// DogIDLt is the resolver for the dogIDLT field.
func (r *dogProfileBreedWhereInputResolver) DogIDLt(ctx context.Context, obj *ent.DogProfileBreedWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: DogIDLt - dogIDLT"))
}

// DogIDLte is the resolver for the dogIDLTE field.
func (r *dogProfileBreedWhereInputResolver) DogIDLte(ctx context.Context, obj *ent.DogProfileBreedWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: DogIDLte - dogIDLTE"))
}

// OwnerID is the resolver for the ownerID field.
func (r *dogProfileOwnerWhereInputResolver) OwnerID(ctx context.Context, obj *ent.DogProfileOwnerWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: OwnerID - ownerID"))
}

// OwnerIDNeq is the resolver for the ownerIDNEQ field.
func (r *dogProfileOwnerWhereInputResolver) OwnerIDNeq(ctx context.Context, obj *ent.DogProfileOwnerWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: OwnerIDNeq - ownerIDNEQ"))
}

// OwnerIDIn is the resolver for the ownerIDIn field.
func (r *dogProfileOwnerWhereInputResolver) OwnerIDIn(ctx context.Context, obj *ent.DogProfileOwnerWhereInput, data []string) error {
	panic(fmt.Errorf("not implemented: OwnerIDIn - ownerIDIn"))
}

// OwnerIDNotIn is the resolver for the ownerIDNotIn field.
func (r *dogProfileOwnerWhereInputResolver) OwnerIDNotIn(ctx context.Context, obj *ent.DogProfileOwnerWhereInput, data []string) error {
	panic(fmt.Errorf("not implemented: OwnerIDNotIn - ownerIDNotIn"))
}

// OwnerIDGt is the resolver for the ownerIDGT field.
func (r *dogProfileOwnerWhereInputResolver) OwnerIDGt(ctx context.Context, obj *ent.DogProfileOwnerWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: OwnerIDGt - ownerIDGT"))
}

// OwnerIDGte is the resolver for the ownerIDGTE field.
func (r *dogProfileOwnerWhereInputResolver) OwnerIDGte(ctx context.Context, obj *ent.DogProfileOwnerWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: OwnerIDGte - ownerIDGTE"))
}

// OwnerIDLt is the resolver for the ownerIDLT field.
func (r *dogProfileOwnerWhereInputResolver) OwnerIDLt(ctx context.Context, obj *ent.DogProfileOwnerWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: OwnerIDLt - ownerIDLT"))
}

// OwnerIDLte is the resolver for the ownerIDLTE field.
func (r *dogProfileOwnerWhereInputResolver) OwnerIDLte(ctx context.Context, obj *ent.DogProfileOwnerWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: OwnerIDLte - ownerIDLTE"))
}

// DogID is the resolver for the dogID field.
func (r *dogProfileOwnerWhereInputResolver) DogID(ctx context.Context, obj *ent.DogProfileOwnerWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: DogID - dogID"))
}

// DogIDNeq is the resolver for the dogIDNEQ field.
func (r *dogProfileOwnerWhereInputResolver) DogIDNeq(ctx context.Context, obj *ent.DogProfileOwnerWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: DogIDNeq - dogIDNEQ"))
}

// DogIDIn is the resolver for the dogIDIn field.
func (r *dogProfileOwnerWhereInputResolver) DogIDIn(ctx context.Context, obj *ent.DogProfileOwnerWhereInput, data []string) error {
	panic(fmt.Errorf("not implemented: DogIDIn - dogIDIn"))
}

// DogIDNotIn is the resolver for the dogIDNotIn field.
func (r *dogProfileOwnerWhereInputResolver) DogIDNotIn(ctx context.Context, obj *ent.DogProfileOwnerWhereInput, data []string) error {
	panic(fmt.Errorf("not implemented: DogIDNotIn - dogIDNotIn"))
}

// DogIDGt is the resolver for the dogIDGT field.
func (r *dogProfileOwnerWhereInputResolver) DogIDGt(ctx context.Context, obj *ent.DogProfileOwnerWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: DogIDGt - dogIDGT"))
}

// DogIDGte is the resolver for the dogIDGTE field.
func (r *dogProfileOwnerWhereInputResolver) DogIDGte(ctx context.Context, obj *ent.DogProfileOwnerWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: DogIDGte - dogIDGTE"))
}

// DogIDLt is the resolver for the dogIDLT field.
func (r *dogProfileOwnerWhereInputResolver) DogIDLt(ctx context.Context, obj *ent.DogProfileOwnerWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: DogIDLt - dogIDLT"))
}

// DogIDLte is the resolver for the dogIDLTE field.
func (r *dogProfileOwnerWhereInputResolver) DogIDLte(ctx context.Context, obj *ent.DogProfileOwnerWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: DogIDLte - dogIDLTE"))
}

// DogImgID is the resolver for the dogImgID field.
func (r *dogWhereInputResolver) DogImgID(ctx context.Context, obj *ent.DogWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: DogImgID - dogImgID"))
}

// DogImgIDNeq is the resolver for the dogImgIDNEQ field.
func (r *dogWhereInputResolver) DogImgIDNeq(ctx context.Context, obj *ent.DogWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: DogImgIDNeq - dogImgIDNEQ"))
}

// DogImgIDIn is the resolver for the dogImgIDIn field.
func (r *dogWhereInputResolver) DogImgIDIn(ctx context.Context, obj *ent.DogWhereInput, data []string) error {
	panic(fmt.Errorf("not implemented: DogImgIDIn - dogImgIDIn"))
}

// DogImgIDNotIn is the resolver for the dogImgIDNotIn field.
func (r *dogWhereInputResolver) DogImgIDNotIn(ctx context.Context, obj *ent.DogWhereInput, data []string) error {
	panic(fmt.Errorf("not implemented: DogImgIDNotIn - dogImgIDNotIn"))
}

// DogImgIDGt is the resolver for the dogImgIDGT field.
func (r *dogWhereInputResolver) DogImgIDGt(ctx context.Context, obj *ent.DogWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: DogImgIDGt - dogImgIDGT"))
}

// DogImgIDGte is the resolver for the dogImgIDGTE field.
func (r *dogWhereInputResolver) DogImgIDGte(ctx context.Context, obj *ent.DogWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: DogImgIDGte - dogImgIDGTE"))
}

// DogImgIDLt is the resolver for the dogImgIDLT field.
func (r *dogWhereInputResolver) DogImgIDLt(ctx context.Context, obj *ent.DogWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: DogImgIDLt - dogImgIDLT"))
}

// DogImgIDLte is the resolver for the dogImgIDLTE field.
func (r *dogWhereInputResolver) DogImgIDLte(ctx context.Context, obj *ent.DogWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: DogImgIDLte - dogImgIDLTE"))
}

// UserImageID is the resolver for the userImageID field.
func (r *userWhereInputResolver) UserImageID(ctx context.Context, obj *ent.UserWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: UserImageID - userImageID"))
}

// UserImageIDNeq is the resolver for the userImageIDNEQ field.
func (r *userWhereInputResolver) UserImageIDNeq(ctx context.Context, obj *ent.UserWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: UserImageIDNeq - userImageIDNEQ"))
}

// UserImageIDIn is the resolver for the userImageIDIn field.
func (r *userWhereInputResolver) UserImageIDIn(ctx context.Context, obj *ent.UserWhereInput, data []string) error {
	panic(fmt.Errorf("not implemented: UserImageIDIn - userImageIDIn"))
}

// UserImageIDNotIn is the resolver for the userImageIDNotIn field.
func (r *userWhereInputResolver) UserImageIDNotIn(ctx context.Context, obj *ent.UserWhereInput, data []string) error {
	panic(fmt.Errorf("not implemented: UserImageIDNotIn - userImageIDNotIn"))
}

// UserImageIDGt is the resolver for the userImageIDGT field.
func (r *userWhereInputResolver) UserImageIDGt(ctx context.Context, obj *ent.UserWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: UserImageIDGt - userImageIDGT"))
}

// UserImageIDGte is the resolver for the userImageIDGTE field.
func (r *userWhereInputResolver) UserImageIDGte(ctx context.Context, obj *ent.UserWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: UserImageIDGte - userImageIDGTE"))
}

// UserImageIDLt is the resolver for the userImageIDLT field.
func (r *userWhereInputResolver) UserImageIDLt(ctx context.Context, obj *ent.UserWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: UserImageIDLt - userImageIDLT"))
}

// UserImageIDLte is the resolver for the userImageIDLTE field.
func (r *userWhereInputResolver) UserImageIDLte(ctx context.Context, obj *ent.UserWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: UserImageIDLte - userImageIDLTE"))
}

// ProfileID is the resolver for the profileID field.
func (r *userWhereInputResolver) ProfileID(ctx context.Context, obj *ent.UserWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: ProfileID - profileID"))
}

// ProfileIDNeq is the resolver for the profileIDNEQ field.
func (r *userWhereInputResolver) ProfileIDNeq(ctx context.Context, obj *ent.UserWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: ProfileIDNeq - profileIDNEQ"))
}

// ProfileIDIn is the resolver for the profileIDIn field.
func (r *userWhereInputResolver) ProfileIDIn(ctx context.Context, obj *ent.UserWhereInput, data []string) error {
	panic(fmt.Errorf("not implemented: ProfileIDIn - profileIDIn"))
}

// ProfileIDNotIn is the resolver for the profileIDNotIn field.
func (r *userWhereInputResolver) ProfileIDNotIn(ctx context.Context, obj *ent.UserWhereInput, data []string) error {
	panic(fmt.Errorf("not implemented: ProfileIDNotIn - profileIDNotIn"))
}

// ProfileIDGt is the resolver for the profileIDGT field.
func (r *userWhereInputResolver) ProfileIDGt(ctx context.Context, obj *ent.UserWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: ProfileIDGt - profileIDGT"))
}

// ProfileIDGte is the resolver for the profileIDGTE field.
func (r *userWhereInputResolver) ProfileIDGte(ctx context.Context, obj *ent.UserWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: ProfileIDGte - profileIDGTE"))
}

// ProfileIDLt is the resolver for the profileIDLT field.
func (r *userWhereInputResolver) ProfileIDLt(ctx context.Context, obj *ent.UserWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: ProfileIDLt - profileIDLT"))
}

// ProfileIDLte is the resolver for the profileIDLTE field.
func (r *userWhereInputResolver) ProfileIDLte(ctx context.Context, obj *ent.UserWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: ProfileIDLte - profileIDLTE"))
}

// Dog returns generated.DogResolver implementation.
func (r *Resolver) Dog() generated.DogResolver { return &dogResolver{r} }

// DogProfileBreed returns generated.DogProfileBreedResolver implementation.
func (r *Resolver) DogProfileBreed() generated.DogProfileBreedResolver {
	return &dogProfileBreedResolver{r}
}

// DogProfileOwner returns generated.DogProfileOwnerResolver implementation.
func (r *Resolver) DogProfileOwner() generated.DogProfileOwnerResolver {
	return &dogProfileOwnerResolver{r}
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

// DogProfileBreedWhereInput returns generated.DogProfileBreedWhereInputResolver implementation.
func (r *Resolver) DogProfileBreedWhereInput() generated.DogProfileBreedWhereInputResolver {
	return &dogProfileBreedWhereInputResolver{r}
}

// DogProfileOwnerWhereInput returns generated.DogProfileOwnerWhereInputResolver implementation.
func (r *Resolver) DogProfileOwnerWhereInput() generated.DogProfileOwnerWhereInputResolver {
	return &dogProfileOwnerWhereInputResolver{r}
}

// DogWhereInput returns generated.DogWhereInputResolver implementation.
func (r *Resolver) DogWhereInput() generated.DogWhereInputResolver { return &dogWhereInputResolver{r} }

// UserWhereInput returns generated.UserWhereInputResolver implementation.
func (r *Resolver) UserWhereInput() generated.UserWhereInputResolver {
	return &userWhereInputResolver{r}
}

type dogResolver struct{ *Resolver }
type dogProfileBreedResolver struct{ *Resolver }
type dogProfileOwnerResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
type dogProfileBreedWhereInputResolver struct{ *Resolver }
type dogProfileOwnerWhereInputResolver struct{ *Resolver }
type dogWhereInputResolver struct{ *Resolver }
type userWhereInputResolver struct{ *Resolver }
