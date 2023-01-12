// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/OSBC-LLC/togo-subgraph-main/ent/profile"
	"github.com/OSBC-LLC/togo-subgraph-main/ent/user"
	"github.com/google/uuid"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// FirstName holds the value of the "first_name" field.
	FirstName string `json:"first_name,omitempty"`
	// LastName holds the value of the "last_name" field.
	LastName string `json:"last_name,omitempty"`
	// UserImageID holds the value of the "user_image_id" field.
	UserImageID uuid.UUID `json:"user_image_id,omitempty"`
	// ProfileID holds the value of the "profile_id" field.
	ProfileID uuid.UUID `json:"profile_id,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges UserEdges `json:"edges"`
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// Profile holds the value of the profile edge.
	Profile *Profile `json:"profile,omitempty"`
	// DogProfiles holds the value of the dogProfiles edge.
	DogProfiles []*DogProfileOwner `json:"dogProfiles,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
	// totalCount holds the count of the edges above.
	totalCount [2]*int
}

// ProfileOrErr returns the Profile value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) ProfileOrErr() (*Profile, error) {
	if e.loadedTypes[0] {
		if e.Profile == nil {
			// The edge profile was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: profile.Label}
		}
		return e.Profile, nil
	}
	return nil, &NotLoadedError{edge: "profile"}
}

// DogProfilesOrErr returns the DogProfiles value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) DogProfilesOrErr() ([]*DogProfileOwner, error) {
	if e.loadedTypes[1] {
		return e.DogProfiles, nil
	}
	return nil, &NotLoadedError{edge: "dogProfiles"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case user.FieldFirstName, user.FieldLastName:
			values[i] = new(sql.NullString)
		case user.FieldUpdatedAt, user.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		case user.FieldID, user.FieldUserImageID, user.FieldProfileID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type User", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				u.ID = *value
			}
		case user.FieldFirstName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field first_name", values[i])
			} else if value.Valid {
				u.FirstName = value.String
			}
		case user.FieldLastName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field last_name", values[i])
			} else if value.Valid {
				u.LastName = value.String
			}
		case user.FieldUserImageID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field user_image_id", values[i])
			} else if value != nil {
				u.UserImageID = *value
			}
		case user.FieldProfileID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field profile_id", values[i])
			} else if value != nil {
				u.ProfileID = *value
			}
		case user.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				u.UpdatedAt = value.Time
			}
		case user.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				u.CreatedAt = value.Time
			}
		}
	}
	return nil
}

// QueryProfile queries the "profile" edge of the User entity.
func (u *User) QueryProfile() *ProfileQuery {
	return (&UserClient{config: u.config}).QueryProfile(u)
}

// QueryDogProfiles queries the "dogProfiles" edge of the User entity.
func (u *User) QueryDogProfiles() *DogProfileOwnerQuery {
	return (&UserClient{config: u.config}).QueryDogProfiles(u)
}

// Update returns a builder for updating this User.
// Note that you need to call User.Unwrap() before calling this method if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return (&UserClient{config: u.config}).UpdateOne(u)
}

// Unwrap unwraps the User entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	_tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = _tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v, ", u.ID))
	builder.WriteString("first_name=")
	builder.WriteString(u.FirstName)
	builder.WriteString(", ")
	builder.WriteString("last_name=")
	builder.WriteString(u.LastName)
	builder.WriteString(", ")
	builder.WriteString("user_image_id=")
	builder.WriteString(fmt.Sprintf("%v", u.UserImageID))
	builder.WriteString(", ")
	builder.WriteString("profile_id=")
	builder.WriteString(fmt.Sprintf("%v", u.ProfileID))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(u.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(u.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User

func (u Users) config(cfg config) {
	for _i := range u {
		u[_i].config = cfg
	}
}
