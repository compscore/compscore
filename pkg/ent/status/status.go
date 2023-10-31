// Code generated by ent, DO NOT EDIT.

package status

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the status type in the database.
	Label = "status"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldError holds the string denoting the error field in the database.
	FieldError = "error"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldTime holds the string denoting the time field in the database.
	FieldTime = "time"
	// FieldPoints holds the string denoting the points field in the database.
	FieldPoints = "points"
	// EdgeCheck holds the string denoting the check edge name in mutations.
	EdgeCheck = "check"
	// EdgeTeam holds the string denoting the team edge name in mutations.
	EdgeTeam = "team"
	// EdgeRound holds the string denoting the round edge name in mutations.
	EdgeRound = "round"
	// Table holds the table name of the status in the database.
	Table = "status"
	// CheckTable is the table that holds the check relation/edge.
	CheckTable = "status"
	// CheckInverseTable is the table name for the Check entity.
	// It exists in this package in order to avoid circular dependency with the "check" package.
	CheckInverseTable = "checks"
	// CheckColumn is the table column denoting the check relation/edge.
	CheckColumn = "status_check"
	// TeamTable is the table that holds the team relation/edge.
	TeamTable = "status"
	// TeamInverseTable is the table name for the Team entity.
	// It exists in this package in order to avoid circular dependency with the "team" package.
	TeamInverseTable = "teams"
	// TeamColumn is the table column denoting the team relation/edge.
	TeamColumn = "status_team"
	// RoundTable is the table that holds the round relation/edge.
	RoundTable = "status"
	// RoundInverseTable is the table name for the Round entity.
	// It exists in this package in order to avoid circular dependency with the "round" package.
	RoundInverseTable = "rounds"
	// RoundColumn is the table column denoting the round relation/edge.
	RoundColumn = "status_round"
)

// Columns holds all SQL columns for status fields.
var Columns = []string{
	FieldID,
	FieldError,
	FieldStatus,
	FieldTime,
	FieldPoints,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "status"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"status_check",
	"status_team",
	"status_round",
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
	// DefaultTime holds the default value on creation for the "time" field.
	DefaultTime func() time.Time
	// PointsValidator is a validator for the "points" field. It is called by the builders before save.
	PointsValidator func(int) error
)

// Status defines the type for the "status" enum field.
type Status string

// StatusUnknown is the default value of the Status enum.
const DefaultStatus = StatusUnknown

// Status values.
const (
	StatusUp      Status = "up"
	StatusDown    Status = "down"
	StatusUnknown Status = "unknown"
)

func (s Status) String() string {
	return string(s)
}

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s Status) error {
	switch s {
	case StatusUp, StatusDown, StatusUnknown:
		return nil
	default:
		return fmt.Errorf("status: invalid enum value for status field: %q", s)
	}
}

// OrderOption defines the ordering options for the Status queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByError orders the results by the error field.
func ByError(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldError, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByTime orders the results by the time field.
func ByTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTime, opts...).ToFunc()
}

// ByPoints orders the results by the points field.
func ByPoints(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPoints, opts...).ToFunc()
}

// ByCheckField orders the results by check field.
func ByCheckField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCheckStep(), sql.OrderByField(field, opts...))
	}
}

// ByTeamField orders the results by team field.
func ByTeamField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTeamStep(), sql.OrderByField(field, opts...))
	}
}

// ByRoundField orders the results by round field.
func ByRoundField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newRoundStep(), sql.OrderByField(field, opts...))
	}
}
func newCheckStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CheckInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, CheckTable, CheckColumn),
	)
}
func newTeamStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TeamInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, TeamTable, TeamColumn),
	)
}
func newRoundStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(RoundInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, RoundTable, RoundColumn),
	)
}
