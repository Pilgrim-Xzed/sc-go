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
	"github.com/google/uuid"
	"github.com/stablecog/sc-go/database/ent/predicate"
	"github.com/stablecog/sc-go/database/ent/voiceover"
	"github.com/stablecog/sc-go/database/ent/voiceovermodel"
	"github.com/stablecog/sc-go/database/ent/voiceoverspeaker"
)

// VoiceoverModelQuery is the builder for querying VoiceoverModel entities.
type VoiceoverModelQuery struct {
	config
	ctx                   *QueryContext
	order                 []OrderFunc
	inters                []Interceptor
	predicates            []predicate.VoiceoverModel
	withVoiceovers        *VoiceoverQuery
	withVoiceoverSpeakers *VoiceoverSpeakerQuery
	modifiers             []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the VoiceoverModelQuery builder.
func (vmq *VoiceoverModelQuery) Where(ps ...predicate.VoiceoverModel) *VoiceoverModelQuery {
	vmq.predicates = append(vmq.predicates, ps...)
	return vmq
}

// Limit the number of records to be returned by this query.
func (vmq *VoiceoverModelQuery) Limit(limit int) *VoiceoverModelQuery {
	vmq.ctx.Limit = &limit
	return vmq
}

// Offset to start from.
func (vmq *VoiceoverModelQuery) Offset(offset int) *VoiceoverModelQuery {
	vmq.ctx.Offset = &offset
	return vmq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (vmq *VoiceoverModelQuery) Unique(unique bool) *VoiceoverModelQuery {
	vmq.ctx.Unique = &unique
	return vmq
}

// Order specifies how the records should be ordered.
func (vmq *VoiceoverModelQuery) Order(o ...OrderFunc) *VoiceoverModelQuery {
	vmq.order = append(vmq.order, o...)
	return vmq
}

// QueryVoiceovers chains the current query on the "voiceovers" edge.
func (vmq *VoiceoverModelQuery) QueryVoiceovers() *VoiceoverQuery {
	query := (&VoiceoverClient{config: vmq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := vmq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := vmq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(voiceovermodel.Table, voiceovermodel.FieldID, selector),
			sqlgraph.To(voiceover.Table, voiceover.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, voiceovermodel.VoiceoversTable, voiceovermodel.VoiceoversColumn),
		)
		fromU = sqlgraph.SetNeighbors(vmq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryVoiceoverSpeakers chains the current query on the "voiceover_speakers" edge.
func (vmq *VoiceoverModelQuery) QueryVoiceoverSpeakers() *VoiceoverSpeakerQuery {
	query := (&VoiceoverSpeakerClient{config: vmq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := vmq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := vmq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(voiceovermodel.Table, voiceovermodel.FieldID, selector),
			sqlgraph.To(voiceoverspeaker.Table, voiceoverspeaker.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, voiceovermodel.VoiceoverSpeakersTable, voiceovermodel.VoiceoverSpeakersColumn),
		)
		fromU = sqlgraph.SetNeighbors(vmq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first VoiceoverModel entity from the query.
// Returns a *NotFoundError when no VoiceoverModel was found.
func (vmq *VoiceoverModelQuery) First(ctx context.Context) (*VoiceoverModel, error) {
	nodes, err := vmq.Limit(1).All(setContextOp(ctx, vmq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{voiceovermodel.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (vmq *VoiceoverModelQuery) FirstX(ctx context.Context) *VoiceoverModel {
	node, err := vmq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first VoiceoverModel ID from the query.
// Returns a *NotFoundError when no VoiceoverModel ID was found.
func (vmq *VoiceoverModelQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = vmq.Limit(1).IDs(setContextOp(ctx, vmq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{voiceovermodel.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (vmq *VoiceoverModelQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := vmq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single VoiceoverModel entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one VoiceoverModel entity is found.
// Returns a *NotFoundError when no VoiceoverModel entities are found.
func (vmq *VoiceoverModelQuery) Only(ctx context.Context) (*VoiceoverModel, error) {
	nodes, err := vmq.Limit(2).All(setContextOp(ctx, vmq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{voiceovermodel.Label}
	default:
		return nil, &NotSingularError{voiceovermodel.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (vmq *VoiceoverModelQuery) OnlyX(ctx context.Context) *VoiceoverModel {
	node, err := vmq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only VoiceoverModel ID in the query.
// Returns a *NotSingularError when more than one VoiceoverModel ID is found.
// Returns a *NotFoundError when no entities are found.
func (vmq *VoiceoverModelQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = vmq.Limit(2).IDs(setContextOp(ctx, vmq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{voiceovermodel.Label}
	default:
		err = &NotSingularError{voiceovermodel.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (vmq *VoiceoverModelQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := vmq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of VoiceoverModels.
func (vmq *VoiceoverModelQuery) All(ctx context.Context) ([]*VoiceoverModel, error) {
	ctx = setContextOp(ctx, vmq.ctx, "All")
	if err := vmq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*VoiceoverModel, *VoiceoverModelQuery]()
	return withInterceptors[[]*VoiceoverModel](ctx, vmq, qr, vmq.inters)
}

// AllX is like All, but panics if an error occurs.
func (vmq *VoiceoverModelQuery) AllX(ctx context.Context) []*VoiceoverModel {
	nodes, err := vmq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of VoiceoverModel IDs.
func (vmq *VoiceoverModelQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	ctx = setContextOp(ctx, vmq.ctx, "IDs")
	if err := vmq.Select(voiceovermodel.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (vmq *VoiceoverModelQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := vmq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (vmq *VoiceoverModelQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, vmq.ctx, "Count")
	if err := vmq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, vmq, querierCount[*VoiceoverModelQuery](), vmq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (vmq *VoiceoverModelQuery) CountX(ctx context.Context) int {
	count, err := vmq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (vmq *VoiceoverModelQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, vmq.ctx, "Exist")
	switch _, err := vmq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (vmq *VoiceoverModelQuery) ExistX(ctx context.Context) bool {
	exist, err := vmq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the VoiceoverModelQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (vmq *VoiceoverModelQuery) Clone() *VoiceoverModelQuery {
	if vmq == nil {
		return nil
	}
	return &VoiceoverModelQuery{
		config:                vmq.config,
		ctx:                   vmq.ctx.Clone(),
		order:                 append([]OrderFunc{}, vmq.order...),
		inters:                append([]Interceptor{}, vmq.inters...),
		predicates:            append([]predicate.VoiceoverModel{}, vmq.predicates...),
		withVoiceovers:        vmq.withVoiceovers.Clone(),
		withVoiceoverSpeakers: vmq.withVoiceoverSpeakers.Clone(),
		// clone intermediate query.
		sql:  vmq.sql.Clone(),
		path: vmq.path,
	}
}

// WithVoiceovers tells the query-builder to eager-load the nodes that are connected to
// the "voiceovers" edge. The optional arguments are used to configure the query builder of the edge.
func (vmq *VoiceoverModelQuery) WithVoiceovers(opts ...func(*VoiceoverQuery)) *VoiceoverModelQuery {
	query := (&VoiceoverClient{config: vmq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	vmq.withVoiceovers = query
	return vmq
}

// WithVoiceoverSpeakers tells the query-builder to eager-load the nodes that are connected to
// the "voiceover_speakers" edge. The optional arguments are used to configure the query builder of the edge.
func (vmq *VoiceoverModelQuery) WithVoiceoverSpeakers(opts ...func(*VoiceoverSpeakerQuery)) *VoiceoverModelQuery {
	query := (&VoiceoverSpeakerClient{config: vmq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	vmq.withVoiceoverSpeakers = query
	return vmq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		NameInWorker string `json:"name_in_worker,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.VoiceoverModel.Query().
//		GroupBy(voiceovermodel.FieldNameInWorker).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (vmq *VoiceoverModelQuery) GroupBy(field string, fields ...string) *VoiceoverModelGroupBy {
	vmq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &VoiceoverModelGroupBy{build: vmq}
	grbuild.flds = &vmq.ctx.Fields
	grbuild.label = voiceovermodel.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		NameInWorker string `json:"name_in_worker,omitempty"`
//	}
//
//	client.VoiceoverModel.Query().
//		Select(voiceovermodel.FieldNameInWorker).
//		Scan(ctx, &v)
func (vmq *VoiceoverModelQuery) Select(fields ...string) *VoiceoverModelSelect {
	vmq.ctx.Fields = append(vmq.ctx.Fields, fields...)
	sbuild := &VoiceoverModelSelect{VoiceoverModelQuery: vmq}
	sbuild.label = voiceovermodel.Label
	sbuild.flds, sbuild.scan = &vmq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a VoiceoverModelSelect configured with the given aggregations.
func (vmq *VoiceoverModelQuery) Aggregate(fns ...AggregateFunc) *VoiceoverModelSelect {
	return vmq.Select().Aggregate(fns...)
}

func (vmq *VoiceoverModelQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range vmq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, vmq); err != nil {
				return err
			}
		}
	}
	for _, f := range vmq.ctx.Fields {
		if !voiceovermodel.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if vmq.path != nil {
		prev, err := vmq.path(ctx)
		if err != nil {
			return err
		}
		vmq.sql = prev
	}
	return nil
}

func (vmq *VoiceoverModelQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*VoiceoverModel, error) {
	var (
		nodes       = []*VoiceoverModel{}
		_spec       = vmq.querySpec()
		loadedTypes = [2]bool{
			vmq.withVoiceovers != nil,
			vmq.withVoiceoverSpeakers != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*VoiceoverModel).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &VoiceoverModel{config: vmq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(vmq.modifiers) > 0 {
		_spec.Modifiers = vmq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, vmq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := vmq.withVoiceovers; query != nil {
		if err := vmq.loadVoiceovers(ctx, query, nodes,
			func(n *VoiceoverModel) { n.Edges.Voiceovers = []*Voiceover{} },
			func(n *VoiceoverModel, e *Voiceover) { n.Edges.Voiceovers = append(n.Edges.Voiceovers, e) }); err != nil {
			return nil, err
		}
	}
	if query := vmq.withVoiceoverSpeakers; query != nil {
		if err := vmq.loadVoiceoverSpeakers(ctx, query, nodes,
			func(n *VoiceoverModel) { n.Edges.VoiceoverSpeakers = []*VoiceoverSpeaker{} },
			func(n *VoiceoverModel, e *VoiceoverSpeaker) {
				n.Edges.VoiceoverSpeakers = append(n.Edges.VoiceoverSpeakers, e)
			}); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (vmq *VoiceoverModelQuery) loadVoiceovers(ctx context.Context, query *VoiceoverQuery, nodes []*VoiceoverModel, init func(*VoiceoverModel), assign func(*VoiceoverModel, *Voiceover)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*VoiceoverModel)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.Where(predicate.Voiceover(func(s *sql.Selector) {
		s.Where(sql.InValues(voiceovermodel.VoiceoversColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.ModelID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "model_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (vmq *VoiceoverModelQuery) loadVoiceoverSpeakers(ctx context.Context, query *VoiceoverSpeakerQuery, nodes []*VoiceoverModel, init func(*VoiceoverModel), assign func(*VoiceoverModel, *VoiceoverSpeaker)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*VoiceoverModel)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.Where(predicate.VoiceoverSpeaker(func(s *sql.Selector) {
		s.Where(sql.InValues(voiceovermodel.VoiceoverSpeakersColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.ModelID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "model_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (vmq *VoiceoverModelQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := vmq.querySpec()
	if len(vmq.modifiers) > 0 {
		_spec.Modifiers = vmq.modifiers
	}
	_spec.Node.Columns = vmq.ctx.Fields
	if len(vmq.ctx.Fields) > 0 {
		_spec.Unique = vmq.ctx.Unique != nil && *vmq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, vmq.driver, _spec)
}

func (vmq *VoiceoverModelQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   voiceovermodel.Table,
			Columns: voiceovermodel.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: voiceovermodel.FieldID,
			},
		},
		From:   vmq.sql,
		Unique: true,
	}
	if unique := vmq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := vmq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, voiceovermodel.FieldID)
		for i := range fields {
			if fields[i] != voiceovermodel.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := vmq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := vmq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := vmq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := vmq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (vmq *VoiceoverModelQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(vmq.driver.Dialect())
	t1 := builder.Table(voiceovermodel.Table)
	columns := vmq.ctx.Fields
	if len(columns) == 0 {
		columns = voiceovermodel.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if vmq.sql != nil {
		selector = vmq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if vmq.ctx.Unique != nil && *vmq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range vmq.modifiers {
		m(selector)
	}
	for _, p := range vmq.predicates {
		p(selector)
	}
	for _, p := range vmq.order {
		p(selector)
	}
	if offset := vmq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := vmq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (vmq *VoiceoverModelQuery) Modify(modifiers ...func(s *sql.Selector)) *VoiceoverModelSelect {
	vmq.modifiers = append(vmq.modifiers, modifiers...)
	return vmq.Select()
}

// VoiceoverModelGroupBy is the group-by builder for VoiceoverModel entities.
type VoiceoverModelGroupBy struct {
	selector
	build *VoiceoverModelQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (vmgb *VoiceoverModelGroupBy) Aggregate(fns ...AggregateFunc) *VoiceoverModelGroupBy {
	vmgb.fns = append(vmgb.fns, fns...)
	return vmgb
}

// Scan applies the selector query and scans the result into the given value.
func (vmgb *VoiceoverModelGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, vmgb.build.ctx, "GroupBy")
	if err := vmgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*VoiceoverModelQuery, *VoiceoverModelGroupBy](ctx, vmgb.build, vmgb, vmgb.build.inters, v)
}

func (vmgb *VoiceoverModelGroupBy) sqlScan(ctx context.Context, root *VoiceoverModelQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(vmgb.fns))
	for _, fn := range vmgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*vmgb.flds)+len(vmgb.fns))
		for _, f := range *vmgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*vmgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := vmgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// VoiceoverModelSelect is the builder for selecting fields of VoiceoverModel entities.
type VoiceoverModelSelect struct {
	*VoiceoverModelQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (vms *VoiceoverModelSelect) Aggregate(fns ...AggregateFunc) *VoiceoverModelSelect {
	vms.fns = append(vms.fns, fns...)
	return vms
}

// Scan applies the selector query and scans the result into the given value.
func (vms *VoiceoverModelSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, vms.ctx, "Select")
	if err := vms.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*VoiceoverModelQuery, *VoiceoverModelSelect](ctx, vms.VoiceoverModelQuery, vms, vms.inters, v)
}

func (vms *VoiceoverModelSelect) sqlScan(ctx context.Context, root *VoiceoverModelQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(vms.fns))
	for _, fn := range vms.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*vms.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := vms.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (vms *VoiceoverModelSelect) Modify(modifiers ...func(s *sql.Selector)) *VoiceoverModelSelect {
	vms.modifiers = append(vms.modifiers, modifiers...)
	return vms
}
