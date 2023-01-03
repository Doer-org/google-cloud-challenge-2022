// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/Doer-org/google-cloud-challenge-2022/infrastructure/ent/ecomment"
	"github.com/Doer-org/google-cloud-challenge-2022/infrastructure/ent/event"
	"github.com/Doer-org/google-cloud-challenge-2022/infrastructure/ent/user"
	"github.com/google/uuid"
)

// Ecomment is the model entity for the Ecomment schema.
type Ecomment struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Body holds the value of the "body" field.
	Body string `json:"body,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the EcommentQuery when eager-loading is set.
	Edges          EcommentEdges `json:"edges"`
	ecomment_event *uuid.UUID
	ecomment_user  *uuid.UUID
}

// EcommentEdges holds the relations/edges for other nodes in the graph.
type EcommentEdges struct {
	// Event holds the value of the event edge.
	Event *Event `json:"event,omitempty"`
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// EventOrErr returns the Event value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e EcommentEdges) EventOrErr() (*Event, error) {
	if e.loadedTypes[0] {
		if e.Event == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: event.Label}
		}
		return e.Event, nil
	}
	return nil, &NotLoadedError{edge: "event"}
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e EcommentEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[1] {
		if e.User == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Ecomment) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case ecomment.FieldBody:
			values[i] = new(sql.NullString)
		case ecomment.FieldID:
			values[i] = new(uuid.UUID)
		case ecomment.ForeignKeys[0]: // ecomment_event
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		case ecomment.ForeignKeys[1]: // ecomment_user
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Ecomment", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Ecomment fields.
func (e *Ecomment) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case ecomment.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				e.ID = *value
			}
		case ecomment.FieldBody:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field body", values[i])
			} else if value.Valid {
				e.Body = value.String
			}
		case ecomment.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field ecomment_event", values[i])
			} else if value.Valid {
				e.ecomment_event = new(uuid.UUID)
				*e.ecomment_event = *value.S.(*uuid.UUID)
			}
		case ecomment.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field ecomment_user", values[i])
			} else if value.Valid {
				e.ecomment_user = new(uuid.UUID)
				*e.ecomment_user = *value.S.(*uuid.UUID)
			}
		}
	}
	return nil
}

// QueryEvent queries the "event" edge of the Ecomment entity.
func (e *Ecomment) QueryEvent() *EventQuery {
	return (&EcommentClient{config: e.config}).QueryEvent(e)
}

// QueryUser queries the "user" edge of the Ecomment entity.
func (e *Ecomment) QueryUser() *UserQuery {
	return (&EcommentClient{config: e.config}).QueryUser(e)
}

// Update returns a builder for updating this Ecomment.
// Note that you need to call Ecomment.Unwrap() before calling this method if this Ecomment
// was returned from a transaction, and the transaction was committed or rolled back.
func (e *Ecomment) Update() *EcommentUpdateOne {
	return (&EcommentClient{config: e.config}).UpdateOne(e)
}

// Unwrap unwraps the Ecomment entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (e *Ecomment) Unwrap() *Ecomment {
	_tx, ok := e.config.driver.(*txDriver)
	if !ok {
		panic("ent: Ecomment is not a transactional entity")
	}
	e.config.driver = _tx.drv
	return e
}

// String implements the fmt.Stringer.
func (e *Ecomment) String() string {
	var builder strings.Builder
	builder.WriteString("Ecomment(")
	builder.WriteString(fmt.Sprintf("id=%v, ", e.ID))
	builder.WriteString("body=")
	builder.WriteString(e.Body)
	builder.WriteByte(')')
	return builder.String()
}

// Ecomments is a parsable slice of Ecomment.
type Ecomments []*Ecomment

func (e Ecomments) config(cfg config) {
	for _i := range e {
		e[_i].config = cfg
	}
}
