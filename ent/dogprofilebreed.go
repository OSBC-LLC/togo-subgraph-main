// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/OSBC-LLC/togo-subgraph-main/ent/dogprofilebreed"
)

// DogProfileBreed is the model entity for the DogProfileBreed schema.
type DogProfileBreed struct {
	config
	// ID of the ent.
	ID int `json:"id,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*DogProfileBreed) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case dogprofilebreed.FieldID:
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type DogProfileBreed", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the DogProfileBreed fields.
func (dpb *DogProfileBreed) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case dogprofilebreed.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			dpb.ID = int(value.Int64)
		}
	}
	return nil
}

// Update returns a builder for updating this DogProfileBreed.
// Note that you need to call DogProfileBreed.Unwrap() before calling this method if this DogProfileBreed
// was returned from a transaction, and the transaction was committed or rolled back.
func (dpb *DogProfileBreed) Update() *DogProfileBreedUpdateOne {
	return (&DogProfileBreedClient{config: dpb.config}).UpdateOne(dpb)
}

// Unwrap unwraps the DogProfileBreed entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (dpb *DogProfileBreed) Unwrap() *DogProfileBreed {
	_tx, ok := dpb.config.driver.(*txDriver)
	if !ok {
		panic("ent: DogProfileBreed is not a transactional entity")
	}
	dpb.config.driver = _tx.drv
	return dpb
}

// String implements the fmt.Stringer.
func (dpb *DogProfileBreed) String() string {
	var builder strings.Builder
	builder.WriteString("DogProfileBreed(")
	builder.WriteString(fmt.Sprintf("id=%v", dpb.ID))
	builder.WriteByte(')')
	return builder.String()
}

// DogProfileBreeds is a parsable slice of DogProfileBreed.
type DogProfileBreeds []*DogProfileBreed

func (dpb DogProfileBreeds) config(cfg config) {
	for _i := range dpb {
		dpb[_i].config = cfg
	}
}