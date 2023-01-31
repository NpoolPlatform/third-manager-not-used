//nolint:dupl
package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/message/npool/third/mgr/v1/usedfor"
	"github.com/NpoolPlatform/third-manager/pkg/db/mixin"

	"github.com/google/uuid"
)

// SMSTemplate holds the schema definition for the SMSTemplate entity.
type SMSTemplate struct {
	ent.Schema
}

func (SMSTemplate) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

// Fields of the AppSMSTemplate.
func (SMSTemplate) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.UUID("app_id", uuid.UUID{}),
		field.UUID("lang_id", uuid.UUID{}),
		field.String("used_for").
			Optional().
			Default(usedfor.UsedFor_DefaultUsedFor.String()),
		field.String("subject").
			Optional().
			Default(""),
		field.String("message").
			Optional().
			Default(""),
	}
}
