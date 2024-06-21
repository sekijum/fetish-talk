// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"
	"server/infrastructure/ent/forum"
	"server/infrastructure/ent/predicate"
	"server/infrastructure/ent/user"
	"server/infrastructure/ent/userforumsubscription"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// UserForumSubscriptionQuery is the builder for querying UserForumSubscription entities.
type UserForumSubscriptionQuery struct {
	config
	ctx        *QueryContext
	order      []userforumsubscription.OrderOption
	inters     []Interceptor
	predicates []predicate.UserForumSubscription
	withUser   *UserQuery
	withForum  *ForumQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the UserForumSubscriptionQuery builder.
func (ufsq *UserForumSubscriptionQuery) Where(ps ...predicate.UserForumSubscription) *UserForumSubscriptionQuery {
	ufsq.predicates = append(ufsq.predicates, ps...)
	return ufsq
}

// Limit the number of records to be returned by this query.
func (ufsq *UserForumSubscriptionQuery) Limit(limit int) *UserForumSubscriptionQuery {
	ufsq.ctx.Limit = &limit
	return ufsq
}

// Offset to start from.
func (ufsq *UserForumSubscriptionQuery) Offset(offset int) *UserForumSubscriptionQuery {
	ufsq.ctx.Offset = &offset
	return ufsq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ufsq *UserForumSubscriptionQuery) Unique(unique bool) *UserForumSubscriptionQuery {
	ufsq.ctx.Unique = &unique
	return ufsq
}

// Order specifies how the records should be ordered.
func (ufsq *UserForumSubscriptionQuery) Order(o ...userforumsubscription.OrderOption) *UserForumSubscriptionQuery {
	ufsq.order = append(ufsq.order, o...)
	return ufsq
}

// QueryUser chains the current query on the "user" edge.
func (ufsq *UserForumSubscriptionQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: ufsq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ufsq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ufsq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(userforumsubscription.Table, userforumsubscription.UserColumn, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, userforumsubscription.UserTable, userforumsubscription.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(ufsq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryForum chains the current query on the "forum" edge.
func (ufsq *UserForumSubscriptionQuery) QueryForum() *ForumQuery {
	query := (&ForumClient{config: ufsq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ufsq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ufsq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(userforumsubscription.Table, userforumsubscription.ForumColumn, selector),
			sqlgraph.To(forum.Table, forum.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, userforumsubscription.ForumTable, userforumsubscription.ForumColumn),
		)
		fromU = sqlgraph.SetNeighbors(ufsq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first UserForumSubscription entity from the query.
// Returns a *NotFoundError when no UserForumSubscription was found.
func (ufsq *UserForumSubscriptionQuery) First(ctx context.Context) (*UserForumSubscription, error) {
	nodes, err := ufsq.Limit(1).All(setContextOp(ctx, ufsq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{userforumsubscription.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ufsq *UserForumSubscriptionQuery) FirstX(ctx context.Context) *UserForumSubscription {
	node, err := ufsq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// Only returns a single UserForumSubscription entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one UserForumSubscription entity is found.
// Returns a *NotFoundError when no UserForumSubscription entities are found.
func (ufsq *UserForumSubscriptionQuery) Only(ctx context.Context) (*UserForumSubscription, error) {
	nodes, err := ufsq.Limit(2).All(setContextOp(ctx, ufsq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{userforumsubscription.Label}
	default:
		return nil, &NotSingularError{userforumsubscription.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ufsq *UserForumSubscriptionQuery) OnlyX(ctx context.Context) *UserForumSubscription {
	node, err := ufsq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// All executes the query and returns a list of UserForumSubscriptions.
func (ufsq *UserForumSubscriptionQuery) All(ctx context.Context) ([]*UserForumSubscription, error) {
	ctx = setContextOp(ctx, ufsq.ctx, "All")
	if err := ufsq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*UserForumSubscription, *UserForumSubscriptionQuery]()
	return withInterceptors[[]*UserForumSubscription](ctx, ufsq, qr, ufsq.inters)
}

// AllX is like All, but panics if an error occurs.
func (ufsq *UserForumSubscriptionQuery) AllX(ctx context.Context) []*UserForumSubscription {
	nodes, err := ufsq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// Count returns the count of the given query.
func (ufsq *UserForumSubscriptionQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, ufsq.ctx, "Count")
	if err := ufsq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, ufsq, querierCount[*UserForumSubscriptionQuery](), ufsq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (ufsq *UserForumSubscriptionQuery) CountX(ctx context.Context) int {
	count, err := ufsq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ufsq *UserForumSubscriptionQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, ufsq.ctx, "Exist")
	switch _, err := ufsq.First(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (ufsq *UserForumSubscriptionQuery) ExistX(ctx context.Context) bool {
	exist, err := ufsq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the UserForumSubscriptionQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ufsq *UserForumSubscriptionQuery) Clone() *UserForumSubscriptionQuery {
	if ufsq == nil {
		return nil
	}
	return &UserForumSubscriptionQuery{
		config:     ufsq.config,
		ctx:        ufsq.ctx.Clone(),
		order:      append([]userforumsubscription.OrderOption{}, ufsq.order...),
		inters:     append([]Interceptor{}, ufsq.inters...),
		predicates: append([]predicate.UserForumSubscription{}, ufsq.predicates...),
		withUser:   ufsq.withUser.Clone(),
		withForum:  ufsq.withForum.Clone(),
		// clone intermediate query.
		sql:  ufsq.sql.Clone(),
		path: ufsq.path,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (ufsq *UserForumSubscriptionQuery) WithUser(opts ...func(*UserQuery)) *UserForumSubscriptionQuery {
	query := (&UserClient{config: ufsq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ufsq.withUser = query
	return ufsq
}

// WithForum tells the query-builder to eager-load the nodes that are connected to
// the "forum" edge. The optional arguments are used to configure the query builder of the edge.
func (ufsq *UserForumSubscriptionQuery) WithForum(opts ...func(*ForumQuery)) *UserForumSubscriptionQuery {
	query := (&ForumClient{config: ufsq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ufsq.withForum = query
	return ufsq
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
//	client.UserForumSubscription.Query().
//		GroupBy(userforumsubscription.FieldUserId).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (ufsq *UserForumSubscriptionQuery) GroupBy(field string, fields ...string) *UserForumSubscriptionGroupBy {
	ufsq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &UserForumSubscriptionGroupBy{build: ufsq}
	grbuild.flds = &ufsq.ctx.Fields
	grbuild.label = userforumsubscription.Label
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
//	client.UserForumSubscription.Query().
//		Select(userforumsubscription.FieldUserId).
//		Scan(ctx, &v)
func (ufsq *UserForumSubscriptionQuery) Select(fields ...string) *UserForumSubscriptionSelect {
	ufsq.ctx.Fields = append(ufsq.ctx.Fields, fields...)
	sbuild := &UserForumSubscriptionSelect{UserForumSubscriptionQuery: ufsq}
	sbuild.label = userforumsubscription.Label
	sbuild.flds, sbuild.scan = &ufsq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a UserForumSubscriptionSelect configured with the given aggregations.
func (ufsq *UserForumSubscriptionQuery) Aggregate(fns ...AggregateFunc) *UserForumSubscriptionSelect {
	return ufsq.Select().Aggregate(fns...)
}

func (ufsq *UserForumSubscriptionQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range ufsq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, ufsq); err != nil {
				return err
			}
		}
	}
	for _, f := range ufsq.ctx.Fields {
		if !userforumsubscription.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ufsq.path != nil {
		prev, err := ufsq.path(ctx)
		if err != nil {
			return err
		}
		ufsq.sql = prev
	}
	return nil
}

func (ufsq *UserForumSubscriptionQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*UserForumSubscription, error) {
	var (
		nodes       = []*UserForumSubscription{}
		_spec       = ufsq.querySpec()
		loadedTypes = [2]bool{
			ufsq.withUser != nil,
			ufsq.withForum != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*UserForumSubscription).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &UserForumSubscription{config: ufsq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ufsq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := ufsq.withUser; query != nil {
		if err := ufsq.loadUser(ctx, query, nodes, nil,
			func(n *UserForumSubscription, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	if query := ufsq.withForum; query != nil {
		if err := ufsq.loadForum(ctx, query, nodes, nil,
			func(n *UserForumSubscription, e *Forum) { n.Edges.Forum = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (ufsq *UserForumSubscriptionQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*UserForumSubscription, init func(*UserForumSubscription), assign func(*UserForumSubscription, *User)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*UserForumSubscription)
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
func (ufsq *UserForumSubscriptionQuery) loadForum(ctx context.Context, query *ForumQuery, nodes []*UserForumSubscription, init func(*UserForumSubscription), assign func(*UserForumSubscription, *Forum)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*UserForumSubscription)
	for i := range nodes {
		fk := nodes[i].ForumId
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(forum.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "forumId" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (ufsq *UserForumSubscriptionQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ufsq.querySpec()
	_spec.Unique = false
	_spec.Node.Columns = nil
	return sqlgraph.CountNodes(ctx, ufsq.driver, _spec)
}

func (ufsq *UserForumSubscriptionQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(userforumsubscription.Table, userforumsubscription.Columns, nil)
	_spec.From = ufsq.sql
	if unique := ufsq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if ufsq.path != nil {
		_spec.Unique = true
	}
	if fields := ufsq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		for i := range fields {
			_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
		}
		if ufsq.withUser != nil {
			_spec.Node.AddColumnOnce(userforumsubscription.FieldUserId)
		}
		if ufsq.withForum != nil {
			_spec.Node.AddColumnOnce(userforumsubscription.FieldForumId)
		}
	}
	if ps := ufsq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ufsq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ufsq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ufsq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ufsq *UserForumSubscriptionQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ufsq.driver.Dialect())
	t1 := builder.Table(userforumsubscription.Table)
	columns := ufsq.ctx.Fields
	if len(columns) == 0 {
		columns = userforumsubscription.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ufsq.sql != nil {
		selector = ufsq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ufsq.ctx.Unique != nil && *ufsq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range ufsq.predicates {
		p(selector)
	}
	for _, p := range ufsq.order {
		p(selector)
	}
	if offset := ufsq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ufsq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// UserForumSubscriptionGroupBy is the group-by builder for UserForumSubscription entities.
type UserForumSubscriptionGroupBy struct {
	selector
	build *UserForumSubscriptionQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ufsgb *UserForumSubscriptionGroupBy) Aggregate(fns ...AggregateFunc) *UserForumSubscriptionGroupBy {
	ufsgb.fns = append(ufsgb.fns, fns...)
	return ufsgb
}

// Scan applies the selector query and scans the result into the given value.
func (ufsgb *UserForumSubscriptionGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ufsgb.build.ctx, "GroupBy")
	if err := ufsgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UserForumSubscriptionQuery, *UserForumSubscriptionGroupBy](ctx, ufsgb.build, ufsgb, ufsgb.build.inters, v)
}

func (ufsgb *UserForumSubscriptionGroupBy) sqlScan(ctx context.Context, root *UserForumSubscriptionQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ufsgb.fns))
	for _, fn := range ufsgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ufsgb.flds)+len(ufsgb.fns))
		for _, f := range *ufsgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ufsgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ufsgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// UserForumSubscriptionSelect is the builder for selecting fields of UserForumSubscription entities.
type UserForumSubscriptionSelect struct {
	*UserForumSubscriptionQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ufss *UserForumSubscriptionSelect) Aggregate(fns ...AggregateFunc) *UserForumSubscriptionSelect {
	ufss.fns = append(ufss.fns, fns...)
	return ufss
}

// Scan applies the selector query and scans the result into the given value.
func (ufss *UserForumSubscriptionSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ufss.ctx, "Select")
	if err := ufss.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UserForumSubscriptionQuery, *UserForumSubscriptionSelect](ctx, ufss.UserForumSubscriptionQuery, ufss, ufss.inters, v)
}

func (ufss *UserForumSubscriptionSelect) sqlScan(ctx context.Context, root *UserForumSubscriptionQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ufss.fns))
	for _, fn := range ufss.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ufss.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ufss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
