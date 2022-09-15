// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/third-manager/pkg/db/ent/appcontact"
	"github.com/NpoolPlatform/third-manager/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// AppContactUpdate is the builder for updating AppContact entities.
type AppContactUpdate struct {
	config
	hooks     []Hook
	mutation  *AppContactMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the AppContactUpdate builder.
func (acu *AppContactUpdate) Where(ps ...predicate.AppContact) *AppContactUpdate {
	acu.mutation.Where(ps...)
	return acu
}

// SetCreatedAt sets the "created_at" field.
func (acu *AppContactUpdate) SetCreatedAt(u uint32) *AppContactUpdate {
	acu.mutation.ResetCreatedAt()
	acu.mutation.SetCreatedAt(u)
	return acu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (acu *AppContactUpdate) SetNillableCreatedAt(u *uint32) *AppContactUpdate {
	if u != nil {
		acu.SetCreatedAt(*u)
	}
	return acu
}

// AddCreatedAt adds u to the "created_at" field.
func (acu *AppContactUpdate) AddCreatedAt(u int32) *AppContactUpdate {
	acu.mutation.AddCreatedAt(u)
	return acu
}

// SetUpdatedAt sets the "updated_at" field.
func (acu *AppContactUpdate) SetUpdatedAt(u uint32) *AppContactUpdate {
	acu.mutation.ResetUpdatedAt()
	acu.mutation.SetUpdatedAt(u)
	return acu
}

// AddUpdatedAt adds u to the "updated_at" field.
func (acu *AppContactUpdate) AddUpdatedAt(u int32) *AppContactUpdate {
	acu.mutation.AddUpdatedAt(u)
	return acu
}

// SetDeletedAt sets the "deleted_at" field.
func (acu *AppContactUpdate) SetDeletedAt(u uint32) *AppContactUpdate {
	acu.mutation.ResetDeletedAt()
	acu.mutation.SetDeletedAt(u)
	return acu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (acu *AppContactUpdate) SetNillableDeletedAt(u *uint32) *AppContactUpdate {
	if u != nil {
		acu.SetDeletedAt(*u)
	}
	return acu
}

// AddDeletedAt adds u to the "deleted_at" field.
func (acu *AppContactUpdate) AddDeletedAt(u int32) *AppContactUpdate {
	acu.mutation.AddDeletedAt(u)
	return acu
}

// SetAppID sets the "app_id" field.
func (acu *AppContactUpdate) SetAppID(u uuid.UUID) *AppContactUpdate {
	acu.mutation.SetAppID(u)
	return acu
}

// SetUsedFor sets the "used_for" field.
func (acu *AppContactUpdate) SetUsedFor(s string) *AppContactUpdate {
	acu.mutation.SetUsedFor(s)
	return acu
}

// SetNillableUsedFor sets the "used_for" field if the given value is not nil.
func (acu *AppContactUpdate) SetNillableUsedFor(s *string) *AppContactUpdate {
	if s != nil {
		acu.SetUsedFor(*s)
	}
	return acu
}

// SetSender sets the "sender" field.
func (acu *AppContactUpdate) SetSender(s string) *AppContactUpdate {
	acu.mutation.SetSender(s)
	return acu
}

// SetNillableSender sets the "sender" field if the given value is not nil.
func (acu *AppContactUpdate) SetNillableSender(s *string) *AppContactUpdate {
	if s != nil {
		acu.SetSender(*s)
	}
	return acu
}

// SetAccount sets the "account" field.
func (acu *AppContactUpdate) SetAccount(s string) *AppContactUpdate {
	acu.mutation.SetAccount(s)
	return acu
}

// SetNillableAccount sets the "account" field if the given value is not nil.
func (acu *AppContactUpdate) SetNillableAccount(s *string) *AppContactUpdate {
	if s != nil {
		acu.SetAccount(*s)
	}
	return acu
}

// SetAccountType sets the "account_type" field.
func (acu *AppContactUpdate) SetAccountType(s string) *AppContactUpdate {
	acu.mutation.SetAccountType(s)
	return acu
}

// SetNillableAccountType sets the "account_type" field if the given value is not nil.
func (acu *AppContactUpdate) SetNillableAccountType(s *string) *AppContactUpdate {
	if s != nil {
		acu.SetAccountType(*s)
	}
	return acu
}

// Mutation returns the AppContactMutation object of the builder.
func (acu *AppContactUpdate) Mutation() *AppContactMutation {
	return acu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (acu *AppContactUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if err := acu.defaults(); err != nil {
		return 0, err
	}
	if len(acu.hooks) == 0 {
		if err = acu.check(); err != nil {
			return 0, err
		}
		affected, err = acu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AppContactMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = acu.check(); err != nil {
				return 0, err
			}
			acu.mutation = mutation
			affected, err = acu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(acu.hooks) - 1; i >= 0; i-- {
			if acu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = acu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, acu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (acu *AppContactUpdate) SaveX(ctx context.Context) int {
	affected, err := acu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (acu *AppContactUpdate) Exec(ctx context.Context) error {
	_, err := acu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acu *AppContactUpdate) ExecX(ctx context.Context) {
	if err := acu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (acu *AppContactUpdate) defaults() error {
	if _, ok := acu.mutation.UpdatedAt(); !ok {
		if appcontact.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized appcontact.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := appcontact.UpdateDefaultUpdatedAt()
		acu.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (acu *AppContactUpdate) check() error {
	if v, ok := acu.mutation.UsedFor(); ok {
		if err := appcontact.UsedForValidator(v); err != nil {
			return &ValidationError{Name: "used_for", err: fmt.Errorf(`ent: validator failed for field "AppContact.used_for": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (acu *AppContactUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *AppContactUpdate {
	acu.modifiers = append(acu.modifiers, modifiers...)
	return acu
}

func (acu *AppContactUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   appcontact.Table,
			Columns: appcontact.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: appcontact.FieldID,
			},
		},
	}
	if ps := acu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := acu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appcontact.FieldCreatedAt,
		})
	}
	if value, ok := acu.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appcontact.FieldCreatedAt,
		})
	}
	if value, ok := acu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appcontact.FieldUpdatedAt,
		})
	}
	if value, ok := acu.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appcontact.FieldUpdatedAt,
		})
	}
	if value, ok := acu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appcontact.FieldDeletedAt,
		})
	}
	if value, ok := acu.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appcontact.FieldDeletedAt,
		})
	}
	if value, ok := acu.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: appcontact.FieldAppID,
		})
	}
	if value, ok := acu.mutation.UsedFor(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: appcontact.FieldUsedFor,
		})
	}
	if value, ok := acu.mutation.Sender(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: appcontact.FieldSender,
		})
	}
	if value, ok := acu.mutation.Account(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: appcontact.FieldAccount,
		})
	}
	if value, ok := acu.mutation.AccountType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: appcontact.FieldAccountType,
		})
	}
	_spec.Modifiers = acu.modifiers
	if n, err = sqlgraph.UpdateNodes(ctx, acu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{appcontact.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// AppContactUpdateOne is the builder for updating a single AppContact entity.
type AppContactUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *AppContactMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedAt sets the "created_at" field.
func (acuo *AppContactUpdateOne) SetCreatedAt(u uint32) *AppContactUpdateOne {
	acuo.mutation.ResetCreatedAt()
	acuo.mutation.SetCreatedAt(u)
	return acuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (acuo *AppContactUpdateOne) SetNillableCreatedAt(u *uint32) *AppContactUpdateOne {
	if u != nil {
		acuo.SetCreatedAt(*u)
	}
	return acuo
}

// AddCreatedAt adds u to the "created_at" field.
func (acuo *AppContactUpdateOne) AddCreatedAt(u int32) *AppContactUpdateOne {
	acuo.mutation.AddCreatedAt(u)
	return acuo
}

// SetUpdatedAt sets the "updated_at" field.
func (acuo *AppContactUpdateOne) SetUpdatedAt(u uint32) *AppContactUpdateOne {
	acuo.mutation.ResetUpdatedAt()
	acuo.mutation.SetUpdatedAt(u)
	return acuo
}

// AddUpdatedAt adds u to the "updated_at" field.
func (acuo *AppContactUpdateOne) AddUpdatedAt(u int32) *AppContactUpdateOne {
	acuo.mutation.AddUpdatedAt(u)
	return acuo
}

// SetDeletedAt sets the "deleted_at" field.
func (acuo *AppContactUpdateOne) SetDeletedAt(u uint32) *AppContactUpdateOne {
	acuo.mutation.ResetDeletedAt()
	acuo.mutation.SetDeletedAt(u)
	return acuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (acuo *AppContactUpdateOne) SetNillableDeletedAt(u *uint32) *AppContactUpdateOne {
	if u != nil {
		acuo.SetDeletedAt(*u)
	}
	return acuo
}

// AddDeletedAt adds u to the "deleted_at" field.
func (acuo *AppContactUpdateOne) AddDeletedAt(u int32) *AppContactUpdateOne {
	acuo.mutation.AddDeletedAt(u)
	return acuo
}

// SetAppID sets the "app_id" field.
func (acuo *AppContactUpdateOne) SetAppID(u uuid.UUID) *AppContactUpdateOne {
	acuo.mutation.SetAppID(u)
	return acuo
}

// SetUsedFor sets the "used_for" field.
func (acuo *AppContactUpdateOne) SetUsedFor(s string) *AppContactUpdateOne {
	acuo.mutation.SetUsedFor(s)
	return acuo
}

// SetNillableUsedFor sets the "used_for" field if the given value is not nil.
func (acuo *AppContactUpdateOne) SetNillableUsedFor(s *string) *AppContactUpdateOne {
	if s != nil {
		acuo.SetUsedFor(*s)
	}
	return acuo
}

// SetSender sets the "sender" field.
func (acuo *AppContactUpdateOne) SetSender(s string) *AppContactUpdateOne {
	acuo.mutation.SetSender(s)
	return acuo
}

// SetNillableSender sets the "sender" field if the given value is not nil.
func (acuo *AppContactUpdateOne) SetNillableSender(s *string) *AppContactUpdateOne {
	if s != nil {
		acuo.SetSender(*s)
	}
	return acuo
}

// SetAccount sets the "account" field.
func (acuo *AppContactUpdateOne) SetAccount(s string) *AppContactUpdateOne {
	acuo.mutation.SetAccount(s)
	return acuo
}

// SetNillableAccount sets the "account" field if the given value is not nil.
func (acuo *AppContactUpdateOne) SetNillableAccount(s *string) *AppContactUpdateOne {
	if s != nil {
		acuo.SetAccount(*s)
	}
	return acuo
}

// SetAccountType sets the "account_type" field.
func (acuo *AppContactUpdateOne) SetAccountType(s string) *AppContactUpdateOne {
	acuo.mutation.SetAccountType(s)
	return acuo
}

// SetNillableAccountType sets the "account_type" field if the given value is not nil.
func (acuo *AppContactUpdateOne) SetNillableAccountType(s *string) *AppContactUpdateOne {
	if s != nil {
		acuo.SetAccountType(*s)
	}
	return acuo
}

// Mutation returns the AppContactMutation object of the builder.
func (acuo *AppContactUpdateOne) Mutation() *AppContactMutation {
	return acuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (acuo *AppContactUpdateOne) Select(field string, fields ...string) *AppContactUpdateOne {
	acuo.fields = append([]string{field}, fields...)
	return acuo
}

// Save executes the query and returns the updated AppContact entity.
func (acuo *AppContactUpdateOne) Save(ctx context.Context) (*AppContact, error) {
	var (
		err  error
		node *AppContact
	)
	if err := acuo.defaults(); err != nil {
		return nil, err
	}
	if len(acuo.hooks) == 0 {
		if err = acuo.check(); err != nil {
			return nil, err
		}
		node, err = acuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AppContactMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = acuo.check(); err != nil {
				return nil, err
			}
			acuo.mutation = mutation
			node, err = acuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(acuo.hooks) - 1; i >= 0; i-- {
			if acuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = acuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, acuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*AppContact)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from AppContactMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (acuo *AppContactUpdateOne) SaveX(ctx context.Context) *AppContact {
	node, err := acuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (acuo *AppContactUpdateOne) Exec(ctx context.Context) error {
	_, err := acuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acuo *AppContactUpdateOne) ExecX(ctx context.Context) {
	if err := acuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (acuo *AppContactUpdateOne) defaults() error {
	if _, ok := acuo.mutation.UpdatedAt(); !ok {
		if appcontact.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized appcontact.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := appcontact.UpdateDefaultUpdatedAt()
		acuo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (acuo *AppContactUpdateOne) check() error {
	if v, ok := acuo.mutation.UsedFor(); ok {
		if err := appcontact.UsedForValidator(v); err != nil {
			return &ValidationError{Name: "used_for", err: fmt.Errorf(`ent: validator failed for field "AppContact.used_for": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (acuo *AppContactUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *AppContactUpdateOne {
	acuo.modifiers = append(acuo.modifiers, modifiers...)
	return acuo
}

func (acuo *AppContactUpdateOne) sqlSave(ctx context.Context) (_node *AppContact, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   appcontact.Table,
			Columns: appcontact.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: appcontact.FieldID,
			},
		},
	}
	id, ok := acuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "AppContact.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := acuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, appcontact.FieldID)
		for _, f := range fields {
			if !appcontact.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != appcontact.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := acuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := acuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appcontact.FieldCreatedAt,
		})
	}
	if value, ok := acuo.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appcontact.FieldCreatedAt,
		})
	}
	if value, ok := acuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appcontact.FieldUpdatedAt,
		})
	}
	if value, ok := acuo.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appcontact.FieldUpdatedAt,
		})
	}
	if value, ok := acuo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appcontact.FieldDeletedAt,
		})
	}
	if value, ok := acuo.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appcontact.FieldDeletedAt,
		})
	}
	if value, ok := acuo.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: appcontact.FieldAppID,
		})
	}
	if value, ok := acuo.mutation.UsedFor(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: appcontact.FieldUsedFor,
		})
	}
	if value, ok := acuo.mutation.Sender(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: appcontact.FieldSender,
		})
	}
	if value, ok := acuo.mutation.Account(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: appcontact.FieldAccount,
		})
	}
	if value, ok := acuo.mutation.AccountType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: appcontact.FieldAccountType,
		})
	}
	_spec.Modifiers = acuo.modifiers
	_node = &AppContact{config: acuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, acuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{appcontact.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
