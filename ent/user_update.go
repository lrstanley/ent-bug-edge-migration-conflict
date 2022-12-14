// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/lrstanley/ent-bug-edge-migration-conflict/ent/predicate"
	"github.com/lrstanley/ent-bug-edge-migration-conflict/ent/user"
	"github.com/lrstanley/ent-bug-edge-migration-conflict/ent/userguild"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

// Where appends a list predicates to the UserUpdate builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.mutation.Where(ps...)
	return uu
}

// SetUpdateTime sets the "update_time" field.
func (uu *UserUpdate) SetUpdateTime(t time.Time) *UserUpdate {
	uu.mutation.SetUpdateTime(t)
	return uu
}

// SetAdmin sets the "admin" field.
func (uu *UserUpdate) SetAdmin(b bool) *UserUpdate {
	uu.mutation.SetAdmin(b)
	return uu
}

// SetNillableAdmin sets the "admin" field if the given value is not nil.
func (uu *UserUpdate) SetNillableAdmin(b *bool) *UserUpdate {
	if b != nil {
		uu.SetAdmin(*b)
	}
	return uu
}

// ClearAdmin clears the value of the "admin" field.
func (uu *UserUpdate) ClearAdmin() *UserUpdate {
	uu.mutation.ClearAdmin()
	return uu
}

// SetUsername sets the "username" field.
func (uu *UserUpdate) SetUsername(s string) *UserUpdate {
	uu.mutation.SetUsername(s)
	return uu
}

// SetDiscriminator sets the "discriminator" field.
func (uu *UserUpdate) SetDiscriminator(s string) *UserUpdate {
	uu.mutation.SetDiscriminator(s)
	return uu
}

// SetEmail sets the "email" field.
func (uu *UserUpdate) SetEmail(s string) *UserUpdate {
	uu.mutation.SetEmail(s)
	return uu
}

// SetAvatarHash sets the "avatar_hash" field.
func (uu *UserUpdate) SetAvatarHash(s string) *UserUpdate {
	uu.mutation.SetAvatarHash(s)
	return uu
}

// SetNillableAvatarHash sets the "avatar_hash" field if the given value is not nil.
func (uu *UserUpdate) SetNillableAvatarHash(s *string) *UserUpdate {
	if s != nil {
		uu.SetAvatarHash(*s)
	}
	return uu
}

// ClearAvatarHash clears the value of the "avatar_hash" field.
func (uu *UserUpdate) ClearAvatarHash() *UserUpdate {
	uu.mutation.ClearAvatarHash()
	return uu
}

// SetAvatarURL sets the "avatar_url" field.
func (uu *UserUpdate) SetAvatarURL(s string) *UserUpdate {
	uu.mutation.SetAvatarURL(s)
	return uu
}

// SetLocale sets the "locale" field.
func (uu *UserUpdate) SetLocale(s string) *UserUpdate {
	uu.mutation.SetLocale(s)
	return uu
}

// SetNillableLocale sets the "locale" field if the given value is not nil.
func (uu *UserUpdate) SetNillableLocale(s *string) *UserUpdate {
	if s != nil {
		uu.SetLocale(*s)
	}
	return uu
}

// ClearLocale clears the value of the "locale" field.
func (uu *UserUpdate) ClearLocale() *UserUpdate {
	uu.mutation.ClearLocale()
	return uu
}

// SetBot sets the "bot" field.
func (uu *UserUpdate) SetBot(b bool) *UserUpdate {
	uu.mutation.SetBot(b)
	return uu
}

// SetNillableBot sets the "bot" field if the given value is not nil.
func (uu *UserUpdate) SetNillableBot(b *bool) *UserUpdate {
	if b != nil {
		uu.SetBot(*b)
	}
	return uu
}

// ClearBot clears the value of the "bot" field.
func (uu *UserUpdate) ClearBot() *UserUpdate {
	uu.mutation.ClearBot()
	return uu
}

// SetSystem sets the "system" field.
func (uu *UserUpdate) SetSystem(b bool) *UserUpdate {
	uu.mutation.SetSystem(b)
	return uu
}

// SetNillableSystem sets the "system" field if the given value is not nil.
func (uu *UserUpdate) SetNillableSystem(b *bool) *UserUpdate {
	if b != nil {
		uu.SetSystem(*b)
	}
	return uu
}

// ClearSystem clears the value of the "system" field.
func (uu *UserUpdate) ClearSystem() *UserUpdate {
	uu.mutation.ClearSystem()
	return uu
}

// SetMfaEnabled sets the "mfa_enabled" field.
func (uu *UserUpdate) SetMfaEnabled(b bool) *UserUpdate {
	uu.mutation.SetMfaEnabled(b)
	return uu
}

// SetNillableMfaEnabled sets the "mfa_enabled" field if the given value is not nil.
func (uu *UserUpdate) SetNillableMfaEnabled(b *bool) *UserUpdate {
	if b != nil {
		uu.SetMfaEnabled(*b)
	}
	return uu
}

// ClearMfaEnabled clears the value of the "mfa_enabled" field.
func (uu *UserUpdate) ClearMfaEnabled() *UserUpdate {
	uu.mutation.ClearMfaEnabled()
	return uu
}

// SetVerified sets the "verified" field.
func (uu *UserUpdate) SetVerified(b bool) *UserUpdate {
	uu.mutation.SetVerified(b)
	return uu
}

// SetNillableVerified sets the "verified" field if the given value is not nil.
func (uu *UserUpdate) SetNillableVerified(b *bool) *UserUpdate {
	if b != nil {
		uu.SetVerified(*b)
	}
	return uu
}

// ClearVerified clears the value of the "verified" field.
func (uu *UserUpdate) ClearVerified() *UserUpdate {
	uu.mutation.ClearVerified()
	return uu
}

// SetFlags sets the "flags" field.
func (uu *UserUpdate) SetFlags(i int) *UserUpdate {
	uu.mutation.ResetFlags()
	uu.mutation.SetFlags(i)
	return uu
}

// SetNillableFlags sets the "flags" field if the given value is not nil.
func (uu *UserUpdate) SetNillableFlags(i *int) *UserUpdate {
	if i != nil {
		uu.SetFlags(*i)
	}
	return uu
}

// AddFlags adds i to the "flags" field.
func (uu *UserUpdate) AddFlags(i int) *UserUpdate {
	uu.mutation.AddFlags(i)
	return uu
}

// ClearFlags clears the value of the "flags" field.
func (uu *UserUpdate) ClearFlags() *UserUpdate {
	uu.mutation.ClearFlags()
	return uu
}

// SetPremiumType sets the "premium_type" field.
func (uu *UserUpdate) SetPremiumType(i int) *UserUpdate {
	uu.mutation.ResetPremiumType()
	uu.mutation.SetPremiumType(i)
	return uu
}

// SetNillablePremiumType sets the "premium_type" field if the given value is not nil.
func (uu *UserUpdate) SetNillablePremiumType(i *int) *UserUpdate {
	if i != nil {
		uu.SetPremiumType(*i)
	}
	return uu
}

// AddPremiumType adds i to the "premium_type" field.
func (uu *UserUpdate) AddPremiumType(i int) *UserUpdate {
	uu.mutation.AddPremiumType(i)
	return uu
}

// ClearPremiumType clears the value of the "premium_type" field.
func (uu *UserUpdate) ClearPremiumType() *UserUpdate {
	uu.mutation.ClearPremiumType()
	return uu
}

// SetPublicFlags sets the "public_flags" field.
func (uu *UserUpdate) SetPublicFlags(i int) *UserUpdate {
	uu.mutation.ResetPublicFlags()
	uu.mutation.SetPublicFlags(i)
	return uu
}

// SetNillablePublicFlags sets the "public_flags" field if the given value is not nil.
func (uu *UserUpdate) SetNillablePublicFlags(i *int) *UserUpdate {
	if i != nil {
		uu.SetPublicFlags(*i)
	}
	return uu
}

// AddPublicFlags adds i to the "public_flags" field.
func (uu *UserUpdate) AddPublicFlags(i int) *UserUpdate {
	uu.mutation.AddPublicFlags(i)
	return uu
}

// ClearPublicFlags clears the value of the "public_flags" field.
func (uu *UserUpdate) ClearPublicFlags() *UserUpdate {
	uu.mutation.ClearPublicFlags()
	return uu
}

// AddGuildIDs adds the "guilds" edge to the UserGuild entity by IDs.
func (uu *UserUpdate) AddGuildIDs(ids ...int) *UserUpdate {
	uu.mutation.AddGuildIDs(ids...)
	return uu
}

// AddGuilds adds the "guilds" edges to the UserGuild entity.
func (uu *UserUpdate) AddGuilds(u ...*UserGuild) *UserUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uu.AddGuildIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// ClearGuilds clears all "guilds" edges to the UserGuild entity.
func (uu *UserUpdate) ClearGuilds() *UserUpdate {
	uu.mutation.ClearGuilds()
	return uu
}

// RemoveGuildIDs removes the "guilds" edge to UserGuild entities by IDs.
func (uu *UserUpdate) RemoveGuildIDs(ids ...int) *UserUpdate {
	uu.mutation.RemoveGuildIDs(ids...)
	return uu
}

// RemoveGuilds removes "guilds" edges to UserGuild entities.
func (uu *UserUpdate) RemoveGuilds(u ...*UserGuild) *UserUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uu.RemoveGuildIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	uu.defaults()
	if len(uu.hooks) == 0 {
		if err = uu.check(); err != nil {
			return 0, err
		}
		affected, err = uu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = uu.check(); err != nil {
				return 0, err
			}
			uu.mutation = mutation
			affected, err = uu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(uu.hooks) - 1; i >= 0; i-- {
			if uu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uu *UserUpdate) defaults() {
	if _, ok := uu.mutation.UpdateTime(); !ok {
		v := user.UpdateDefaultUpdateTime()
		uu.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uu *UserUpdate) check() error {
	if v, ok := uu.mutation.Discriminator(); ok {
		if err := user.DiscriminatorValidator(v); err != nil {
			return &ValidationError{Name: "discriminator", err: fmt.Errorf(`ent: validator failed for field "User.discriminator": %w`, err)}
		}
	}
	if v, ok := uu.mutation.Email(); ok {
		if err := user.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "User.email": %w`, err)}
		}
	}
	if v, ok := uu.mutation.AvatarHash(); ok {
		if err := user.AvatarHashValidator(v); err != nil {
			return &ValidationError{Name: "avatar_hash", err: fmt.Errorf(`ent: validator failed for field "User.avatar_hash": %w`, err)}
		}
	}
	if v, ok := uu.mutation.AvatarURL(); ok {
		if err := user.AvatarURLValidator(v); err != nil {
			return &ValidationError{Name: "avatar_url", err: fmt.Errorf(`ent: validator failed for field "User.avatar_url": %w`, err)}
		}
	}
	if v, ok := uu.mutation.Locale(); ok {
		if err := user.LocaleValidator(v); err != nil {
			return &ValidationError{Name: "locale", err: fmt.Errorf(`ent: validator failed for field "User.locale": %w`, err)}
		}
	}
	return nil
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   user.Table,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: user.FieldID,
			},
		},
	}
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: user.FieldUpdateTime,
		})
	}
	if value, ok := uu.mutation.Admin(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: user.FieldAdmin,
		})
	}
	if uu.mutation.AdminCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Column: user.FieldAdmin,
		})
	}
	if value, ok := uu.mutation.Username(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldUsername,
		})
	}
	if value, ok := uu.mutation.Discriminator(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldDiscriminator,
		})
	}
	if value, ok := uu.mutation.Email(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldEmail,
		})
	}
	if value, ok := uu.mutation.AvatarHash(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldAvatarHash,
		})
	}
	if uu.mutation.AvatarHashCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldAvatarHash,
		})
	}
	if value, ok := uu.mutation.AvatarURL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldAvatarURL,
		})
	}
	if value, ok := uu.mutation.Locale(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldLocale,
		})
	}
	if uu.mutation.LocaleCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldLocale,
		})
	}
	if value, ok := uu.mutation.Bot(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: user.FieldBot,
		})
	}
	if uu.mutation.BotCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Column: user.FieldBot,
		})
	}
	if value, ok := uu.mutation.System(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: user.FieldSystem,
		})
	}
	if uu.mutation.SystemCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Column: user.FieldSystem,
		})
	}
	if value, ok := uu.mutation.MfaEnabled(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: user.FieldMfaEnabled,
		})
	}
	if uu.mutation.MfaEnabledCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Column: user.FieldMfaEnabled,
		})
	}
	if value, ok := uu.mutation.Verified(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: user.FieldVerified,
		})
	}
	if uu.mutation.VerifiedCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Column: user.FieldVerified,
		})
	}
	if value, ok := uu.mutation.Flags(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: user.FieldFlags,
		})
	}
	if value, ok := uu.mutation.AddedFlags(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: user.FieldFlags,
		})
	}
	if uu.mutation.FlagsCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: user.FieldFlags,
		})
	}
	if value, ok := uu.mutation.PremiumType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: user.FieldPremiumType,
		})
	}
	if value, ok := uu.mutation.AddedPremiumType(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: user.FieldPremiumType,
		})
	}
	if uu.mutation.PremiumTypeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: user.FieldPremiumType,
		})
	}
	if value, ok := uu.mutation.PublicFlags(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: user.FieldPublicFlags,
		})
	}
	if value, ok := uu.mutation.AddedPublicFlags(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: user.FieldPublicFlags,
		})
	}
	if uu.mutation.PublicFlagsCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: user.FieldPublicFlags,
		})
	}
	if uu.mutation.GuildsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.GuildsTable,
			Columns: user.GuildsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: userguild.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedGuildsIDs(); len(nodes) > 0 && !uu.mutation.GuildsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.GuildsTable,
			Columns: user.GuildsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: userguild.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.GuildsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.GuildsTable,
			Columns: user.GuildsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: userguild.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserMutation
}

// SetUpdateTime sets the "update_time" field.
func (uuo *UserUpdateOne) SetUpdateTime(t time.Time) *UserUpdateOne {
	uuo.mutation.SetUpdateTime(t)
	return uuo
}

// SetAdmin sets the "admin" field.
func (uuo *UserUpdateOne) SetAdmin(b bool) *UserUpdateOne {
	uuo.mutation.SetAdmin(b)
	return uuo
}

// SetNillableAdmin sets the "admin" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableAdmin(b *bool) *UserUpdateOne {
	if b != nil {
		uuo.SetAdmin(*b)
	}
	return uuo
}

// ClearAdmin clears the value of the "admin" field.
func (uuo *UserUpdateOne) ClearAdmin() *UserUpdateOne {
	uuo.mutation.ClearAdmin()
	return uuo
}

// SetUsername sets the "username" field.
func (uuo *UserUpdateOne) SetUsername(s string) *UserUpdateOne {
	uuo.mutation.SetUsername(s)
	return uuo
}

// SetDiscriminator sets the "discriminator" field.
func (uuo *UserUpdateOne) SetDiscriminator(s string) *UserUpdateOne {
	uuo.mutation.SetDiscriminator(s)
	return uuo
}

// SetEmail sets the "email" field.
func (uuo *UserUpdateOne) SetEmail(s string) *UserUpdateOne {
	uuo.mutation.SetEmail(s)
	return uuo
}

// SetAvatarHash sets the "avatar_hash" field.
func (uuo *UserUpdateOne) SetAvatarHash(s string) *UserUpdateOne {
	uuo.mutation.SetAvatarHash(s)
	return uuo
}

// SetNillableAvatarHash sets the "avatar_hash" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableAvatarHash(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetAvatarHash(*s)
	}
	return uuo
}

// ClearAvatarHash clears the value of the "avatar_hash" field.
func (uuo *UserUpdateOne) ClearAvatarHash() *UserUpdateOne {
	uuo.mutation.ClearAvatarHash()
	return uuo
}

// SetAvatarURL sets the "avatar_url" field.
func (uuo *UserUpdateOne) SetAvatarURL(s string) *UserUpdateOne {
	uuo.mutation.SetAvatarURL(s)
	return uuo
}

// SetLocale sets the "locale" field.
func (uuo *UserUpdateOne) SetLocale(s string) *UserUpdateOne {
	uuo.mutation.SetLocale(s)
	return uuo
}

// SetNillableLocale sets the "locale" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableLocale(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetLocale(*s)
	}
	return uuo
}

// ClearLocale clears the value of the "locale" field.
func (uuo *UserUpdateOne) ClearLocale() *UserUpdateOne {
	uuo.mutation.ClearLocale()
	return uuo
}

// SetBot sets the "bot" field.
func (uuo *UserUpdateOne) SetBot(b bool) *UserUpdateOne {
	uuo.mutation.SetBot(b)
	return uuo
}

// SetNillableBot sets the "bot" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableBot(b *bool) *UserUpdateOne {
	if b != nil {
		uuo.SetBot(*b)
	}
	return uuo
}

// ClearBot clears the value of the "bot" field.
func (uuo *UserUpdateOne) ClearBot() *UserUpdateOne {
	uuo.mutation.ClearBot()
	return uuo
}

// SetSystem sets the "system" field.
func (uuo *UserUpdateOne) SetSystem(b bool) *UserUpdateOne {
	uuo.mutation.SetSystem(b)
	return uuo
}

// SetNillableSystem sets the "system" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableSystem(b *bool) *UserUpdateOne {
	if b != nil {
		uuo.SetSystem(*b)
	}
	return uuo
}

// ClearSystem clears the value of the "system" field.
func (uuo *UserUpdateOne) ClearSystem() *UserUpdateOne {
	uuo.mutation.ClearSystem()
	return uuo
}

// SetMfaEnabled sets the "mfa_enabled" field.
func (uuo *UserUpdateOne) SetMfaEnabled(b bool) *UserUpdateOne {
	uuo.mutation.SetMfaEnabled(b)
	return uuo
}

// SetNillableMfaEnabled sets the "mfa_enabled" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableMfaEnabled(b *bool) *UserUpdateOne {
	if b != nil {
		uuo.SetMfaEnabled(*b)
	}
	return uuo
}

// ClearMfaEnabled clears the value of the "mfa_enabled" field.
func (uuo *UserUpdateOne) ClearMfaEnabled() *UserUpdateOne {
	uuo.mutation.ClearMfaEnabled()
	return uuo
}

// SetVerified sets the "verified" field.
func (uuo *UserUpdateOne) SetVerified(b bool) *UserUpdateOne {
	uuo.mutation.SetVerified(b)
	return uuo
}

// SetNillableVerified sets the "verified" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableVerified(b *bool) *UserUpdateOne {
	if b != nil {
		uuo.SetVerified(*b)
	}
	return uuo
}

// ClearVerified clears the value of the "verified" field.
func (uuo *UserUpdateOne) ClearVerified() *UserUpdateOne {
	uuo.mutation.ClearVerified()
	return uuo
}

// SetFlags sets the "flags" field.
func (uuo *UserUpdateOne) SetFlags(i int) *UserUpdateOne {
	uuo.mutation.ResetFlags()
	uuo.mutation.SetFlags(i)
	return uuo
}

// SetNillableFlags sets the "flags" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableFlags(i *int) *UserUpdateOne {
	if i != nil {
		uuo.SetFlags(*i)
	}
	return uuo
}

// AddFlags adds i to the "flags" field.
func (uuo *UserUpdateOne) AddFlags(i int) *UserUpdateOne {
	uuo.mutation.AddFlags(i)
	return uuo
}

// ClearFlags clears the value of the "flags" field.
func (uuo *UserUpdateOne) ClearFlags() *UserUpdateOne {
	uuo.mutation.ClearFlags()
	return uuo
}

// SetPremiumType sets the "premium_type" field.
func (uuo *UserUpdateOne) SetPremiumType(i int) *UserUpdateOne {
	uuo.mutation.ResetPremiumType()
	uuo.mutation.SetPremiumType(i)
	return uuo
}

// SetNillablePremiumType sets the "premium_type" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillablePremiumType(i *int) *UserUpdateOne {
	if i != nil {
		uuo.SetPremiumType(*i)
	}
	return uuo
}

// AddPremiumType adds i to the "premium_type" field.
func (uuo *UserUpdateOne) AddPremiumType(i int) *UserUpdateOne {
	uuo.mutation.AddPremiumType(i)
	return uuo
}

// ClearPremiumType clears the value of the "premium_type" field.
func (uuo *UserUpdateOne) ClearPremiumType() *UserUpdateOne {
	uuo.mutation.ClearPremiumType()
	return uuo
}

// SetPublicFlags sets the "public_flags" field.
func (uuo *UserUpdateOne) SetPublicFlags(i int) *UserUpdateOne {
	uuo.mutation.ResetPublicFlags()
	uuo.mutation.SetPublicFlags(i)
	return uuo
}

// SetNillablePublicFlags sets the "public_flags" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillablePublicFlags(i *int) *UserUpdateOne {
	if i != nil {
		uuo.SetPublicFlags(*i)
	}
	return uuo
}

// AddPublicFlags adds i to the "public_flags" field.
func (uuo *UserUpdateOne) AddPublicFlags(i int) *UserUpdateOne {
	uuo.mutation.AddPublicFlags(i)
	return uuo
}

// ClearPublicFlags clears the value of the "public_flags" field.
func (uuo *UserUpdateOne) ClearPublicFlags() *UserUpdateOne {
	uuo.mutation.ClearPublicFlags()
	return uuo
}

// AddGuildIDs adds the "guilds" edge to the UserGuild entity by IDs.
func (uuo *UserUpdateOne) AddGuildIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddGuildIDs(ids...)
	return uuo
}

// AddGuilds adds the "guilds" edges to the UserGuild entity.
func (uuo *UserUpdateOne) AddGuilds(u ...*UserGuild) *UserUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uuo.AddGuildIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// ClearGuilds clears all "guilds" edges to the UserGuild entity.
func (uuo *UserUpdateOne) ClearGuilds() *UserUpdateOne {
	uuo.mutation.ClearGuilds()
	return uuo
}

// RemoveGuildIDs removes the "guilds" edge to UserGuild entities by IDs.
func (uuo *UserUpdateOne) RemoveGuildIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemoveGuildIDs(ids...)
	return uuo
}

// RemoveGuilds removes "guilds" edges to UserGuild entities.
func (uuo *UserUpdateOne) RemoveGuilds(u ...*UserGuild) *UserUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uuo.RemoveGuildIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *UserUpdateOne) Select(field string, fields ...string) *UserUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated User entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	var (
		err  error
		node *User
	)
	uuo.defaults()
	if len(uuo.hooks) == 0 {
		if err = uuo.check(); err != nil {
			return nil, err
		}
		node, err = uuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = uuo.check(); err != nil {
				return nil, err
			}
			uuo.mutation = mutation
			node, err = uuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(uuo.hooks) - 1; i >= 0; i-- {
			if uuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, uuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*User)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from UserMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uuo *UserUpdateOne) defaults() {
	if _, ok := uuo.mutation.UpdateTime(); !ok {
		v := user.UpdateDefaultUpdateTime()
		uuo.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uuo *UserUpdateOne) check() error {
	if v, ok := uuo.mutation.Discriminator(); ok {
		if err := user.DiscriminatorValidator(v); err != nil {
			return &ValidationError{Name: "discriminator", err: fmt.Errorf(`ent: validator failed for field "User.discriminator": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.Email(); ok {
		if err := user.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "User.email": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.AvatarHash(); ok {
		if err := user.AvatarHashValidator(v); err != nil {
			return &ValidationError{Name: "avatar_hash", err: fmt.Errorf(`ent: validator failed for field "User.avatar_hash": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.AvatarURL(); ok {
		if err := user.AvatarURLValidator(v); err != nil {
			return &ValidationError{Name: "avatar_url", err: fmt.Errorf(`ent: validator failed for field "User.avatar_url": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.Locale(); ok {
		if err := user.LocaleValidator(v); err != nil {
			return &ValidationError{Name: "locale", err: fmt.Errorf(`ent: validator failed for field "User.locale": %w`, err)}
		}
	}
	return nil
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (_node *User, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   user.Table,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: user.FieldID,
			},
		},
	}
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "User.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user.FieldID)
		for _, f := range fields {
			if !user.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != user.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uuo.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: user.FieldUpdateTime,
		})
	}
	if value, ok := uuo.mutation.Admin(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: user.FieldAdmin,
		})
	}
	if uuo.mutation.AdminCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Column: user.FieldAdmin,
		})
	}
	if value, ok := uuo.mutation.Username(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldUsername,
		})
	}
	if value, ok := uuo.mutation.Discriminator(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldDiscriminator,
		})
	}
	if value, ok := uuo.mutation.Email(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldEmail,
		})
	}
	if value, ok := uuo.mutation.AvatarHash(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldAvatarHash,
		})
	}
	if uuo.mutation.AvatarHashCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldAvatarHash,
		})
	}
	if value, ok := uuo.mutation.AvatarURL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldAvatarURL,
		})
	}
	if value, ok := uuo.mutation.Locale(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldLocale,
		})
	}
	if uuo.mutation.LocaleCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldLocale,
		})
	}
	if value, ok := uuo.mutation.Bot(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: user.FieldBot,
		})
	}
	if uuo.mutation.BotCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Column: user.FieldBot,
		})
	}
	if value, ok := uuo.mutation.System(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: user.FieldSystem,
		})
	}
	if uuo.mutation.SystemCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Column: user.FieldSystem,
		})
	}
	if value, ok := uuo.mutation.MfaEnabled(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: user.FieldMfaEnabled,
		})
	}
	if uuo.mutation.MfaEnabledCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Column: user.FieldMfaEnabled,
		})
	}
	if value, ok := uuo.mutation.Verified(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: user.FieldVerified,
		})
	}
	if uuo.mutation.VerifiedCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Column: user.FieldVerified,
		})
	}
	if value, ok := uuo.mutation.Flags(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: user.FieldFlags,
		})
	}
	if value, ok := uuo.mutation.AddedFlags(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: user.FieldFlags,
		})
	}
	if uuo.mutation.FlagsCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: user.FieldFlags,
		})
	}
	if value, ok := uuo.mutation.PremiumType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: user.FieldPremiumType,
		})
	}
	if value, ok := uuo.mutation.AddedPremiumType(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: user.FieldPremiumType,
		})
	}
	if uuo.mutation.PremiumTypeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: user.FieldPremiumType,
		})
	}
	if value, ok := uuo.mutation.PublicFlags(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: user.FieldPublicFlags,
		})
	}
	if value, ok := uuo.mutation.AddedPublicFlags(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: user.FieldPublicFlags,
		})
	}
	if uuo.mutation.PublicFlagsCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: user.FieldPublicFlags,
		})
	}
	if uuo.mutation.GuildsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.GuildsTable,
			Columns: user.GuildsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: userguild.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedGuildsIDs(); len(nodes) > 0 && !uuo.mutation.GuildsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.GuildsTable,
			Columns: user.GuildsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: userguild.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.GuildsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.GuildsTable,
			Columns: user.GuildsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: userguild.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &User{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
