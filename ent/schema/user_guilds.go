// Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
// of this source code is governed by the MIT license that can be found in
// the LICENSE file.

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type UserGuild struct {
	ent.Schema
}

func (UserGuild) Fields() []ent.Field {
	return []ent.Field{
		field.String("guild_id").Immutable().Comment("Guild id."),
		field.String("name").MinLen(2).MaxLen(100).Comment("Guild name (2-100 chars, excl. trailing/leading spaces)."),
		field.Bool("owner").Optional().Default(false).Comment("True if the user is the owner of the guild."),
		field.Bool("admin").Optional().Default(false).Comment("True if the user has admin permissions in the guild (pulled from permissions)."),
		field.Strings("features").Optional().Optional().Default([]string{}).Comment("Enabled guild features."),
		field.String("icon_hash").Optional().MaxLen(2048).Comment("Icon hash."),
		field.String("icon_url").MaxLen(2048),
		field.Uint64("permissions").Optional().Default(0).Comment("Permissions for the user (excludes overrides)."),
	}
}

func (UserGuild) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

func (UserGuild) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("admins", User.Type).Ref("guilds").
			Required().
			Comment("The users that are an admin (or owner) of this server."),
	}
}
