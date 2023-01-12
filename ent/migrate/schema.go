// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// BreedsColumns holds the columns for the "breeds" table.
	BreedsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "created_at", Type: field.TypeTime},
	}
	// BreedsTable holds the schema information for the "breeds" table.
	BreedsTable = &schema.Table{
		Name:       "breeds",
		Columns:    BreedsColumns,
		PrimaryKey: []*schema.Column{BreedsColumns[0]},
	}
	// DogsColumns holds the columns for the "dogs" table.
	DogsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "full_name", Type: field.TypeString},
		{Name: "age", Type: field.TypeInt},
		{Name: "weight_lbs", Type: field.TypeFloat64},
		{Name: "weight_kgs", Type: field.TypeFloat64},
		{Name: "size", Type: field.TypeString},
		{Name: "birthday", Type: field.TypeTime},
		{Name: "dog_img_id", Type: field.TypeUUID},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "created_at", Type: field.TypeTime},
	}
	// DogsTable holds the schema information for the "dogs" table.
	DogsTable = &schema.Table{
		Name:       "dogs",
		Columns:    DogsColumns,
		PrimaryKey: []*schema.Column{DogsColumns[0]},
	}
	// DogProfileBreedsColumns holds the columns for the "dog_profile_breeds" table.
	DogProfileBreedsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "breed_id", Type: field.TypeUUID},
		{Name: "dog_id", Type: field.TypeUUID},
		{Name: "percentage", Type: field.TypeFloat64},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "created_at", Type: field.TypeTime},
	}
	// DogProfileBreedsTable holds the schema information for the "dog_profile_breeds" table.
	DogProfileBreedsTable = &schema.Table{
		Name:       "dog_profile_breeds",
		Columns:    DogProfileBreedsColumns,
		PrimaryKey: []*schema.Column{DogProfileBreedsColumns[0]},
	}
	// DogProfileOwnersColumns holds the columns for the "dog_profile_owners" table.
	DogProfileOwnersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "dog_id", Type: field.TypeUUID},
		{Name: "owner_id", Type: field.TypeUUID},
	}
	// DogProfileOwnersTable holds the schema information for the "dog_profile_owners" table.
	DogProfileOwnersTable = &schema.Table{
		Name:       "dog_profile_owners",
		Columns:    DogProfileOwnersColumns,
		PrimaryKey: []*schema.Column{DogProfileOwnersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "dog_profile_owners_dogs_ownerProfiles",
				Columns:    []*schema.Column{DogProfileOwnersColumns[3]},
				RefColumns: []*schema.Column{DogsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "dog_profile_owners_users_dogProfiles",
				Columns:    []*schema.Column{DogProfileOwnersColumns[4]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// ImagesColumns holds the columns for the "images" table.
	ImagesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "url", Type: field.TypeString},
		{Name: "width", Type: field.TypeInt},
		{Name: "height", Type: field.TypeInt},
		{Name: "type", Type: field.TypeString},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "created_at", Type: field.TypeTime},
	}
	// ImagesTable holds the schema information for the "images" table.
	ImagesTable = &schema.Table{
		Name:       "images",
		Columns:    ImagesColumns,
		PrimaryKey: []*schema.Column{ImagesColumns[0]},
	}
	// ProfilesColumns holds the columns for the "profiles" table.
	ProfilesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "description", Type: field.TypeString},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "created_at", Type: field.TypeTime},
	}
	// ProfilesTable holds the schema information for the "profiles" table.
	ProfilesTable = &schema.Table{
		Name:       "profiles",
		Columns:    ProfilesColumns,
		PrimaryKey: []*schema.Column{ProfilesColumns[0]},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "first_name", Type: field.TypeString},
		{Name: "last_name", Type: field.TypeString},
		{Name: "user_image_id", Type: field.TypeUUID},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "profile_id", Type: field.TypeUUID},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "users_profiles_users",
				Columns:    []*schema.Column{UsersColumns[6]},
				RefColumns: []*schema.Column{ProfilesColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		BreedsTable,
		DogsTable,
		DogProfileBreedsTable,
		DogProfileOwnersTable,
		ImagesTable,
		ProfilesTable,
		UsersTable,
	}
)

func init() {
	DogProfileOwnersTable.ForeignKeys[0].RefTable = DogsTable
	DogProfileOwnersTable.ForeignKeys[1].RefTable = UsersTable
	UsersTable.ForeignKeys[0].RefTable = ProfilesTable
}
