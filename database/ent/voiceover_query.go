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
	"github.com/stablecog/sc-go/database/ent/apitoken"
	"github.com/stablecog/sc-go/database/ent/deviceinfo"
	"github.com/stablecog/sc-go/database/ent/predicate"
	"github.com/stablecog/sc-go/database/ent/prompt"
	"github.com/stablecog/sc-go/database/ent/user"
	"github.com/stablecog/sc-go/database/ent/voiceover"
	"github.com/stablecog/sc-go/database/ent/voiceovermodel"
	"github.com/stablecog/sc-go/database/ent/voiceoveroutput"
	"github.com/stablecog/sc-go/database/ent/voiceoverspeaker"
)

// VoiceoverQuery is the builder for querying Voiceover entities.
type VoiceoverQuery struct {
	config
	ctx                   *QueryContext
	order                 []OrderFunc
	inters                []Interceptor
	predicates            []predicate.Voiceover
	withUser              *UserQuery
	withPrompt            *PromptQuery
	withDeviceInfo        *DeviceInfoQuery
	withVoiceoverModels   *VoiceoverModelQuery
	withVoiceoverSpeakers *VoiceoverSpeakerQuery
	withAPITokens         *ApiTokenQuery
	withVoiceoverOutputs  *VoiceoverOutputQuery
	modifiers             []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the VoiceoverQuery builder.
func (vq *VoiceoverQuery) Where(ps ...predicate.Voiceover) *VoiceoverQuery {
	vq.predicates = append(vq.predicates, ps...)
	return vq
}

// Limit the number of records to be returned by this query.
func (vq *VoiceoverQuery) Limit(limit int) *VoiceoverQuery {
	vq.ctx.Limit = &limit
	return vq
}

// Offset to start from.
func (vq *VoiceoverQuery) Offset(offset int) *VoiceoverQuery {
	vq.ctx.Offset = &offset
	return vq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (vq *VoiceoverQuery) Unique(unique bool) *VoiceoverQuery {
	vq.ctx.Unique = &unique
	return vq
}

// Order specifies how the records should be ordered.
func (vq *VoiceoverQuery) Order(o ...OrderFunc) *VoiceoverQuery {
	vq.order = append(vq.order, o...)
	return vq
}

// QueryUser chains the current query on the "user" edge.
func (vq *VoiceoverQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: vq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := vq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := vq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(voiceover.Table, voiceover.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, voiceover.UserTable, voiceover.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(vq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryPrompt chains the current query on the "prompt" edge.
func (vq *VoiceoverQuery) QueryPrompt() *PromptQuery {
	query := (&PromptClient{config: vq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := vq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := vq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(voiceover.Table, voiceover.FieldID, selector),
			sqlgraph.To(prompt.Table, prompt.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, voiceover.PromptTable, voiceover.PromptColumn),
		)
		fromU = sqlgraph.SetNeighbors(vq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryDeviceInfo chains the current query on the "device_info" edge.
func (vq *VoiceoverQuery) QueryDeviceInfo() *DeviceInfoQuery {
	query := (&DeviceInfoClient{config: vq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := vq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := vq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(voiceover.Table, voiceover.FieldID, selector),
			sqlgraph.To(deviceinfo.Table, deviceinfo.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, voiceover.DeviceInfoTable, voiceover.DeviceInfoColumn),
		)
		fromU = sqlgraph.SetNeighbors(vq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryVoiceoverModels chains the current query on the "voiceover_models" edge.
func (vq *VoiceoverQuery) QueryVoiceoverModels() *VoiceoverModelQuery {
	query := (&VoiceoverModelClient{config: vq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := vq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := vq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(voiceover.Table, voiceover.FieldID, selector),
			sqlgraph.To(voiceovermodel.Table, voiceovermodel.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, voiceover.VoiceoverModelsTable, voiceover.VoiceoverModelsColumn),
		)
		fromU = sqlgraph.SetNeighbors(vq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryVoiceoverSpeakers chains the current query on the "voiceover_speakers" edge.
func (vq *VoiceoverQuery) QueryVoiceoverSpeakers() *VoiceoverSpeakerQuery {
	query := (&VoiceoverSpeakerClient{config: vq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := vq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := vq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(voiceover.Table, voiceover.FieldID, selector),
			sqlgraph.To(voiceoverspeaker.Table, voiceoverspeaker.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, voiceover.VoiceoverSpeakersTable, voiceover.VoiceoverSpeakersColumn),
		)
		fromU = sqlgraph.SetNeighbors(vq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryAPITokens chains the current query on the "api_tokens" edge.
func (vq *VoiceoverQuery) QueryAPITokens() *ApiTokenQuery {
	query := (&ApiTokenClient{config: vq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := vq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := vq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(voiceover.Table, voiceover.FieldID, selector),
			sqlgraph.To(apitoken.Table, apitoken.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, voiceover.APITokensTable, voiceover.APITokensColumn),
		)
		fromU = sqlgraph.SetNeighbors(vq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryVoiceoverOutputs chains the current query on the "voiceover_outputs" edge.
func (vq *VoiceoverQuery) QueryVoiceoverOutputs() *VoiceoverOutputQuery {
	query := (&VoiceoverOutputClient{config: vq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := vq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := vq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(voiceover.Table, voiceover.FieldID, selector),
			sqlgraph.To(voiceoveroutput.Table, voiceoveroutput.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, voiceover.VoiceoverOutputsTable, voiceover.VoiceoverOutputsColumn),
		)
		fromU = sqlgraph.SetNeighbors(vq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Voiceover entity from the query.
// Returns a *NotFoundError when no Voiceover was found.
func (vq *VoiceoverQuery) First(ctx context.Context) (*Voiceover, error) {
	nodes, err := vq.Limit(1).All(setContextOp(ctx, vq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{voiceover.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (vq *VoiceoverQuery) FirstX(ctx context.Context) *Voiceover {
	node, err := vq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Voiceover ID from the query.
// Returns a *NotFoundError when no Voiceover ID was found.
func (vq *VoiceoverQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = vq.Limit(1).IDs(setContextOp(ctx, vq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{voiceover.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (vq *VoiceoverQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := vq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Voiceover entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Voiceover entity is found.
// Returns a *NotFoundError when no Voiceover entities are found.
func (vq *VoiceoverQuery) Only(ctx context.Context) (*Voiceover, error) {
	nodes, err := vq.Limit(2).All(setContextOp(ctx, vq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{voiceover.Label}
	default:
		return nil, &NotSingularError{voiceover.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (vq *VoiceoverQuery) OnlyX(ctx context.Context) *Voiceover {
	node, err := vq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Voiceover ID in the query.
// Returns a *NotSingularError when more than one Voiceover ID is found.
// Returns a *NotFoundError when no entities are found.
func (vq *VoiceoverQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = vq.Limit(2).IDs(setContextOp(ctx, vq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{voiceover.Label}
	default:
		err = &NotSingularError{voiceover.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (vq *VoiceoverQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := vq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Voiceovers.
func (vq *VoiceoverQuery) All(ctx context.Context) ([]*Voiceover, error) {
	ctx = setContextOp(ctx, vq.ctx, "All")
	if err := vq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Voiceover, *VoiceoverQuery]()
	return withInterceptors[[]*Voiceover](ctx, vq, qr, vq.inters)
}

// AllX is like All, but panics if an error occurs.
func (vq *VoiceoverQuery) AllX(ctx context.Context) []*Voiceover {
	nodes, err := vq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Voiceover IDs.
func (vq *VoiceoverQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	ctx = setContextOp(ctx, vq.ctx, "IDs")
	if err := vq.Select(voiceover.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (vq *VoiceoverQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := vq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (vq *VoiceoverQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, vq.ctx, "Count")
	if err := vq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, vq, querierCount[*VoiceoverQuery](), vq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (vq *VoiceoverQuery) CountX(ctx context.Context) int {
	count, err := vq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (vq *VoiceoverQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, vq.ctx, "Exist")
	switch _, err := vq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (vq *VoiceoverQuery) ExistX(ctx context.Context) bool {
	exist, err := vq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the VoiceoverQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (vq *VoiceoverQuery) Clone() *VoiceoverQuery {
	if vq == nil {
		return nil
	}
	return &VoiceoverQuery{
		config:                vq.config,
		ctx:                   vq.ctx.Clone(),
		order:                 append([]OrderFunc{}, vq.order...),
		inters:                append([]Interceptor{}, vq.inters...),
		predicates:            append([]predicate.Voiceover{}, vq.predicates...),
		withUser:              vq.withUser.Clone(),
		withPrompt:            vq.withPrompt.Clone(),
		withDeviceInfo:        vq.withDeviceInfo.Clone(),
		withVoiceoverModels:   vq.withVoiceoverModels.Clone(),
		withVoiceoverSpeakers: vq.withVoiceoverSpeakers.Clone(),
		withAPITokens:         vq.withAPITokens.Clone(),
		withVoiceoverOutputs:  vq.withVoiceoverOutputs.Clone(),
		// clone intermediate query.
		sql:  vq.sql.Clone(),
		path: vq.path,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (vq *VoiceoverQuery) WithUser(opts ...func(*UserQuery)) *VoiceoverQuery {
	query := (&UserClient{config: vq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	vq.withUser = query
	return vq
}

// WithPrompt tells the query-builder to eager-load the nodes that are connected to
// the "prompt" edge. The optional arguments are used to configure the query builder of the edge.
func (vq *VoiceoverQuery) WithPrompt(opts ...func(*PromptQuery)) *VoiceoverQuery {
	query := (&PromptClient{config: vq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	vq.withPrompt = query
	return vq
}

// WithDeviceInfo tells the query-builder to eager-load the nodes that are connected to
// the "device_info" edge. The optional arguments are used to configure the query builder of the edge.
func (vq *VoiceoverQuery) WithDeviceInfo(opts ...func(*DeviceInfoQuery)) *VoiceoverQuery {
	query := (&DeviceInfoClient{config: vq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	vq.withDeviceInfo = query
	return vq
}

// WithVoiceoverModels tells the query-builder to eager-load the nodes that are connected to
// the "voiceover_models" edge. The optional arguments are used to configure the query builder of the edge.
func (vq *VoiceoverQuery) WithVoiceoverModels(opts ...func(*VoiceoverModelQuery)) *VoiceoverQuery {
	query := (&VoiceoverModelClient{config: vq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	vq.withVoiceoverModels = query
	return vq
}

// WithVoiceoverSpeakers tells the query-builder to eager-load the nodes that are connected to
// the "voiceover_speakers" edge. The optional arguments are used to configure the query builder of the edge.
func (vq *VoiceoverQuery) WithVoiceoverSpeakers(opts ...func(*VoiceoverSpeakerQuery)) *VoiceoverQuery {
	query := (&VoiceoverSpeakerClient{config: vq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	vq.withVoiceoverSpeakers = query
	return vq
}

// WithAPITokens tells the query-builder to eager-load the nodes that are connected to
// the "api_tokens" edge. The optional arguments are used to configure the query builder of the edge.
func (vq *VoiceoverQuery) WithAPITokens(opts ...func(*ApiTokenQuery)) *VoiceoverQuery {
	query := (&ApiTokenClient{config: vq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	vq.withAPITokens = query
	return vq
}

// WithVoiceoverOutputs tells the query-builder to eager-load the nodes that are connected to
// the "voiceover_outputs" edge. The optional arguments are used to configure the query builder of the edge.
func (vq *VoiceoverQuery) WithVoiceoverOutputs(opts ...func(*VoiceoverOutputQuery)) *VoiceoverQuery {
	query := (&VoiceoverOutputClient{config: vq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	vq.withVoiceoverOutputs = query
	return vq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CountryCode string `json:"country_code,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Voiceover.Query().
//		GroupBy(voiceover.FieldCountryCode).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (vq *VoiceoverQuery) GroupBy(field string, fields ...string) *VoiceoverGroupBy {
	vq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &VoiceoverGroupBy{build: vq}
	grbuild.flds = &vq.ctx.Fields
	grbuild.label = voiceover.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CountryCode string `json:"country_code,omitempty"`
//	}
//
//	client.Voiceover.Query().
//		Select(voiceover.FieldCountryCode).
//		Scan(ctx, &v)
func (vq *VoiceoverQuery) Select(fields ...string) *VoiceoverSelect {
	vq.ctx.Fields = append(vq.ctx.Fields, fields...)
	sbuild := &VoiceoverSelect{VoiceoverQuery: vq}
	sbuild.label = voiceover.Label
	sbuild.flds, sbuild.scan = &vq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a VoiceoverSelect configured with the given aggregations.
func (vq *VoiceoverQuery) Aggregate(fns ...AggregateFunc) *VoiceoverSelect {
	return vq.Select().Aggregate(fns...)
}

func (vq *VoiceoverQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range vq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, vq); err != nil {
				return err
			}
		}
	}
	for _, f := range vq.ctx.Fields {
		if !voiceover.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if vq.path != nil {
		prev, err := vq.path(ctx)
		if err != nil {
			return err
		}
		vq.sql = prev
	}
	return nil
}

func (vq *VoiceoverQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Voiceover, error) {
	var (
		nodes       = []*Voiceover{}
		_spec       = vq.querySpec()
		loadedTypes = [7]bool{
			vq.withUser != nil,
			vq.withPrompt != nil,
			vq.withDeviceInfo != nil,
			vq.withVoiceoverModels != nil,
			vq.withVoiceoverSpeakers != nil,
			vq.withAPITokens != nil,
			vq.withVoiceoverOutputs != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Voiceover).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Voiceover{config: vq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(vq.modifiers) > 0 {
		_spec.Modifiers = vq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, vq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := vq.withUser; query != nil {
		if err := vq.loadUser(ctx, query, nodes, nil,
			func(n *Voiceover, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	if query := vq.withPrompt; query != nil {
		if err := vq.loadPrompt(ctx, query, nodes, nil,
			func(n *Voiceover, e *Prompt) { n.Edges.Prompt = e }); err != nil {
			return nil, err
		}
	}
	if query := vq.withDeviceInfo; query != nil {
		if err := vq.loadDeviceInfo(ctx, query, nodes, nil,
			func(n *Voiceover, e *DeviceInfo) { n.Edges.DeviceInfo = e }); err != nil {
			return nil, err
		}
	}
	if query := vq.withVoiceoverModels; query != nil {
		if err := vq.loadVoiceoverModels(ctx, query, nodes, nil,
			func(n *Voiceover, e *VoiceoverModel) { n.Edges.VoiceoverModels = e }); err != nil {
			return nil, err
		}
	}
	if query := vq.withVoiceoverSpeakers; query != nil {
		if err := vq.loadVoiceoverSpeakers(ctx, query, nodes, nil,
			func(n *Voiceover, e *VoiceoverSpeaker) { n.Edges.VoiceoverSpeakers = e }); err != nil {
			return nil, err
		}
	}
	if query := vq.withAPITokens; query != nil {
		if err := vq.loadAPITokens(ctx, query, nodes, nil,
			func(n *Voiceover, e *ApiToken) { n.Edges.APITokens = e }); err != nil {
			return nil, err
		}
	}
	if query := vq.withVoiceoverOutputs; query != nil {
		if err := vq.loadVoiceoverOutputs(ctx, query, nodes,
			func(n *Voiceover) { n.Edges.VoiceoverOutputs = []*VoiceoverOutput{} },
			func(n *Voiceover, e *VoiceoverOutput) { n.Edges.VoiceoverOutputs = append(n.Edges.VoiceoverOutputs, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (vq *VoiceoverQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*Voiceover, init func(*Voiceover), assign func(*Voiceover, *User)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Voiceover)
	for i := range nodes {
		fk := nodes[i].UserID
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
			return fmt.Errorf(`unexpected foreign-key "user_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (vq *VoiceoverQuery) loadPrompt(ctx context.Context, query *PromptQuery, nodes []*Voiceover, init func(*Voiceover), assign func(*Voiceover, *Prompt)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Voiceover)
	for i := range nodes {
		if nodes[i].PromptID == nil {
			continue
		}
		fk := *nodes[i].PromptID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(prompt.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "prompt_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (vq *VoiceoverQuery) loadDeviceInfo(ctx context.Context, query *DeviceInfoQuery, nodes []*Voiceover, init func(*Voiceover), assign func(*Voiceover, *DeviceInfo)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Voiceover)
	for i := range nodes {
		fk := nodes[i].DeviceInfoID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(deviceinfo.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "device_info_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (vq *VoiceoverQuery) loadVoiceoverModels(ctx context.Context, query *VoiceoverModelQuery, nodes []*Voiceover, init func(*Voiceover), assign func(*Voiceover, *VoiceoverModel)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Voiceover)
	for i := range nodes {
		fk := nodes[i].ModelID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(voiceovermodel.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "model_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (vq *VoiceoverQuery) loadVoiceoverSpeakers(ctx context.Context, query *VoiceoverSpeakerQuery, nodes []*Voiceover, init func(*Voiceover), assign func(*Voiceover, *VoiceoverSpeaker)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Voiceover)
	for i := range nodes {
		fk := nodes[i].SpeakerID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(voiceoverspeaker.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "speaker_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (vq *VoiceoverQuery) loadAPITokens(ctx context.Context, query *ApiTokenQuery, nodes []*Voiceover, init func(*Voiceover), assign func(*Voiceover, *ApiToken)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Voiceover)
	for i := range nodes {
		if nodes[i].APITokenID == nil {
			continue
		}
		fk := *nodes[i].APITokenID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(apitoken.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "api_token_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (vq *VoiceoverQuery) loadVoiceoverOutputs(ctx context.Context, query *VoiceoverOutputQuery, nodes []*Voiceover, init func(*Voiceover), assign func(*Voiceover, *VoiceoverOutput)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*Voiceover)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.Where(predicate.VoiceoverOutput(func(s *sql.Selector) {
		s.Where(sql.InValues(voiceover.VoiceoverOutputsColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.VoiceoverID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "voiceover_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (vq *VoiceoverQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := vq.querySpec()
	if len(vq.modifiers) > 0 {
		_spec.Modifiers = vq.modifiers
	}
	_spec.Node.Columns = vq.ctx.Fields
	if len(vq.ctx.Fields) > 0 {
		_spec.Unique = vq.ctx.Unique != nil && *vq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, vq.driver, _spec)
}

func (vq *VoiceoverQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   voiceover.Table,
			Columns: voiceover.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: voiceover.FieldID,
			},
		},
		From:   vq.sql,
		Unique: true,
	}
	if unique := vq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := vq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, voiceover.FieldID)
		for i := range fields {
			if fields[i] != voiceover.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := vq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := vq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := vq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := vq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (vq *VoiceoverQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(vq.driver.Dialect())
	t1 := builder.Table(voiceover.Table)
	columns := vq.ctx.Fields
	if len(columns) == 0 {
		columns = voiceover.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if vq.sql != nil {
		selector = vq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if vq.ctx.Unique != nil && *vq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range vq.modifiers {
		m(selector)
	}
	for _, p := range vq.predicates {
		p(selector)
	}
	for _, p := range vq.order {
		p(selector)
	}
	if offset := vq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := vq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (vq *VoiceoverQuery) Modify(modifiers ...func(s *sql.Selector)) *VoiceoverSelect {
	vq.modifiers = append(vq.modifiers, modifiers...)
	return vq.Select()
}

// VoiceoverGroupBy is the group-by builder for Voiceover entities.
type VoiceoverGroupBy struct {
	selector
	build *VoiceoverQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (vgb *VoiceoverGroupBy) Aggregate(fns ...AggregateFunc) *VoiceoverGroupBy {
	vgb.fns = append(vgb.fns, fns...)
	return vgb
}

// Scan applies the selector query and scans the result into the given value.
func (vgb *VoiceoverGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, vgb.build.ctx, "GroupBy")
	if err := vgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*VoiceoverQuery, *VoiceoverGroupBy](ctx, vgb.build, vgb, vgb.build.inters, v)
}

func (vgb *VoiceoverGroupBy) sqlScan(ctx context.Context, root *VoiceoverQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(vgb.fns))
	for _, fn := range vgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*vgb.flds)+len(vgb.fns))
		for _, f := range *vgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*vgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := vgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// VoiceoverSelect is the builder for selecting fields of Voiceover entities.
type VoiceoverSelect struct {
	*VoiceoverQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (vs *VoiceoverSelect) Aggregate(fns ...AggregateFunc) *VoiceoverSelect {
	vs.fns = append(vs.fns, fns...)
	return vs
}

// Scan applies the selector query and scans the result into the given value.
func (vs *VoiceoverSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, vs.ctx, "Select")
	if err := vs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*VoiceoverQuery, *VoiceoverSelect](ctx, vs.VoiceoverQuery, vs, vs.inters, v)
}

func (vs *VoiceoverSelect) sqlScan(ctx context.Context, root *VoiceoverQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(vs.fns))
	for _, fn := range vs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*vs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := vs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (vs *VoiceoverSelect) Modify(modifiers ...func(s *sql.Selector)) *VoiceoverSelect {
	vs.modifiers = append(vs.modifiers, modifiers...)
	return vs
}
