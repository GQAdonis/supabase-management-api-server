// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"tribemedia.io/m/ent/function"
	"tribemedia.io/m/ent/predicate"
	"tribemedia.io/m/ent/project"
)

// FunctionQuery is the builder for querying Function entities.
type FunctionQuery struct {
	config
	ctx         *QueryContext
	order       []function.OrderOption
	inters      []Interceptor
	predicates  []predicate.Function
	withProject *ProjectQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the FunctionQuery builder.
func (fq *FunctionQuery) Where(ps ...predicate.Function) *FunctionQuery {
	fq.predicates = append(fq.predicates, ps...)
	return fq
}

// Limit the number of records to be returned by this query.
func (fq *FunctionQuery) Limit(limit int) *FunctionQuery {
	fq.ctx.Limit = &limit
	return fq
}

// Offset to start from.
func (fq *FunctionQuery) Offset(offset int) *FunctionQuery {
	fq.ctx.Offset = &offset
	return fq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (fq *FunctionQuery) Unique(unique bool) *FunctionQuery {
	fq.ctx.Unique = &unique
	return fq
}

// Order specifies how the records should be ordered.
func (fq *FunctionQuery) Order(o ...function.OrderOption) *FunctionQuery {
	fq.order = append(fq.order, o...)
	return fq
}

// QueryProject chains the current query on the "project" edge.
func (fq *FunctionQuery) QueryProject() *ProjectQuery {
	query := (&ProjectClient{config: fq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := fq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := fq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(function.Table, function.FieldID, selector),
			sqlgraph.To(project.Table, project.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, function.ProjectTable, function.ProjectColumn),
		)
		fromU = sqlgraph.SetNeighbors(fq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Function entity from the query.
// Returns a *NotFoundError when no Function was found.
func (fq *FunctionQuery) First(ctx context.Context) (*Function, error) {
	nodes, err := fq.Limit(1).All(setContextOp(ctx, fq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{function.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (fq *FunctionQuery) FirstX(ctx context.Context) *Function {
	node, err := fq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Function ID from the query.
// Returns a *NotFoundError when no Function ID was found.
func (fq *FunctionQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = fq.Limit(1).IDs(setContextOp(ctx, fq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{function.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (fq *FunctionQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := fq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Function entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Function entity is found.
// Returns a *NotFoundError when no Function entities are found.
func (fq *FunctionQuery) Only(ctx context.Context) (*Function, error) {
	nodes, err := fq.Limit(2).All(setContextOp(ctx, fq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{function.Label}
	default:
		return nil, &NotSingularError{function.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (fq *FunctionQuery) OnlyX(ctx context.Context) *Function {
	node, err := fq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Function ID in the query.
// Returns a *NotSingularError when more than one Function ID is found.
// Returns a *NotFoundError when no entities are found.
func (fq *FunctionQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = fq.Limit(2).IDs(setContextOp(ctx, fq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{function.Label}
	default:
		err = &NotSingularError{function.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (fq *FunctionQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := fq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Functions.
func (fq *FunctionQuery) All(ctx context.Context) ([]*Function, error) {
	ctx = setContextOp(ctx, fq.ctx, "All")
	if err := fq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Function, *FunctionQuery]()
	return withInterceptors[[]*Function](ctx, fq, qr, fq.inters)
}

// AllX is like All, but panics if an error occurs.
func (fq *FunctionQuery) AllX(ctx context.Context) []*Function {
	nodes, err := fq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Function IDs.
func (fq *FunctionQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if fq.ctx.Unique == nil && fq.path != nil {
		fq.Unique(true)
	}
	ctx = setContextOp(ctx, fq.ctx, "IDs")
	if err = fq.Select(function.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (fq *FunctionQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := fq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (fq *FunctionQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, fq.ctx, "Count")
	if err := fq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, fq, querierCount[*FunctionQuery](), fq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (fq *FunctionQuery) CountX(ctx context.Context) int {
	count, err := fq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (fq *FunctionQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, fq.ctx, "Exist")
	switch _, err := fq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (fq *FunctionQuery) ExistX(ctx context.Context) bool {
	exist, err := fq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the FunctionQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (fq *FunctionQuery) Clone() *FunctionQuery {
	if fq == nil {
		return nil
	}
	return &FunctionQuery{
		config:      fq.config,
		ctx:         fq.ctx.Clone(),
		order:       append([]function.OrderOption{}, fq.order...),
		inters:      append([]Interceptor{}, fq.inters...),
		predicates:  append([]predicate.Function{}, fq.predicates...),
		withProject: fq.withProject.Clone(),
		// clone intermediate query.
		sql:  fq.sql.Clone(),
		path: fq.path,
	}
}

// WithProject tells the query-builder to eager-load the nodes that are connected to
// the "project" edge. The optional arguments are used to configure the query builder of the edge.
func (fq *FunctionQuery) WithProject(opts ...func(*ProjectQuery)) *FunctionQuery {
	query := (&ProjectClient{config: fq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	fq.withProject = query
	return fq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		ProjectID uuid.UUID `json:"project_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Function.Query().
//		GroupBy(function.FieldProjectID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (fq *FunctionQuery) GroupBy(field string, fields ...string) *FunctionGroupBy {
	fq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &FunctionGroupBy{build: fq}
	grbuild.flds = &fq.ctx.Fields
	grbuild.label = function.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		ProjectID uuid.UUID `json:"project_id,omitempty"`
//	}
//
//	client.Function.Query().
//		Select(function.FieldProjectID).
//		Scan(ctx, &v)
func (fq *FunctionQuery) Select(fields ...string) *FunctionSelect {
	fq.ctx.Fields = append(fq.ctx.Fields, fields...)
	sbuild := &FunctionSelect{FunctionQuery: fq}
	sbuild.label = function.Label
	sbuild.flds, sbuild.scan = &fq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a FunctionSelect configured with the given aggregations.
func (fq *FunctionQuery) Aggregate(fns ...AggregateFunc) *FunctionSelect {
	return fq.Select().Aggregate(fns...)
}

func (fq *FunctionQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range fq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, fq); err != nil {
				return err
			}
		}
	}
	for _, f := range fq.ctx.Fields {
		if !function.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if fq.path != nil {
		prev, err := fq.path(ctx)
		if err != nil {
			return err
		}
		fq.sql = prev
	}
	return nil
}

func (fq *FunctionQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Function, error) {
	var (
		nodes       = []*Function{}
		_spec       = fq.querySpec()
		loadedTypes = [1]bool{
			fq.withProject != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Function).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Function{config: fq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, fq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := fq.withProject; query != nil {
		if err := fq.loadProject(ctx, query, nodes, nil,
			func(n *Function, e *Project) { n.Edges.Project = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (fq *FunctionQuery) loadProject(ctx context.Context, query *ProjectQuery, nodes []*Function, init func(*Function), assign func(*Function, *Project)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Function)
	for i := range nodes {
		fk := nodes[i].ProjectID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(project.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "project_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (fq *FunctionQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := fq.querySpec()
	_spec.Node.Columns = fq.ctx.Fields
	if len(fq.ctx.Fields) > 0 {
		_spec.Unique = fq.ctx.Unique != nil && *fq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, fq.driver, _spec)
}

func (fq *FunctionQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(function.Table, function.Columns, sqlgraph.NewFieldSpec(function.FieldID, field.TypeUUID))
	_spec.From = fq.sql
	if unique := fq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if fq.path != nil {
		_spec.Unique = true
	}
	if fields := fq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, function.FieldID)
		for i := range fields {
			if fields[i] != function.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if fq.withProject != nil {
			_spec.Node.AddColumnOnce(function.FieldProjectID)
		}
	}
	if ps := fq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := fq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := fq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := fq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (fq *FunctionQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(fq.driver.Dialect())
	t1 := builder.Table(function.Table)
	columns := fq.ctx.Fields
	if len(columns) == 0 {
		columns = function.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if fq.sql != nil {
		selector = fq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if fq.ctx.Unique != nil && *fq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range fq.predicates {
		p(selector)
	}
	for _, p := range fq.order {
		p(selector)
	}
	if offset := fq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := fq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// FunctionGroupBy is the group-by builder for Function entities.
type FunctionGroupBy struct {
	selector
	build *FunctionQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (fgb *FunctionGroupBy) Aggregate(fns ...AggregateFunc) *FunctionGroupBy {
	fgb.fns = append(fgb.fns, fns...)
	return fgb
}

// Scan applies the selector query and scans the result into the given value.
func (fgb *FunctionGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, fgb.build.ctx, "GroupBy")
	if err := fgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*FunctionQuery, *FunctionGroupBy](ctx, fgb.build, fgb, fgb.build.inters, v)
}

func (fgb *FunctionGroupBy) sqlScan(ctx context.Context, root *FunctionQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(fgb.fns))
	for _, fn := range fgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*fgb.flds)+len(fgb.fns))
		for _, f := range *fgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*fgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := fgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// FunctionSelect is the builder for selecting fields of Function entities.
type FunctionSelect struct {
	*FunctionQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (fs *FunctionSelect) Aggregate(fns ...AggregateFunc) *FunctionSelect {
	fs.fns = append(fs.fns, fns...)
	return fs
}

// Scan applies the selector query and scans the result into the given value.
func (fs *FunctionSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, fs.ctx, "Select")
	if err := fs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*FunctionQuery, *FunctionSelect](ctx, fs.FunctionQuery, fs, fs.inters, v)
}

func (fs *FunctionSelect) sqlScan(ctx context.Context, root *FunctionQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(fs.fns))
	for _, fn := range fs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*fs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := fs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
