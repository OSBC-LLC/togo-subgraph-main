package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type DogSize int

const (
	XSmall DogSize = iota
	Small
	Medium
	Large
	XLarge
)

func (i DogSize) String() string {
	switch i {
	case XSmall:
		return "XSMALL"
	case Small:
		return "SMALL"
	case Medium:
		return "MEDIUM"
	case Large:
		return "LARGE"
	case XLarge:
		return "XLARGE"
	}
	return "undefined"
}

// Dog holds the schema definition for the Dog entity.
type Dog struct {
	ent.Schema
}

// Fields of the Dog.
func (Dog) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Unique(),
		field.String("full_name"),
		field.Int("age").Positive(),
		field.Float("weight_lbs").Positive(),
		field.Float("weight_kgs").Positive(),
		field.String("size"),
		field.Time("birthday"),
		field.UUID("dog_img_id", uuid.UUID{}),
		field.Time("updated_at").Default(time.Now),
		field.Time("created_at").Default(time.Now),
	}
}

// Edges of the Dog.
func (Dog) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("image", Image.Type).
			Ref("dogs").
			Field("dog_img_id").
			Required().
			Unique(),

		edge.To("ownerProfiles", DogProfileOwner.Type),
		edge.To("breedProfiles", DogProfileBreed.Type),
	}
}
