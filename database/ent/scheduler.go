// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/stablecog/sc-go/database/ent/scheduler"
)

// Scheduler is the model entity for the Scheduler schema.
type Scheduler struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// NameInWorker holds the value of the "name_in_worker" field.
	NameInWorker string `json:"name_in_worker,omitempty"`
	// IsActive holds the value of the "is_active" field.
	IsActive bool `json:"is_active,omitempty"`
	// IsDefault holds the value of the "is_default" field.
	IsDefault bool `json:"is_default,omitempty"`
	// IsHidden holds the value of the "is_hidden" field.
	IsHidden bool `json:"is_hidden,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SchedulerQuery when eager-loading is set.
	Edges SchedulerEdges `json:"edges"`
}

// SchedulerEdges holds the relations/edges for other nodes in the graph.
type SchedulerEdges struct {
	// Generations holds the value of the generations edge.
	Generations []*Generation `json:"generations,omitempty"`
	// GenerationModels holds the value of the generation_models edge.
	GenerationModels []*GenerationModel `json:"generation_models,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// GenerationsOrErr returns the Generations value or an error if the edge
// was not loaded in eager-loading.
func (e SchedulerEdges) GenerationsOrErr() ([]*Generation, error) {
	if e.loadedTypes[0] {
		return e.Generations, nil
	}
	return nil, &NotLoadedError{edge: "generations"}
}

// GenerationModelsOrErr returns the GenerationModels value or an error if the edge
// was not loaded in eager-loading.
func (e SchedulerEdges) GenerationModelsOrErr() ([]*GenerationModel, error) {
	if e.loadedTypes[1] {
		return e.GenerationModels, nil
	}
	return nil, &NotLoadedError{edge: "generation_models"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Scheduler) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case scheduler.FieldIsActive, scheduler.FieldIsDefault, scheduler.FieldIsHidden:
			values[i] = new(sql.NullBool)
		case scheduler.FieldNameInWorker:
			values[i] = new(sql.NullString)
		case scheduler.FieldCreatedAt, scheduler.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case scheduler.FieldID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Scheduler", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Scheduler fields.
func (s *Scheduler) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case scheduler.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				s.ID = *value
			}
		case scheduler.FieldNameInWorker:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name_in_worker", values[i])
			} else if value.Valid {
				s.NameInWorker = value.String
			}
		case scheduler.FieldIsActive:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_active", values[i])
			} else if value.Valid {
				s.IsActive = value.Bool
			}
		case scheduler.FieldIsDefault:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_default", values[i])
			} else if value.Valid {
				s.IsDefault = value.Bool
			}
		case scheduler.FieldIsHidden:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_hidden", values[i])
			} else if value.Valid {
				s.IsHidden = value.Bool
			}
		case scheduler.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				s.CreatedAt = value.Time
			}
		case scheduler.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				s.UpdatedAt = value.Time
			}
		}
	}
	return nil
}

// QueryGenerations queries the "generations" edge of the Scheduler entity.
func (s *Scheduler) QueryGenerations() *GenerationQuery {
	return NewSchedulerClient(s.config).QueryGenerations(s)
}

// QueryGenerationModels queries the "generation_models" edge of the Scheduler entity.
func (s *Scheduler) QueryGenerationModels() *GenerationModelQuery {
	return NewSchedulerClient(s.config).QueryGenerationModels(s)
}

// Update returns a builder for updating this Scheduler.
// Note that you need to call Scheduler.Unwrap() before calling this method if this Scheduler
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Scheduler) Update() *SchedulerUpdateOne {
	return NewSchedulerClient(s.config).UpdateOne(s)
}

// Unwrap unwraps the Scheduler entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Scheduler) Unwrap() *Scheduler {
	_tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Scheduler is not a transactional entity")
	}
	s.config.driver = _tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Scheduler) String() string {
	var builder strings.Builder
	builder.WriteString("Scheduler(")
	builder.WriteString(fmt.Sprintf("id=%v, ", s.ID))
	builder.WriteString("name_in_worker=")
	builder.WriteString(s.NameInWorker)
	builder.WriteString(", ")
	builder.WriteString("is_active=")
	builder.WriteString(fmt.Sprintf("%v", s.IsActive))
	builder.WriteString(", ")
	builder.WriteString("is_default=")
	builder.WriteString(fmt.Sprintf("%v", s.IsDefault))
	builder.WriteString(", ")
	builder.WriteString("is_hidden=")
	builder.WriteString(fmt.Sprintf("%v", s.IsHidden))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(s.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(s.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Schedulers is a parsable slice of Scheduler.
type Schedulers []*Scheduler

func (s Schedulers) config(cfg config) {
	for _i := range s {
		s[_i].config = cfg
	}
}
