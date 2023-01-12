package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// DogProfileBreed holds the schema definition for the DogProfileBreed entity.
type DogProfileBreed struct {
	ent.Schema
}

// Fields of the DogProfileBreed.
func (DogProfileBreed) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Unique(),
		field.UUID("breed_id", uuid.UUID{}),
		field.UUID("dog_id", uuid.UUID{}),
		field.Float("percentage").Positive(),
		field.Time("updated_at").Default(time.Now),
		field.Time("created_at").Default(time.Now),
	}
}

// Edges of the DogProfileBreed.
func (DogProfileBreed) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("dog", Dog.Type).
			Ref("breedProfiles").
			Field("dog_id").
			Required().
			Unique(),

		edge.From("breed", Breed.Type).
			Ref("dogProfiles").
			Field("breed_id").
			Required().
			Unique(),
	}
}
