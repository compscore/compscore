// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/compscore/compscore/pkg/ent/round"
	"github.com/compscore/compscore/pkg/ent/score"
	"github.com/compscore/compscore/pkg/ent/user"
	"github.com/google/uuid"
)

// Score is the model entity for the Score schema.
type Score struct {
	config `json:"-"`
	// ID of the ent.
	// ID of the score
	ID uuid.UUID `json:"id"`
	// Score for team at a given round
	Score int `json:"score"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ScoreQuery when eager-loading is set.
	Edges        ScoreEdges `json:"edges"`
	round_scores *uuid.UUID
	user_scores  *uuid.UUID
	selectValues sql.SelectValues
}

// ScoreEdges holds the relations/edges for other nodes in the graph.
type ScoreEdges struct {
	// Round for the score
	Round *Round `json:"round,omitempty"`
	// Team for the score
	User *User `json:"team,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// RoundOrErr returns the Round value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ScoreEdges) RoundOrErr() (*Round, error) {
	if e.loadedTypes[0] {
		if e.Round == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: round.Label}
		}
		return e.Round, nil
	}
	return nil, &NotLoadedError{edge: "round"}
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ScoreEdges) UserOrErr() (*User, error) {
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
func (*Score) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case score.FieldScore:
			values[i] = new(sql.NullInt64)
		case score.FieldID:
			values[i] = new(uuid.UUID)
		case score.ForeignKeys[0]: // round_scores
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		case score.ForeignKeys[1]: // user_scores
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Score fields.
func (s *Score) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case score.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				s.ID = *value
			}
		case score.FieldScore:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field score", values[i])
			} else if value.Valid {
				s.Score = int(value.Int64)
			}
		case score.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field round_scores", values[i])
			} else if value.Valid {
				s.round_scores = new(uuid.UUID)
				*s.round_scores = *value.S.(*uuid.UUID)
			}
		case score.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field user_scores", values[i])
			} else if value.Valid {
				s.user_scores = new(uuid.UUID)
				*s.user_scores = *value.S.(*uuid.UUID)
			}
		default:
			s.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Score.
// This includes values selected through modifiers, order, etc.
func (s *Score) Value(name string) (ent.Value, error) {
	return s.selectValues.Get(name)
}

// QueryRound queries the "round" edge of the Score entity.
func (s *Score) QueryRound() *RoundQuery {
	return NewScoreClient(s.config).QueryRound(s)
}

// QueryUser queries the "user" edge of the Score entity.
func (s *Score) QueryUser() *UserQuery {
	return NewScoreClient(s.config).QueryUser(s)
}

// Update returns a builder for updating this Score.
// Note that you need to call Score.Unwrap() before calling this method if this Score
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Score) Update() *ScoreUpdateOne {
	return NewScoreClient(s.config).UpdateOne(s)
}

// Unwrap unwraps the Score entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Score) Unwrap() *Score {
	_tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Score is not a transactional entity")
	}
	s.config.driver = _tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Score) String() string {
	var builder strings.Builder
	builder.WriteString("Score(")
	builder.WriteString(fmt.Sprintf("id=%v, ", s.ID))
	builder.WriteString("score=")
	builder.WriteString(fmt.Sprintf("%v", s.Score))
	builder.WriteByte(')')
	return builder.String()
}

// Scores is a parsable slice of Score.
type Scores []*Score
