// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/third-manager/pkg/db/ent/notiftemplate"
	"github.com/NpoolPlatform/third-manager/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// NotifTemplateUpdate is the builder for updating NotifTemplate entities.
type NotifTemplateUpdate struct {
	config
	hooks     []Hook
	mutation  *NotifTemplateMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the NotifTemplateUpdate builder.
func (ntu *NotifTemplateUpdate) Where(ps ...predicate.NotifTemplate) *NotifTemplateUpdate {
	ntu.mutation.Where(ps...)
	return ntu
}

// SetCreatedAt sets the "created_at" field.
func (ntu *NotifTemplateUpdate) SetCreatedAt(u uint32) *NotifTemplateUpdate {
	ntu.mutation.ResetCreatedAt()
	ntu.mutation.SetCreatedAt(u)
	return ntu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ntu *NotifTemplateUpdate) SetNillableCreatedAt(u *uint32) *NotifTemplateUpdate {
	if u != nil {
		ntu.SetCreatedAt(*u)
	}
	return ntu
}

// AddCreatedAt adds u to the "created_at" field.
func (ntu *NotifTemplateUpdate) AddCreatedAt(u int32) *NotifTemplateUpdate {
	ntu.mutation.AddCreatedAt(u)
	return ntu
}

// SetUpdatedAt sets the "updated_at" field.
func (ntu *NotifTemplateUpdate) SetUpdatedAt(u uint32) *NotifTemplateUpdate {
	ntu.mutation.ResetUpdatedAt()
	ntu.mutation.SetUpdatedAt(u)
	return ntu
}

// AddUpdatedAt adds u to the "updated_at" field.
func (ntu *NotifTemplateUpdate) AddUpdatedAt(u int32) *NotifTemplateUpdate {
	ntu.mutation.AddUpdatedAt(u)
	return ntu
}

// SetDeletedAt sets the "deleted_at" field.
func (ntu *NotifTemplateUpdate) SetDeletedAt(u uint32) *NotifTemplateUpdate {
	ntu.mutation.ResetDeletedAt()
	ntu.mutation.SetDeletedAt(u)
	return ntu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ntu *NotifTemplateUpdate) SetNillableDeletedAt(u *uint32) *NotifTemplateUpdate {
	if u != nil {
		ntu.SetDeletedAt(*u)
	}
	return ntu
}

// AddDeletedAt adds u to the "deleted_at" field.
func (ntu *NotifTemplateUpdate) AddDeletedAt(u int32) *NotifTemplateUpdate {
	ntu.mutation.AddDeletedAt(u)
	return ntu
}

// SetAppID sets the "app_id" field.
func (ntu *NotifTemplateUpdate) SetAppID(u uuid.UUID) *NotifTemplateUpdate {
	ntu.mutation.SetAppID(u)
	return ntu
}

// SetLangID sets the "lang_id" field.
func (ntu *NotifTemplateUpdate) SetLangID(u uuid.UUID) *NotifTemplateUpdate {
	ntu.mutation.SetLangID(u)
	return ntu
}

// SetUsedFor sets the "used_for" field.
func (ntu *NotifTemplateUpdate) SetUsedFor(s string) *NotifTemplateUpdate {
	ntu.mutation.SetUsedFor(s)
	return ntu
}

// SetNillableUsedFor sets the "used_for" field if the given value is not nil.
func (ntu *NotifTemplateUpdate) SetNillableUsedFor(s *string) *NotifTemplateUpdate {
	if s != nil {
		ntu.SetUsedFor(*s)
	}
	return ntu
}

// ClearUsedFor clears the value of the "used_for" field.
func (ntu *NotifTemplateUpdate) ClearUsedFor() *NotifTemplateUpdate {
	ntu.mutation.ClearUsedFor()
	return ntu
}

// SetTitle sets the "title" field.
func (ntu *NotifTemplateUpdate) SetTitle(s string) *NotifTemplateUpdate {
	ntu.mutation.SetTitle(s)
	return ntu
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (ntu *NotifTemplateUpdate) SetNillableTitle(s *string) *NotifTemplateUpdate {
	if s != nil {
		ntu.SetTitle(*s)
	}
	return ntu
}

// ClearTitle clears the value of the "title" field.
func (ntu *NotifTemplateUpdate) ClearTitle() *NotifTemplateUpdate {
	ntu.mutation.ClearTitle()
	return ntu
}

// SetContent sets the "content" field.
func (ntu *NotifTemplateUpdate) SetContent(s string) *NotifTemplateUpdate {
	ntu.mutation.SetContent(s)
	return ntu
}

// SetNillableContent sets the "content" field if the given value is not nil.
func (ntu *NotifTemplateUpdate) SetNillableContent(s *string) *NotifTemplateUpdate {
	if s != nil {
		ntu.SetContent(*s)
	}
	return ntu
}

// ClearContent clears the value of the "content" field.
func (ntu *NotifTemplateUpdate) ClearContent() *NotifTemplateUpdate {
	ntu.mutation.ClearContent()
	return ntu
}

// Mutation returns the NotifTemplateMutation object of the builder.
func (ntu *NotifTemplateUpdate) Mutation() *NotifTemplateMutation {
	return ntu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ntu *NotifTemplateUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if err := ntu.defaults(); err != nil {
		return 0, err
	}
	if len(ntu.hooks) == 0 {
		affected, err = ntu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*NotifTemplateMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ntu.mutation = mutation
			affected, err = ntu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ntu.hooks) - 1; i >= 0; i-- {
			if ntu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ntu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ntu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ntu *NotifTemplateUpdate) SaveX(ctx context.Context) int {
	affected, err := ntu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ntu *NotifTemplateUpdate) Exec(ctx context.Context) error {
	_, err := ntu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ntu *NotifTemplateUpdate) ExecX(ctx context.Context) {
	if err := ntu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ntu *NotifTemplateUpdate) defaults() error {
	if _, ok := ntu.mutation.UpdatedAt(); !ok {
		if notiftemplate.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized notiftemplate.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := notiftemplate.UpdateDefaultUpdatedAt()
		ntu.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (ntu *NotifTemplateUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *NotifTemplateUpdate {
	ntu.modifiers = append(ntu.modifiers, modifiers...)
	return ntu
}

func (ntu *NotifTemplateUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   notiftemplate.Table,
			Columns: notiftemplate.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: notiftemplate.FieldID,
			},
		},
	}
	if ps := ntu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ntu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: notiftemplate.FieldCreatedAt,
		})
	}
	if value, ok := ntu.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: notiftemplate.FieldCreatedAt,
		})
	}
	if value, ok := ntu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: notiftemplate.FieldUpdatedAt,
		})
	}
	if value, ok := ntu.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: notiftemplate.FieldUpdatedAt,
		})
	}
	if value, ok := ntu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: notiftemplate.FieldDeletedAt,
		})
	}
	if value, ok := ntu.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: notiftemplate.FieldDeletedAt,
		})
	}
	if value, ok := ntu.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: notiftemplate.FieldAppID,
		})
	}
	if value, ok := ntu.mutation.LangID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: notiftemplate.FieldLangID,
		})
	}
	if value, ok := ntu.mutation.UsedFor(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: notiftemplate.FieldUsedFor,
		})
	}
	if ntu.mutation.UsedForCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: notiftemplate.FieldUsedFor,
		})
	}
	if value, ok := ntu.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: notiftemplate.FieldTitle,
		})
	}
	if ntu.mutation.TitleCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: notiftemplate.FieldTitle,
		})
	}
	if value, ok := ntu.mutation.Content(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: notiftemplate.FieldContent,
		})
	}
	if ntu.mutation.ContentCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: notiftemplate.FieldContent,
		})
	}
	_spec.Modifiers = ntu.modifiers
	if n, err = sqlgraph.UpdateNodes(ctx, ntu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{notiftemplate.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// NotifTemplateUpdateOne is the builder for updating a single NotifTemplate entity.
type NotifTemplateUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *NotifTemplateMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedAt sets the "created_at" field.
func (ntuo *NotifTemplateUpdateOne) SetCreatedAt(u uint32) *NotifTemplateUpdateOne {
	ntuo.mutation.ResetCreatedAt()
	ntuo.mutation.SetCreatedAt(u)
	return ntuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ntuo *NotifTemplateUpdateOne) SetNillableCreatedAt(u *uint32) *NotifTemplateUpdateOne {
	if u != nil {
		ntuo.SetCreatedAt(*u)
	}
	return ntuo
}

// AddCreatedAt adds u to the "created_at" field.
func (ntuo *NotifTemplateUpdateOne) AddCreatedAt(u int32) *NotifTemplateUpdateOne {
	ntuo.mutation.AddCreatedAt(u)
	return ntuo
}

// SetUpdatedAt sets the "updated_at" field.
func (ntuo *NotifTemplateUpdateOne) SetUpdatedAt(u uint32) *NotifTemplateUpdateOne {
	ntuo.mutation.ResetUpdatedAt()
	ntuo.mutation.SetUpdatedAt(u)
	return ntuo
}

// AddUpdatedAt adds u to the "updated_at" field.
func (ntuo *NotifTemplateUpdateOne) AddUpdatedAt(u int32) *NotifTemplateUpdateOne {
	ntuo.mutation.AddUpdatedAt(u)
	return ntuo
}

// SetDeletedAt sets the "deleted_at" field.
func (ntuo *NotifTemplateUpdateOne) SetDeletedAt(u uint32) *NotifTemplateUpdateOne {
	ntuo.mutation.ResetDeletedAt()
	ntuo.mutation.SetDeletedAt(u)
	return ntuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ntuo *NotifTemplateUpdateOne) SetNillableDeletedAt(u *uint32) *NotifTemplateUpdateOne {
	if u != nil {
		ntuo.SetDeletedAt(*u)
	}
	return ntuo
}

// AddDeletedAt adds u to the "deleted_at" field.
func (ntuo *NotifTemplateUpdateOne) AddDeletedAt(u int32) *NotifTemplateUpdateOne {
	ntuo.mutation.AddDeletedAt(u)
	return ntuo
}

// SetAppID sets the "app_id" field.
func (ntuo *NotifTemplateUpdateOne) SetAppID(u uuid.UUID) *NotifTemplateUpdateOne {
	ntuo.mutation.SetAppID(u)
	return ntuo
}

// SetLangID sets the "lang_id" field.
func (ntuo *NotifTemplateUpdateOne) SetLangID(u uuid.UUID) *NotifTemplateUpdateOne {
	ntuo.mutation.SetLangID(u)
	return ntuo
}

// SetUsedFor sets the "used_for" field.
func (ntuo *NotifTemplateUpdateOne) SetUsedFor(s string) *NotifTemplateUpdateOne {
	ntuo.mutation.SetUsedFor(s)
	return ntuo
}

// SetNillableUsedFor sets the "used_for" field if the given value is not nil.
func (ntuo *NotifTemplateUpdateOne) SetNillableUsedFor(s *string) *NotifTemplateUpdateOne {
	if s != nil {
		ntuo.SetUsedFor(*s)
	}
	return ntuo
}

// ClearUsedFor clears the value of the "used_for" field.
func (ntuo *NotifTemplateUpdateOne) ClearUsedFor() *NotifTemplateUpdateOne {
	ntuo.mutation.ClearUsedFor()
	return ntuo
}

// SetTitle sets the "title" field.
func (ntuo *NotifTemplateUpdateOne) SetTitle(s string) *NotifTemplateUpdateOne {
	ntuo.mutation.SetTitle(s)
	return ntuo
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (ntuo *NotifTemplateUpdateOne) SetNillableTitle(s *string) *NotifTemplateUpdateOne {
	if s != nil {
		ntuo.SetTitle(*s)
	}
	return ntuo
}

// ClearTitle clears the value of the "title" field.
func (ntuo *NotifTemplateUpdateOne) ClearTitle() *NotifTemplateUpdateOne {
	ntuo.mutation.ClearTitle()
	return ntuo
}

// SetContent sets the "content" field.
func (ntuo *NotifTemplateUpdateOne) SetContent(s string) *NotifTemplateUpdateOne {
	ntuo.mutation.SetContent(s)
	return ntuo
}

// SetNillableContent sets the "content" field if the given value is not nil.
func (ntuo *NotifTemplateUpdateOne) SetNillableContent(s *string) *NotifTemplateUpdateOne {
	if s != nil {
		ntuo.SetContent(*s)
	}
	return ntuo
}

// ClearContent clears the value of the "content" field.
func (ntuo *NotifTemplateUpdateOne) ClearContent() *NotifTemplateUpdateOne {
	ntuo.mutation.ClearContent()
	return ntuo
}

// Mutation returns the NotifTemplateMutation object of the builder.
func (ntuo *NotifTemplateUpdateOne) Mutation() *NotifTemplateMutation {
	return ntuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ntuo *NotifTemplateUpdateOne) Select(field string, fields ...string) *NotifTemplateUpdateOne {
	ntuo.fields = append([]string{field}, fields...)
	return ntuo
}

// Save executes the query and returns the updated NotifTemplate entity.
func (ntuo *NotifTemplateUpdateOne) Save(ctx context.Context) (*NotifTemplate, error) {
	var (
		err  error
		node *NotifTemplate
	)
	if err := ntuo.defaults(); err != nil {
		return nil, err
	}
	if len(ntuo.hooks) == 0 {
		node, err = ntuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*NotifTemplateMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ntuo.mutation = mutation
			node, err = ntuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ntuo.hooks) - 1; i >= 0; i-- {
			if ntuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ntuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, ntuo.mutation)
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

// SaveX is like Save, but panics if an error occurs.
func (ntuo *NotifTemplateUpdateOne) SaveX(ctx context.Context) *NotifTemplate {
	node, err := ntuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ntuo *NotifTemplateUpdateOne) Exec(ctx context.Context) error {
	_, err := ntuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ntuo *NotifTemplateUpdateOne) ExecX(ctx context.Context) {
	if err := ntuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ntuo *NotifTemplateUpdateOne) defaults() error {
	if _, ok := ntuo.mutation.UpdatedAt(); !ok {
		if notiftemplate.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized notiftemplate.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := notiftemplate.UpdateDefaultUpdatedAt()
		ntuo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (ntuo *NotifTemplateUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *NotifTemplateUpdateOne {
	ntuo.modifiers = append(ntuo.modifiers, modifiers...)
	return ntuo
}

func (ntuo *NotifTemplateUpdateOne) sqlSave(ctx context.Context) (_node *NotifTemplate, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   notiftemplate.Table,
			Columns: notiftemplate.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: notiftemplate.FieldID,
			},
		},
	}
	id, ok := ntuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "NotifTemplate.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ntuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, notiftemplate.FieldID)
		for _, f := range fields {
			if !notiftemplate.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != notiftemplate.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ntuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ntuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: notiftemplate.FieldCreatedAt,
		})
	}
	if value, ok := ntuo.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: notiftemplate.FieldCreatedAt,
		})
	}
	if value, ok := ntuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: notiftemplate.FieldUpdatedAt,
		})
	}
	if value, ok := ntuo.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: notiftemplate.FieldUpdatedAt,
		})
	}
	if value, ok := ntuo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: notiftemplate.FieldDeletedAt,
		})
	}
	if value, ok := ntuo.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: notiftemplate.FieldDeletedAt,
		})
	}
	if value, ok := ntuo.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: notiftemplate.FieldAppID,
		})
	}
	if value, ok := ntuo.mutation.LangID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: notiftemplate.FieldLangID,
		})
	}
	if value, ok := ntuo.mutation.UsedFor(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: notiftemplate.FieldUsedFor,
		})
	}
	if ntuo.mutation.UsedForCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: notiftemplate.FieldUsedFor,
		})
	}
	if value, ok := ntuo.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: notiftemplate.FieldTitle,
		})
	}
	if ntuo.mutation.TitleCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: notiftemplate.FieldTitle,
		})
	}
	if value, ok := ntuo.mutation.Content(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: notiftemplate.FieldContent,
		})
	}
	if ntuo.mutation.ContentCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: notiftemplate.FieldContent,
		})
	}
	_spec.Modifiers = ntuo.modifiers
	_node = &NotifTemplate{config: ntuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ntuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{notiftemplate.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
