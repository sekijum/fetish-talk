// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"
	"server/infrastructure/ent/board"
	"server/infrastructure/ent/predicate"
	"server/infrastructure/ent/user"
	"server/infrastructure/ent/userboardlike"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// UserBoardLikeQuery is the builder for querying UserBoardLike entities.
type UserBoardLikeQuery struct {
	config
	ctx        *QueryContext
	order      []userboardlike.OrderOption
	inters     []Interceptor
	predicates []predicate.UserBoardLike
	withUser   *UserQuery
	withBoard  *BoardQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the UserBoardLikeQuery builder.
func (ublq *UserBoardLikeQuery) Where(ps ...predicate.UserBoardLike) *UserBoardLikeQuery {
	ublq.predicates = append(ublq.predicates, ps...)
	return ublq
}

// Limit the number of records to be returned by this query.
func (ublq *UserBoardLikeQuery) Limit(limit int) *UserBoardLikeQuery {
	ublq.ctx.Limit = &limit
	return ublq
}

// Offset to start from.
func (ublq *UserBoardLikeQuery) Offset(offset int) *UserBoardLikeQuery {
	ublq.ctx.Offset = &offset
	return ublq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ublq *UserBoardLikeQuery) Unique(unique bool) *UserBoardLikeQuery {
	ublq.ctx.Unique = &unique
	return ublq
}

// Order specifies how the records should be ordered.
func (ublq *UserBoardLikeQuery) Order(o ...userboardlike.OrderOption) *UserBoardLikeQuery {
	ublq.order = append(ublq.order, o...)
	return ublq
}

// QueryUser chains the current query on the "user" edge.
func (ublq *UserBoardLikeQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: ublq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ublq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ublq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(userboardlike.Table, userboardlike.UserColumn, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, userboardlike.UserTable, userboardlike.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(ublq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryBoard chains the current query on the "board" edge.
func (ublq *UserBoardLikeQuery) QueryBoard() *BoardQuery {
	query := (&BoardClient{config: ublq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ublq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ublq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(userboardlike.Table, userboardlike.BoardColumn, selector),
			sqlgraph.To(board.Table, board.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, userboardlike.BoardTable, userboardlike.BoardColumn),
		)
		fromU = sqlgraph.SetNeighbors(ublq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first UserBoardLike entity from the query.
// Returns a *NotFoundError when no UserBoardLike was found.
func (ublq *UserBoardLikeQuery) First(ctx context.Context) (*UserBoardLike, error) {
	nodes, err := ublq.Limit(1).All(setContextOp(ctx, ublq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{userboardlike.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ublq *UserBoardLikeQuery) FirstX(ctx context.Context) *UserBoardLike {
	node, err := ublq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// Only returns a single UserBoardLike entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one UserBoardLike entity is found.
// Returns a *NotFoundError when no UserBoardLike entities are found.
func (ublq *UserBoardLikeQuery) Only(ctx context.Context) (*UserBoardLike, error) {
	nodes, err := ublq.Limit(2).All(setContextOp(ctx, ublq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{userboardlike.Label}
	default:
		return nil, &NotSingularError{userboardlike.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ublq *UserBoardLikeQuery) OnlyX(ctx context.Context) *UserBoardLike {
	node, err := ublq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// All executes the query and returns a list of UserBoardLikes.
func (ublq *UserBoardLikeQuery) All(ctx context.Context) ([]*UserBoardLike, error) {
	ctx = setContextOp(ctx, ublq.ctx, "All")
	if err := ublq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*UserBoardLike, *UserBoardLikeQuery]()
	return withInterceptors[[]*UserBoardLike](ctx, ublq, qr, ublq.inters)
}

// AllX is like All, but panics if an error occurs.
func (ublq *UserBoardLikeQuery) AllX(ctx context.Context) []*UserBoardLike {
	nodes, err := ublq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// Count returns the count of the given query.
func (ublq *UserBoardLikeQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, ublq.ctx, "Count")
	if err := ublq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, ublq, querierCount[*UserBoardLikeQuery](), ublq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (ublq *UserBoardLikeQuery) CountX(ctx context.Context) int {
	count, err := ublq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ublq *UserBoardLikeQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, ublq.ctx, "Exist")
	switch _, err := ublq.First(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (ublq *UserBoardLikeQuery) ExistX(ctx context.Context) bool {
	exist, err := ublq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the UserBoardLikeQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ublq *UserBoardLikeQuery) Clone() *UserBoardLikeQuery {
	if ublq == nil {
		return nil
	}
	return &UserBoardLikeQuery{
		config:     ublq.config,
		ctx:        ublq.ctx.Clone(),
		order:      append([]userboardlike.OrderOption{}, ublq.order...),
		inters:     append([]Interceptor{}, ublq.inters...),
		predicates: append([]predicate.UserBoardLike{}, ublq.predicates...),
		withUser:   ublq.withUser.Clone(),
		withBoard:  ublq.withBoard.Clone(),
		// clone intermediate query.
		sql:  ublq.sql.Clone(),
		path: ublq.path,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (ublq *UserBoardLikeQuery) WithUser(opts ...func(*UserQuery)) *UserBoardLikeQuery {
	query := (&UserClient{config: ublq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ublq.withUser = query
	return ublq
}

// WithBoard tells the query-builder to eager-load the nodes that are connected to
// the "board" edge. The optional arguments are used to configure the query builder of the edge.
func (ublq *UserBoardLikeQuery) WithBoard(opts ...func(*BoardQuery)) *UserBoardLikeQuery {
	query := (&BoardClient{config: ublq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ublq.withBoard = query
	return ublq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		UserId int `json:"userId,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.UserBoardLike.Query().
//		GroupBy(userboardlike.FieldUserId).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (ublq *UserBoardLikeQuery) GroupBy(field string, fields ...string) *UserBoardLikeGroupBy {
	ublq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &UserBoardLikeGroupBy{build: ublq}
	grbuild.flds = &ublq.ctx.Fields
	grbuild.label = userboardlike.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		UserId int `json:"userId,omitempty"`
//	}
//
//	client.UserBoardLike.Query().
//		Select(userboardlike.FieldUserId).
//		Scan(ctx, &v)
func (ublq *UserBoardLikeQuery) Select(fields ...string) *UserBoardLikeSelect {
	ublq.ctx.Fields = append(ublq.ctx.Fields, fields...)
	sbuild := &UserBoardLikeSelect{UserBoardLikeQuery: ublq}
	sbuild.label = userboardlike.Label
	sbuild.flds, sbuild.scan = &ublq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a UserBoardLikeSelect configured with the given aggregations.
func (ublq *UserBoardLikeQuery) Aggregate(fns ...AggregateFunc) *UserBoardLikeSelect {
	return ublq.Select().Aggregate(fns...)
}

func (ublq *UserBoardLikeQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range ublq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, ublq); err != nil {
				return err
			}
		}
	}
	for _, f := range ublq.ctx.Fields {
		if !userboardlike.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ublq.path != nil {
		prev, err := ublq.path(ctx)
		if err != nil {
			return err
		}
		ublq.sql = prev
	}
	return nil
}

func (ublq *UserBoardLikeQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*UserBoardLike, error) {
	var (
		nodes       = []*UserBoardLike{}
		_spec       = ublq.querySpec()
		loadedTypes = [2]bool{
			ublq.withUser != nil,
			ublq.withBoard != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*UserBoardLike).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &UserBoardLike{config: ublq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ublq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := ublq.withUser; query != nil {
		if err := ublq.loadUser(ctx, query, nodes, nil,
			func(n *UserBoardLike, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	if query := ublq.withBoard; query != nil {
		if err := ublq.loadBoard(ctx, query, nodes, nil,
			func(n *UserBoardLike, e *Board) { n.Edges.Board = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (ublq *UserBoardLikeQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*UserBoardLike, init func(*UserBoardLike), assign func(*UserBoardLike, *User)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*UserBoardLike)
	for i := range nodes {
		fk := nodes[i].UserId
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "userId" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (ublq *UserBoardLikeQuery) loadBoard(ctx context.Context, query *BoardQuery, nodes []*UserBoardLike, init func(*UserBoardLike), assign func(*UserBoardLike, *Board)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*UserBoardLike)
	for i := range nodes {
		fk := nodes[i].BoardId
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(board.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "boardId" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (ublq *UserBoardLikeQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ublq.querySpec()
	_spec.Unique = false
	_spec.Node.Columns = nil
	return sqlgraph.CountNodes(ctx, ublq.driver, _spec)
}

func (ublq *UserBoardLikeQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(userboardlike.Table, userboardlike.Columns, nil)
	_spec.From = ublq.sql
	if unique := ublq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if ublq.path != nil {
		_spec.Unique = true
	}
	if fields := ublq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		for i := range fields {
			_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
		}
		if ublq.withUser != nil {
			_spec.Node.AddColumnOnce(userboardlike.FieldUserId)
		}
		if ublq.withBoard != nil {
			_spec.Node.AddColumnOnce(userboardlike.FieldBoardId)
		}
	}
	if ps := ublq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ublq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ublq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ublq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ublq *UserBoardLikeQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ublq.driver.Dialect())
	t1 := builder.Table(userboardlike.Table)
	columns := ublq.ctx.Fields
	if len(columns) == 0 {
		columns = userboardlike.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ublq.sql != nil {
		selector = ublq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ublq.ctx.Unique != nil && *ublq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range ublq.predicates {
		p(selector)
	}
	for _, p := range ublq.order {
		p(selector)
	}
	if offset := ublq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ublq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// UserBoardLikeGroupBy is the group-by builder for UserBoardLike entities.
type UserBoardLikeGroupBy struct {
	selector
	build *UserBoardLikeQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ublgb *UserBoardLikeGroupBy) Aggregate(fns ...AggregateFunc) *UserBoardLikeGroupBy {
	ublgb.fns = append(ublgb.fns, fns...)
	return ublgb
}

// Scan applies the selector query and scans the result into the given value.
func (ublgb *UserBoardLikeGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ublgb.build.ctx, "GroupBy")
	if err := ublgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UserBoardLikeQuery, *UserBoardLikeGroupBy](ctx, ublgb.build, ublgb, ublgb.build.inters, v)
}

func (ublgb *UserBoardLikeGroupBy) sqlScan(ctx context.Context, root *UserBoardLikeQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ublgb.fns))
	for _, fn := range ublgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ublgb.flds)+len(ublgb.fns))
		for _, f := range *ublgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ublgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ublgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// UserBoardLikeSelect is the builder for selecting fields of UserBoardLike entities.
type UserBoardLikeSelect struct {
	*UserBoardLikeQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ubls *UserBoardLikeSelect) Aggregate(fns ...AggregateFunc) *UserBoardLikeSelect {
	ubls.fns = append(ubls.fns, fns...)
	return ubls
}

// Scan applies the selector query and scans the result into the given value.
func (ubls *UserBoardLikeSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ubls.ctx, "Select")
	if err := ubls.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UserBoardLikeQuery, *UserBoardLikeSelect](ctx, ubls.UserBoardLikeQuery, ubls, ubls.inters, v)
}

func (ubls *UserBoardLikeSelect) sqlScan(ctx context.Context, root *UserBoardLikeQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ubls.fns))
	for _, fn := range ubls.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ubls.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ubls.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
