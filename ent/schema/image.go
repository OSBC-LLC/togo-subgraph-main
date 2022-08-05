package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type ImageType int

const (
	DogImgType ImageType = iota
	UserImgType
)

func (i ImageType) String() string {
	switch i {
	case DogImgType:
		return "DOG"
	case UserImgType:
		return "USER"
	}
	return "undefined"
}

// Image holds the schema definition for the Image entity.
type Image struct {
	ent.Schema
}

// Fields of the Image.
func (Image) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Unique(),
		field.String("url"),
		field.Int("width").Positive(),
		field.Int("height").Positive(),
		field.String("type"),
		field.Time("updated_at").Default(time.Now),
		field.Time("created_at").Default(time.Now),
	}
}

// Edges of the Image.
func (Image) Edges() []ent.Edge {
	return nil
}
