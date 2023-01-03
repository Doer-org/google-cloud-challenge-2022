// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/Doer-org/google-cloud-challenge-2022/infrastructure/ent/etype"
	"github.com/Doer-org/google-cloud-challenge-2022/infrastructure/ent/event"
	"github.com/google/uuid"
)

// EType is the model entity for the EType schema.
type EType struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ETypeQuery when eager-loading is set.
	Edges      ETypeEdges `json:"edges"`
	event_type *uuid.UUID
}

// ETypeEdges holds the relations/edges for other nodes in the graph.
type ETypeEdges struct {
	// Event holds the value of the event edge.
	Event *Event `json:"event,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// EventOrErr returns the Event value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ETypeEdges) EventOrErr() (*Event, error) {
	if e.loadedTypes[0] {
		if e.Event == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: event.Label}
		}
		return e.Event, nil
	}
	return nil, &NotLoadedError{edge: "event"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*EType) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case etype.FieldName:
			values[i] = new(sql.NullString)
		case etype.FieldID:
			values[i] = new(uuid.UUID)
		case etype.ForeignKeys[0]: // event_type
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			return nil, fmt.Errorf("unexpected column %q for type EType", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the EType fields.
func (e *EType) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case etype.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				e.ID = *value
			}
		case etype.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				e.Name = value.String
			}
		case etype.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field event_type", values[i])
			} else if value.Valid {
				e.event_type = new(uuid.UUID)
				*e.event_type = *value.S.(*uuid.UUID)
			}
		}
	}
	return nil
}

// QueryEvent queries the "event" edge of the EType entity.
func (e *EType) QueryEvent() *EventQuery {
	return (&ETypeClient{config: e.config}).QueryEvent(e)
}

// Update returns a builder for updating this EType.
// Note that you need to call EType.Unwrap() before calling this method if this EType
// was returned from a transaction, and the transaction was committed or rolled back.
func (e *EType) Update() *ETypeUpdateOne {
	return (&ETypeClient{config: e.config}).UpdateOne(e)
}

// Unwrap unwraps the EType entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (e *EType) Unwrap() *EType {
	_tx, ok := e.config.driver.(*txDriver)
	if !ok {
		panic("ent: EType is not a transactional entity")
	}
	e.config.driver = _tx.drv
	return e
}

// String implements the fmt.Stringer.
func (e *EType) String() string {
	var builder strings.Builder
	builder.WriteString("EType(")
	builder.WriteString(fmt.Sprintf("id=%v, ", e.ID))
	builder.WriteString("name=")
	builder.WriteString(e.Name)
	builder.WriteByte(')')
	return builder.String()
}

// ETypes is a parsable slice of EType.
type ETypes []*EType

func (e ETypes) config(cfg config) {
	for _i := range e {
		e[_i].config = cfg
	}
}
