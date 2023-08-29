// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/compscore/compscore/pkg/ent/round"
)

// Round is the model entity for the Round schema.
type Round struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Round number
	Number int `json:"number,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the RoundQuery when eager-loading is set.
	Edges        RoundEdges `json:"edges"`
	selectValues sql.SelectValues
}

// RoundEdges holds the relations/edges for other nodes in the graph.
type RoundEdges struct {
	// Check statuses
	Status []*Status `json:"status,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// StatusOrErr returns the Status value or an error if the edge
// was not loaded in eager-loading.
func (e RoundEdges) StatusOrErr() ([]*Status, error) {
	if e.loadedTypes[0] {
		return e.Status, nil
	}
	return nil, &NotLoadedError{edge: "status"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Round) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case round.FieldID, round.FieldNumber:
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Round fields.
func (r *Round) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case round.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			r.ID = int(value.Int64)
		case round.FieldNumber:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field number", values[i])
			} else if value.Valid {
				r.Number = int(value.Int64)
			}
		default:
			r.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Round.
// This includes values selected through modifiers, order, etc.
func (r *Round) Value(name string) (ent.Value, error) {
	return r.selectValues.Get(name)
}

// QueryStatus queries the "status" edge of the Round entity.
func (r *Round) QueryStatus() *StatusQuery {
	return NewRoundClient(r.config).QueryStatus(r)
}

// Update returns a builder for updating this Round.
// Note that you need to call Round.Unwrap() before calling this method if this Round
// was returned from a transaction, and the transaction was committed or rolled back.
func (r *Round) Update() *RoundUpdateOne {
	return NewRoundClient(r.config).UpdateOne(r)
}

// Unwrap unwraps the Round entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (r *Round) Unwrap() *Round {
	_tx, ok := r.config.driver.(*txDriver)
	if !ok {
		panic("ent: Round is not a transactional entity")
	}
	r.config.driver = _tx.drv
	return r
}

// String implements the fmt.Stringer.
func (r *Round) String() string {
	var builder strings.Builder
	builder.WriteString("Round(")
	builder.WriteString(fmt.Sprintf("id=%v, ", r.ID))
	builder.WriteString("number=")
	builder.WriteString(fmt.Sprintf("%v", r.Number))
	builder.WriteByte(')')
	return builder.String()
}

// Rounds is a parsable slice of Round.
type Rounds []*Round
