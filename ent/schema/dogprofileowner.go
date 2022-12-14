package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// DogProfileOwner holds the schema definition for the DogProfileOwner entity.
type DogProfileOwner struct {
	ent.Schema
}

// Fields of the DogProfileOwner.
func (DogProfileOwner) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Unique(),
		field.UUID("owner_id", uuid.UUID{}),
		field.UUID("dog_id", uuid.UUID{}),
		field.Time("updated_at").Default(time.Now),
		field.Time("created_at").Default(time.Now),
	}
}

// Edges of the DogProfileOwner.
func (DogProfileOwner) Edges() []ent.Edge {
	return nil
}
