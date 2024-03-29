// Code generated by ent, DO NOT EDIT.

package image

import (
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the image type in the database.
	Label = "image"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldURL holds the string denoting the url field in the database.
	FieldURL = "url"
	// FieldWidth holds the string denoting the width field in the database.
	FieldWidth = "width"
	// FieldHeight holds the string denoting the height field in the database.
	FieldHeight = "height"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// EdgeUsers holds the string denoting the users edge name in mutations.
	EdgeUsers = "users"
	// EdgeDogs holds the string denoting the dogs edge name in mutations.
	EdgeDogs = "dogs"
	// Table holds the table name of the image in the database.
	Table = "images"
	// UsersTable is the table that holds the users relation/edge.
	UsersTable = "users"
	// UsersInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UsersInverseTable = "users"
	// UsersColumn is the table column denoting the users relation/edge.
	UsersColumn = "user_image_id"
	// DogsTable is the table that holds the dogs relation/edge.
	DogsTable = "dogs"
	// DogsInverseTable is the table name for the Dog entity.
	// It exists in this package in order to avoid circular dependency with the "dog" package.
	DogsInverseTable = "dogs"
	// DogsColumn is the table column denoting the dogs relation/edge.
	DogsColumn = "dog_img_id"
)

// Columns holds all SQL columns for image fields.
var Columns = []string{
	FieldID,
	FieldURL,
	FieldWidth,
	FieldHeight,
	FieldType,
	FieldUpdatedAt,
	FieldCreatedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// WidthValidator is a validator for the "width" field. It is called by the builders before save.
	WidthValidator func(int) error
	// HeightValidator is a validator for the "height" field. It is called by the builders before save.
	HeightValidator func(int) error
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
