// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// ContactsColumns holds the columns for the "contacts" table.
	ContactsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "app_id", Type: field.TypeUUID},
		{Name: "used_for", Type: field.TypeString, Nullable: true, Default: "DefaultUsedFor"},
		{Name: "sender", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "account", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "account_type", Type: field.TypeString, Nullable: true, Default: ""},
	}
	// ContactsTable holds the schema information for the "contacts" table.
	ContactsTable = &schema.Table{
		Name:       "contacts",
		Columns:    ContactsColumns,
		PrimaryKey: []*schema.Column{ContactsColumns[0]},
	}
	// EmailTemplatesColumns holds the columns for the "email_templates" table.
	EmailTemplatesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "app_id", Type: field.TypeUUID},
		{Name: "lang_id", Type: field.TypeUUID},
		{Name: "default_to_username", Type: field.TypeString},
		{Name: "used_for", Type: field.TypeString, Nullable: true, Default: "DefaultUsedFor"},
		{Name: "sender", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "reply_tos", Type: field.TypeJSON, Nullable: true},
		{Name: "cc_tos", Type: field.TypeJSON, Nullable: true},
		{Name: "subject", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "body", Type: field.TypeString, Nullable: true, Size: 8192, Default: ""},
	}
	// EmailTemplatesTable holds the schema information for the "email_templates" table.
	EmailTemplatesTable = &schema.Table{
		Name:       "email_templates",
		Columns:    EmailTemplatesColumns,
		PrimaryKey: []*schema.Column{EmailTemplatesColumns[0]},
	}
	// FrontendTemplatesColumns holds the columns for the "frontend_templates" table.
	FrontendTemplatesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "app_id", Type: field.TypeUUID},
		{Name: "lang_id", Type: field.TypeUUID},
		{Name: "used_for", Type: field.TypeString, Nullable: true, Default: "DefaultUsedFor"},
		{Name: "title", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "content", Type: field.TypeString, Nullable: true, Size: 2147483647, Default: ""},
		{Name: "sender", Type: field.TypeString, Nullable: true, Default: ""},
	}
	// FrontendTemplatesTable holds the schema information for the "frontend_templates" table.
	FrontendTemplatesTable = &schema.Table{
		Name:       "frontend_templates",
		Columns:    FrontendTemplatesColumns,
		PrimaryKey: []*schema.Column{FrontendTemplatesColumns[0]},
	}
	// SmsTemplatesColumns holds the columns for the "sms_templates" table.
	SmsTemplatesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "app_id", Type: field.TypeUUID},
		{Name: "lang_id", Type: field.TypeUUID},
		{Name: "used_for", Type: field.TypeString, Nullable: true, Default: "DefaultUsedFor"},
		{Name: "subject", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "message", Type: field.TypeString, Nullable: true, Default: ""},
	}
	// SmsTemplatesTable holds the schema information for the "sms_templates" table.
	SmsTemplatesTable = &schema.Table{
		Name:       "sms_templates",
		Columns:    SmsTemplatesColumns,
		PrimaryKey: []*schema.Column{SmsTemplatesColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ContactsTable,
		EmailTemplatesTable,
		FrontendTemplatesTable,
		SmsTemplatesTable,
	}
)

func init() {
}
