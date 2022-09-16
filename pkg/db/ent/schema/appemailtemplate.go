package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/NpoolPlatform/third-manager/pkg/db/mixin"

	"github.com/google/uuid"
)

// AppEmailTemplate holds the schema definition for the AppEmailTemplate entity.
type AppEmailTemplate struct {
	ent.Schema
}

func (AppEmailTemplate) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

// Fields of the AppEmailTemplate.
func (AppEmailTemplate) Fields() []ent.Field {
	var maxLen = 8192
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.UUID("app_id", uuid.UUID{}),
		field.UUID("lang_id", uuid.UUID{}),
		field.String("default_to_username"),
		field.String("used_for"),
		field.String("sender"),
		field.JSON("reply_tos", []string{}),
		field.JSON("cc_tos", []string{}),
		field.String("subject"),
		field.String("body").MaxLen(maxLen),
	}
}

// Edges of the AppEmailTemplate.
func (AppEmailTemplate) Edges() []ent.Edge {
	return nil
}

func (AppEmailTemplate) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("app_id", "lang_id", "used_for").
			Unique(),
	}
}
