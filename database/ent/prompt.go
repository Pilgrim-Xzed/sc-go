// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/stablecog/sc-go/database/ent/prompt"
)

// Prompt is the model entity for the Prompt schema.
type Prompt struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Text holds the value of the "text" field.
	Text string `json:"text,omitempty"`
	// Type holds the value of the "type" field.
	Type prompt.Type `json:"type,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PromptQuery when eager-loading is set.
	Edges PromptEdges `json:"edges"`
}

// PromptEdges holds the relations/edges for other nodes in the graph.
type PromptEdges struct {
	// Generations holds the value of the generations edge.
	Generations []*Generation `json:"generations,omitempty"`
	// Voiceovers holds the value of the voiceovers edge.
	Voiceovers []*Voiceover `json:"voiceovers,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// GenerationsOrErr returns the Generations value or an error if the edge
// was not loaded in eager-loading.
func (e PromptEdges) GenerationsOrErr() ([]*Generation, error) {
	if e.loadedTypes[0] {
		return e.Generations, nil
	}
	return nil, &NotLoadedError{edge: "generations"}
}

// VoiceoversOrErr returns the Voiceovers value or an error if the edge
// was not loaded in eager-loading.
func (e PromptEdges) VoiceoversOrErr() ([]*Voiceover, error) {
	if e.loadedTypes[1] {
		return e.Voiceovers, nil
	}
	return nil, &NotLoadedError{edge: "voiceovers"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Prompt) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case prompt.FieldText, prompt.FieldType:
			values[i] = new(sql.NullString)
		case prompt.FieldCreatedAt, prompt.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case prompt.FieldID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Prompt", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Prompt fields.
func (pr *Prompt) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case prompt.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				pr.ID = *value
			}
		case prompt.FieldText:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field text", values[i])
			} else if value.Valid {
				pr.Text = value.String
			}
		case prompt.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				pr.Type = prompt.Type(value.String)
			}
		case prompt.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				pr.CreatedAt = value.Time
			}
		case prompt.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				pr.UpdatedAt = value.Time
			}
		}
	}
	return nil
}

// QueryGenerations queries the "generations" edge of the Prompt entity.
func (pr *Prompt) QueryGenerations() *GenerationQuery {
	return NewPromptClient(pr.config).QueryGenerations(pr)
}

// QueryVoiceovers queries the "voiceovers" edge of the Prompt entity.
func (pr *Prompt) QueryVoiceovers() *VoiceoverQuery {
	return NewPromptClient(pr.config).QueryVoiceovers(pr)
}

// Update returns a builder for updating this Prompt.
// Note that you need to call Prompt.Unwrap() before calling this method if this Prompt
// was returned from a transaction, and the transaction was committed or rolled back.
func (pr *Prompt) Update() *PromptUpdateOne {
	return NewPromptClient(pr.config).UpdateOne(pr)
}

// Unwrap unwraps the Prompt entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pr *Prompt) Unwrap() *Prompt {
	_tx, ok := pr.config.driver.(*txDriver)
	if !ok {
		panic("ent: Prompt is not a transactional entity")
	}
	pr.config.driver = _tx.drv
	return pr
}

// String implements the fmt.Stringer.
func (pr *Prompt) String() string {
	var builder strings.Builder
	builder.WriteString("Prompt(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pr.ID))
	builder.WriteString("text=")
	builder.WriteString(pr.Text)
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(fmt.Sprintf("%v", pr.Type))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(pr.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(pr.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Prompts is a parsable slice of Prompt.
type Prompts []*Prompt

func (pr Prompts) config(cfg config) {
	for _i := range pr {
		pr[_i].config = cfg
	}
}
