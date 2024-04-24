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
	"tribemedia.io/m/ent/pgsodiumconfig"
	"tribemedia.io/m/ent/predicate"
	"tribemedia.io/m/ent/project"
)

// PgsodiumConfigQuery is the builder for querying PgsodiumConfig entities.
type PgsodiumConfigQuery struct {
	config
	ctx         *QueryContext
	order       []pgsodiumconfig.OrderOption
	inters      []Interceptor
	predicates  []predicate.PgsodiumConfig
	withProject *ProjectQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the PgsodiumConfigQuery builder.
func (pcq *PgsodiumConfigQuery) Where(ps ...predicate.PgsodiumConfig) *PgsodiumConfigQuery {
	pcq.predicates = append(pcq.predicates, ps...)
	return pcq
}

// Limit the number of records to be returned by this query.
func (pcq *PgsodiumConfigQuery) Limit(limit int) *PgsodiumConfigQuery {
	pcq.ctx.Limit = &limit
	return pcq
}

// Offset to start from.
func (pcq *PgsodiumConfigQuery) Offset(offset int) *PgsodiumConfigQuery {
	pcq.ctx.Offset = &offset
	return pcq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (pcq *PgsodiumConfigQuery) Unique(unique bool) *PgsodiumConfigQuery {
	pcq.ctx.Unique = &unique
	return pcq
}

// Order specifies how the records should be ordered.
func (pcq *PgsodiumConfigQuery) Order(o ...pgsodiumconfig.OrderOption) *PgsodiumConfigQuery {
	pcq.order = append(pcq.order, o...)
	return pcq
}

// QueryProject chains the current query on the "project" edge.
func (pcq *PgsodiumConfigQuery) QueryProject() *ProjectQuery {
	query := (&ProjectClient{config: pcq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pcq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(pgsodiumconfig.Table, pgsodiumconfig.FieldID, selector),
			sqlgraph.To(project.Table, project.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, pgsodiumconfig.ProjectTable, pgsodiumconfig.ProjectColumn),
		)
		fromU = sqlgraph.SetNeighbors(pcq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first PgsodiumConfig entity from the query.
// Returns a *NotFoundError when no PgsodiumConfig was found.
func (pcq *PgsodiumConfigQuery) First(ctx context.Context) (*PgsodiumConfig, error) {
	nodes, err := pcq.Limit(1).All(setContextOp(ctx, pcq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{pgsodiumconfig.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (pcq *PgsodiumConfigQuery) FirstX(ctx context.Context) *PgsodiumConfig {
	node, err := pcq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first PgsodiumConfig ID from the query.
// Returns a *NotFoundError when no PgsodiumConfig ID was found.
func (pcq *PgsodiumConfigQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = pcq.Limit(1).IDs(setContextOp(ctx, pcq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{pgsodiumconfig.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (pcq *PgsodiumConfigQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := pcq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single PgsodiumConfig entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one PgsodiumConfig entity is found.
// Returns a *NotFoundError when no PgsodiumConfig entities are found.
func (pcq *PgsodiumConfigQuery) Only(ctx context.Context) (*PgsodiumConfig, error) {
	nodes, err := pcq.Limit(2).All(setContextOp(ctx, pcq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{pgsodiumconfig.Label}
	default:
		return nil, &NotSingularError{pgsodiumconfig.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (pcq *PgsodiumConfigQuery) OnlyX(ctx context.Context) *PgsodiumConfig {
	node, err := pcq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only PgsodiumConfig ID in the query.
// Returns a *NotSingularError when more than one PgsodiumConfig ID is found.
// Returns a *NotFoundError when no entities are found.
func (pcq *PgsodiumConfigQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = pcq.Limit(2).IDs(setContextOp(ctx, pcq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{pgsodiumconfig.Label}
	default:
		err = &NotSingularError{pgsodiumconfig.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (pcq *PgsodiumConfigQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := pcq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of PgsodiumConfigs.
func (pcq *PgsodiumConfigQuery) All(ctx context.Context) ([]*PgsodiumConfig, error) {
	ctx = setContextOp(ctx, pcq.ctx, "All")
	if err := pcq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*PgsodiumConfig, *PgsodiumConfigQuery]()
	return withInterceptors[[]*PgsodiumConfig](ctx, pcq, qr, pcq.inters)
}

// AllX is like All, but panics if an error occurs.
func (pcq *PgsodiumConfigQuery) AllX(ctx context.Context) []*PgsodiumConfig {
	nodes, err := pcq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of PgsodiumConfig IDs.
func (pcq *PgsodiumConfigQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if pcq.ctx.Unique == nil && pcq.path != nil {
		pcq.Unique(true)
	}
	ctx = setContextOp(ctx, pcq.ctx, "IDs")
	if err = pcq.Select(pgsodiumconfig.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (pcq *PgsodiumConfigQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := pcq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (pcq *PgsodiumConfigQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, pcq.ctx, "Count")
	if err := pcq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, pcq, querierCount[*PgsodiumConfigQuery](), pcq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (pcq *PgsodiumConfigQuery) CountX(ctx context.Context) int {
	count, err := pcq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (pcq *PgsodiumConfigQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, pcq.ctx, "Exist")
	switch _, err := pcq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (pcq *PgsodiumConfigQuery) ExistX(ctx context.Context) bool {
	exist, err := pcq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the PgsodiumConfigQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (pcq *PgsodiumConfigQuery) Clone() *PgsodiumConfigQuery {
	if pcq == nil {
		return nil
	}
	return &PgsodiumConfigQuery{
		config:      pcq.config,
		ctx:         pcq.ctx.Clone(),
		order:       append([]pgsodiumconfig.OrderOption{}, pcq.order...),
		inters:      append([]Interceptor{}, pcq.inters...),
		predicates:  append([]predicate.PgsodiumConfig{}, pcq.predicates...),
		withProject: pcq.withProject.Clone(),
		// clone intermediate query.
		sql:  pcq.sql.Clone(),
		path: pcq.path,
	}
}

// WithProject tells the query-builder to eager-load the nodes that are connected to
// the "project" edge. The optional arguments are used to configure the query builder of the edge.
func (pcq *PgsodiumConfigQuery) WithProject(opts ...func(*ProjectQuery)) *PgsodiumConfigQuery {
	query := (&ProjectClient{config: pcq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pcq.withProject = query
	return pcq
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
//	client.PgsodiumConfig.Query().
//		GroupBy(pgsodiumconfig.FieldProjectID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (pcq *PgsodiumConfigQuery) GroupBy(field string, fields ...string) *PgsodiumConfigGroupBy {
	pcq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &PgsodiumConfigGroupBy{build: pcq}
	grbuild.flds = &pcq.ctx.Fields
	grbuild.label = pgsodiumconfig.Label
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
//	client.PgsodiumConfig.Query().
//		Select(pgsodiumconfig.FieldProjectID).
//		Scan(ctx, &v)
func (pcq *PgsodiumConfigQuery) Select(fields ...string) *PgsodiumConfigSelect {
	pcq.ctx.Fields = append(pcq.ctx.Fields, fields...)
	sbuild := &PgsodiumConfigSelect{PgsodiumConfigQuery: pcq}
	sbuild.label = pgsodiumconfig.Label
	sbuild.flds, sbuild.scan = &pcq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a PgsodiumConfigSelect configured with the given aggregations.
func (pcq *PgsodiumConfigQuery) Aggregate(fns ...AggregateFunc) *PgsodiumConfigSelect {
	return pcq.Select().Aggregate(fns...)
}

func (pcq *PgsodiumConfigQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range pcq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, pcq); err != nil {
				return err
			}
		}
	}
	for _, f := range pcq.ctx.Fields {
		if !pgsodiumconfig.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if pcq.path != nil {
		prev, err := pcq.path(ctx)
		if err != nil {
			return err
		}
		pcq.sql = prev
	}
	return nil
}

func (pcq *PgsodiumConfigQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*PgsodiumConfig, error) {
	var (
		nodes       = []*PgsodiumConfig{}
		_spec       = pcq.querySpec()
		loadedTypes = [1]bool{
			pcq.withProject != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*PgsodiumConfig).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &PgsodiumConfig{config: pcq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, pcq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := pcq.withProject; query != nil {
		if err := pcq.loadProject(ctx, query, nodes, nil,
			func(n *PgsodiumConfig, e *Project) { n.Edges.Project = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (pcq *PgsodiumConfigQuery) loadProject(ctx context.Context, query *ProjectQuery, nodes []*PgsodiumConfig, init func(*PgsodiumConfig), assign func(*PgsodiumConfig, *Project)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*PgsodiumConfig)
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

func (pcq *PgsodiumConfigQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := pcq.querySpec()
	_spec.Node.Columns = pcq.ctx.Fields
	if len(pcq.ctx.Fields) > 0 {
		_spec.Unique = pcq.ctx.Unique != nil && *pcq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, pcq.driver, _spec)
}

func (pcq *PgsodiumConfigQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(pgsodiumconfig.Table, pgsodiumconfig.Columns, sqlgraph.NewFieldSpec(pgsodiumconfig.FieldID, field.TypeUUID))
	_spec.From = pcq.sql
	if unique := pcq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if pcq.path != nil {
		_spec.Unique = true
	}
	if fields := pcq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, pgsodiumconfig.FieldID)
		for i := range fields {
			if fields[i] != pgsodiumconfig.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if pcq.withProject != nil {
			_spec.Node.AddColumnOnce(pgsodiumconfig.FieldProjectID)
		}
	}
	if ps := pcq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := pcq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := pcq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := pcq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (pcq *PgsodiumConfigQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(pcq.driver.Dialect())
	t1 := builder.Table(pgsodiumconfig.Table)
	columns := pcq.ctx.Fields
	if len(columns) == 0 {
		columns = pgsodiumconfig.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if pcq.sql != nil {
		selector = pcq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if pcq.ctx.Unique != nil && *pcq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range pcq.predicates {
		p(selector)
	}
	for _, p := range pcq.order {
		p(selector)
	}
	if offset := pcq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := pcq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// PgsodiumConfigGroupBy is the group-by builder for PgsodiumConfig entities.
type PgsodiumConfigGroupBy struct {
	selector
	build *PgsodiumConfigQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pcgb *PgsodiumConfigGroupBy) Aggregate(fns ...AggregateFunc) *PgsodiumConfigGroupBy {
	pcgb.fns = append(pcgb.fns, fns...)
	return pcgb
}

// Scan applies the selector query and scans the result into the given value.
func (pcgb *PgsodiumConfigGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pcgb.build.ctx, "GroupBy")
	if err := pcgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PgsodiumConfigQuery, *PgsodiumConfigGroupBy](ctx, pcgb.build, pcgb, pcgb.build.inters, v)
}

func (pcgb *PgsodiumConfigGroupBy) sqlScan(ctx context.Context, root *PgsodiumConfigQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(pcgb.fns))
	for _, fn := range pcgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*pcgb.flds)+len(pcgb.fns))
		for _, f := range *pcgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*pcgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pcgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// PgsodiumConfigSelect is the builder for selecting fields of PgsodiumConfig entities.
type PgsodiumConfigSelect struct {
	*PgsodiumConfigQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (pcs *PgsodiumConfigSelect) Aggregate(fns ...AggregateFunc) *PgsodiumConfigSelect {
	pcs.fns = append(pcs.fns, fns...)
	return pcs
}

// Scan applies the selector query and scans the result into the given value.
func (pcs *PgsodiumConfigSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pcs.ctx, "Select")
	if err := pcs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PgsodiumConfigQuery, *PgsodiumConfigSelect](ctx, pcs.PgsodiumConfigQuery, pcs, pcs.inters, v)
}

func (pcs *PgsodiumConfigSelect) sqlScan(ctx context.Context, root *PgsodiumConfigQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(pcs.fns))
	for _, fn := range pcs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*pcs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pcs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}