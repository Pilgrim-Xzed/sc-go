// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/stablecog/sc-go/database/ent/generationoutput"
	"github.com/stablecog/sc-go/database/ent/upscale"
	"github.com/stablecog/sc-go/database/ent/upscaleoutput"
)

// UpscaleOutputCreate is the builder for creating a UpscaleOutput entity.
type UpscaleOutputCreate struct {
	config
	mutation *UpscaleOutputMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetImagePath sets the "image_path" field.
func (uoc *UpscaleOutputCreate) SetImagePath(s string) *UpscaleOutputCreate {
	uoc.mutation.SetImagePath(s)
	return uoc
}

// SetInputImageURL sets the "input_image_url" field.
func (uoc *UpscaleOutputCreate) SetInputImageURL(s string) *UpscaleOutputCreate {
	uoc.mutation.SetInputImageURL(s)
	return uoc
}

// SetNillableInputImageURL sets the "input_image_url" field if the given value is not nil.
func (uoc *UpscaleOutputCreate) SetNillableInputImageURL(s *string) *UpscaleOutputCreate {
	if s != nil {
		uoc.SetInputImageURL(*s)
	}
	return uoc
}

// SetUpscaleID sets the "upscale_id" field.
func (uoc *UpscaleOutputCreate) SetUpscaleID(u uuid.UUID) *UpscaleOutputCreate {
	uoc.mutation.SetUpscaleID(u)
	return uoc
}

// SetGenerationOutputID sets the "generation_output_id" field.
func (uoc *UpscaleOutputCreate) SetGenerationOutputID(u uuid.UUID) *UpscaleOutputCreate {
	uoc.mutation.SetGenerationOutputID(u)
	return uoc
}

// SetNillableGenerationOutputID sets the "generation_output_id" field if the given value is not nil.
func (uoc *UpscaleOutputCreate) SetNillableGenerationOutputID(u *uuid.UUID) *UpscaleOutputCreate {
	if u != nil {
		uoc.SetGenerationOutputID(*u)
	}
	return uoc
}

// SetDeletedAt sets the "deleted_at" field.
func (uoc *UpscaleOutputCreate) SetDeletedAt(t time.Time) *UpscaleOutputCreate {
	uoc.mutation.SetDeletedAt(t)
	return uoc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (uoc *UpscaleOutputCreate) SetNillableDeletedAt(t *time.Time) *UpscaleOutputCreate {
	if t != nil {
		uoc.SetDeletedAt(*t)
	}
	return uoc
}

// SetCreatedAt sets the "created_at" field.
func (uoc *UpscaleOutputCreate) SetCreatedAt(t time.Time) *UpscaleOutputCreate {
	uoc.mutation.SetCreatedAt(t)
	return uoc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (uoc *UpscaleOutputCreate) SetNillableCreatedAt(t *time.Time) *UpscaleOutputCreate {
	if t != nil {
		uoc.SetCreatedAt(*t)
	}
	return uoc
}

// SetUpdatedAt sets the "updated_at" field.
func (uoc *UpscaleOutputCreate) SetUpdatedAt(t time.Time) *UpscaleOutputCreate {
	uoc.mutation.SetUpdatedAt(t)
	return uoc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (uoc *UpscaleOutputCreate) SetNillableUpdatedAt(t *time.Time) *UpscaleOutputCreate {
	if t != nil {
		uoc.SetUpdatedAt(*t)
	}
	return uoc
}

// SetID sets the "id" field.
func (uoc *UpscaleOutputCreate) SetID(u uuid.UUID) *UpscaleOutputCreate {
	uoc.mutation.SetID(u)
	return uoc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (uoc *UpscaleOutputCreate) SetNillableID(u *uuid.UUID) *UpscaleOutputCreate {
	if u != nil {
		uoc.SetID(*u)
	}
	return uoc
}

// SetUpscalesID sets the "upscales" edge to the Upscale entity by ID.
func (uoc *UpscaleOutputCreate) SetUpscalesID(id uuid.UUID) *UpscaleOutputCreate {
	uoc.mutation.SetUpscalesID(id)
	return uoc
}

// SetUpscales sets the "upscales" edge to the Upscale entity.
func (uoc *UpscaleOutputCreate) SetUpscales(u *Upscale) *UpscaleOutputCreate {
	return uoc.SetUpscalesID(u.ID)
}

// SetGenerationOutput sets the "generation_output" edge to the GenerationOutput entity.
func (uoc *UpscaleOutputCreate) SetGenerationOutput(g *GenerationOutput) *UpscaleOutputCreate {
	return uoc.SetGenerationOutputID(g.ID)
}

// Mutation returns the UpscaleOutputMutation object of the builder.
func (uoc *UpscaleOutputCreate) Mutation() *UpscaleOutputMutation {
	return uoc.mutation
}

// Save creates the UpscaleOutput in the database.
func (uoc *UpscaleOutputCreate) Save(ctx context.Context) (*UpscaleOutput, error) {
	uoc.defaults()
	return withHooks[*UpscaleOutput, UpscaleOutputMutation](ctx, uoc.sqlSave, uoc.mutation, uoc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (uoc *UpscaleOutputCreate) SaveX(ctx context.Context) *UpscaleOutput {
	v, err := uoc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uoc *UpscaleOutputCreate) Exec(ctx context.Context) error {
	_, err := uoc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uoc *UpscaleOutputCreate) ExecX(ctx context.Context) {
	if err := uoc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uoc *UpscaleOutputCreate) defaults() {
	if _, ok := uoc.mutation.CreatedAt(); !ok {
		v := upscaleoutput.DefaultCreatedAt()
		uoc.mutation.SetCreatedAt(v)
	}
	if _, ok := uoc.mutation.UpdatedAt(); !ok {
		v := upscaleoutput.DefaultUpdatedAt()
		uoc.mutation.SetUpdatedAt(v)
	}
	if _, ok := uoc.mutation.ID(); !ok {
		v := upscaleoutput.DefaultID()
		uoc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uoc *UpscaleOutputCreate) check() error {
	if _, ok := uoc.mutation.ImagePath(); !ok {
		return &ValidationError{Name: "image_path", err: errors.New(`ent: missing required field "UpscaleOutput.image_path"`)}
	}
	if _, ok := uoc.mutation.UpscaleID(); !ok {
		return &ValidationError{Name: "upscale_id", err: errors.New(`ent: missing required field "UpscaleOutput.upscale_id"`)}
	}
	if _, ok := uoc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "UpscaleOutput.created_at"`)}
	}
	if _, ok := uoc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "UpscaleOutput.updated_at"`)}
	}
	if _, ok := uoc.mutation.UpscalesID(); !ok {
		return &ValidationError{Name: "upscales", err: errors.New(`ent: missing required edge "UpscaleOutput.upscales"`)}
	}
	return nil
}

func (uoc *UpscaleOutputCreate) sqlSave(ctx context.Context) (*UpscaleOutput, error) {
	if err := uoc.check(); err != nil {
		return nil, err
	}
	_node, _spec := uoc.createSpec()
	if err := sqlgraph.CreateNode(ctx, uoc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	uoc.mutation.id = &_node.ID
	uoc.mutation.done = true
	return _node, nil
}

func (uoc *UpscaleOutputCreate) createSpec() (*UpscaleOutput, *sqlgraph.CreateSpec) {
	var (
		_node = &UpscaleOutput{config: uoc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: upscaleoutput.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: upscaleoutput.FieldID,
			},
		}
	)
	_spec.OnConflict = uoc.conflict
	if id, ok := uoc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := uoc.mutation.ImagePath(); ok {
		_spec.SetField(upscaleoutput.FieldImagePath, field.TypeString, value)
		_node.ImagePath = value
	}
	if value, ok := uoc.mutation.InputImageURL(); ok {
		_spec.SetField(upscaleoutput.FieldInputImageURL, field.TypeString, value)
		_node.InputImageURL = &value
	}
	if value, ok := uoc.mutation.DeletedAt(); ok {
		_spec.SetField(upscaleoutput.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = &value
	}
	if value, ok := uoc.mutation.CreatedAt(); ok {
		_spec.SetField(upscaleoutput.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := uoc.mutation.UpdatedAt(); ok {
		_spec.SetField(upscaleoutput.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := uoc.mutation.UpscalesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   upscaleoutput.UpscalesTable,
			Columns: []string{upscaleoutput.UpscalesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: upscale.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.UpscaleID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uoc.mutation.GenerationOutputIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   upscaleoutput.GenerationOutputTable,
			Columns: []string{upscaleoutput.GenerationOutputColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: generationoutput.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.GenerationOutputID = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.UpscaleOutput.Create().
//		SetImagePath(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.UpscaleOutputUpsert) {
//			SetImagePath(v+v).
//		}).
//		Exec(ctx)
func (uoc *UpscaleOutputCreate) OnConflict(opts ...sql.ConflictOption) *UpscaleOutputUpsertOne {
	uoc.conflict = opts
	return &UpscaleOutputUpsertOne{
		create: uoc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.UpscaleOutput.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (uoc *UpscaleOutputCreate) OnConflictColumns(columns ...string) *UpscaleOutputUpsertOne {
	uoc.conflict = append(uoc.conflict, sql.ConflictColumns(columns...))
	return &UpscaleOutputUpsertOne{
		create: uoc,
	}
}

type (
	// UpscaleOutputUpsertOne is the builder for "upsert"-ing
	//  one UpscaleOutput node.
	UpscaleOutputUpsertOne struct {
		create *UpscaleOutputCreate
	}

	// UpscaleOutputUpsert is the "OnConflict" setter.
	UpscaleOutputUpsert struct {
		*sql.UpdateSet
	}
)

// SetImagePath sets the "image_path" field.
func (u *UpscaleOutputUpsert) SetImagePath(v string) *UpscaleOutputUpsert {
	u.Set(upscaleoutput.FieldImagePath, v)
	return u
}

// UpdateImagePath sets the "image_path" field to the value that was provided on create.
func (u *UpscaleOutputUpsert) UpdateImagePath() *UpscaleOutputUpsert {
	u.SetExcluded(upscaleoutput.FieldImagePath)
	return u
}

// SetInputImageURL sets the "input_image_url" field.
func (u *UpscaleOutputUpsert) SetInputImageURL(v string) *UpscaleOutputUpsert {
	u.Set(upscaleoutput.FieldInputImageURL, v)
	return u
}

// UpdateInputImageURL sets the "input_image_url" field to the value that was provided on create.
func (u *UpscaleOutputUpsert) UpdateInputImageURL() *UpscaleOutputUpsert {
	u.SetExcluded(upscaleoutput.FieldInputImageURL)
	return u
}

// ClearInputImageURL clears the value of the "input_image_url" field.
func (u *UpscaleOutputUpsert) ClearInputImageURL() *UpscaleOutputUpsert {
	u.SetNull(upscaleoutput.FieldInputImageURL)
	return u
}

// SetUpscaleID sets the "upscale_id" field.
func (u *UpscaleOutputUpsert) SetUpscaleID(v uuid.UUID) *UpscaleOutputUpsert {
	u.Set(upscaleoutput.FieldUpscaleID, v)
	return u
}

// UpdateUpscaleID sets the "upscale_id" field to the value that was provided on create.
func (u *UpscaleOutputUpsert) UpdateUpscaleID() *UpscaleOutputUpsert {
	u.SetExcluded(upscaleoutput.FieldUpscaleID)
	return u
}

// SetGenerationOutputID sets the "generation_output_id" field.
func (u *UpscaleOutputUpsert) SetGenerationOutputID(v uuid.UUID) *UpscaleOutputUpsert {
	u.Set(upscaleoutput.FieldGenerationOutputID, v)
	return u
}

// UpdateGenerationOutputID sets the "generation_output_id" field to the value that was provided on create.
func (u *UpscaleOutputUpsert) UpdateGenerationOutputID() *UpscaleOutputUpsert {
	u.SetExcluded(upscaleoutput.FieldGenerationOutputID)
	return u
}

// ClearGenerationOutputID clears the value of the "generation_output_id" field.
func (u *UpscaleOutputUpsert) ClearGenerationOutputID() *UpscaleOutputUpsert {
	u.SetNull(upscaleoutput.FieldGenerationOutputID)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *UpscaleOutputUpsert) SetDeletedAt(v time.Time) *UpscaleOutputUpsert {
	u.Set(upscaleoutput.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *UpscaleOutputUpsert) UpdateDeletedAt() *UpscaleOutputUpsert {
	u.SetExcluded(upscaleoutput.FieldDeletedAt)
	return u
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *UpscaleOutputUpsert) ClearDeletedAt() *UpscaleOutputUpsert {
	u.SetNull(upscaleoutput.FieldDeletedAt)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *UpscaleOutputUpsert) SetUpdatedAt(v time.Time) *UpscaleOutputUpsert {
	u.Set(upscaleoutput.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *UpscaleOutputUpsert) UpdateUpdatedAt() *UpscaleOutputUpsert {
	u.SetExcluded(upscaleoutput.FieldUpdatedAt)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.UpscaleOutput.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(upscaleoutput.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *UpscaleOutputUpsertOne) UpdateNewValues() *UpscaleOutputUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(upscaleoutput.FieldID)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(upscaleoutput.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.UpscaleOutput.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *UpscaleOutputUpsertOne) Ignore() *UpscaleOutputUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *UpscaleOutputUpsertOne) DoNothing() *UpscaleOutputUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the UpscaleOutputCreate.OnConflict
// documentation for more info.
func (u *UpscaleOutputUpsertOne) Update(set func(*UpscaleOutputUpsert)) *UpscaleOutputUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&UpscaleOutputUpsert{UpdateSet: update})
	}))
	return u
}

// SetImagePath sets the "image_path" field.
func (u *UpscaleOutputUpsertOne) SetImagePath(v string) *UpscaleOutputUpsertOne {
	return u.Update(func(s *UpscaleOutputUpsert) {
		s.SetImagePath(v)
	})
}

// UpdateImagePath sets the "image_path" field to the value that was provided on create.
func (u *UpscaleOutputUpsertOne) UpdateImagePath() *UpscaleOutputUpsertOne {
	return u.Update(func(s *UpscaleOutputUpsert) {
		s.UpdateImagePath()
	})
}

// SetInputImageURL sets the "input_image_url" field.
func (u *UpscaleOutputUpsertOne) SetInputImageURL(v string) *UpscaleOutputUpsertOne {
	return u.Update(func(s *UpscaleOutputUpsert) {
		s.SetInputImageURL(v)
	})
}

// UpdateInputImageURL sets the "input_image_url" field to the value that was provided on create.
func (u *UpscaleOutputUpsertOne) UpdateInputImageURL() *UpscaleOutputUpsertOne {
	return u.Update(func(s *UpscaleOutputUpsert) {
		s.UpdateInputImageURL()
	})
}

// ClearInputImageURL clears the value of the "input_image_url" field.
func (u *UpscaleOutputUpsertOne) ClearInputImageURL() *UpscaleOutputUpsertOne {
	return u.Update(func(s *UpscaleOutputUpsert) {
		s.ClearInputImageURL()
	})
}

// SetUpscaleID sets the "upscale_id" field.
func (u *UpscaleOutputUpsertOne) SetUpscaleID(v uuid.UUID) *UpscaleOutputUpsertOne {
	return u.Update(func(s *UpscaleOutputUpsert) {
		s.SetUpscaleID(v)
	})
}

// UpdateUpscaleID sets the "upscale_id" field to the value that was provided on create.
func (u *UpscaleOutputUpsertOne) UpdateUpscaleID() *UpscaleOutputUpsertOne {
	return u.Update(func(s *UpscaleOutputUpsert) {
		s.UpdateUpscaleID()
	})
}

// SetGenerationOutputID sets the "generation_output_id" field.
func (u *UpscaleOutputUpsertOne) SetGenerationOutputID(v uuid.UUID) *UpscaleOutputUpsertOne {
	return u.Update(func(s *UpscaleOutputUpsert) {
		s.SetGenerationOutputID(v)
	})
}

// UpdateGenerationOutputID sets the "generation_output_id" field to the value that was provided on create.
func (u *UpscaleOutputUpsertOne) UpdateGenerationOutputID() *UpscaleOutputUpsertOne {
	return u.Update(func(s *UpscaleOutputUpsert) {
		s.UpdateGenerationOutputID()
	})
}

// ClearGenerationOutputID clears the value of the "generation_output_id" field.
func (u *UpscaleOutputUpsertOne) ClearGenerationOutputID() *UpscaleOutputUpsertOne {
	return u.Update(func(s *UpscaleOutputUpsert) {
		s.ClearGenerationOutputID()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *UpscaleOutputUpsertOne) SetDeletedAt(v time.Time) *UpscaleOutputUpsertOne {
	return u.Update(func(s *UpscaleOutputUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *UpscaleOutputUpsertOne) UpdateDeletedAt() *UpscaleOutputUpsertOne {
	return u.Update(func(s *UpscaleOutputUpsert) {
		s.UpdateDeletedAt()
	})
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *UpscaleOutputUpsertOne) ClearDeletedAt() *UpscaleOutputUpsertOne {
	return u.Update(func(s *UpscaleOutputUpsert) {
		s.ClearDeletedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *UpscaleOutputUpsertOne) SetUpdatedAt(v time.Time) *UpscaleOutputUpsertOne {
	return u.Update(func(s *UpscaleOutputUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *UpscaleOutputUpsertOne) UpdateUpdatedAt() *UpscaleOutputUpsertOne {
	return u.Update(func(s *UpscaleOutputUpsert) {
		s.UpdateUpdatedAt()
	})
}

// Exec executes the query.
func (u *UpscaleOutputUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for UpscaleOutputCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *UpscaleOutputUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *UpscaleOutputUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: UpscaleOutputUpsertOne.ID is not supported by MySQL driver. Use UpscaleOutputUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *UpscaleOutputUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// UpscaleOutputCreateBulk is the builder for creating many UpscaleOutput entities in bulk.
type UpscaleOutputCreateBulk struct {
	config
	builders []*UpscaleOutputCreate
	conflict []sql.ConflictOption
}

// Save creates the UpscaleOutput entities in the database.
func (uocb *UpscaleOutputCreateBulk) Save(ctx context.Context) ([]*UpscaleOutput, error) {
	specs := make([]*sqlgraph.CreateSpec, len(uocb.builders))
	nodes := make([]*UpscaleOutput, len(uocb.builders))
	mutators := make([]Mutator, len(uocb.builders))
	for i := range uocb.builders {
		func(i int, root context.Context) {
			builder := uocb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UpscaleOutputMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, uocb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = uocb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, uocb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, uocb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (uocb *UpscaleOutputCreateBulk) SaveX(ctx context.Context) []*UpscaleOutput {
	v, err := uocb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uocb *UpscaleOutputCreateBulk) Exec(ctx context.Context) error {
	_, err := uocb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uocb *UpscaleOutputCreateBulk) ExecX(ctx context.Context) {
	if err := uocb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.UpscaleOutput.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.UpscaleOutputUpsert) {
//			SetImagePath(v+v).
//		}).
//		Exec(ctx)
func (uocb *UpscaleOutputCreateBulk) OnConflict(opts ...sql.ConflictOption) *UpscaleOutputUpsertBulk {
	uocb.conflict = opts
	return &UpscaleOutputUpsertBulk{
		create: uocb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.UpscaleOutput.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (uocb *UpscaleOutputCreateBulk) OnConflictColumns(columns ...string) *UpscaleOutputUpsertBulk {
	uocb.conflict = append(uocb.conflict, sql.ConflictColumns(columns...))
	return &UpscaleOutputUpsertBulk{
		create: uocb,
	}
}

// UpscaleOutputUpsertBulk is the builder for "upsert"-ing
// a bulk of UpscaleOutput nodes.
type UpscaleOutputUpsertBulk struct {
	create *UpscaleOutputCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.UpscaleOutput.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(upscaleoutput.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *UpscaleOutputUpsertBulk) UpdateNewValues() *UpscaleOutputUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(upscaleoutput.FieldID)
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(upscaleoutput.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.UpscaleOutput.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *UpscaleOutputUpsertBulk) Ignore() *UpscaleOutputUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *UpscaleOutputUpsertBulk) DoNothing() *UpscaleOutputUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the UpscaleOutputCreateBulk.OnConflict
// documentation for more info.
func (u *UpscaleOutputUpsertBulk) Update(set func(*UpscaleOutputUpsert)) *UpscaleOutputUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&UpscaleOutputUpsert{UpdateSet: update})
	}))
	return u
}

// SetImagePath sets the "image_path" field.
func (u *UpscaleOutputUpsertBulk) SetImagePath(v string) *UpscaleOutputUpsertBulk {
	return u.Update(func(s *UpscaleOutputUpsert) {
		s.SetImagePath(v)
	})
}

// UpdateImagePath sets the "image_path" field to the value that was provided on create.
func (u *UpscaleOutputUpsertBulk) UpdateImagePath() *UpscaleOutputUpsertBulk {
	return u.Update(func(s *UpscaleOutputUpsert) {
		s.UpdateImagePath()
	})
}

// SetInputImageURL sets the "input_image_url" field.
func (u *UpscaleOutputUpsertBulk) SetInputImageURL(v string) *UpscaleOutputUpsertBulk {
	return u.Update(func(s *UpscaleOutputUpsert) {
		s.SetInputImageURL(v)
	})
}

// UpdateInputImageURL sets the "input_image_url" field to the value that was provided on create.
func (u *UpscaleOutputUpsertBulk) UpdateInputImageURL() *UpscaleOutputUpsertBulk {
	return u.Update(func(s *UpscaleOutputUpsert) {
		s.UpdateInputImageURL()
	})
}

// ClearInputImageURL clears the value of the "input_image_url" field.
func (u *UpscaleOutputUpsertBulk) ClearInputImageURL() *UpscaleOutputUpsertBulk {
	return u.Update(func(s *UpscaleOutputUpsert) {
		s.ClearInputImageURL()
	})
}

// SetUpscaleID sets the "upscale_id" field.
func (u *UpscaleOutputUpsertBulk) SetUpscaleID(v uuid.UUID) *UpscaleOutputUpsertBulk {
	return u.Update(func(s *UpscaleOutputUpsert) {
		s.SetUpscaleID(v)
	})
}

// UpdateUpscaleID sets the "upscale_id" field to the value that was provided on create.
func (u *UpscaleOutputUpsertBulk) UpdateUpscaleID() *UpscaleOutputUpsertBulk {
	return u.Update(func(s *UpscaleOutputUpsert) {
		s.UpdateUpscaleID()
	})
}

// SetGenerationOutputID sets the "generation_output_id" field.
func (u *UpscaleOutputUpsertBulk) SetGenerationOutputID(v uuid.UUID) *UpscaleOutputUpsertBulk {
	return u.Update(func(s *UpscaleOutputUpsert) {
		s.SetGenerationOutputID(v)
	})
}

// UpdateGenerationOutputID sets the "generation_output_id" field to the value that was provided on create.
func (u *UpscaleOutputUpsertBulk) UpdateGenerationOutputID() *UpscaleOutputUpsertBulk {
	return u.Update(func(s *UpscaleOutputUpsert) {
		s.UpdateGenerationOutputID()
	})
}

// ClearGenerationOutputID clears the value of the "generation_output_id" field.
func (u *UpscaleOutputUpsertBulk) ClearGenerationOutputID() *UpscaleOutputUpsertBulk {
	return u.Update(func(s *UpscaleOutputUpsert) {
		s.ClearGenerationOutputID()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *UpscaleOutputUpsertBulk) SetDeletedAt(v time.Time) *UpscaleOutputUpsertBulk {
	return u.Update(func(s *UpscaleOutputUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *UpscaleOutputUpsertBulk) UpdateDeletedAt() *UpscaleOutputUpsertBulk {
	return u.Update(func(s *UpscaleOutputUpsert) {
		s.UpdateDeletedAt()
	})
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *UpscaleOutputUpsertBulk) ClearDeletedAt() *UpscaleOutputUpsertBulk {
	return u.Update(func(s *UpscaleOutputUpsert) {
		s.ClearDeletedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *UpscaleOutputUpsertBulk) SetUpdatedAt(v time.Time) *UpscaleOutputUpsertBulk {
	return u.Update(func(s *UpscaleOutputUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *UpscaleOutputUpsertBulk) UpdateUpdatedAt() *UpscaleOutputUpsertBulk {
	return u.Update(func(s *UpscaleOutputUpsert) {
		s.UpdateUpdatedAt()
	})
}

// Exec executes the query.
func (u *UpscaleOutputUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the UpscaleOutputCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for UpscaleOutputCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *UpscaleOutputUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
