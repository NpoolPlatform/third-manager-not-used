package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/message/npool/notif/mgr/v1/notif"
	"github.com/NpoolPlatform/third-manager/pkg/db/mixin"

	"github.com/google/uuid"
)

// NotifTemplate holds the schema definition for the NotifTemplate entity.
type NotifTemplate struct {
	ent.Schema
}

func (NotifTemplate) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

// Fields of the NotifTemplate.
func (NotifTemplate) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.UUID("app_id", uuid.UUID{}),
		field.UUID("lang_id", uuid.UUID{}),
		field.String("used_for").
			Optional().
			Default(notif.EventType_DefaultEventType.String()),
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
