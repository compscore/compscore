// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/compscore/compscore/pkg/ent/team"
)

// Team is the model entity for the Team schema.
type Team struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"-"`
	// Team number
	Number int `json:"number"`
	// Team name
	Name string `json:"name"`
	// Team password
	Password string `json:"-"`
	// User Permissions
	Role team.Role `json:"role"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TeamQuery when eager-loading is set.
	Edges        TeamEdges `json:"edges"`
	selectValues sql.SelectValues
}

// TeamEdges holds the relations/edges for other nodes in the graph.
type TeamEdges struct {
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
func (e TeamEdges) StatusOrErr() ([]*Status, error) {
	if e.loadedTypes[0] {
		return e.Status, nil
	}
	return nil, &NotLoadedError{edge: "status"}
}

// CredentialOrErr returns the Credential value or an error if the edge
// was not loaded in eager-loading.
func (e TeamEdges) CredentialOrErr() ([]*Credential, error) {
	if e.loadedTypes[1] {
		return e.Credential, nil
	}
	return nil, &NotLoadedError{edge: "credential"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Team) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case team.FieldID, team.FieldNumber:
			values[i] = new(sql.NullInt64)
		case team.FieldName, team.FieldPassword, team.FieldRole:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Team fields.
func (t *Team) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case team.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			t.ID = int(value.Int64)
		case team.FieldNumber:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field number", values[i])
			} else if value.Valid {
				t.Number = int(value.Int64)
			}
		case team.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				t.Name = value.String
			}
		case team.FieldPassword:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field password", values[i])
			} else if value.Valid {
				t.Password = value.String
			}
		case team.FieldRole:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field role", values[i])
			} else if value.Valid {
				t.Role = team.Role(value.String)
			}
		default:
			t.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Team.
// This includes values selected through modifiers, order, etc.
func (t *Team) Value(name string) (ent.Value, error) {
	return t.selectValues.Get(name)
}

// QueryStatus queries the "status" edge of the Team entity.
func (t *Team) QueryStatus() *StatusQuery {
	return NewTeamClient(t.config).QueryStatus(t)
}

// QueryCredential queries the "credential" edge of the Team entity.
func (t *Team) QueryCredential() *CredentialQuery {
	return NewTeamClient(t.config).QueryCredential(t)
}

// Update returns a builder for updating this Team.
// Note that you need to call Team.Unwrap() before calling this method if this Team
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Team) Update() *TeamUpdateOne {
	return NewTeamClient(t.config).UpdateOne(t)
}

// Unwrap unwraps the Team entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *Team) Unwrap() *Team {
	_tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Team is not a transactional entity")
	}
	t.config.driver = _tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Team) String() string {
	var builder strings.Builder
	builder.WriteString("Team(")
	builder.WriteString(fmt.Sprintf("id=%v, ", t.ID))
	builder.WriteString("number=")
	builder.WriteString(fmt.Sprintf("%v", t.Number))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(t.Name)
	builder.WriteString(", ")
	builder.WriteString("password=<sensitive>")
	builder.WriteString(", ")
	builder.WriteString("role=")
	builder.WriteString(fmt.Sprintf("%v", t.Role))
	builder.WriteByte(')')
	return builder.String()
}

// Teams is a parsable slice of Team.
type Teams []*Team
