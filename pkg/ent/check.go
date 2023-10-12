// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/compscore/compscore/pkg/ent/check"
)

// Check is the model entity for the Check schema.
type Check struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"-"`
	// Check name
	Name string `json:"name"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CheckQuery when eager-loading is set.
	Edges        CheckEdges `json:"edges"`
	selectValues sql.SelectValues
}

// CheckEdges holds the relations/edges for other nodes in the graph.
type CheckEdges struct {
	// Check statuses
	Status []*Status `json:"status,omitempty"`
	// Check credential
	Credential []*Credential `json:"credential,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// StatusOrErr returns the Status value or an error if the edge
// was not loaded in eager-loading.
func (e CheckEdges) StatusOrErr() ([]*Status, error) {
	if e.loadedTypes[0] {
		return e.Status, nil
	}
	return nil, &NotLoadedError{edge: "status"}
}

// CredentialOrErr returns the Credential value or an error if the edge
// was not loaded in eager-loading.
func (e CheckEdges) CredentialOrErr() ([]*Credential, error) {
	if e.loadedTypes[1] {
		return e.Credential, nil
	}
	return nil, &NotLoadedError{edge: "credential"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Check) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case check.FieldID:
			values[i] = new(sql.NullInt64)
		case check.FieldName:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Check fields.
func (c *Check) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case check.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = int(value.Int64)
		case check.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				c.Name = value.String
			}
		default:
			c.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Check.
// This includes values selected through modifiers, order, etc.
func (c *Check) Value(name string) (ent.Value, error) {
	return c.selectValues.Get(name)
}

// QueryStatus queries the "status" edge of the Check entity.
func (c *Check) QueryStatus() *StatusQuery {
	return NewCheckClient(c.config).QueryStatus(c)
}

// QueryCredential queries the "credential" edge of the Check entity.
func (c *Check) QueryCredential() *CredentialQuery {
	return NewCheckClient(c.config).QueryCredential(c)
}

// Update returns a builder for updating this Check.
// Note that you need to call Check.Unwrap() before calling this method if this Check
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Check) Update() *CheckUpdateOne {
	return NewCheckClient(c.config).UpdateOne(c)
}

// Unwrap unwraps the Check entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Check) Unwrap() *Check {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Check is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Check) String() string {
	var builder strings.Builder
	builder.WriteString("Check(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("name=")
	builder.WriteString(c.Name)
	builder.WriteByte(')')
	return builder.String()
}

// Checks is a parsable slice of Check.
type Checks []*Check
