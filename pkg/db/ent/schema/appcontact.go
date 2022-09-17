package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/third-manager/pkg/db/mixin"

	"github.com/google/uuid"

	"github.com/NpoolPlatform/message/npool/third/mgr/v1/usedfor"
)

// AppContact holds the schema definition for the AppContact entity.
type AppContact struct {
	ent.Schema
}

func (AppContact) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

// Fields of the AppContact.
func (AppContact) Fields() []ent.Field {
	var maxLen = 32
	return []ent.Field{
		field.
			UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.
			UUID("app_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.UUID{}
			}),
		field.
			String("used_for").
			Optional().
			MaxLen(maxLen).
			Default(usedfor.UsedFor_DefaultUsedFor.String()),
		field.
			String("sender").
			Optional().
			Default(""),
		field.
			String("account").
			Optional().
			Default(""),
		field.
			String("account_type").
			Optional().
			Default(""),
	}
}

// Edges of the AppContact.
func (AppContact) Edges() []ent.Edge {
	return nil
}
