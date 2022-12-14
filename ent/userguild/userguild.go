// Code generated by ent, DO NOT EDIT.

package userguild

import (
	"time"
)

const (
	// Label holds the string label denoting the userguild type in the database.
	Label = "user_guild"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldGuildID holds the string denoting the guild_id field in the database.
	FieldGuildID = "guild_id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldOwner holds the string denoting the owner field in the database.
	FieldOwner = "owner"
	// FieldAdmin holds the string denoting the admin field in the database.
	FieldAdmin = "admin"
	// FieldFeatures holds the string denoting the features field in the database.
	FieldFeatures = "features"
	// FieldIconHash holds the string denoting the icon_hash field in the database.
	FieldIconHash = "icon_hash"
	// FieldIconURL holds the string denoting the icon_url field in the database.
	FieldIconURL = "icon_url"
	// FieldPermissions holds the string denoting the permissions field in the database.
	FieldPermissions = "permissions"
	// EdgeAdmins holds the string denoting the admins edge name in mutations.
	EdgeAdmins = "admins"
	// Table holds the table name of the userguild in the database.
	Table = "user_guilds"
	// AdminsTable is the table that holds the admins relation/edge. The primary key declared below.
	AdminsTable = "user_guilds"
	// AdminsInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	AdminsInverseTable = "users"
)

// Columns holds all SQL columns for userguild fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldGuildID,
	FieldName,
	FieldOwner,
	FieldAdmin,
	FieldFeatures,
	FieldIconHash,
	FieldIconURL,
	FieldPermissions,
}

var (
	// AdminsPrimaryKey and AdminsColumn2 are the table columns denoting the
	// primary key for the admins relation (M2M).
	AdminsPrimaryKey = []string{"user_id", "user_guild_id"}
)

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
	// DefaultCreateTime holds the default value on creation for the "create_time" field.
	DefaultCreateTime func() time.Time
	// DefaultUpdateTime holds the default value on creation for the "update_time" field.
	DefaultUpdateTime func() time.Time
	// UpdateDefaultUpdateTime holds the default value on update for the "update_time" field.
	UpdateDefaultUpdateTime func() time.Time
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// DefaultOwner holds the default value on creation for the "owner" field.
	DefaultOwner bool
	// DefaultAdmin holds the default value on creation for the "admin" field.
	DefaultAdmin bool
	// DefaultFeatures holds the default value on creation for the "features" field.
	DefaultFeatures []string
	// IconHashValidator is a validator for the "icon_hash" field. It is called by the builders before save.
	IconHashValidator func(string) error
	// IconURLValidator is a validator for the "icon_url" field. It is called by the builders before save.
	IconURLValidator func(string) error
	// DefaultPermissions holds the default value on creation for the "permissions" field.
	DefaultPermissions uint64
)
