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
	"tribemedia.io/m/ent/authconfig"
	"tribemedia.io/m/ent/predicate"
	"tribemedia.io/m/ent/project"
)

// AuthConfigQuery is the builder for querying AuthConfig entities.
type AuthConfigQuery struct {
	config
	ctx         *QueryContext
	order       []authconfig.OrderOption
	inters      []Interceptor
	predicates  []predicate.AuthConfig
	withProject *ProjectQuery
	withFKs     bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AuthConfigQuery builder.
func (acq *AuthConfigQuery) Where(ps ...predicate.AuthConfig) *AuthConfigQuery {
	acq.predicates = append(acq.predicates, ps...)
	return acq
}

// Limit the number of records to be returned by this query.
func (acq *AuthConfigQuery) Limit(limit int) *AuthConfigQuery {
	acq.ctx.Limit = &limit
	return acq
}

// Offset to start from.
func (acq *AuthConfigQuery) Offset(offset int) *AuthConfigQuery {
	acq.ctx.Offset = &offset
	return acq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (acq *AuthConfigQuery) Unique(unique bool) *AuthConfigQuery {
	acq.ctx.Unique = &unique
	return acq
}

// Order specifies how the records should be ordered.
func (acq *AuthConfigQuery) Order(o ...authconfig.OrderOption) *AuthConfigQuery {
	acq.order = append(acq.order, o...)
	return acq
}

// QueryProject chains the current query on the "project" edge.
func (acq *AuthConfigQuery) QueryProject() *ProjectQuery {
	query := (&ProjectClient{config: acq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := acq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := acq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(authconfig.Table, authconfig.FieldID, selector),
			sqlgraph.To(project.Table, project.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, authconfig.ProjectTable, authconfig.ProjectColumn),
		)
		fromU = sqlgraph.SetNeighbors(acq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first AuthConfig entity from the query.
// Returns a *NotFoundError when no AuthConfig was found.
func (acq *AuthConfigQuery) First(ctx context.Context) (*AuthConfig, error) {
	nodes, err := acq.Limit(1).All(setContextOp(ctx, acq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{authconfig.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (acq *AuthConfigQuery) FirstX(ctx context.Context) *AuthConfig {
	node, err := acq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first AuthConfig ID from the query.
// Returns a *NotFoundError when no AuthConfig ID was found.
func (acq *AuthConfigQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = acq.Limit(1).IDs(setContextOp(ctx, acq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{authconfig.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (acq *AuthConfigQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := acq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single AuthConfig entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one AuthConfig entity is found.
// Returns a *NotFoundError when no AuthConfig entities are found.
func (acq *AuthConfigQuery) Only(ctx context.Context) (*AuthConfig, error) {
	nodes, err := acq.Limit(2).All(setContextOp(ctx, acq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{authconfig.Label}
	default:
		return nil, &NotSingularError{authconfig.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (acq *AuthConfigQuery) OnlyX(ctx context.Context) *AuthConfig {
	node, err := acq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only AuthConfig ID in the query.
// Returns a *NotSingularError when more than one AuthConfig ID is found.
// Returns a *NotFoundError when no entities are found.
func (acq *AuthConfigQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = acq.Limit(2).IDs(setContextOp(ctx, acq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{authconfig.Label}
	default:
		err = &NotSingularError{authconfig.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (acq *AuthConfigQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := acq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of AuthConfigs.
func (acq *AuthConfigQuery) All(ctx context.Context) ([]*AuthConfig, error) {
	ctx = setContextOp(ctx, acq.ctx, "All")
	if err := acq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*AuthConfig, *AuthConfigQuery]()
	return withInterceptors[[]*AuthConfig](ctx, acq, qr, acq.inters)
}

// AllX is like All, but panics if an error occurs.
func (acq *AuthConfigQuery) AllX(ctx context.Context) []*AuthConfig {
	nodes, err := acq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of AuthConfig IDs.
func (acq *AuthConfigQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if acq.ctx.Unique == nil && acq.path != nil {
		acq.Unique(true)
	}
	ctx = setContextOp(ctx, acq.ctx, "IDs")
	if err = acq.Select(authconfig.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (acq *AuthConfigQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := acq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (acq *AuthConfigQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, acq.ctx, "Count")
	if err := acq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, acq, querierCount[*AuthConfigQuery](), acq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (acq *AuthConfigQuery) CountX(ctx context.Context) int {
	count, err := acq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (acq *AuthConfigQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, acq.ctx, "Exist")
	switch _, err := acq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (acq *AuthConfigQuery) ExistX(ctx context.Context) bool {
	exist, err := acq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AuthConfigQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (acq *AuthConfigQuery) Clone() *AuthConfigQuery {
	if acq == nil {
		return nil
	}
	return &AuthConfigQuery{
		config:      acq.config,
		ctx:         acq.ctx.Clone(),
		order:       append([]authconfig.OrderOption{}, acq.order...),
		inters:      append([]Interceptor{}, acq.inters...),
		predicates:  append([]predicate.AuthConfig{}, acq.predicates...),
		withProject: acq.withProject.Clone(),
		// clone intermediate query.
		sql:  acq.sql.Clone(),
		path: acq.path,
	}
}

// WithProject tells the query-builder to eager-load the nodes that are connected to
// the "project" edge. The optional arguments are used to configure the query builder of the edge.
func (acq *AuthConfigQuery) WithProject(opts ...func(*ProjectQuery)) *AuthConfigQuery {
	query := (&ProjectClient{config: acq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	acq.withProject = query
	return acq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		DisableSignup bool `json:"disable_signup,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.AuthConfig.Query().
//		GroupBy(authconfig.FieldDisableSignup).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (acq *AuthConfigQuery) GroupBy(field string, fields ...string) *AuthConfigGroupBy {
	acq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &AuthConfigGroupBy{build: acq}
	grbuild.flds = &acq.ctx.Fields
	grbuild.label = authconfig.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		DisableSignup bool `json:"disable_signup,omitempty"`
//	}
//
//	client.AuthConfig.Query().
//		Select(authconfig.FieldDisableSignup).
//		Scan(ctx, &v)
func (acq *AuthConfigQuery) Select(fields ...string) *AuthConfigSelect {
	acq.ctx.Fields = append(acq.ctx.Fields, fields...)
	sbuild := &AuthConfigSelect{AuthConfigQuery: acq}
	sbuild.label = authconfig.Label
	sbuild.flds, sbuild.scan = &acq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a AuthConfigSelect configured with the given aggregations.
func (acq *AuthConfigQuery) Aggregate(fns ...AggregateFunc) *AuthConfigSelect {
	return acq.Select().Aggregate(fns...)
}

func (acq *AuthConfigQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range acq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, acq); err != nil {
				return err
			}
		}
	}
	for _, f := range acq.ctx.Fields {
		if !authconfig.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if acq.path != nil {
		prev, err := acq.path(ctx)
		if err != nil {
			return err
		}
		acq.sql = prev
	}
	return nil
}

func (acq *AuthConfigQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*AuthConfig, error) {
	var (
		nodes       = []*AuthConfig{}
		withFKs     = acq.withFKs
		_spec       = acq.querySpec()
		loadedTypes = [1]bool{
			acq.withProject != nil,
		}
	)
	if acq.withProject != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, authconfig.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*AuthConfig).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &AuthConfig{config: acq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, acq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := acq.withProject; query != nil {
		if err := acq.loadProject(ctx, query, nodes, nil,
			func(n *AuthConfig, e *Project) { n.Edges.Project = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (acq *AuthConfigQuery) loadProject(ctx context.Context, query *ProjectQuery, nodes []*AuthConfig, init func(*AuthConfig), assign func(*AuthConfig, *Project)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*AuthConfig)
	for i := range nodes {
		if nodes[i].project_auth_config == nil {
			continue
		}
		fk := *nodes[i].project_auth_config
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
			return fmt.Errorf(`unexpected foreign-key "project_auth_config" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (acq *AuthConfigQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := acq.querySpec()
	_spec.Node.Columns = acq.ctx.Fields
	if len(acq.ctx.Fields) > 0 {
		_spec.Unique = acq.ctx.Unique != nil && *acq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, acq.driver, _spec)
}

func (acq *AuthConfigQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(authconfig.Table, authconfig.Columns, sqlgraph.NewFieldSpec(authconfig.FieldID, field.TypeUUID))
	_spec.From = acq.sql
	if unique := acq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if acq.path != nil {
		_spec.Unique = true
	}
	if fields := acq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, authconfig.FieldID)
		for i := range fields {
			if fields[i] != authconfig.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := acq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := acq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := acq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := acq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (acq *AuthConfigQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(acq.driver.Dialect())
	t1 := builder.Table(authconfig.Table)
	columns := acq.ctx.Fields
	if len(columns) == 0 {
		columns = authconfig.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if acq.sql != nil {
		selector = acq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if acq.ctx.Unique != nil && *acq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range acq.predicates {
		p(selector)
	}
	for _, p := range acq.order {
		p(selector)
	}
	if offset := acq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := acq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// AuthConfigGroupBy is the group-by builder for AuthConfig entities.
type AuthConfigGroupBy struct {
	selector
	build *AuthConfigQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (acgb *AuthConfigGroupBy) Aggregate(fns ...AggregateFunc) *AuthConfigGroupBy {
	acgb.fns = append(acgb.fns, fns...)
	return acgb
}

// Scan applies the selector query and scans the result into the given value.
func (acgb *AuthConfigGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, acgb.build.ctx, "GroupBy")
	if err := acgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AuthConfigQuery, *AuthConfigGroupBy](ctx, acgb.build, acgb, acgb.build.inters, v)
}

func (acgb *AuthConfigGroupBy) sqlScan(ctx context.Context, root *AuthConfigQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(acgb.fns))
	for _, fn := range acgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*acgb.flds)+len(acgb.fns))
		for _, f := range *acgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*acgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := acgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// AuthConfigSelect is the builder for selecting fields of AuthConfig entities.
type AuthConfigSelect struct {
	*AuthConfigQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (acs *AuthConfigSelect) Aggregate(fns ...AggregateFunc) *AuthConfigSelect {
	acs.fns = append(acs.fns, fns...)
	return acs
}

// Scan applies the selector query and scans the result into the given value.
func (acs *AuthConfigSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, acs.ctx, "Select")
	if err := acs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AuthConfigQuery, *AuthConfigSelect](ctx, acs.AuthConfigQuery, acs, acs.inters, v)
}

func (acs *AuthConfigSelect) sqlScan(ctx context.Context, root *AuthConfigQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(acs.fns))
	for _, fn := range acs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*acs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := acs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}