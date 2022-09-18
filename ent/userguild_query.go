// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/lrstanley/ent-bug-edge-migration-conflict/ent/predicate"
	"github.com/lrstanley/ent-bug-edge-migration-conflict/ent/user"
	"github.com/lrstanley/ent-bug-edge-migration-conflict/ent/userguild"
)

// UserGuildQuery is the builder for querying UserGuild entities.
type UserGuildQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.UserGuild
	withAdmins *UserQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the UserGuildQuery builder.
func (ugq *UserGuildQuery) Where(ps ...predicate.UserGuild) *UserGuildQuery {
	ugq.predicates = append(ugq.predicates, ps...)
	return ugq
}

// Limit adds a limit step to the query.
func (ugq *UserGuildQuery) Limit(limit int) *UserGuildQuery {
	ugq.limit = &limit
	return ugq
}

// Offset adds an offset step to the query.
func (ugq *UserGuildQuery) Offset(offset int) *UserGuildQuery {
	ugq.offset = &offset
	return ugq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ugq *UserGuildQuery) Unique(unique bool) *UserGuildQuery {
	ugq.unique = &unique
	return ugq
}

// Order adds an order step to the query.
func (ugq *UserGuildQuery) Order(o ...OrderFunc) *UserGuildQuery {
	ugq.order = append(ugq.order, o...)
	return ugq
}

// QueryAdmins chains the current query on the "admins" edge.
func (ugq *UserGuildQuery) QueryAdmins() *UserQuery {
	query := &UserQuery{config: ugq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ugq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ugq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(userguild.Table, userguild.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, userguild.AdminsTable, userguild.AdminsPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(ugq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first UserGuild entity from the query.
// Returns a *NotFoundError when no UserGuild was found.
func (ugq *UserGuildQuery) First(ctx context.Context) (*UserGuild, error) {
	nodes, err := ugq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{userguild.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ugq *UserGuildQuery) FirstX(ctx context.Context) *UserGuild {
	node, err := ugq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first UserGuild ID from the query.
// Returns a *NotFoundError when no UserGuild ID was found.
func (ugq *UserGuildQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ugq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{userguild.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ugq *UserGuildQuery) FirstIDX(ctx context.Context) int {
	id, err := ugq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single UserGuild entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one UserGuild entity is found.
// Returns a *NotFoundError when no UserGuild entities are found.
func (ugq *UserGuildQuery) Only(ctx context.Context) (*UserGuild, error) {
	nodes, err := ugq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{userguild.Label}
	default:
		return nil, &NotSingularError{userguild.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ugq *UserGuildQuery) OnlyX(ctx context.Context) *UserGuild {
	node, err := ugq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only UserGuild ID in the query.
// Returns a *NotSingularError when more than one UserGuild ID is found.
// Returns a *NotFoundError when no entities are found.
func (ugq *UserGuildQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ugq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{userguild.Label}
	default:
		err = &NotSingularError{userguild.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ugq *UserGuildQuery) OnlyIDX(ctx context.Context) int {
	id, err := ugq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of UserGuilds.
func (ugq *UserGuildQuery) All(ctx context.Context) ([]*UserGuild, error) {
	if err := ugq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return ugq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (ugq *UserGuildQuery) AllX(ctx context.Context) []*UserGuild {
	nodes, err := ugq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of UserGuild IDs.
func (ugq *UserGuildQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := ugq.Select(userguild.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ugq *UserGuildQuery) IDsX(ctx context.Context) []int {
	ids, err := ugq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ugq *UserGuildQuery) Count(ctx context.Context) (int, error) {
	if err := ugq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return ugq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (ugq *UserGuildQuery) CountX(ctx context.Context) int {
	count, err := ugq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ugq *UserGuildQuery) Exist(ctx context.Context) (bool, error) {
	if err := ugq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return ugq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (ugq *UserGuildQuery) ExistX(ctx context.Context) bool {
	exist, err := ugq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the UserGuildQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ugq *UserGuildQuery) Clone() *UserGuildQuery {
	if ugq == nil {
		return nil
	}
	return &UserGuildQuery{
		config:     ugq.config,
		limit:      ugq.limit,
		offset:     ugq.offset,
		order:      append([]OrderFunc{}, ugq.order...),
		predicates: append([]predicate.UserGuild{}, ugq.predicates...),
		withAdmins: ugq.withAdmins.Clone(),
		// clone intermediate query.
		sql:    ugq.sql.Clone(),
		path:   ugq.path,
		unique: ugq.unique,
	}
}

// WithAdmins tells the query-builder to eager-load the nodes that are connected to
// the "admins" edge. The optional arguments are used to configure the query builder of the edge.
func (ugq *UserGuildQuery) WithAdmins(opts ...func(*UserQuery)) *UserGuildQuery {
	query := &UserQuery{config: ugq.config}
	for _, opt := range opts {
		opt(query)
	}
	ugq.withAdmins = query
	return ugq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.UserGuild.Query().
//		GroupBy(userguild.FieldCreateTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (ugq *UserGuildQuery) GroupBy(field string, fields ...string) *UserGuildGroupBy {
	grbuild := &UserGuildGroupBy{config: ugq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := ugq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return ugq.sqlQuery(ctx), nil
	}
	grbuild.label = userguild.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time,omitempty"`
//	}
//
//	client.UserGuild.Query().
//		Select(userguild.FieldCreateTime).
//		Scan(ctx, &v)
func (ugq *UserGuildQuery) Select(fields ...string) *UserGuildSelect {
	ugq.fields = append(ugq.fields, fields...)
	selbuild := &UserGuildSelect{UserGuildQuery: ugq}
	selbuild.label = userguild.Label
	selbuild.flds, selbuild.scan = &ugq.fields, selbuild.Scan
	return selbuild
}

func (ugq *UserGuildQuery) prepareQuery(ctx context.Context) error {
	for _, f := range ugq.fields {
		if !userguild.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ugq.path != nil {
		prev, err := ugq.path(ctx)
		if err != nil {
			return err
		}
		ugq.sql = prev
	}
	return nil
}

func (ugq *UserGuildQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*UserGuild, error) {
	var (
		nodes       = []*UserGuild{}
		_spec       = ugq.querySpec()
		loadedTypes = [1]bool{
			ugq.withAdmins != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*UserGuild).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &UserGuild{config: ugq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ugq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := ugq.withAdmins; query != nil {
		if err := ugq.loadAdmins(ctx, query, nodes,
			func(n *UserGuild) { n.Edges.Admins = []*User{} },
			func(n *UserGuild, e *User) { n.Edges.Admins = append(n.Edges.Admins, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (ugq *UserGuildQuery) loadAdmins(ctx context.Context, query *UserQuery, nodes []*UserGuild, init func(*UserGuild), assign func(*UserGuild, *User)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*UserGuild)
	nids := make(map[int]map[*UserGuild]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(userguild.AdminsTable)
		s.Join(joinT).On(s.C(user.FieldID), joinT.C(userguild.AdminsPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(userguild.AdminsPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(userguild.AdminsPrimaryKey[1]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	neighbors, err := query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
		assign := spec.Assign
		values := spec.ScanValues
		spec.ScanValues = func(columns []string) ([]any, error) {
			values, err := values(columns[1:])
			if err != nil {
				return nil, err
			}
			return append([]any{new(sql.NullInt64)}, values...), nil
		}
		spec.Assign = func(columns []string, values []any) error {
			outValue := int(values[0].(*sql.NullInt64).Int64)
			inValue := int(values[1].(*sql.NullInt64).Int64)
			if nids[inValue] == nil {
				nids[inValue] = map[*UserGuild]struct{}{byID[outValue]: struct{}{}}
				return assign(columns[1:], values[1:])
			}
			nids[inValue][byID[outValue]] = struct{}{}
			return nil
		}
	})
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "admins" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}

func (ugq *UserGuildQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ugq.querySpec()
	_spec.Node.Columns = ugq.fields
	if len(ugq.fields) > 0 {
		_spec.Unique = ugq.unique != nil && *ugq.unique
	}
	return sqlgraph.CountNodes(ctx, ugq.driver, _spec)
}

func (ugq *UserGuildQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := ugq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

func (ugq *UserGuildQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   userguild.Table,
			Columns: userguild.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: userguild.FieldID,
			},
		},
		From:   ugq.sql,
		Unique: true,
	}
	if unique := ugq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := ugq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, userguild.FieldID)
		for i := range fields {
			if fields[i] != userguild.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ugq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ugq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ugq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ugq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ugq *UserGuildQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ugq.driver.Dialect())
	t1 := builder.Table(userguild.Table)
	columns := ugq.fields
	if len(columns) == 0 {
		columns = userguild.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ugq.sql != nil {
		selector = ugq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ugq.unique != nil && *ugq.unique {
		selector.Distinct()
	}
	for _, p := range ugq.predicates {
		p(selector)
	}
	for _, p := range ugq.order {
		p(selector)
	}
	if offset := ugq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ugq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// UserGuildGroupBy is the group-by builder for UserGuild entities.
type UserGuildGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (uggb *UserGuildGroupBy) Aggregate(fns ...AggregateFunc) *UserGuildGroupBy {
	uggb.fns = append(uggb.fns, fns...)
	return uggb
}

// Scan applies the group-by query and scans the result into the given value.
func (uggb *UserGuildGroupBy) Scan(ctx context.Context, v any) error {
	query, err := uggb.path(ctx)
	if err != nil {
		return err
	}
	uggb.sql = query
	return uggb.sqlScan(ctx, v)
}

func (uggb *UserGuildGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range uggb.fields {
		if !userguild.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := uggb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := uggb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (uggb *UserGuildGroupBy) sqlQuery() *sql.Selector {
	selector := uggb.sql.Select()
	aggregation := make([]string, 0, len(uggb.fns))
	for _, fn := range uggb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(uggb.fields)+len(uggb.fns))
		for _, f := range uggb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(uggb.fields...)...)
}

// UserGuildSelect is the builder for selecting fields of UserGuild entities.
type UserGuildSelect struct {
	*UserGuildQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ugs *UserGuildSelect) Scan(ctx context.Context, v any) error {
	if err := ugs.prepareQuery(ctx); err != nil {
		return err
	}
	ugs.sql = ugs.UserGuildQuery.sqlQuery(ctx)
	return ugs.sqlScan(ctx, v)
}

func (ugs *UserGuildSelect) sqlScan(ctx context.Context, v any) error {
	rows := &sql.Rows{}
	query, args := ugs.sql.Query()
	if err := ugs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
