// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/third-manager/pkg/db/ent/notiftemplate"
	"github.com/google/uuid"
)

// NotifTemplateCreate is the builder for creating a NotifTemplate entity.
type NotifTemplateCreate struct {
	config
	mutation *NotifTemplateMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (ntc *NotifTemplateCreate) SetCreatedAt(u uint32) *NotifTemplateCreate {
	ntc.mutation.SetCreatedAt(u)
	return ntc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ntc *NotifTemplateCreate) SetNillableCreatedAt(u *uint32) *NotifTemplateCreate {
	if u != nil {
		ntc.SetCreatedAt(*u)
	}
	return ntc
}

// SetUpdatedAt sets the "updated_at" field.
func (ntc *NotifTemplateCreate) SetUpdatedAt(u uint32) *NotifTemplateCreate {
	ntc.mutation.SetUpdatedAt(u)
	return ntc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ntc *NotifTemplateCreate) SetNillableUpdatedAt(u *uint32) *NotifTemplateCreate {
	if u != nil {
		ntc.SetUpdatedAt(*u)
	}
	return ntc
}

// SetDeletedAt sets the "deleted_at" field.
func (ntc *NotifTemplateCreate) SetDeletedAt(u uint32) *NotifTemplateCreate {
	ntc.mutation.SetDeletedAt(u)
	return ntc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ntc *NotifTemplateCreate) SetNillableDeletedAt(u *uint32) *NotifTemplateCreate {
	if u != nil {
		ntc.SetDeletedAt(*u)
	}
	return ntc
}

// SetAppID sets the "app_id" field.
func (ntc *NotifTemplateCreate) SetAppID(u uuid.UUID) *NotifTemplateCreate {
	ntc.mutation.SetAppID(u)
	return ntc
}

// SetLangID sets the "lang_id" field.
func (ntc *NotifTemplateCreate) SetLangID(u uuid.UUID) *NotifTemplateCreate {
	ntc.mutation.SetLangID(u)
	return ntc
}

// SetUsedFor sets the "used_for" field.
func (ntc *NotifTemplateCreate) SetUsedFor(s string) *NotifTemplateCreate {
	ntc.mutation.SetUsedFor(s)
	return ntc
}

// SetNillableUsedFor sets the "used_for" field if the given value is not nil.
func (ntc *NotifTemplateCreate) SetNillableUsedFor(s *string) *NotifTemplateCreate {
	if s != nil {
		ntc.SetUsedFor(*s)
	}
	return ntc
}

// SetTitle sets the "title" field.
func (ntc *NotifTemplateCreate) SetTitle(s string) *NotifTemplateCreate {
	ntc.mutation.SetTitle(s)
	return ntc
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (ntc *NotifTemplateCreate) SetNillableTitle(s *string) *NotifTemplateCreate {
	if s != nil {
		ntc.SetTitle(*s)
	}
	return ntc
}

// SetContent sets the "content" field.
func (ntc *NotifTemplateCreate) SetContent(s string) *NotifTemplateCreate {
	ntc.mutation.SetContent(s)
	return ntc
}

// SetNillableContent sets the "content" field if the given value is not nil.
func (ntc *NotifTemplateCreate) SetNillableContent(s *string) *NotifTemplateCreate {
	if s != nil {
		ntc.SetContent(*s)
	}
	return ntc
}

// SetID sets the "id" field.
func (ntc *NotifTemplateCreate) SetID(u uuid.UUID) *NotifTemplateCreate {
	ntc.mutation.SetID(u)
	return ntc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (ntc *NotifTemplateCreate) SetNillableID(u *uuid.UUID) *NotifTemplateCreate {
	if u != nil {
		ntc.SetID(*u)
	}
	return ntc
}

// Mutation returns the NotifTemplateMutation object of the builder.
func (ntc *NotifTemplateCreate) Mutation() *NotifTemplateMutation {
	return ntc.mutation
}

// Save creates the NotifTemplate in the database.
func (ntc *NotifTemplateCreate) Save(ctx context.Context) (*NotifTemplate, error) {
	var (
		err  error
		node *NotifTemplate
	)
	if err := ntc.defaults(); err != nil {
		return nil, err
	}
	if len(ntc.hooks) == 0 {
		if err = ntc.check(); err != nil {
			return nil, err
		}
		node, err = ntc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*NotifTemplateMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ntc.check(); err != nil {
				return nil, err
			}
			ntc.mutation = mutation
			if node, err = ntc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ntc.hooks) - 1; i >= 0; i-- {
			if ntc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ntc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, ntc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*NotifTemplate)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from NotifTemplateMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ntc *NotifTemplateCreate) SaveX(ctx context.Context) *NotifTemplate {
	v, err := ntc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ntc *NotifTemplateCreate) Exec(ctx context.Context) error {
	_, err := ntc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ntc *NotifTemplateCreate) ExecX(ctx context.Context) {
	if err := ntc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ntc *NotifTemplateCreate) defaults() error {
	if _, ok := ntc.mutation.CreatedAt(); !ok {
		if notiftemplate.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized notiftemplate.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := notiftemplate.DefaultCreatedAt()
		ntc.mutation.SetCreatedAt(v)
	}
	if _, ok := ntc.mutation.UpdatedAt(); !ok {
		if notiftemplate.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized notiftemplate.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := notiftemplate.DefaultUpdatedAt()
		ntc.mutation.SetUpdatedAt(v)
	}
	if _, ok := ntc.mutation.DeletedAt(); !ok {
		if notiftemplate.DefaultDeletedAt == nil {
			return fmt.Errorf("ent: uninitialized notiftemplate.DefaultDeletedAt (forgotten import ent/runtime?)")
		}
		v := notiftemplate.DefaultDeletedAt()
		ntc.mutation.SetDeletedAt(v)
	}
	if _, ok := ntc.mutation.UsedFor(); !ok {
		v := notiftemplate.DefaultUsedFor
		ntc.mutation.SetUsedFor(v)
	}
	if _, ok := ntc.mutation.Title(); !ok {
		v := notiftemplate.DefaultTitle
		ntc.mutation.SetTitle(v)
	}
	if _, ok := ntc.mutation.Content(); !ok {
		v := notiftemplate.DefaultContent
		ntc.mutation.SetContent(v)
	}
	if _, ok := ntc.mutation.ID(); !ok {
		if notiftemplate.DefaultID == nil {
			return fmt.Errorf("ent: uninitialized notiftemplate.DefaultID (forgotten import ent/runtime?)")
		}
		v := notiftemplate.DefaultID()
		ntc.mutation.SetID(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (ntc *NotifTemplateCreate) check() error {
	if _, ok := ntc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "NotifTemplate.created_at"`)}
	}
	if _, ok := ntc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "NotifTemplate.updated_at"`)}
	}
	if _, ok := ntc.mutation.DeletedAt(); !ok {
		return &ValidationError{Name: "deleted_at", err: errors.New(`ent: missing required field "NotifTemplate.deleted_at"`)}
	}
	if _, ok := ntc.mutation.AppID(); !ok {
		return &ValidationError{Name: "app_id", err: errors.New(`ent: missing required field "NotifTemplate.app_id"`)}
	}
	if _, ok := ntc.mutation.LangID(); !ok {
		return &ValidationError{Name: "lang_id", err: errors.New(`ent: missing required field "NotifTemplate.lang_id"`)}
	}
	return nil
}

func (ntc *NotifTemplateCreate) sqlSave(ctx context.Context) (*NotifTemplate, error) {
	_node, _spec := ntc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ntc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	return _node, nil
}

func (ntc *NotifTemplateCreate) createSpec() (*NotifTemplate, *sqlgraph.CreateSpec) {
	var (
		_node = &NotifTemplate{config: ntc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: notiftemplate.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: notiftemplate.FieldID,
			},
		}
	)
	_spec.OnConflict = ntc.conflict
	if id, ok := ntc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := ntc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: notiftemplate.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := ntc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: notiftemplate.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := ntc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: notiftemplate.FieldDeletedAt,
		})
		_node.DeletedAt = value
	}
	if value, ok := ntc.mutation.AppID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: notiftemplate.FieldAppID,
		})
		_node.AppID = value
	}
	if value, ok := ntc.mutation.LangID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: notiftemplate.FieldLangID,
		})
		_node.LangID = value
	}
	if value, ok := ntc.mutation.UsedFor(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: notiftemplate.FieldUsedFor,
		})
		_node.UsedFor = value
	}
	if value, ok := ntc.mutation.Title(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: notiftemplate.FieldTitle,
		})
		_node.Title = value
	}
	if value, ok := ntc.mutation.Content(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: notiftemplate.FieldContent,
		})
		_node.Content = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.NotifTemplate.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.NotifTemplateUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (ntc *NotifTemplateCreate) OnConflict(opts ...sql.ConflictOption) *NotifTemplateUpsertOne {
	ntc.conflict = opts
	return &NotifTemplateUpsertOne{
		create: ntc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.NotifTemplate.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (ntc *NotifTemplateCreate) OnConflictColumns(columns ...string) *NotifTemplateUpsertOne {
	ntc.conflict = append(ntc.conflict, sql.ConflictColumns(columns...))
	return &NotifTemplateUpsertOne{
		create: ntc,
	}
}

type (
	// NotifTemplateUpsertOne is the builder for "upsert"-ing
	//  one NotifTemplate node.
	NotifTemplateUpsertOne struct {
		create *NotifTemplateCreate
	}

	// NotifTemplateUpsert is the "OnConflict" setter.
	NotifTemplateUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedAt sets the "created_at" field.
func (u *NotifTemplateUpsert) SetCreatedAt(v uint32) *NotifTemplateUpsert {
	u.Set(notiftemplate.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *NotifTemplateUpsert) UpdateCreatedAt() *NotifTemplateUpsert {
	u.SetExcluded(notiftemplate.FieldCreatedAt)
	return u
}

// AddCreatedAt adds v to the "created_at" field.
func (u *NotifTemplateUpsert) AddCreatedAt(v uint32) *NotifTemplateUpsert {
	u.Add(notiftemplate.FieldCreatedAt, v)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *NotifTemplateUpsert) SetUpdatedAt(v uint32) *NotifTemplateUpsert {
	u.Set(notiftemplate.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *NotifTemplateUpsert) UpdateUpdatedAt() *NotifTemplateUpsert {
	u.SetExcluded(notiftemplate.FieldUpdatedAt)
	return u
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *NotifTemplateUpsert) AddUpdatedAt(v uint32) *NotifTemplateUpsert {
	u.Add(notiftemplate.FieldUpdatedAt, v)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *NotifTemplateUpsert) SetDeletedAt(v uint32) *NotifTemplateUpsert {
	u.Set(notiftemplate.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *NotifTemplateUpsert) UpdateDeletedAt() *NotifTemplateUpsert {
	u.SetExcluded(notiftemplate.FieldDeletedAt)
	return u
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *NotifTemplateUpsert) AddDeletedAt(v uint32) *NotifTemplateUpsert {
	u.Add(notiftemplate.FieldDeletedAt, v)
	return u
}

// SetAppID sets the "app_id" field.
func (u *NotifTemplateUpsert) SetAppID(v uuid.UUID) *NotifTemplateUpsert {
	u.Set(notiftemplate.FieldAppID, v)
	return u
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *NotifTemplateUpsert) UpdateAppID() *NotifTemplateUpsert {
	u.SetExcluded(notiftemplate.FieldAppID)
	return u
}

// SetLangID sets the "lang_id" field.
func (u *NotifTemplateUpsert) SetLangID(v uuid.UUID) *NotifTemplateUpsert {
	u.Set(notiftemplate.FieldLangID, v)
	return u
}

// UpdateLangID sets the "lang_id" field to the value that was provided on create.
func (u *NotifTemplateUpsert) UpdateLangID() *NotifTemplateUpsert {
	u.SetExcluded(notiftemplate.FieldLangID)
	return u
}

// SetUsedFor sets the "used_for" field.
func (u *NotifTemplateUpsert) SetUsedFor(v string) *NotifTemplateUpsert {
	u.Set(notiftemplate.FieldUsedFor, v)
	return u
}

// UpdateUsedFor sets the "used_for" field to the value that was provided on create.
func (u *NotifTemplateUpsert) UpdateUsedFor() *NotifTemplateUpsert {
	u.SetExcluded(notiftemplate.FieldUsedFor)
	return u
}

// ClearUsedFor clears the value of the "used_for" field.
func (u *NotifTemplateUpsert) ClearUsedFor() *NotifTemplateUpsert {
	u.SetNull(notiftemplate.FieldUsedFor)
	return u
}

// SetTitle sets the "title" field.
func (u *NotifTemplateUpsert) SetTitle(v string) *NotifTemplateUpsert {
	u.Set(notiftemplate.FieldTitle, v)
	return u
}

// UpdateTitle sets the "title" field to the value that was provided on create.
func (u *NotifTemplateUpsert) UpdateTitle() *NotifTemplateUpsert {
	u.SetExcluded(notiftemplate.FieldTitle)
	return u
}

// ClearTitle clears the value of the "title" field.
func (u *NotifTemplateUpsert) ClearTitle() *NotifTemplateUpsert {
	u.SetNull(notiftemplate.FieldTitle)
	return u
}

// SetContent sets the "content" field.
func (u *NotifTemplateUpsert) SetContent(v string) *NotifTemplateUpsert {
	u.Set(notiftemplate.FieldContent, v)
	return u
}

// UpdateContent sets the "content" field to the value that was provided on create.
func (u *NotifTemplateUpsert) UpdateContent() *NotifTemplateUpsert {
	u.SetExcluded(notiftemplate.FieldContent)
	return u
}

// ClearContent clears the value of the "content" field.
func (u *NotifTemplateUpsert) ClearContent() *NotifTemplateUpsert {
	u.SetNull(notiftemplate.FieldContent)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.NotifTemplate.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(notiftemplate.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *NotifTemplateUpsertOne) UpdateNewValues() *NotifTemplateUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(notiftemplate.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.NotifTemplate.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *NotifTemplateUpsertOne) Ignore() *NotifTemplateUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *NotifTemplateUpsertOne) DoNothing() *NotifTemplateUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the NotifTemplateCreate.OnConflict
// documentation for more info.
func (u *NotifTemplateUpsertOne) Update(set func(*NotifTemplateUpsert)) *NotifTemplateUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&NotifTemplateUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *NotifTemplateUpsertOne) SetCreatedAt(v uint32) *NotifTemplateUpsertOne {
	return u.Update(func(s *NotifTemplateUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *NotifTemplateUpsertOne) AddCreatedAt(v uint32) *NotifTemplateUpsertOne {
	return u.Update(func(s *NotifTemplateUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *NotifTemplateUpsertOne) UpdateCreatedAt() *NotifTemplateUpsertOne {
	return u.Update(func(s *NotifTemplateUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *NotifTemplateUpsertOne) SetUpdatedAt(v uint32) *NotifTemplateUpsertOne {
	return u.Update(func(s *NotifTemplateUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *NotifTemplateUpsertOne) AddUpdatedAt(v uint32) *NotifTemplateUpsertOne {
	return u.Update(func(s *NotifTemplateUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *NotifTemplateUpsertOne) UpdateUpdatedAt() *NotifTemplateUpsertOne {
	return u.Update(func(s *NotifTemplateUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *NotifTemplateUpsertOne) SetDeletedAt(v uint32) *NotifTemplateUpsertOne {
	return u.Update(func(s *NotifTemplateUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *NotifTemplateUpsertOne) AddDeletedAt(v uint32) *NotifTemplateUpsertOne {
	return u.Update(func(s *NotifTemplateUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *NotifTemplateUpsertOne) UpdateDeletedAt() *NotifTemplateUpsertOne {
	return u.Update(func(s *NotifTemplateUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetAppID sets the "app_id" field.
func (u *NotifTemplateUpsertOne) SetAppID(v uuid.UUID) *NotifTemplateUpsertOne {
	return u.Update(func(s *NotifTemplateUpsert) {
		s.SetAppID(v)
	})
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *NotifTemplateUpsertOne) UpdateAppID() *NotifTemplateUpsertOne {
	return u.Update(func(s *NotifTemplateUpsert) {
		s.UpdateAppID()
	})
}

// SetLangID sets the "lang_id" field.
func (u *NotifTemplateUpsertOne) SetLangID(v uuid.UUID) *NotifTemplateUpsertOne {
	return u.Update(func(s *NotifTemplateUpsert) {
		s.SetLangID(v)
	})
}

// UpdateLangID sets the "lang_id" field to the value that was provided on create.
func (u *NotifTemplateUpsertOne) UpdateLangID() *NotifTemplateUpsertOne {
	return u.Update(func(s *NotifTemplateUpsert) {
		s.UpdateLangID()
	})
}

// SetUsedFor sets the "used_for" field.
func (u *NotifTemplateUpsertOne) SetUsedFor(v string) *NotifTemplateUpsertOne {
	return u.Update(func(s *NotifTemplateUpsert) {
		s.SetUsedFor(v)
	})
}

// UpdateUsedFor sets the "used_for" field to the value that was provided on create.
func (u *NotifTemplateUpsertOne) UpdateUsedFor() *NotifTemplateUpsertOne {
	return u.Update(func(s *NotifTemplateUpsert) {
		s.UpdateUsedFor()
	})
}

// ClearUsedFor clears the value of the "used_for" field.
func (u *NotifTemplateUpsertOne) ClearUsedFor() *NotifTemplateUpsertOne {
	return u.Update(func(s *NotifTemplateUpsert) {
		s.ClearUsedFor()
	})
}

// SetTitle sets the "title" field.
func (u *NotifTemplateUpsertOne) SetTitle(v string) *NotifTemplateUpsertOne {
	return u.Update(func(s *NotifTemplateUpsert) {
		s.SetTitle(v)
	})
}

// UpdateTitle sets the "title" field to the value that was provided on create.
func (u *NotifTemplateUpsertOne) UpdateTitle() *NotifTemplateUpsertOne {
	return u.Update(func(s *NotifTemplateUpsert) {
		s.UpdateTitle()
	})
}

// ClearTitle clears the value of the "title" field.
func (u *NotifTemplateUpsertOne) ClearTitle() *NotifTemplateUpsertOne {
	return u.Update(func(s *NotifTemplateUpsert) {
		s.ClearTitle()
	})
}

// SetContent sets the "content" field.
func (u *NotifTemplateUpsertOne) SetContent(v string) *NotifTemplateUpsertOne {
	return u.Update(func(s *NotifTemplateUpsert) {
		s.SetContent(v)
	})
}

// UpdateContent sets the "content" field to the value that was provided on create.
func (u *NotifTemplateUpsertOne) UpdateContent() *NotifTemplateUpsertOne {
	return u.Update(func(s *NotifTemplateUpsert) {
		s.UpdateContent()
	})
}

// ClearContent clears the value of the "content" field.
func (u *NotifTemplateUpsertOne) ClearContent() *NotifTemplateUpsertOne {
	return u.Update(func(s *NotifTemplateUpsert) {
		s.ClearContent()
	})
}

// Exec executes the query.
func (u *NotifTemplateUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for NotifTemplateCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *NotifTemplateUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *NotifTemplateUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: NotifTemplateUpsertOne.ID is not supported by MySQL driver. Use NotifTemplateUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *NotifTemplateUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// NotifTemplateCreateBulk is the builder for creating many NotifTemplate entities in bulk.
type NotifTemplateCreateBulk struct {
	config
	builders []*NotifTemplateCreate
	conflict []sql.ConflictOption
}

// Save creates the NotifTemplate entities in the database.
func (ntcb *NotifTemplateCreateBulk) Save(ctx context.Context) ([]*NotifTemplate, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ntcb.builders))
	nodes := make([]*NotifTemplate, len(ntcb.builders))
	mutators := make([]Mutator, len(ntcb.builders))
	for i := range ntcb.builders {
		func(i int, root context.Context) {
			builder := ntcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*NotifTemplateMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ntcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = ntcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ntcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ntcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ntcb *NotifTemplateCreateBulk) SaveX(ctx context.Context) []*NotifTemplate {
	v, err := ntcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ntcb *NotifTemplateCreateBulk) Exec(ctx context.Context) error {
	_, err := ntcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ntcb *NotifTemplateCreateBulk) ExecX(ctx context.Context) {
	if err := ntcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.NotifTemplate.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.NotifTemplateUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (ntcb *NotifTemplateCreateBulk) OnConflict(opts ...sql.ConflictOption) *NotifTemplateUpsertBulk {
	ntcb.conflict = opts
	return &NotifTemplateUpsertBulk{
		create: ntcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.NotifTemplate.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (ntcb *NotifTemplateCreateBulk) OnConflictColumns(columns ...string) *NotifTemplateUpsertBulk {
	ntcb.conflict = append(ntcb.conflict, sql.ConflictColumns(columns...))
	return &NotifTemplateUpsertBulk{
		create: ntcb,
	}
}

// NotifTemplateUpsertBulk is the builder for "upsert"-ing
// a bulk of NotifTemplate nodes.
type NotifTemplateUpsertBulk struct {
	create *NotifTemplateCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.NotifTemplate.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(notiftemplate.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *NotifTemplateUpsertBulk) UpdateNewValues() *NotifTemplateUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(notiftemplate.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.NotifTemplate.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *NotifTemplateUpsertBulk) Ignore() *NotifTemplateUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *NotifTemplateUpsertBulk) DoNothing() *NotifTemplateUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the NotifTemplateCreateBulk.OnConflict
// documentation for more info.
func (u *NotifTemplateUpsertBulk) Update(set func(*NotifTemplateUpsert)) *NotifTemplateUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&NotifTemplateUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *NotifTemplateUpsertBulk) SetCreatedAt(v uint32) *NotifTemplateUpsertBulk {
	return u.Update(func(s *NotifTemplateUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *NotifTemplateUpsertBulk) AddCreatedAt(v uint32) *NotifTemplateUpsertBulk {
	return u.Update(func(s *NotifTemplateUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *NotifTemplateUpsertBulk) UpdateCreatedAt() *NotifTemplateUpsertBulk {
	return u.Update(func(s *NotifTemplateUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *NotifTemplateUpsertBulk) SetUpdatedAt(v uint32) *NotifTemplateUpsertBulk {
	return u.Update(func(s *NotifTemplateUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *NotifTemplateUpsertBulk) AddUpdatedAt(v uint32) *NotifTemplateUpsertBulk {
	return u.Update(func(s *NotifTemplateUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *NotifTemplateUpsertBulk) UpdateUpdatedAt() *NotifTemplateUpsertBulk {
	return u.Update(func(s *NotifTemplateUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *NotifTemplateUpsertBulk) SetDeletedAt(v uint32) *NotifTemplateUpsertBulk {
	return u.Update(func(s *NotifTemplateUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *NotifTemplateUpsertBulk) AddDeletedAt(v uint32) *NotifTemplateUpsertBulk {
	return u.Update(func(s *NotifTemplateUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *NotifTemplateUpsertBulk) UpdateDeletedAt() *NotifTemplateUpsertBulk {
	return u.Update(func(s *NotifTemplateUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetAppID sets the "app_id" field.
func (u *NotifTemplateUpsertBulk) SetAppID(v uuid.UUID) *NotifTemplateUpsertBulk {
	return u.Update(func(s *NotifTemplateUpsert) {
		s.SetAppID(v)
	})
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *NotifTemplateUpsertBulk) UpdateAppID() *NotifTemplateUpsertBulk {
	return u.Update(func(s *NotifTemplateUpsert) {
		s.UpdateAppID()
	})
}

// SetLangID sets the "lang_id" field.
func (u *NotifTemplateUpsertBulk) SetLangID(v uuid.UUID) *NotifTemplateUpsertBulk {
	return u.Update(func(s *NotifTemplateUpsert) {
		s.SetLangID(v)
	})
}

// UpdateLangID sets the "lang_id" field to the value that was provided on create.
func (u *NotifTemplateUpsertBulk) UpdateLangID() *NotifTemplateUpsertBulk {
	return u.Update(func(s *NotifTemplateUpsert) {
		s.UpdateLangID()
	})
}

// SetUsedFor sets the "used_for" field.
func (u *NotifTemplateUpsertBulk) SetUsedFor(v string) *NotifTemplateUpsertBulk {
	return u.Update(func(s *NotifTemplateUpsert) {
		s.SetUsedFor(v)
	})
}

// UpdateUsedFor sets the "used_for" field to the value that was provided on create.
func (u *NotifTemplateUpsertBulk) UpdateUsedFor() *NotifTemplateUpsertBulk {
	return u.Update(func(s *NotifTemplateUpsert) {
		s.UpdateUsedFor()
	})
}

// ClearUsedFor clears the value of the "used_for" field.
func (u *NotifTemplateUpsertBulk) ClearUsedFor() *NotifTemplateUpsertBulk {
	return u.Update(func(s *NotifTemplateUpsert) {
		s.ClearUsedFor()
	})
}

// SetTitle sets the "title" field.
func (u *NotifTemplateUpsertBulk) SetTitle(v string) *NotifTemplateUpsertBulk {
	return u.Update(func(s *NotifTemplateUpsert) {
		s.SetTitle(v)
	})
}

// UpdateTitle sets the "title" field to the value that was provided on create.
func (u *NotifTemplateUpsertBulk) UpdateTitle() *NotifTemplateUpsertBulk {
	return u.Update(func(s *NotifTemplateUpsert) {
		s.UpdateTitle()
	})
}

// ClearTitle clears the value of the "title" field.
func (u *NotifTemplateUpsertBulk) ClearTitle() *NotifTemplateUpsertBulk {
	return u.Update(func(s *NotifTemplateUpsert) {
		s.ClearTitle()
	})
}

// SetContent sets the "content" field.
func (u *NotifTemplateUpsertBulk) SetContent(v string) *NotifTemplateUpsertBulk {
	return u.Update(func(s *NotifTemplateUpsert) {
		s.SetContent(v)
	})
}

// UpdateContent sets the "content" field to the value that was provided on create.
func (u *NotifTemplateUpsertBulk) UpdateContent() *NotifTemplateUpsertBulk {
	return u.Update(func(s *NotifTemplateUpsert) {
		s.UpdateContent()
	})
}

// ClearContent clears the value of the "content" field.
func (u *NotifTemplateUpsertBulk) ClearContent() *NotifTemplateUpsertBulk {
	return u.Update(func(s *NotifTemplateUpsert) {
		s.ClearContent()
	})
}

// Exec executes the query.
func (u *NotifTemplateUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the NotifTemplateCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for NotifTemplateCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *NotifTemplateUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
