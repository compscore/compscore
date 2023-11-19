// Code generated by ent, DO NOT EDIT.

package score

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the score type in the database.
	Label = "score"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldScore holds the string denoting the score field in the database.
	FieldScore = "score"
	// EdgeRound holds the string denoting the round edge name in mutations.
	EdgeRound = "round"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// Table holds the table name of the score in the database.
	Table = "scores"
	// RoundTable is the table that holds the round relation/edge.
	RoundTable = "scores"
	// RoundInverseTable is the table name for the Round entity.
	// It exists in this package in order to avoid circular dependency with the "round" package.
	RoundInverseTable = "rounds"
	// RoundColumn is the table column denoting the round relation/edge.
	RoundColumn = "round_scores"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "scores"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_scores"
)

// Columns holds all SQL columns for score fields.
var Columns = []string{
	FieldID,
	FieldScore,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "scores"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"round_scores",
	"user_scores",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// ScoreValidator is a validator for the "score" field. It is called by the builders before save.
	ScoreValidator func(int) error
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(int) error
)

// OrderOption defines the ordering options for the Score queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByScore orders the results by the score field.
func ByScore(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldScore, opts...).ToFunc()
}

// ByRoundField orders the results by round field.
func ByRoundField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newRoundStep(), sql.OrderByField(field, opts...))
	}
}

// ByUserField orders the results by user field.
func ByUserField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserStep(), sql.OrderByField(field, opts...))
	}
}
func newRoundStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(RoundInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, RoundTable, RoundColumn),
	)
}
func newUserStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
	)
}
