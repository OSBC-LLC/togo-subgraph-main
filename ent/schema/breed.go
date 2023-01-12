package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Breed holds the schema definition for the Breed entity.
type Breed struct {
	ent.Schema
}

// Fields of the Breed.
func (Breed) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Unique(),
		field.String("name").Unique(),
		field.Time("updated_at").Default(time.Now),
		field.Time("created_at").Default(time.Now),
	}
}

// Edges of the Breed.
func (Breed) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("dogProfiles", DogProfileBreed.Type),
	}
}
