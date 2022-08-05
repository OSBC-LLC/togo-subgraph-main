// Code generated by ent, DO NOT EDIT.

package dog

const (
	// Label holds the string label denoting the dog type in the database.
	Label = "dog"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// Table holds the table name of the dog in the database.
	Table = "dogs"
)

// Columns holds all SQL columns for dog fields.
var Columns = []string{
	FieldID,
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
