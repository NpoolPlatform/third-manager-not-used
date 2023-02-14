package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/message/npool/third/mgr/v1/usedfor"
	"github.com/NpoolPlatform/third-manager/pkg/db/mixin"

	"github.com/google/uuid"
)

// FrontendTemplate holds the schema definition for the FrontendTemplate entity.
type FrontendTemplate struct {
	ent.Schema
}

func (FrontendTemplate) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

// Fields of the FrontendTemplate.
func (FrontendTemplate) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.UUID("app_id", uuid.UUID{}),
		field.UUID("lang_id", uuid.UUID{}),
		field.String("used_for").
			Optional().
			Default(usedfor.UsedFor_DefaultUsedFor.String()),
		field.String("title").
			Optional().
			Default(""),
		field.Text("content").
			Optional().
			Default(""),
		field.String("sender").
			Optional().
			Default(""),
	}
}
