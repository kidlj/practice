// Code generated by entc, DO NOT EDIT.

package user

import (
	"time"

	"github.com/facebookincubator/ent"
	"github.com/kidlj/demo/ent/mixins/ent/schema"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at vertex property in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at vertex property in the database.
	FieldUpdatedAt = "updated_at"
	// FieldAge holds the string denoting the age vertex property in the database.
	FieldAge = "age"
	// FieldName holds the string denoting the name vertex property in the database.
	FieldName = "name"
	// FieldNickname holds the string denoting the nickname vertex property in the database.
	FieldNickname = "nickname"

	// Table holds the table name of the user in the database.
	Table = "users"
)

// Columns holds all SQL columns are user fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldAge,
	FieldName,
	FieldNickname,
}

var (
	mixin       = schema.User{}.Mixin()
	mixinFields = [...][]ent.Field{
		mixin[0].Fields(),
		mixin[1].Fields(),
	}
	fields = schema.User{}.Fields()

	// descCreatedAt is the schema descriptor for created_at field.
	descCreatedAt = mixinFields[0][0].Descriptor()
	// DefaultCreatedAt holds the default value on creation for the created_at field.
	DefaultCreatedAt = descCreatedAt.Default.(func() time.Time)

	// descUpdatedAt is the schema descriptor for updated_at field.
	descUpdatedAt = mixinFields[0][1].Descriptor()
	// DefaultUpdatedAt holds the default value on creation for the updated_at field.
	DefaultUpdatedAt = descUpdatedAt.Default.(func() time.Time)
	// UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	UpdateDefaultUpdatedAt = descUpdatedAt.UpdateDefault.(func() time.Time)

	// descAge is the schema descriptor for age field.
	descAge = mixinFields[1][0].Descriptor()
	// AgeValidator is a validator for the "age" field. It is called by the builders before save.
	AgeValidator = descAge.Validators[0].(func(int) error)

	// descName is the schema descriptor for name field.
	descName = mixinFields[1][1].Descriptor()
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator = descName.Validators[0].(func(string) error)
)