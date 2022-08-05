package schema

import "entgo.io/ent"

// DogProfileBreed holds the schema definition for the DogProfileBreed entity.
type DogProfileBreed struct {
	ent.Schema
}

// Fields of the DogProfileBreed.
func (DogProfileBreed) Fields() []ent.Field {
	return nil
}

// Edges of the DogProfileBreed.
func (DogProfileBreed) Edges() []ent.Edge {
	return nil
}
