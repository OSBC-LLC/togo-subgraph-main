// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/OSBC-LLC/togo-subgraph-main/ent/dog"
	"github.com/OSBC-LLC/togo-subgraph-main/ent/dogprofileowner"
	"github.com/OSBC-LLC/togo-subgraph-main/ent/user"
	"github.com/google/uuid"
)

// DogProfileOwner is the model entity for the DogProfileOwner schema.
type DogProfileOwner struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// OwnerID holds the value of the "owner_id" field.
	OwnerID uuid.UUID `json:"owner_id,omitempty"`
	// DogID holds the value of the "dog_id" field.
	DogID uuid.UUID `json:"dog_id,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the DogProfileOwnerQuery when eager-loading is set.
	Edges DogProfileOwnerEdges `json:"edges"`
}

// DogProfileOwnerEdges holds the relations/edges for other nodes in the graph.
type DogProfileOwnerEdges struct {
	// Owner holds the value of the owner edge.
	Owner *User `json:"owner,omitempty"`
	// Dog holds the value of the dog edge.
	Dog *Dog `json:"dog,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
	// totalCount holds the count of the edges above.
	totalCount [1]*int
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e DogProfileOwnerEdges) OwnerOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.Owner == nil {
			// The edge owner was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Owner, nil
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// DogOrErr returns the Dog value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e DogProfileOwnerEdges) DogOrErr() (*Dog, error) {
	if e.loadedTypes[1] {
		if e.Dog == nil {
			// The edge dog was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: dog.Label}
		}
		return e.Dog, nil
	}
	return nil, &NotLoadedError{edge: "dog"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*DogProfileOwner) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case dogprofileowner.FieldUpdatedAt, dogprofileowner.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		case dogprofileowner.FieldID, dogprofileowner.FieldOwnerID, dogprofileowner.FieldDogID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type DogProfileOwner", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the DogProfileOwner fields.
func (dpo *DogProfileOwner) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case dogprofileowner.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				dpo.ID = *value
			}
		case dogprofileowner.FieldOwnerID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field owner_id", values[i])
			} else if value != nil {
				dpo.OwnerID = *value
			}
		case dogprofileowner.FieldDogID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field dog_id", values[i])
			} else if value != nil {
				dpo.DogID = *value
			}
		case dogprofileowner.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				dpo.UpdatedAt = value.Time
			}
		case dogprofileowner.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				dpo.CreatedAt = value.Time
			}
		}
	}
	return nil
}

// QueryOwner queries the "owner" edge of the DogProfileOwner entity.
func (dpo *DogProfileOwner) QueryOwner() *UserQuery {
	return (&DogProfileOwnerClient{config: dpo.config}).QueryOwner(dpo)
}

// QueryDog queries the "dog" edge of the DogProfileOwner entity.
func (dpo *DogProfileOwner) QueryDog() *DogQuery {
	return (&DogProfileOwnerClient{config: dpo.config}).QueryDog(dpo)
}

// Update returns a builder for updating this DogProfileOwner.
// Note that you need to call DogProfileOwner.Unwrap() before calling this method if this DogProfileOwner
// was returned from a transaction, and the transaction was committed or rolled back.
func (dpo *DogProfileOwner) Update() *DogProfileOwnerUpdateOne {
	return (&DogProfileOwnerClient{config: dpo.config}).UpdateOne(dpo)
}

// Unwrap unwraps the DogProfileOwner entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (dpo *DogProfileOwner) Unwrap() *DogProfileOwner {
	_tx, ok := dpo.config.driver.(*txDriver)
	if !ok {
		panic("ent: DogProfileOwner is not a transactional entity")
	}
	dpo.config.driver = _tx.drv
	return dpo
}

// String implements the fmt.Stringer.
func (dpo *DogProfileOwner) String() string {
	var builder strings.Builder
	builder.WriteString("DogProfileOwner(")
	builder.WriteString(fmt.Sprintf("id=%v, ", dpo.ID))
	builder.WriteString("owner_id=")
	builder.WriteString(fmt.Sprintf("%v", dpo.OwnerID))
	builder.WriteString(", ")
	builder.WriteString("dog_id=")
	builder.WriteString(fmt.Sprintf("%v", dpo.DogID))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(dpo.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(dpo.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// DogProfileOwners is a parsable slice of DogProfileOwner.
type DogProfileOwners []*DogProfileOwner

func (dpo DogProfileOwners) config(cfg config) {
	for _i := range dpo {
		dpo[_i].config = cfg
	}
}
