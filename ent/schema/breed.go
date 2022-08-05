package schema

import "entgo.io/ent"

// Breed holds the schema definition for the Breed entity.
type Breed struct {
	ent.Schema
}

// Fields of the Breed.
func (Breed) Fields() []ent.Field {
	return nil
}

// Edges of the Breed.
func (Breed) Edges() []ent.Edge {
	return nil
}
