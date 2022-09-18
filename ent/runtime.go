// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/lrstanley/ent-bug-edge-migration-conflict/ent/schema"
	"github.com/lrstanley/ent-bug-edge-migration-conflict/ent/user"
	"github.com/lrstanley/ent-bug-edge-migration-conflict/ent/userguild"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	userMixin := schema.User{}.Mixin()
	userMixinFields0 := userMixin[0].Fields()
	_ = userMixinFields0
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescCreateTime is the schema descriptor for create_time field.
	userDescCreateTime := userMixinFields0[0].Descriptor()
	// user.DefaultCreateTime holds the default value on creation for the create_time field.
	user.DefaultCreateTime = userDescCreateTime.Default.(func() time.Time)
	// userDescUpdateTime is the schema descriptor for update_time field.
	userDescUpdateTime := userMixinFields0[1].Descriptor()
	// user.DefaultUpdateTime holds the default value on creation for the update_time field.
	user.DefaultUpdateTime = userDescUpdateTime.Default.(func() time.Time)
	// user.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	user.UpdateDefaultUpdateTime = userDescUpdateTime.UpdateDefault.(func() time.Time)
	// userDescAdmin is the schema descriptor for admin field.
	userDescAdmin := userFields[1].Descriptor()
	// user.DefaultAdmin holds the default value on creation for the admin field.
	user.DefaultAdmin = userDescAdmin.Default.(bool)
	// userDescDiscriminator is the schema descriptor for discriminator field.
	userDescDiscriminator := userFields[3].Descriptor()
	// user.DiscriminatorValidator is a validator for the "discriminator" field. It is called by the builders before save.
	user.DiscriminatorValidator = userDescDiscriminator.Validators[0].(func(string) error)
	// userDescEmail is the schema descriptor for email field.
	userDescEmail := userFields[4].Descriptor()
	// user.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	user.EmailValidator = userDescEmail.Validators[0].(func(string) error)
	// userDescAvatarHash is the schema descriptor for avatar_hash field.
	userDescAvatarHash := userFields[5].Descriptor()
	// user.AvatarHashValidator is a validator for the "avatar_hash" field. It is called by the builders before save.
	user.AvatarHashValidator = userDescAvatarHash.Validators[0].(func(string) error)
	// userDescAvatarURL is the schema descriptor for avatar_url field.
	userDescAvatarURL := userFields[6].Descriptor()
	// user.AvatarURLValidator is a validator for the "avatar_url" field. It is called by the builders before save.
	user.AvatarURLValidator = userDescAvatarURL.Validators[0].(func(string) error)
	// userDescLocale is the schema descriptor for locale field.
	userDescLocale := userFields[7].Descriptor()
	// user.LocaleValidator is a validator for the "locale" field. It is called by the builders before save.
	user.LocaleValidator = userDescLocale.Validators[0].(func(string) error)
	// userDescBot is the schema descriptor for bot field.
	userDescBot := userFields[8].Descriptor()
	// user.DefaultBot holds the default value on creation for the bot field.
	user.DefaultBot = userDescBot.Default.(bool)
	// userDescSystem is the schema descriptor for system field.
	userDescSystem := userFields[9].Descriptor()
	// user.DefaultSystem holds the default value on creation for the system field.
	user.DefaultSystem = userDescSystem.Default.(bool)
	// userDescMfaEnabled is the schema descriptor for mfa_enabled field.
	userDescMfaEnabled := userFields[10].Descriptor()
	// user.DefaultMfaEnabled holds the default value on creation for the mfa_enabled field.
	user.DefaultMfaEnabled = userDescMfaEnabled.Default.(bool)
	// userDescVerified is the schema descriptor for verified field.
	userDescVerified := userFields[11].Descriptor()
	// user.DefaultVerified holds the default value on creation for the verified field.
	user.DefaultVerified = userDescVerified.Default.(bool)
	// userDescFlags is the schema descriptor for flags field.
	userDescFlags := userFields[12].Descriptor()
	// user.DefaultFlags holds the default value on creation for the flags field.
	user.DefaultFlags = userDescFlags.Default.(int)
	// userDescPremiumType is the schema descriptor for premium_type field.
	userDescPremiumType := userFields[13].Descriptor()
	// user.DefaultPremiumType holds the default value on creation for the premium_type field.
	user.DefaultPremiumType = userDescPremiumType.Default.(int)
	// userDescPublicFlags is the schema descriptor for public_flags field.
	userDescPublicFlags := userFields[14].Descriptor()
	// user.DefaultPublicFlags holds the default value on creation for the public_flags field.
	user.DefaultPublicFlags = userDescPublicFlags.Default.(int)
	userguildMixin := schema.UserGuild{}.Mixin()
	userguildMixinFields0 := userguildMixin[0].Fields()
	_ = userguildMixinFields0
	userguildFields := schema.UserGuild{}.Fields()
	_ = userguildFields
	// userguildDescCreateTime is the schema descriptor for create_time field.
	userguildDescCreateTime := userguildMixinFields0[0].Descriptor()
	// userguild.DefaultCreateTime holds the default value on creation for the create_time field.
	userguild.DefaultCreateTime = userguildDescCreateTime.Default.(func() time.Time)
	// userguildDescUpdateTime is the schema descriptor for update_time field.
	userguildDescUpdateTime := userguildMixinFields0[1].Descriptor()
	// userguild.DefaultUpdateTime holds the default value on creation for the update_time field.
	userguild.DefaultUpdateTime = userguildDescUpdateTime.Default.(func() time.Time)
	// userguild.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	userguild.UpdateDefaultUpdateTime = userguildDescUpdateTime.UpdateDefault.(func() time.Time)
	// userguildDescName is the schema descriptor for name field.
	userguildDescName := userguildFields[1].Descriptor()
	// userguild.NameValidator is a validator for the "name" field. It is called by the builders before save.
	userguild.NameValidator = func() func(string) error {
		validators := userguildDescName.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(name string) error {
			for _, fn := range fns {
				if err := fn(name); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// userguildDescOwner is the schema descriptor for owner field.
	userguildDescOwner := userguildFields[2].Descriptor()
	// userguild.DefaultOwner holds the default value on creation for the owner field.
	userguild.DefaultOwner = userguildDescOwner.Default.(bool)
	// userguildDescAdmin is the schema descriptor for admin field.
	userguildDescAdmin := userguildFields[3].Descriptor()
	// userguild.DefaultAdmin holds the default value on creation for the admin field.
	userguild.DefaultAdmin = userguildDescAdmin.Default.(bool)
	// userguildDescFeatures is the schema descriptor for features field.
	userguildDescFeatures := userguildFields[4].Descriptor()
	// userguild.DefaultFeatures holds the default value on creation for the features field.
	userguild.DefaultFeatures = userguildDescFeatures.Default.([]string)
	// userguildDescIconHash is the schema descriptor for icon_hash field.
	userguildDescIconHash := userguildFields[5].Descriptor()
	// userguild.IconHashValidator is a validator for the "icon_hash" field. It is called by the builders before save.
	userguild.IconHashValidator = userguildDescIconHash.Validators[0].(func(string) error)
	// userguildDescIconURL is the schema descriptor for icon_url field.
	userguildDescIconURL := userguildFields[6].Descriptor()
	// userguild.IconURLValidator is a validator for the "icon_url" field. It is called by the builders before save.
	userguild.IconURLValidator = userguildDescIconURL.Validators[0].(func(string) error)
	// userguildDescPermissions is the schema descriptor for permissions field.
	userguildDescPermissions := userguildFields[7].Descriptor()
	// userguild.DefaultPermissions holds the default value on creation for the permissions field.
	userguild.DefaultPermissions = userguildDescPermissions.Default.(uint64)
}
