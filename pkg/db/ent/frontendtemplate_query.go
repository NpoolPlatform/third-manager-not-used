// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/third-manager/pkg/db/ent/frontendtemplate"
	"github.com/NpoolPlatform/third-manager/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// FrontendTemplateQuery is the builder for querying FrontendTemplate entities.
type FrontendTemplateQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.FrontendTemplate
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the FrontendTemplateQuery builder.
func (ftq *FrontendTemplateQuery) Where(ps ...predicate.FrontendTemplate) *FrontendTemplateQuery {
	ftq.predicates = append(ftq.predicates, ps...)
	return ftq
}

// Limit adds a limit step to the query.
func (ftq *FrontendTemplateQuery) Limit(limit int) *FrontendTemplateQuery {
	ftq.limit = &limit
	return ftq
}

// Offset adds an offset step to the query.
func (ftq *FrontendTemplateQuery) Offset(offset int) *FrontendTemplateQuery {
	ftq.offset = &offset
	return ftq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ftq *FrontendTemplateQuery) Unique(unique bool) *FrontendTemplateQuery {
	ftq.unique = &unique
	return ftq
}

// Order adds an order step to the query.
func (ftq *FrontendTemplateQuery) Order(o ...OrderFunc) *FrontendTemplateQuery {
	ftq.order = append(ftq.order, o...)
	return ftq
}

// First returns the first FrontendTemplate entity from the query.
// Returns a *NotFoundError when no FrontendTemplate was found.
func (ftq *FrontendTemplateQuery) First(ctx context.Context) (*FrontendTemplate, error) {
	nodes, err := ftq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{frontendtemplate.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ftq *FrontendTemplateQuery) FirstX(ctx context.Context) *FrontendTemplate {
	node, err := ftq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first FrontendTemplate ID from the query.
// Returns a *NotFoundError when no FrontendTemplate ID was found.
func (ftq *FrontendTemplateQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = ftq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{frontendtemplate.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ftq *FrontendTemplateQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := ftq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single FrontendTemplate entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one FrontendTemplate entity is found.
// Returns a *NotFoundError when no FrontendTemplate entities are found.
func (ftq *FrontendTemplateQuery) Only(ctx context.Context) (*FrontendTemplate, error) {
	nodes, err := ftq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{frontendtemplate.Label}
	default:
		return nil, &NotSingularError{frontendtemplate.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ftq *FrontendTemplateQuery) OnlyX(ctx context.Context) *FrontendTemplate {
	node, err := ftq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only FrontendTemplate ID in the query.
// Returns a *NotSingularError when more than one FrontendTemplate ID is found.
// Returns a *NotFoundError when no entities are found.
func (ftq *FrontendTemplateQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = ftq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{frontendtemplate.Label}
	default:
		err = &NotSingularError{frontendtemplate.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ftq *FrontendTemplateQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := ftq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of FrontendTemplates.
func (ftq *FrontendTemplateQuery) All(ctx context.Context) ([]*FrontendTemplate, error) {
	if err := ftq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return ftq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (ftq *FrontendTemplateQuery) AllX(ctx context.Context) []*FrontendTemplate {
	nodes, err := ftq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of FrontendTemplate IDs.
func (ftq *FrontendTemplateQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := ftq.Select(frontendtemplate.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ftq *FrontendTemplateQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := ftq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ftq *FrontendTemplateQuery) Count(ctx context.Context) (int, error) {
	if err := ftq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return ftq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (ftq *FrontendTemplateQuery) CountX(ctx context.Context) int {
	count, err := ftq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ftq *FrontendTemplateQuery) Exist(ctx context.Context) (bool, error) {
	if err := ftq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return ftq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (ftq *FrontendTemplateQuery) ExistX(ctx context.Context) bool {
	exist, err := ftq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the FrontendTemplateQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ftq *FrontendTemplateQuery) Clone() *FrontendTemplateQuery {
	if ftq == nil {
		return nil
	}
	return &FrontendTemplateQuery{
		config:     ftq.config,
		limit:      ftq.limit,
		offset:     ftq.offset,
		order:      append([]OrderFunc{}, ftq.order...),
		predicates: append([]predicate.FrontendTemplate{}, ftq.predicates...),
		// clone intermediate query.
		sql:    ftq.sql.Clone(),
		path:   ftq.path,
		unique: ftq.unique,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt uint32 `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.FrontendTemplate.Query().
//		GroupBy(frontendtemplate.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (ftq *FrontendTemplateQuery) GroupBy(field string, fields ...string) *FrontendTemplateGroupBy {
	grbuild := &FrontendTemplateGroupBy{config: ftq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := ftq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return ftq.sqlQuery(ctx), nil
	}
	grbuild.label = frontendtemplate.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt uint32 `json:"created_at,omitempty"`
//	}
//
//	client.FrontendTemplate.Query().
//		Select(frontendtemplate.FieldCreatedAt).
//		Scan(ctx, &v)
//
func (ftq *FrontendTemplateQuery) Select(fields ...string) *FrontendTemplateSelect {
	ftq.fields = append(ftq.fields, fields...)
	selbuild := &FrontendTemplateSelect{FrontendTemplateQuery: ftq}
	selbuild.label = frontendtemplate.Label
	selbuild.flds, selbuild.scan = &ftq.fields, selbuild.Scan
	return selbuild
}

func (ftq *FrontendTemplateQuery) prepareQuery(ctx context.Context) error {
	for _, f := range ftq.fields {
		if !frontendtemplate.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ftq.path != nil {
		prev, err := ftq.path(ctx)
		if err != nil {
			return err
		}
		ftq.sql = prev
	}
	if frontendtemplate.Policy == nil {
		return errors.New("ent: uninitialized frontendtemplate.Policy (forgotten import ent/runtime?)")
	}
	if err := frontendtemplate.Policy.EvalQuery(ctx, ftq); err != nil {
		return err
	}
	return nil
}

func (ftq *FrontendTemplateQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*FrontendTemplate, error) {
	var (
		nodes = []*FrontendTemplate{}
		_spec = ftq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*FrontendTemplate).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &FrontendTemplate{config: ftq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(ftq.modifiers) > 0 {
		_spec.Modifiers = ftq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ftq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (ftq *FrontendTemplateQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ftq.querySpec()
	if len(ftq.modifiers) > 0 {
		_spec.Modifiers = ftq.modifiers
	}
	_spec.Node.Columns = ftq.fields
	if len(ftq.fields) > 0 {
		_spec.Unique = ftq.unique != nil && *ftq.unique
	}
	return sqlgraph.CountNodes(ctx, ftq.driver, _spec)
}

func (ftq *FrontendTemplateQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := ftq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (ftq *FrontendTemplateQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   frontendtemplate.Table,
			Columns: frontendtemplate.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: frontendtemplate.FieldID,
			},
		},
		From:   ftq.sql,
		Unique: true,
	}
	if unique := ftq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := ftq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, frontendtemplate.FieldID)
		for i := range fields {
			if fields[i] != frontendtemplate.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ftq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ftq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ftq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ftq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ftq *FrontendTemplateQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ftq.driver.Dialect())
	t1 := builder.Table(frontendtemplate.Table)
	columns := ftq.fields
	if len(columns) == 0 {
		columns = frontendtemplate.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ftq.sql != nil {
		selector = ftq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ftq.unique != nil && *ftq.unique {
		selector.Distinct()
	}
	for _, m := range ftq.modifiers {
		m(selector)
	}
	for _, p := range ftq.predicates {
		p(selector)
	}
	for _, p := range ftq.order {
		p(selector)
	}
	if offset := ftq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ftq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (ftq *FrontendTemplateQuery) ForUpdate(opts ...sql.LockOption) *FrontendTemplateQuery {
	if ftq.driver.Dialect() == dialect.Postgres {
		ftq.Unique(false)
	}
	ftq.modifiers = append(ftq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return ftq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (ftq *FrontendTemplateQuery) ForShare(opts ...sql.LockOption) *FrontendTemplateQuery {
	if ftq.driver.Dialect() == dialect.Postgres {
		ftq.Unique(false)
	}
	ftq.modifiers = append(ftq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return ftq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (ftq *FrontendTemplateQuery) Modify(modifiers ...func(s *sql.Selector)) *FrontendTemplateSelect {
	ftq.modifiers = append(ftq.modifiers, modifiers...)
	return ftq.Select()
}

// FrontendTemplateGroupBy is the group-by builder for FrontendTemplate entities.
type FrontendTemplateGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ftgb *FrontendTemplateGroupBy) Aggregate(fns ...AggregateFunc) *FrontendTemplateGroupBy {
	ftgb.fns = append(ftgb.fns, fns...)
	return ftgb
}

// Scan applies the group-by query and scans the result into the given value.
func (ftgb *FrontendTemplateGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := ftgb.path(ctx)
	if err != nil {
		return err
	}
	ftgb.sql = query
	return ftgb.sqlScan(ctx, v)
}

func (ftgb *FrontendTemplateGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range ftgb.fields {
		if !frontendtemplate.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := ftgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ftgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ftgb *FrontendTemplateGroupBy) sqlQuery() *sql.Selector {
	selector := ftgb.sql.Select()
	aggregation := make([]string, 0, len(ftgb.fns))
	for _, fn := range ftgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(ftgb.fields)+len(ftgb.fns))
		for _, f := range ftgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(ftgb.fields...)...)
}

// FrontendTemplateSelect is the builder for selecting fields of FrontendTemplate entities.
type FrontendTemplateSelect struct {
	*FrontendTemplateQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (fts *FrontendTemplateSelect) Scan(ctx context.Context, v interface{}) error {
	if err := fts.prepareQuery(ctx); err != nil {
		return err
	}
	fts.sql = fts.FrontendTemplateQuery.sqlQuery(ctx)
	return fts.sqlScan(ctx, v)
}

func (fts *FrontendTemplateSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := fts.sql.Query()
	if err := fts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (fts *FrontendTemplateSelect) Modify(modifiers ...func(s *sql.Selector)) *FrontendTemplateSelect {
	fts.modifiers = append(fts.modifiers, modifiers...)
	return fts
}
