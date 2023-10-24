// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/stablecog/sc-go/database/ent/ipblacklist"
	"github.com/stablecog/sc-go/database/ent/predicate"
)

// IPBlackListUpdate is the builder for updating IPBlackList entities.
type IPBlackListUpdate struct {
	config
	hooks     []Hook
	mutation  *IPBlackListMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the IPBlackListUpdate builder.
func (iblu *IPBlackListUpdate) Where(ps ...predicate.IPBlackList) *IPBlackListUpdate {
	iblu.mutation.Where(ps...)
	return iblu
}

// SetIP sets the "ip" field.
func (iblu *IPBlackListUpdate) SetIP(s string) *IPBlackListUpdate {
	iblu.mutation.SetIP(s)
	return iblu
}

// SetUpdatedAt sets the "updated_at" field.
func (iblu *IPBlackListUpdate) SetUpdatedAt(t time.Time) *IPBlackListUpdate {
	iblu.mutation.SetUpdatedAt(t)
	return iblu
}

// Mutation returns the IPBlackListMutation object of the builder.
func (iblu *IPBlackListUpdate) Mutation() *IPBlackListMutation {
	return iblu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (iblu *IPBlackListUpdate) Save(ctx context.Context) (int, error) {
	iblu.defaults()
	return withHooks[int, IPBlackListMutation](ctx, iblu.sqlSave, iblu.mutation, iblu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (iblu *IPBlackListUpdate) SaveX(ctx context.Context) int {
	affected, err := iblu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (iblu *IPBlackListUpdate) Exec(ctx context.Context) error {
	_, err := iblu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iblu *IPBlackListUpdate) ExecX(ctx context.Context) {
	if err := iblu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (iblu *IPBlackListUpdate) defaults() {
	if _, ok := iblu.mutation.UpdatedAt(); !ok {
		v := ipblacklist.UpdateDefaultUpdatedAt()
		iblu.mutation.SetUpdatedAt(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (iblu *IPBlackListUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *IPBlackListUpdate {
	iblu.modifiers = append(iblu.modifiers, modifiers...)
	return iblu
}

func (iblu *IPBlackListUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   ipblacklist.Table,
			Columns: ipblacklist.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: ipblacklist.FieldID,
			},
		},
	}
	if ps := iblu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := iblu.mutation.IP(); ok {
		_spec.SetField(ipblacklist.FieldIP, field.TypeString, value)
	}
	if value, ok := iblu.mutation.UpdatedAt(); ok {
		_spec.SetField(ipblacklist.FieldUpdatedAt, field.TypeTime, value)
	}
	_spec.AddModifiers(iblu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, iblu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{ipblacklist.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	iblu.mutation.done = true
	return n, nil
}

// IPBlackListUpdateOne is the builder for updating a single IPBlackList entity.
type IPBlackListUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *IPBlackListMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetIP sets the "ip" field.
func (ibluo *IPBlackListUpdateOne) SetIP(s string) *IPBlackListUpdateOne {
	ibluo.mutation.SetIP(s)
	return ibluo
}

// SetUpdatedAt sets the "updated_at" field.
func (ibluo *IPBlackListUpdateOne) SetUpdatedAt(t time.Time) *IPBlackListUpdateOne {
	ibluo.mutation.SetUpdatedAt(t)
	return ibluo
}

// Mutation returns the IPBlackListMutation object of the builder.
func (ibluo *IPBlackListUpdateOne) Mutation() *IPBlackListMutation {
	return ibluo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ibluo *IPBlackListUpdateOne) Select(field string, fields ...string) *IPBlackListUpdateOne {
	ibluo.fields = append([]string{field}, fields...)
	return ibluo
}

// Save executes the query and returns the updated IPBlackList entity.
func (ibluo *IPBlackListUpdateOne) Save(ctx context.Context) (*IPBlackList, error) {
	ibluo.defaults()
	return withHooks[*IPBlackList, IPBlackListMutation](ctx, ibluo.sqlSave, ibluo.mutation, ibluo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ibluo *IPBlackListUpdateOne) SaveX(ctx context.Context) *IPBlackList {
	node, err := ibluo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ibluo *IPBlackListUpdateOne) Exec(ctx context.Context) error {
	_, err := ibluo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ibluo *IPBlackListUpdateOne) ExecX(ctx context.Context) {
	if err := ibluo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ibluo *IPBlackListUpdateOne) defaults() {
	if _, ok := ibluo.mutation.UpdatedAt(); !ok {
		v := ipblacklist.UpdateDefaultUpdatedAt()
		ibluo.mutation.SetUpdatedAt(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (ibluo *IPBlackListUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *IPBlackListUpdateOne {
	ibluo.modifiers = append(ibluo.modifiers, modifiers...)
	return ibluo
}

func (ibluo *IPBlackListUpdateOne) sqlSave(ctx context.Context) (_node *IPBlackList, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   ipblacklist.Table,
			Columns: ipblacklist.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: ipblacklist.FieldID,
			},
		},
	}
	id, ok := ibluo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "IPBlackList.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ibluo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, ipblacklist.FieldID)
		for _, f := range fields {
			if !ipblacklist.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != ipblacklist.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ibluo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ibluo.mutation.IP(); ok {
		_spec.SetField(ipblacklist.FieldIP, field.TypeString, value)
	}
	if value, ok := ibluo.mutation.UpdatedAt(); ok {
		_spec.SetField(ipblacklist.FieldUpdatedAt, field.TypeTime, value)
	}
	_spec.AddModifiers(ibluo.modifiers...)
	_node = &IPBlackList{config: ibluo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ibluo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{ipblacklist.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ibluo.mutation.done = true
	return _node, nil
}
