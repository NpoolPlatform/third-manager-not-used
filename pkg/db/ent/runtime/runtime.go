// Code generated by ent, DO NOT EDIT.

package runtime

import (
	"context"

	"github.com/NpoolPlatform/third-manager/pkg/db/ent/contact"
	"github.com/NpoolPlatform/third-manager/pkg/db/ent/emailtemplate"
	"github.com/NpoolPlatform/third-manager/pkg/db/ent/notiftemplate"
	"github.com/NpoolPlatform/third-manager/pkg/db/ent/schema"
	"github.com/NpoolPlatform/third-manager/pkg/db/ent/smstemplate"
	"github.com/google/uuid"

	"entgo.io/ent"
	"entgo.io/ent/privacy"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	contactMixin := schema.Contact{}.Mixin()
	contact.Policy = privacy.NewPolicies(contactMixin[0], schema.Contact{})
	contact.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := contact.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	contactMixinFields0 := contactMixin[0].Fields()
	_ = contactMixinFields0
	contactFields := schema.Contact{}.Fields()
	_ = contactFields
	// contactDescCreatedAt is the schema descriptor for created_at field.
	contactDescCreatedAt := contactMixinFields0[0].Descriptor()
	// contact.DefaultCreatedAt holds the default value on creation for the created_at field.
	contact.DefaultCreatedAt = contactDescCreatedAt.Default.(func() uint32)
	// contactDescUpdatedAt is the schema descriptor for updated_at field.
	contactDescUpdatedAt := contactMixinFields0[1].Descriptor()
	// contact.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	contact.DefaultUpdatedAt = contactDescUpdatedAt.Default.(func() uint32)
	// contact.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	contact.UpdateDefaultUpdatedAt = contactDescUpdatedAt.UpdateDefault.(func() uint32)
	// contactDescDeletedAt is the schema descriptor for deleted_at field.
	contactDescDeletedAt := contactMixinFields0[2].Descriptor()
	// contact.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	contact.DefaultDeletedAt = contactDescDeletedAt.Default.(func() uint32)
	// contactDescUsedFor is the schema descriptor for used_for field.
	contactDescUsedFor := contactFields[2].Descriptor()
	// contact.DefaultUsedFor holds the default value on creation for the used_for field.
	contact.DefaultUsedFor = contactDescUsedFor.Default.(string)
	// contactDescSender is the schema descriptor for sender field.
	contactDescSender := contactFields[3].Descriptor()
	// contact.DefaultSender holds the default value on creation for the sender field.
	contact.DefaultSender = contactDescSender.Default.(string)
	// contactDescAccount is the schema descriptor for account field.
	contactDescAccount := contactFields[4].Descriptor()
	// contact.DefaultAccount holds the default value on creation for the account field.
	contact.DefaultAccount = contactDescAccount.Default.(string)
	// contactDescAccountType is the schema descriptor for account_type field.
	contactDescAccountType := contactFields[5].Descriptor()
	// contact.DefaultAccountType holds the default value on creation for the account_type field.
	contact.DefaultAccountType = contactDescAccountType.Default.(string)
	// contactDescID is the schema descriptor for id field.
	contactDescID := contactFields[0].Descriptor()
	// contact.DefaultID holds the default value on creation for the id field.
	contact.DefaultID = contactDescID.Default.(func() uuid.UUID)
	emailtemplateMixin := schema.EmailTemplate{}.Mixin()
	emailtemplate.Policy = privacy.NewPolicies(emailtemplateMixin[0], schema.EmailTemplate{})
	emailtemplate.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := emailtemplate.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	emailtemplateMixinFields0 := emailtemplateMixin[0].Fields()
	_ = emailtemplateMixinFields0
	emailtemplateFields := schema.EmailTemplate{}.Fields()
	_ = emailtemplateFields
	// emailtemplateDescCreatedAt is the schema descriptor for created_at field.
	emailtemplateDescCreatedAt := emailtemplateMixinFields0[0].Descriptor()
	// emailtemplate.DefaultCreatedAt holds the default value on creation for the created_at field.
	emailtemplate.DefaultCreatedAt = emailtemplateDescCreatedAt.Default.(func() uint32)
	// emailtemplateDescUpdatedAt is the schema descriptor for updated_at field.
	emailtemplateDescUpdatedAt := emailtemplateMixinFields0[1].Descriptor()
	// emailtemplate.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	emailtemplate.DefaultUpdatedAt = emailtemplateDescUpdatedAt.Default.(func() uint32)
	// emailtemplate.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	emailtemplate.UpdateDefaultUpdatedAt = emailtemplateDescUpdatedAt.UpdateDefault.(func() uint32)
	// emailtemplateDescDeletedAt is the schema descriptor for deleted_at field.
	emailtemplateDescDeletedAt := emailtemplateMixinFields0[2].Descriptor()
	// emailtemplate.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	emailtemplate.DefaultDeletedAt = emailtemplateDescDeletedAt.Default.(func() uint32)
	// emailtemplateDescUsedFor is the schema descriptor for used_for field.
	emailtemplateDescUsedFor := emailtemplateFields[4].Descriptor()
	// emailtemplate.DefaultUsedFor holds the default value on creation for the used_for field.
	emailtemplate.DefaultUsedFor = emailtemplateDescUsedFor.Default.(string)
	// emailtemplateDescSender is the schema descriptor for sender field.
	emailtemplateDescSender := emailtemplateFields[5].Descriptor()
	// emailtemplate.DefaultSender holds the default value on creation for the sender field.
	emailtemplate.DefaultSender = emailtemplateDescSender.Default.(string)
	// emailtemplateDescReplyTos is the schema descriptor for reply_tos field.
	emailtemplateDescReplyTos := emailtemplateFields[6].Descriptor()
	// emailtemplate.DefaultReplyTos holds the default value on creation for the reply_tos field.
	emailtemplate.DefaultReplyTos = emailtemplateDescReplyTos.Default.([]string)
	// emailtemplateDescCcTos is the schema descriptor for cc_tos field.
	emailtemplateDescCcTos := emailtemplateFields[7].Descriptor()
	// emailtemplate.DefaultCcTos holds the default value on creation for the cc_tos field.
	emailtemplate.DefaultCcTos = emailtemplateDescCcTos.Default.([]string)
	// emailtemplateDescSubject is the schema descriptor for subject field.
	emailtemplateDescSubject := emailtemplateFields[8].Descriptor()
	// emailtemplate.DefaultSubject holds the default value on creation for the subject field.
	emailtemplate.DefaultSubject = emailtemplateDescSubject.Default.(string)
	// emailtemplateDescBody is the schema descriptor for body field.
	emailtemplateDescBody := emailtemplateFields[9].Descriptor()
	// emailtemplate.DefaultBody holds the default value on creation for the body field.
	emailtemplate.DefaultBody = emailtemplateDescBody.Default.(string)
	// emailtemplate.BodyValidator is a validator for the "body" field. It is called by the builders before save.
	emailtemplate.BodyValidator = emailtemplateDescBody.Validators[0].(func(string) error)
	// emailtemplateDescID is the schema descriptor for id field.
	emailtemplateDescID := emailtemplateFields[0].Descriptor()
	// emailtemplate.DefaultID holds the default value on creation for the id field.
	emailtemplate.DefaultID = emailtemplateDescID.Default.(func() uuid.UUID)
	notiftemplateMixin := schema.NotifTemplate{}.Mixin()
	notiftemplate.Policy = privacy.NewPolicies(notiftemplateMixin[0], schema.NotifTemplate{})
	notiftemplate.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := notiftemplate.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	notiftemplateMixinFields0 := notiftemplateMixin[0].Fields()
	_ = notiftemplateMixinFields0
	notiftemplateFields := schema.NotifTemplate{}.Fields()
	_ = notiftemplateFields
	// notiftemplateDescCreatedAt is the schema descriptor for created_at field.
	notiftemplateDescCreatedAt := notiftemplateMixinFields0[0].Descriptor()
	// notiftemplate.DefaultCreatedAt holds the default value on creation for the created_at field.
	notiftemplate.DefaultCreatedAt = notiftemplateDescCreatedAt.Default.(func() uint32)
	// notiftemplateDescUpdatedAt is the schema descriptor for updated_at field.
	notiftemplateDescUpdatedAt := notiftemplateMixinFields0[1].Descriptor()
	// notiftemplate.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	notiftemplate.DefaultUpdatedAt = notiftemplateDescUpdatedAt.Default.(func() uint32)
	// notiftemplate.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	notiftemplate.UpdateDefaultUpdatedAt = notiftemplateDescUpdatedAt.UpdateDefault.(func() uint32)
	// notiftemplateDescDeletedAt is the schema descriptor for deleted_at field.
	notiftemplateDescDeletedAt := notiftemplateMixinFields0[2].Descriptor()
	// notiftemplate.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	notiftemplate.DefaultDeletedAt = notiftemplateDescDeletedAt.Default.(func() uint32)
	// notiftemplateDescUsedFor is the schema descriptor for used_for field.
	notiftemplateDescUsedFor := notiftemplateFields[3].Descriptor()
	// notiftemplate.DefaultUsedFor holds the default value on creation for the used_for field.
	notiftemplate.DefaultUsedFor = notiftemplateDescUsedFor.Default.(string)
	// notiftemplateDescTitle is the schema descriptor for title field.
	notiftemplateDescTitle := notiftemplateFields[4].Descriptor()
	// notiftemplate.DefaultTitle holds the default value on creation for the title field.
	notiftemplate.DefaultTitle = notiftemplateDescTitle.Default.(string)
	// notiftemplateDescContent is the schema descriptor for content field.
	notiftemplateDescContent := notiftemplateFields[5].Descriptor()
	// notiftemplate.DefaultContent holds the default value on creation for the content field.
	notiftemplate.DefaultContent = notiftemplateDescContent.Default.(string)
	// notiftemplateDescID is the schema descriptor for id field.
	notiftemplateDescID := notiftemplateFields[0].Descriptor()
	// notiftemplate.DefaultID holds the default value on creation for the id field.
	notiftemplate.DefaultID = notiftemplateDescID.Default.(func() uuid.UUID)
	smstemplateMixin := schema.SMSTemplate{}.Mixin()
	smstemplate.Policy = privacy.NewPolicies(smstemplateMixin[0], schema.SMSTemplate{})
	smstemplate.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := smstemplate.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	smstemplateMixinFields0 := smstemplateMixin[0].Fields()
	_ = smstemplateMixinFields0
	smstemplateFields := schema.SMSTemplate{}.Fields()
	_ = smstemplateFields
	// smstemplateDescCreatedAt is the schema descriptor for created_at field.
	smstemplateDescCreatedAt := smstemplateMixinFields0[0].Descriptor()
	// smstemplate.DefaultCreatedAt holds the default value on creation for the created_at field.
	smstemplate.DefaultCreatedAt = smstemplateDescCreatedAt.Default.(func() uint32)
	// smstemplateDescUpdatedAt is the schema descriptor for updated_at field.
	smstemplateDescUpdatedAt := smstemplateMixinFields0[1].Descriptor()
	// smstemplate.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	smstemplate.DefaultUpdatedAt = smstemplateDescUpdatedAt.Default.(func() uint32)
	// smstemplate.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	smstemplate.UpdateDefaultUpdatedAt = smstemplateDescUpdatedAt.UpdateDefault.(func() uint32)
	// smstemplateDescDeletedAt is the schema descriptor for deleted_at field.
	smstemplateDescDeletedAt := smstemplateMixinFields0[2].Descriptor()
	// smstemplate.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	smstemplate.DefaultDeletedAt = smstemplateDescDeletedAt.Default.(func() uint32)
	// smstemplateDescUsedFor is the schema descriptor for used_for field.
	smstemplateDescUsedFor := smstemplateFields[3].Descriptor()
	// smstemplate.DefaultUsedFor holds the default value on creation for the used_for field.
	smstemplate.DefaultUsedFor = smstemplateDescUsedFor.Default.(string)
	// smstemplateDescSubject is the schema descriptor for subject field.
	smstemplateDescSubject := smstemplateFields[4].Descriptor()
	// smstemplate.DefaultSubject holds the default value on creation for the subject field.
	smstemplate.DefaultSubject = smstemplateDescSubject.Default.(string)
	// smstemplateDescMessage is the schema descriptor for message field.
	smstemplateDescMessage := smstemplateFields[5].Descriptor()
	// smstemplate.DefaultMessage holds the default value on creation for the message field.
	smstemplate.DefaultMessage = smstemplateDescMessage.Default.(string)
	// smstemplateDescID is the schema descriptor for id field.
	smstemplateDescID := smstemplateFields[0].Descriptor()
	// smstemplate.DefaultID holds the default value on creation for the id field.
	smstemplate.DefaultID = smstemplateDescID.Default.(func() uuid.UUID)
}

const (
	Version = "v0.11.2" // Version of ent codegen.
)
