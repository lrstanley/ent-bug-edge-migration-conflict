// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/lrstanley/ent-bug-edge-migration-conflict/ent/user"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID string `json:"user_id,omitempty"`
	// Whether or not the user is a spectrograph admin.
	Admin bool `json:"admin,omitempty"`
	// The users username, not unique across the platform.
	Username string `json:"username,omitempty"`
	// The users 4-digit discord-tag.
	Discriminator string `json:"discriminator,omitempty"`
	// The users email address.
	Email string `json:"email,omitempty"`
	// The users avatar hash.
	AvatarHash string `json:"avatar_hash,omitempty"`
	// The users avatar URL (generated if no avatar present).
	AvatarURL string `json:"avatar_url,omitempty"`
	// The users chosen language option.
	Locale string `json:"locale,omitempty"`
	// Whether the user belongs to an OAuth2 application.
	Bot bool `json:"bot,omitempty"`
	// Whether the user is an Official Discord System user (part of the urgent message system).
	System bool `json:"system,omitempty"`
	// Whether the user has two factor enabled on their account.
	MfaEnabled bool `json:"mfa_enabled,omitempty"`
	// Whether the email on this account has been verified.
	Verified bool `json:"verified,omitempty"`
	// The flags on a users account.
	Flags int `json:"flags,omitempty"`
	// The type of Nitro subscription on a users account.
	PremiumType int `json:"premium_type,omitempty"`
	// The public flags on a users account.
	PublicFlags int `json:"public_flags,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges UserEdges `json:"edges"`
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// Guilds that the user is either owner or admin of (and thus can add the connection to the bot).
	Guilds []*UserGuild `json:"guilds,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// GuildsOrErr returns the Guilds value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) GuildsOrErr() ([]*UserGuild, error) {
	if e.loadedTypes[0] {
		return e.Guilds, nil
	}
	return nil, &NotLoadedError{edge: "guilds"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case user.FieldAdmin, user.FieldBot, user.FieldSystem, user.FieldMfaEnabled, user.FieldVerified:
			values[i] = new(sql.NullBool)
		case user.FieldID, user.FieldFlags, user.FieldPremiumType, user.FieldPublicFlags:
			values[i] = new(sql.NullInt64)
		case user.FieldUserID, user.FieldUsername, user.FieldDiscriminator, user.FieldEmail, user.FieldAvatarHash, user.FieldAvatarURL, user.FieldLocale:
			values[i] = new(sql.NullString)
		case user.FieldCreateTime, user.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type User", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			u.ID = int(value.Int64)
		case user.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				u.CreateTime = value.Time
			}
		case user.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				u.UpdateTime = value.Time
			}
		case user.FieldUserID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				u.UserID = value.String
			}
		case user.FieldAdmin:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field admin", values[i])
			} else if value.Valid {
				u.Admin = value.Bool
			}
		case user.FieldUsername:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field username", values[i])
			} else if value.Valid {
				u.Username = value.String
			}
		case user.FieldDiscriminator:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field discriminator", values[i])
			} else if value.Valid {
				u.Discriminator = value.String
			}
		case user.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				u.Email = value.String
			}
		case user.FieldAvatarHash:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field avatar_hash", values[i])
			} else if value.Valid {
				u.AvatarHash = value.String
			}
		case user.FieldAvatarURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field avatar_url", values[i])
			} else if value.Valid {
				u.AvatarURL = value.String
			}
		case user.FieldLocale:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field locale", values[i])
			} else if value.Valid {
				u.Locale = value.String
			}
		case user.FieldBot:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field bot", values[i])
			} else if value.Valid {
				u.Bot = value.Bool
			}
		case user.FieldSystem:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field system", values[i])
			} else if value.Valid {
				u.System = value.Bool
			}
		case user.FieldMfaEnabled:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field mfa_enabled", values[i])
			} else if value.Valid {
				u.MfaEnabled = value.Bool
			}
		case user.FieldVerified:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field verified", values[i])
			} else if value.Valid {
				u.Verified = value.Bool
			}
		case user.FieldFlags:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field flags", values[i])
			} else if value.Valid {
				u.Flags = int(value.Int64)
			}
		case user.FieldPremiumType:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field premium_type", values[i])
			} else if value.Valid {
				u.PremiumType = int(value.Int64)
			}
		case user.FieldPublicFlags:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field public_flags", values[i])
			} else if value.Valid {
				u.PublicFlags = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryGuilds queries the "guilds" edge of the User entity.
func (u *User) QueryGuilds() *UserGuildQuery {
	return (&UserClient{config: u.config}).QueryGuilds(u)
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
	builder.WriteString("create_time=")
	builder.WriteString(u.CreateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("update_time=")
	builder.WriteString(u.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(u.UserID)
	builder.WriteString(", ")
	builder.WriteString("admin=")
	builder.WriteString(fmt.Sprintf("%v", u.Admin))
	builder.WriteString(", ")
	builder.WriteString("username=")
	builder.WriteString(u.Username)
	builder.WriteString(", ")
	builder.WriteString("discriminator=")
	builder.WriteString(u.Discriminator)
	builder.WriteString(", ")
	builder.WriteString("email=")
	builder.WriteString(u.Email)
	builder.WriteString(", ")
	builder.WriteString("avatar_hash=")
	builder.WriteString(u.AvatarHash)
	builder.WriteString(", ")
	builder.WriteString("avatar_url=")
	builder.WriteString(u.AvatarURL)
	builder.WriteString(", ")
	builder.WriteString("locale=")
	builder.WriteString(u.Locale)
	builder.WriteString(", ")
	builder.WriteString("bot=")
	builder.WriteString(fmt.Sprintf("%v", u.Bot))
	builder.WriteString(", ")
	builder.WriteString("system=")
	builder.WriteString(fmt.Sprintf("%v", u.System))
	builder.WriteString(", ")
	builder.WriteString("mfa_enabled=")
	builder.WriteString(fmt.Sprintf("%v", u.MfaEnabled))
	builder.WriteString(", ")
	builder.WriteString("verified=")
	builder.WriteString(fmt.Sprintf("%v", u.Verified))
	builder.WriteString(", ")
	builder.WriteString("flags=")
	builder.WriteString(fmt.Sprintf("%v", u.Flags))
	builder.WriteString(", ")
	builder.WriteString("premium_type=")
	builder.WriteString(fmt.Sprintf("%v", u.PremiumType))
	builder.WriteString(", ")
	builder.WriteString("public_flags=")
	builder.WriteString(fmt.Sprintf("%v", u.PublicFlags))
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
