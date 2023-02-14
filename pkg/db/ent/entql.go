// Code generated by ent, DO NOT EDIT.

package ent

import (
	"github.com/NpoolPlatform/third-manager/pkg/db/ent/contact"
	"github.com/NpoolPlatform/third-manager/pkg/db/ent/emailtemplate"
	"github.com/NpoolPlatform/third-manager/pkg/db/ent/frontendtemplate"
	"github.com/NpoolPlatform/third-manager/pkg/db/ent/smstemplate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entql"
	"entgo.io/ent/schema/field"
)

// schemaGraph holds a representation of ent/schema at runtime.
var schemaGraph = func() *sqlgraph.Schema {
	graph := &sqlgraph.Schema{Nodes: make([]*sqlgraph.Node, 4)}
	graph.Nodes[0] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   contact.Table,
			Columns: contact.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: contact.FieldID,
			},
		},
		Type: "Contact",
		Fields: map[string]*sqlgraph.FieldSpec{
			contact.FieldCreatedAt:   {Type: field.TypeUint32, Column: contact.FieldCreatedAt},
			contact.FieldUpdatedAt:   {Type: field.TypeUint32, Column: contact.FieldUpdatedAt},
			contact.FieldDeletedAt:   {Type: field.TypeUint32, Column: contact.FieldDeletedAt},
			contact.FieldAppID:       {Type: field.TypeUUID, Column: contact.FieldAppID},
			contact.FieldUsedFor:     {Type: field.TypeString, Column: contact.FieldUsedFor},
			contact.FieldSender:      {Type: field.TypeString, Column: contact.FieldSender},
			contact.FieldAccount:     {Type: field.TypeString, Column: contact.FieldAccount},
			contact.FieldAccountType: {Type: field.TypeString, Column: contact.FieldAccountType},
		},
	}
	graph.Nodes[1] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   emailtemplate.Table,
			Columns: emailtemplate.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: emailtemplate.FieldID,
			},
		},
		Type: "EmailTemplate",
		Fields: map[string]*sqlgraph.FieldSpec{
			emailtemplate.FieldCreatedAt:         {Type: field.TypeUint32, Column: emailtemplate.FieldCreatedAt},
			emailtemplate.FieldUpdatedAt:         {Type: field.TypeUint32, Column: emailtemplate.FieldUpdatedAt},
			emailtemplate.FieldDeletedAt:         {Type: field.TypeUint32, Column: emailtemplate.FieldDeletedAt},
			emailtemplate.FieldAppID:             {Type: field.TypeUUID, Column: emailtemplate.FieldAppID},
			emailtemplate.FieldLangID:            {Type: field.TypeUUID, Column: emailtemplate.FieldLangID},
			emailtemplate.FieldDefaultToUsername: {Type: field.TypeString, Column: emailtemplate.FieldDefaultToUsername},
			emailtemplate.FieldUsedFor:           {Type: field.TypeString, Column: emailtemplate.FieldUsedFor},
			emailtemplate.FieldSender:            {Type: field.TypeString, Column: emailtemplate.FieldSender},
			emailtemplate.FieldReplyTos:          {Type: field.TypeJSON, Column: emailtemplate.FieldReplyTos},
			emailtemplate.FieldCcTos:             {Type: field.TypeJSON, Column: emailtemplate.FieldCcTos},
			emailtemplate.FieldSubject:           {Type: field.TypeString, Column: emailtemplate.FieldSubject},
			emailtemplate.FieldBody:              {Type: field.TypeString, Column: emailtemplate.FieldBody},
		},
	}
	graph.Nodes[2] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   frontendtemplate.Table,
			Columns: frontendtemplate.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: frontendtemplate.FieldID,
			},
		},
		Type: "FrontendTemplate",
		Fields: map[string]*sqlgraph.FieldSpec{
			frontendtemplate.FieldCreatedAt: {Type: field.TypeUint32, Column: frontendtemplate.FieldCreatedAt},
			frontendtemplate.FieldUpdatedAt: {Type: field.TypeUint32, Column: frontendtemplate.FieldUpdatedAt},
			frontendtemplate.FieldDeletedAt: {Type: field.TypeUint32, Column: frontendtemplate.FieldDeletedAt},
			frontendtemplate.FieldAppID:     {Type: field.TypeUUID, Column: frontendtemplate.FieldAppID},
			frontendtemplate.FieldLangID:    {Type: field.TypeUUID, Column: frontendtemplate.FieldLangID},
			frontendtemplate.FieldUsedFor:   {Type: field.TypeString, Column: frontendtemplate.FieldUsedFor},
			frontendtemplate.FieldTitle:     {Type: field.TypeString, Column: frontendtemplate.FieldTitle},
			frontendtemplate.FieldContent:   {Type: field.TypeString, Column: frontendtemplate.FieldContent},
			frontendtemplate.FieldSender:    {Type: field.TypeString, Column: frontendtemplate.FieldSender},
		},
	}
	graph.Nodes[3] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   smstemplate.Table,
			Columns: smstemplate.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: smstemplate.FieldID,
			},
		},
		Type: "SMSTemplate",
		Fields: map[string]*sqlgraph.FieldSpec{
			smstemplate.FieldCreatedAt: {Type: field.TypeUint32, Column: smstemplate.FieldCreatedAt},
			smstemplate.FieldUpdatedAt: {Type: field.TypeUint32, Column: smstemplate.FieldUpdatedAt},
			smstemplate.FieldDeletedAt: {Type: field.TypeUint32, Column: smstemplate.FieldDeletedAt},
			smstemplate.FieldAppID:     {Type: field.TypeUUID, Column: smstemplate.FieldAppID},
			smstemplate.FieldLangID:    {Type: field.TypeUUID, Column: smstemplate.FieldLangID},
			smstemplate.FieldUsedFor:   {Type: field.TypeString, Column: smstemplate.FieldUsedFor},
			smstemplate.FieldSubject:   {Type: field.TypeString, Column: smstemplate.FieldSubject},
			smstemplate.FieldMessage:   {Type: field.TypeString, Column: smstemplate.FieldMessage},
		},
	}
	return graph
}()

// predicateAdder wraps the addPredicate method.
// All update, update-one and query builders implement this interface.
type predicateAdder interface {
	addPredicate(func(s *sql.Selector))
}

// addPredicate implements the predicateAdder interface.
func (cq *ContactQuery) addPredicate(pred func(s *sql.Selector)) {
	cq.predicates = append(cq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the ContactQuery builder.
func (cq *ContactQuery) Filter() *ContactFilter {
	return &ContactFilter{config: cq.config, predicateAdder: cq}
}

// addPredicate implements the predicateAdder interface.
func (m *ContactMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the ContactMutation builder.
func (m *ContactMutation) Filter() *ContactFilter {
	return &ContactFilter{config: m.config, predicateAdder: m}
}

// ContactFilter provides a generic filtering capability at runtime for ContactQuery.
type ContactFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *ContactFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[0].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql [16]byte predicate on the id field.
func (f *ContactFilter) WhereID(p entql.ValueP) {
	f.Where(p.Field(contact.FieldID))
}

// WhereCreatedAt applies the entql uint32 predicate on the created_at field.
func (f *ContactFilter) WhereCreatedAt(p entql.Uint32P) {
	f.Where(p.Field(contact.FieldCreatedAt))
}

// WhereUpdatedAt applies the entql uint32 predicate on the updated_at field.
func (f *ContactFilter) WhereUpdatedAt(p entql.Uint32P) {
	f.Where(p.Field(contact.FieldUpdatedAt))
}

// WhereDeletedAt applies the entql uint32 predicate on the deleted_at field.
func (f *ContactFilter) WhereDeletedAt(p entql.Uint32P) {
	f.Where(p.Field(contact.FieldDeletedAt))
}

// WhereAppID applies the entql [16]byte predicate on the app_id field.
func (f *ContactFilter) WhereAppID(p entql.ValueP) {
	f.Where(p.Field(contact.FieldAppID))
}

// WhereUsedFor applies the entql string predicate on the used_for field.
func (f *ContactFilter) WhereUsedFor(p entql.StringP) {
	f.Where(p.Field(contact.FieldUsedFor))
}

// WhereSender applies the entql string predicate on the sender field.
func (f *ContactFilter) WhereSender(p entql.StringP) {
	f.Where(p.Field(contact.FieldSender))
}

// WhereAccount applies the entql string predicate on the account field.
func (f *ContactFilter) WhereAccount(p entql.StringP) {
	f.Where(p.Field(contact.FieldAccount))
}

// WhereAccountType applies the entql string predicate on the account_type field.
func (f *ContactFilter) WhereAccountType(p entql.StringP) {
	f.Where(p.Field(contact.FieldAccountType))
}

// addPredicate implements the predicateAdder interface.
func (etq *EmailTemplateQuery) addPredicate(pred func(s *sql.Selector)) {
	etq.predicates = append(etq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the EmailTemplateQuery builder.
func (etq *EmailTemplateQuery) Filter() *EmailTemplateFilter {
	return &EmailTemplateFilter{config: etq.config, predicateAdder: etq}
}

// addPredicate implements the predicateAdder interface.
func (m *EmailTemplateMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the EmailTemplateMutation builder.
func (m *EmailTemplateMutation) Filter() *EmailTemplateFilter {
	return &EmailTemplateFilter{config: m.config, predicateAdder: m}
}

// EmailTemplateFilter provides a generic filtering capability at runtime for EmailTemplateQuery.
type EmailTemplateFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *EmailTemplateFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[1].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql [16]byte predicate on the id field.
func (f *EmailTemplateFilter) WhereID(p entql.ValueP) {
	f.Where(p.Field(emailtemplate.FieldID))
}

// WhereCreatedAt applies the entql uint32 predicate on the created_at field.
func (f *EmailTemplateFilter) WhereCreatedAt(p entql.Uint32P) {
	f.Where(p.Field(emailtemplate.FieldCreatedAt))
}

// WhereUpdatedAt applies the entql uint32 predicate on the updated_at field.
func (f *EmailTemplateFilter) WhereUpdatedAt(p entql.Uint32P) {
	f.Where(p.Field(emailtemplate.FieldUpdatedAt))
}

// WhereDeletedAt applies the entql uint32 predicate on the deleted_at field.
func (f *EmailTemplateFilter) WhereDeletedAt(p entql.Uint32P) {
	f.Where(p.Field(emailtemplate.FieldDeletedAt))
}

// WhereAppID applies the entql [16]byte predicate on the app_id field.
func (f *EmailTemplateFilter) WhereAppID(p entql.ValueP) {
	f.Where(p.Field(emailtemplate.FieldAppID))
}

// WhereLangID applies the entql [16]byte predicate on the lang_id field.
func (f *EmailTemplateFilter) WhereLangID(p entql.ValueP) {
	f.Where(p.Field(emailtemplate.FieldLangID))
}

// WhereDefaultToUsername applies the entql string predicate on the default_to_username field.
func (f *EmailTemplateFilter) WhereDefaultToUsername(p entql.StringP) {
	f.Where(p.Field(emailtemplate.FieldDefaultToUsername))
}

// WhereUsedFor applies the entql string predicate on the used_for field.
func (f *EmailTemplateFilter) WhereUsedFor(p entql.StringP) {
	f.Where(p.Field(emailtemplate.FieldUsedFor))
}

// WhereSender applies the entql string predicate on the sender field.
func (f *EmailTemplateFilter) WhereSender(p entql.StringP) {
	f.Where(p.Field(emailtemplate.FieldSender))
}

// WhereReplyTos applies the entql json.RawMessage predicate on the reply_tos field.
func (f *EmailTemplateFilter) WhereReplyTos(p entql.BytesP) {
	f.Where(p.Field(emailtemplate.FieldReplyTos))
}

// WhereCcTos applies the entql json.RawMessage predicate on the cc_tos field.
func (f *EmailTemplateFilter) WhereCcTos(p entql.BytesP) {
	f.Where(p.Field(emailtemplate.FieldCcTos))
}

// WhereSubject applies the entql string predicate on the subject field.
func (f *EmailTemplateFilter) WhereSubject(p entql.StringP) {
	f.Where(p.Field(emailtemplate.FieldSubject))
}

// WhereBody applies the entql string predicate on the body field.
func (f *EmailTemplateFilter) WhereBody(p entql.StringP) {
	f.Where(p.Field(emailtemplate.FieldBody))
}

// addPredicate implements the predicateAdder interface.
func (ftq *FrontendTemplateQuery) addPredicate(pred func(s *sql.Selector)) {
	ftq.predicates = append(ftq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the FrontendTemplateQuery builder.
func (ftq *FrontendTemplateQuery) Filter() *FrontendTemplateFilter {
	return &FrontendTemplateFilter{config: ftq.config, predicateAdder: ftq}
}

// addPredicate implements the predicateAdder interface.
func (m *FrontendTemplateMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the FrontendTemplateMutation builder.
func (m *FrontendTemplateMutation) Filter() *FrontendTemplateFilter {
	return &FrontendTemplateFilter{config: m.config, predicateAdder: m}
}

// FrontendTemplateFilter provides a generic filtering capability at runtime for FrontendTemplateQuery.
type FrontendTemplateFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *FrontendTemplateFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[2].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql [16]byte predicate on the id field.
func (f *FrontendTemplateFilter) WhereID(p entql.ValueP) {
	f.Where(p.Field(frontendtemplate.FieldID))
}

// WhereCreatedAt applies the entql uint32 predicate on the created_at field.
func (f *FrontendTemplateFilter) WhereCreatedAt(p entql.Uint32P) {
	f.Where(p.Field(frontendtemplate.FieldCreatedAt))
}

// WhereUpdatedAt applies the entql uint32 predicate on the updated_at field.
func (f *FrontendTemplateFilter) WhereUpdatedAt(p entql.Uint32P) {
	f.Where(p.Field(frontendtemplate.FieldUpdatedAt))
}

// WhereDeletedAt applies the entql uint32 predicate on the deleted_at field.
func (f *FrontendTemplateFilter) WhereDeletedAt(p entql.Uint32P) {
	f.Where(p.Field(frontendtemplate.FieldDeletedAt))
}

// WhereAppID applies the entql [16]byte predicate on the app_id field.
func (f *FrontendTemplateFilter) WhereAppID(p entql.ValueP) {
	f.Where(p.Field(frontendtemplate.FieldAppID))
}

// WhereLangID applies the entql [16]byte predicate on the lang_id field.
func (f *FrontendTemplateFilter) WhereLangID(p entql.ValueP) {
	f.Where(p.Field(frontendtemplate.FieldLangID))
}

// WhereUsedFor applies the entql string predicate on the used_for field.
func (f *FrontendTemplateFilter) WhereUsedFor(p entql.StringP) {
	f.Where(p.Field(frontendtemplate.FieldUsedFor))
}

// WhereTitle applies the entql string predicate on the title field.
func (f *FrontendTemplateFilter) WhereTitle(p entql.StringP) {
	f.Where(p.Field(frontendtemplate.FieldTitle))
}

// WhereContent applies the entql string predicate on the content field.
func (f *FrontendTemplateFilter) WhereContent(p entql.StringP) {
	f.Where(p.Field(frontendtemplate.FieldContent))
}

// WhereSender applies the entql string predicate on the sender field.
func (f *FrontendTemplateFilter) WhereSender(p entql.StringP) {
	f.Where(p.Field(frontendtemplate.FieldSender))
}

// addPredicate implements the predicateAdder interface.
func (stq *SMSTemplateQuery) addPredicate(pred func(s *sql.Selector)) {
	stq.predicates = append(stq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the SMSTemplateQuery builder.
func (stq *SMSTemplateQuery) Filter() *SMSTemplateFilter {
	return &SMSTemplateFilter{config: stq.config, predicateAdder: stq}
}

// addPredicate implements the predicateAdder interface.
func (m *SMSTemplateMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the SMSTemplateMutation builder.
func (m *SMSTemplateMutation) Filter() *SMSTemplateFilter {
	return &SMSTemplateFilter{config: m.config, predicateAdder: m}
}

// SMSTemplateFilter provides a generic filtering capability at runtime for SMSTemplateQuery.
type SMSTemplateFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *SMSTemplateFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[3].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql [16]byte predicate on the id field.
func (f *SMSTemplateFilter) WhereID(p entql.ValueP) {
	f.Where(p.Field(smstemplate.FieldID))
}

// WhereCreatedAt applies the entql uint32 predicate on the created_at field.
func (f *SMSTemplateFilter) WhereCreatedAt(p entql.Uint32P) {
	f.Where(p.Field(smstemplate.FieldCreatedAt))
}

// WhereUpdatedAt applies the entql uint32 predicate on the updated_at field.
func (f *SMSTemplateFilter) WhereUpdatedAt(p entql.Uint32P) {
	f.Where(p.Field(smstemplate.FieldUpdatedAt))
}

// WhereDeletedAt applies the entql uint32 predicate on the deleted_at field.
func (f *SMSTemplateFilter) WhereDeletedAt(p entql.Uint32P) {
	f.Where(p.Field(smstemplate.FieldDeletedAt))
}

// WhereAppID applies the entql [16]byte predicate on the app_id field.
func (f *SMSTemplateFilter) WhereAppID(p entql.ValueP) {
	f.Where(p.Field(smstemplate.FieldAppID))
}

// WhereLangID applies the entql [16]byte predicate on the lang_id field.
func (f *SMSTemplateFilter) WhereLangID(p entql.ValueP) {
	f.Where(p.Field(smstemplate.FieldLangID))
}

// WhereUsedFor applies the entql string predicate on the used_for field.
func (f *SMSTemplateFilter) WhereUsedFor(p entql.StringP) {
	f.Where(p.Field(smstemplate.FieldUsedFor))
}

// WhereSubject applies the entql string predicate on the subject field.
func (f *SMSTemplateFilter) WhereSubject(p entql.StringP) {
	f.Where(p.Field(smstemplate.FieldSubject))
}

// WhereMessage applies the entql string predicate on the message field.
func (f *SMSTemplateFilter) WhereMessage(p entql.StringP) {
	f.Where(p.Field(smstemplate.FieldMessage))
}
