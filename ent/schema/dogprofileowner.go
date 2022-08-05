package schema

import "entgo.io/ent"

// DogProfileOwner holds the schema definition for the DogProfileOwner entity.
type DogProfileOwner struct {
	ent.Schema
}

// Fields of the DogProfileOwner.
func (DogProfileOwner) Fields() []ent.Field {
	return nil
}

// Edges of the DogProfileOwner.
func (DogProfileOwner) Edges() []ent.Edge {
	return nil
}
