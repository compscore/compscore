// Code generated by ent, DO NOT EDIT.

package round

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the round type in the database.
	Label = "round"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldNumber holds the string denoting the number field in the database.
	FieldNumber = "number"
	// FieldCompleted holds the string denoting the completed field in the database.
	FieldCompleted = "completed"
	// EdgeStatus holds the string denoting the status edge name in mutations.
	EdgeStatus = "status"
	// EdgeScores holds the string denoting the scores edge name in mutations.
	EdgeScores = "scores"
	// Table holds the table name of the round in the database.
	Table = "rounds"
	// StatusTable is the table that holds the status relation/edge.
	StatusTable = "status"
	// StatusInverseTable is the table name for the Status entity.
	// It exists in this package in order to avoid circular dependency with the "status" package.
	StatusInverseTable = "status"
	// StatusColumn is the table column denoting the status relation/edge.
	StatusColumn = "status_round"
	// ScoresTable is the table that holds the scores relation/edge.
	ScoresTable = "scores"
	// ScoresInverseTable is the table name for the Score entity.
	// It exists in this package in order to avoid circular dependency with the "score" package.
	ScoresInverseTable = "scores"
	// ScoresColumn is the table column denoting the scores relation/edge.
	ScoresColumn = "round_scores"
)

// Columns holds all SQL columns for round fields.
var Columns = []string{
	FieldID,
	FieldNumber,
	FieldCompleted,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// NumberValidator is a validator for the "number" field. It is called by the builders before save.
	NumberValidator func(int) error
	// DefaultCompleted holds the default value on creation for the "completed" field.
	DefaultCompleted bool
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// OrderOption defines the ordering options for the Round queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByNumber orders the results by the number field.
func ByNumber(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNumber, opts...).ToFunc()
}

// ByCompleted orders the results by the completed field.
func ByCompleted(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCompleted, opts...).ToFunc()
}

// ByStatusCount orders the results by status count.
func ByStatusCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newStatusStep(), opts...)
	}
}

// ByStatus orders the results by status terms.
func ByStatus(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newStatusStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByScoresCount orders the results by scores count.
func ByScoresCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newScoresStep(), opts...)
	}
}

// ByScores orders the results by scores terms.
func ByScores(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newScoresStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newStatusStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(StatusInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, StatusTable, StatusColumn),
	)
}
func newScoresStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ScoresInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ScoresTable, ScoresColumn),
	)
}
