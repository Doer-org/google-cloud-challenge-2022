// Code generated by ent, DO NOT EDIT.

package googleauth

import (
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the googleauth type in the database.
	Label = "google_auth"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldAccessToken holds the string denoting the access_token field in the database.
	FieldAccessToken = "access_token"
	// FieldRefreshToken holds the string denoting the refresh_token field in the database.
	FieldRefreshToken = "refresh_token"
	// FieldExpiry holds the string denoting the expiry field in the database.
	FieldExpiry = "expiry"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// Table holds the table name of the googleauth in the database.
	Table = "google_auths"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "google_auths"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_id"
)

// Columns holds all SQL columns for googleauth fields.
var Columns = []string{
	FieldID,
	FieldUserID,
	FieldAccessToken,
	FieldRefreshToken,
	FieldExpiry,
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
	// DefaultUserID holds the default value on creation for the "user_id" field.
	DefaultUserID func() uuid.UUID
	// AccessTokenValidator is a validator for the "access_token" field. It is called by the builders before save.
	AccessTokenValidator func(string) error
	// RefreshTokenValidator is a validator for the "refresh_token" field. It is called by the builders before save.
	RefreshTokenValidator func(string) error
)
