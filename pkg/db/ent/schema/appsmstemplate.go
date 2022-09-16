package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/NpoolPlatform/third-manager/pkg/db/mixin"

	"github.com/google/uuid"
)

// AppSMSTemplate holds the schema definition for the AppSMSTemplate entity.
type AppSMSTemplate struct {
	ent.Schema
}

func (AppSMSTemplate) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

// Fields of the AppSMSTemplate.
func (AppSMSTemplate) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.UUID("app_id", uuid.UUID{}),
		field.UUID("lang_id", uuid.UUID{}),
		field.String("used_for").Default(""),
		field.String("subject").Default(""),
		field.String("message").Default(""),
	}
}

// Edges of the AppSMSTemplate.
func (AppSMSTemplate) Edges() []ent.Edge {
	return nil
}

func (AppSMSTemplate) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("app_id", "lang_id", "used_for").
			Unique(),
	}
}
