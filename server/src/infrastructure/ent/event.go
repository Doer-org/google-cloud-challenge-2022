// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/Doer-org/google-cloud-challenge-2022/infrastructure/ent/event"
	"github.com/Doer-org/google-cloud-challenge-2022/infrastructure/ent/user"
	"github.com/google/uuid"
)

// Event is the model entity for the Event schema.
type Event struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Detail holds the value of the "detail" field.
	Detail string `json:"detail,omitempty"`
	// Location holds the value of the "location" field.
	Location string `json:"location,omitempty"`
	// Size holds the value of the "size" field.
	Size int `json:"size,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// LimitHour holds the value of the "limit_hour" field.
	LimitHour int `json:"limit_hour,omitempty"`
	// Type holds the value of the "type" field.
	Type string `json:"type,omitempty"`
	// State holds the value of the "state" field.
	State string `json:"state,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the EventQuery when eager-loading is set.
	Edges       EventEdges `json:"edges"`
	event_admin *uuid.UUID
}

// EventEdges holds the relations/edges for other nodes in the graph.
type EventEdges struct {
	// Admin holds the value of the admin edge.
	Admin *User `json:"admin,omitempty"`
	// Users holds the value of the users edge.
	Users []*User `json:"users,omitempty"`
	// Comments holds the value of the comments edge.
	Comments []*Comment `json:"comments,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// AdminOrErr returns the Admin value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e EventEdges) AdminOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.Admin == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Admin, nil
	}
	return nil, &NotLoadedError{edge: "admin"}
}

// UsersOrErr returns the Users value or an error if the edge
// was not loaded in eager-loading.
func (e EventEdges) UsersOrErr() ([]*User, error) {
	if e.loadedTypes[1] {
		return e.Users, nil
	}
	return nil, &NotLoadedError{edge: "users"}
}

// CommentsOrErr returns the Comments value or an error if the edge
// was not loaded in eager-loading.
func (e EventEdges) CommentsOrErr() ([]*Comment, error) {
	if e.loadedTypes[2] {
		return e.Comments, nil
	}
	return nil, &NotLoadedError{edge: "comments"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Event) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case event.FieldSize, event.FieldLimitHour:
			values[i] = new(sql.NullInt64)
		case event.FieldName, event.FieldDetail, event.FieldLocation, event.FieldType, event.FieldState:
			values[i] = new(sql.NullString)
		case event.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		case event.FieldID:
			values[i] = new(uuid.UUID)
		case event.ForeignKeys[0]: // event_admin
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Event", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Event fields.
func (e *Event) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case event.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				e.ID = *value
			}
		case event.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				e.Name = value.String
			}
		case event.FieldDetail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field detail", values[i])
			} else if value.Valid {
				e.Detail = value.String
			}
		case event.FieldLocation:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field location", values[i])
			} else if value.Valid {
				e.Location = value.String
			}
		case event.FieldSize:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field size", values[i])
			} else if value.Valid {
				e.Size = int(value.Int64)
			}
		case event.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				e.CreatedAt = value.Time
			}
		case event.FieldLimitHour:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field limit_hour", values[i])
			} else if value.Valid {
				e.LimitHour = int(value.Int64)
			}
		case event.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				e.Type = value.String
			}
		case event.FieldState:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field state", values[i])
			} else if value.Valid {
				e.State = value.String
			}
		case event.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field event_admin", values[i])
			} else if value.Valid {
				e.event_admin = new(uuid.UUID)
				*e.event_admin = *value.S.(*uuid.UUID)
			}
		}
	}
	return nil
}

// QueryAdmin queries the "admin" edge of the Event entity.
func (e *Event) QueryAdmin() *UserQuery {
	return (&EventClient{config: e.config}).QueryAdmin(e)
}

// QueryUsers queries the "users" edge of the Event entity.
func (e *Event) QueryUsers() *UserQuery {
	return (&EventClient{config: e.config}).QueryUsers(e)
}

// QueryComments queries the "comments" edge of the Event entity.
func (e *Event) QueryComments() *CommentQuery {
	return (&EventClient{config: e.config}).QueryComments(e)
}

// Update returns a builder for updating this Event.
// Note that you need to call Event.Unwrap() before calling this method if this Event
// was returned from a transaction, and the transaction was committed or rolled back.
func (e *Event) Update() *EventUpdateOne {
	return (&EventClient{config: e.config}).UpdateOne(e)
}

// Unwrap unwraps the Event entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (e *Event) Unwrap() *Event {
	_tx, ok := e.config.driver.(*txDriver)
	if !ok {
		panic("ent: Event is not a transactional entity")
	}
	e.config.driver = _tx.drv
	return e
}

// String implements the fmt.Stringer.
func (e *Event) String() string {
	var builder strings.Builder
	builder.WriteString("Event(")
	builder.WriteString(fmt.Sprintf("id=%v, ", e.ID))
	builder.WriteString("name=")
	builder.WriteString(e.Name)
	builder.WriteString(", ")
	builder.WriteString("detail=")
	builder.WriteString(e.Detail)
	builder.WriteString(", ")
	builder.WriteString("location=")
	builder.WriteString(e.Location)
	builder.WriteString(", ")
	builder.WriteString("size=")
	builder.WriteString(fmt.Sprintf("%v", e.Size))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(e.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("limit_hour=")
	builder.WriteString(fmt.Sprintf("%v", e.LimitHour))
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(e.Type)
	builder.WriteString(", ")
	builder.WriteString("state=")
	builder.WriteString(e.State)
	builder.WriteByte(')')
	return builder.String()
}

// Events is a parsable slice of Event.
type Events []*Event

func (e Events) config(cfg config) {
	for _i := range e {
		e[_i].config = cfg
	}
}
