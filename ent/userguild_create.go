// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/lrstanley/ent-bug-edge-migration-conflict/ent/user"
	"github.com/lrstanley/ent-bug-edge-migration-conflict/ent/userguild"
)

// UserGuildCreate is the builder for creating a UserGuild entity.
type UserGuildCreate struct {
	config
	mutation *UserGuildMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (ugc *UserGuildCreate) SetCreateTime(t time.Time) *UserGuildCreate {
	ugc.mutation.SetCreateTime(t)
	return ugc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (ugc *UserGuildCreate) SetNillableCreateTime(t *time.Time) *UserGuildCreate {
	if t != nil {
		ugc.SetCreateTime(*t)
	}
	return ugc
}

// SetUpdateTime sets the "update_time" field.
func (ugc *UserGuildCreate) SetUpdateTime(t time.Time) *UserGuildCreate {
	ugc.mutation.SetUpdateTime(t)
	return ugc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (ugc *UserGuildCreate) SetNillableUpdateTime(t *time.Time) *UserGuildCreate {
	if t != nil {
		ugc.SetUpdateTime(*t)
	}
	return ugc
}

// SetGuildID sets the "guild_id" field.
func (ugc *UserGuildCreate) SetGuildID(s string) *UserGuildCreate {
	ugc.mutation.SetGuildID(s)
	return ugc
}

// SetName sets the "name" field.
func (ugc *UserGuildCreate) SetName(s string) *UserGuildCreate {
	ugc.mutation.SetName(s)
	return ugc
}

// SetOwner sets the "owner" field.
func (ugc *UserGuildCreate) SetOwner(b bool) *UserGuildCreate {
	ugc.mutation.SetOwner(b)
	return ugc
}

// SetNillableOwner sets the "owner" field if the given value is not nil.
func (ugc *UserGuildCreate) SetNillableOwner(b *bool) *UserGuildCreate {
	if b != nil {
		ugc.SetOwner(*b)
	}
	return ugc
}

// SetAdmin sets the "admin" field.
func (ugc *UserGuildCreate) SetAdmin(b bool) *UserGuildCreate {
	ugc.mutation.SetAdmin(b)
	return ugc
}

// SetNillableAdmin sets the "admin" field if the given value is not nil.
func (ugc *UserGuildCreate) SetNillableAdmin(b *bool) *UserGuildCreate {
	if b != nil {
		ugc.SetAdmin(*b)
	}
	return ugc
}

// SetFeatures sets the "features" field.
func (ugc *UserGuildCreate) SetFeatures(s []string) *UserGuildCreate {
	ugc.mutation.SetFeatures(s)
	return ugc
}

// SetIconHash sets the "icon_hash" field.
func (ugc *UserGuildCreate) SetIconHash(s string) *UserGuildCreate {
	ugc.mutation.SetIconHash(s)
	return ugc
}

// SetNillableIconHash sets the "icon_hash" field if the given value is not nil.
func (ugc *UserGuildCreate) SetNillableIconHash(s *string) *UserGuildCreate {
	if s != nil {
		ugc.SetIconHash(*s)
	}
	return ugc
}

// SetIconURL sets the "icon_url" field.
func (ugc *UserGuildCreate) SetIconURL(s string) *UserGuildCreate {
	ugc.mutation.SetIconURL(s)
	return ugc
}

// SetPermissions sets the "permissions" field.
func (ugc *UserGuildCreate) SetPermissions(u uint64) *UserGuildCreate {
	ugc.mutation.SetPermissions(u)
	return ugc
}

// SetNillablePermissions sets the "permissions" field if the given value is not nil.
func (ugc *UserGuildCreate) SetNillablePermissions(u *uint64) *UserGuildCreate {
	if u != nil {
		ugc.SetPermissions(*u)
	}
	return ugc
}

// AddAdminIDs adds the "admins" edge to the User entity by IDs.
func (ugc *UserGuildCreate) AddAdminIDs(ids ...int) *UserGuildCreate {
	ugc.mutation.AddAdminIDs(ids...)
	return ugc
}

// AddAdmins adds the "admins" edges to the User entity.
func (ugc *UserGuildCreate) AddAdmins(u ...*User) *UserGuildCreate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return ugc.AddAdminIDs(ids...)
}

// Mutation returns the UserGuildMutation object of the builder.
func (ugc *UserGuildCreate) Mutation() *UserGuildMutation {
	return ugc.mutation
}

// Save creates the UserGuild in the database.
func (ugc *UserGuildCreate) Save(ctx context.Context) (*UserGuild, error) {
	var (
		err  error
		node *UserGuild
	)
	ugc.defaults()
	if len(ugc.hooks) == 0 {
		if err = ugc.check(); err != nil {
			return nil, err
		}
		node, err = ugc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserGuildMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ugc.check(); err != nil {
				return nil, err
			}
			ugc.mutation = mutation
			if node, err = ugc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ugc.hooks) - 1; i >= 0; i-- {
			if ugc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ugc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, ugc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*UserGuild)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from UserGuildMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ugc *UserGuildCreate) SaveX(ctx context.Context) *UserGuild {
	v, err := ugc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ugc *UserGuildCreate) Exec(ctx context.Context) error {
	_, err := ugc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ugc *UserGuildCreate) ExecX(ctx context.Context) {
	if err := ugc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ugc *UserGuildCreate) defaults() {
	if _, ok := ugc.mutation.CreateTime(); !ok {
		v := userguild.DefaultCreateTime()
		ugc.mutation.SetCreateTime(v)
	}
	if _, ok := ugc.mutation.UpdateTime(); !ok {
		v := userguild.DefaultUpdateTime()
		ugc.mutation.SetUpdateTime(v)
	}
	if _, ok := ugc.mutation.Owner(); !ok {
		v := userguild.DefaultOwner
		ugc.mutation.SetOwner(v)
	}
	if _, ok := ugc.mutation.Admin(); !ok {
		v := userguild.DefaultAdmin
		ugc.mutation.SetAdmin(v)
	}
	if _, ok := ugc.mutation.Features(); !ok {
		v := userguild.DefaultFeatures
		ugc.mutation.SetFeatures(v)
	}
	if _, ok := ugc.mutation.Permissions(); !ok {
		v := userguild.DefaultPermissions
		ugc.mutation.SetPermissions(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ugc *UserGuildCreate) check() error {
	if _, ok := ugc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "UserGuild.create_time"`)}
	}
	if _, ok := ugc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`ent: missing required field "UserGuild.update_time"`)}
	}
	if _, ok := ugc.mutation.GuildID(); !ok {
		return &ValidationError{Name: "guild_id", err: errors.New(`ent: missing required field "UserGuild.guild_id"`)}
	}
	if _, ok := ugc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "UserGuild.name"`)}
	}
	if v, ok := ugc.mutation.Name(); ok {
		if err := userguild.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "UserGuild.name": %w`, err)}
		}
	}
	if v, ok := ugc.mutation.IconHash(); ok {
		if err := userguild.IconHashValidator(v); err != nil {
			return &ValidationError{Name: "icon_hash", err: fmt.Errorf(`ent: validator failed for field "UserGuild.icon_hash": %w`, err)}
		}
	}
	if _, ok := ugc.mutation.IconURL(); !ok {
		return &ValidationError{Name: "icon_url", err: errors.New(`ent: missing required field "UserGuild.icon_url"`)}
	}
	if v, ok := ugc.mutation.IconURL(); ok {
		if err := userguild.IconURLValidator(v); err != nil {
			return &ValidationError{Name: "icon_url", err: fmt.Errorf(`ent: validator failed for field "UserGuild.icon_url": %w`, err)}
		}
	}
	if len(ugc.mutation.AdminsIDs()) == 0 {
		return &ValidationError{Name: "admins", err: errors.New(`ent: missing required edge "UserGuild.admins"`)}
	}
	return nil
}

func (ugc *UserGuildCreate) sqlSave(ctx context.Context) (*UserGuild, error) {
	_node, _spec := ugc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ugc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (ugc *UserGuildCreate) createSpec() (*UserGuild, *sqlgraph.CreateSpec) {
	var (
		_node = &UserGuild{config: ugc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: userguild.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: userguild.FieldID,
			},
		}
	)
	if value, ok := ugc.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: userguild.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := ugc.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: userguild.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	if value, ok := ugc.mutation.GuildID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: userguild.FieldGuildID,
		})
		_node.GuildID = value
	}
	if value, ok := ugc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: userguild.FieldName,
		})
		_node.Name = value
	}
	if value, ok := ugc.mutation.Owner(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: userguild.FieldOwner,
		})
		_node.Owner = value
	}
	if value, ok := ugc.mutation.Admin(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: userguild.FieldAdmin,
		})
		_node.Admin = value
	}
	if value, ok := ugc.mutation.Features(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: userguild.FieldFeatures,
		})
		_node.Features = value
	}
	if value, ok := ugc.mutation.IconHash(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: userguild.FieldIconHash,
		})
		_node.IconHash = value
	}
	if value, ok := ugc.mutation.IconURL(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: userguild.FieldIconURL,
		})
		_node.IconURL = value
	}
	if value, ok := ugc.mutation.Permissions(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: userguild.FieldPermissions,
		})
		_node.Permissions = value
	}
	if nodes := ugc.mutation.AdminsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   userguild.AdminsTable,
			Columns: userguild.AdminsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// UserGuildCreateBulk is the builder for creating many UserGuild entities in bulk.
type UserGuildCreateBulk struct {
	config
	builders []*UserGuildCreate
}

// Save creates the UserGuild entities in the database.
func (ugcb *UserGuildCreateBulk) Save(ctx context.Context) ([]*UserGuild, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ugcb.builders))
	nodes := make([]*UserGuild, len(ugcb.builders))
	mutators := make([]Mutator, len(ugcb.builders))
	for i := range ugcb.builders {
		func(i int, root context.Context) {
			builder := ugcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserGuildMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ugcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ugcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ugcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ugcb *UserGuildCreateBulk) SaveX(ctx context.Context) []*UserGuild {
	v, err := ugcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ugcb *UserGuildCreateBulk) Exec(ctx context.Context) error {
	_, err := ugcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ugcb *UserGuildCreateBulk) ExecX(ctx context.Context) {
	if err := ugcb.Exec(ctx); err != nil {
		panic(err)
	}
}