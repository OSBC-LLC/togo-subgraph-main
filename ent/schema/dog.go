package schema

import "entgo.io/ent"

// Dog holds the schema definition for the Dog entity.
type Dog struct {
	ent.Schema
}

// Fields of the Dog.
func (Dog) Fields() []ent.Field {
	return nil
}

// Edges of the Dog.
func (Dog) Edges() []ent.Edge {
	return nil
}
