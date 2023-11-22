// Code generated by ent, DO NOT EDIT.

package user

import (
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// FieldTeamNumber holds the string denoting the team_number field in the database.
	FieldTeamNumber = "team_number"
	// FieldRole holds the string denoting the role field in the database.
	FieldRole = "role"
	// EdgeCredentials holds the string denoting the credentials edge name in mutations.
	EdgeCredentials = "credentials"
	// EdgeStatuses holds the string denoting the statuses edge name in mutations.
	EdgeStatuses = "statuses"
	// EdgeScores holds the string denoting the scores edge name in mutations.
	EdgeScores = "scores"
	// Table holds the table name of the user in the database.
	Table = "users"
	// CredentialsTable is the table that holds the credentials relation/edge.
	CredentialsTable = "credentials"
	// CredentialsInverseTable is the table name for the Credential entity.
	// It exists in this package in order to avoid circular dependency with the "credential" package.
	CredentialsInverseTable = "credentials"
	// CredentialsColumn is the table column denoting the credentials relation/edge.
	CredentialsColumn = "credential_user"
	// StatusesTable is the table that holds the statuses relation/edge.
	StatusesTable = "status"
	// StatusesInverseTable is the table name for the Status entity.
	// It exists in this package in order to avoid circular dependency with the "status" package.
	StatusesInverseTable = "status"
	// StatusesColumn is the table column denoting the statuses relation/edge.
	StatusesColumn = "status_user"
	// ScoresTable is the table that holds the scores relation/edge.
	ScoresTable = "scores"
	// ScoresInverseTable is the table name for the Score entity.
	// It exists in this package in order to avoid circular dependency with the "score" package.
	ScoresInverseTable = "scores"
	// ScoresColumn is the table column denoting the scores relation/edge.
	ScoresColumn = "user_scores"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldPassword,
	FieldTeamNumber,
	FieldRole,
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
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	PasswordValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// Role defines the type for the "role" enum field.
type Role string

// RoleCompetitor is the default value of the Role enum.
const DefaultRole = RoleCompetitor

// Role values.
const (
	RoleCompetitor Role = "competitor"
	RoleAdmin      Role = "admin"
)

func (r Role) String() string {
	return string(r)
}

// RoleValidator is a validator for the "role" field enum values. It is called by the builders before save.
func RoleValidator(r Role) error {
	switch r {
	case RoleCompetitor, RoleAdmin:
		return nil
	default:
		return fmt.Errorf("user: invalid enum value for role field: %q", r)
	}
}

// OrderOption defines the ordering options for the User queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByPassword orders the results by the password field.
func ByPassword(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPassword, opts...).ToFunc()
}

// ByTeamNumber orders the results by the team_number field.
func ByTeamNumber(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTeamNumber, opts...).ToFunc()
}

// ByRole orders the results by the role field.
func ByRole(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRole, opts...).ToFunc()
}

// ByCredentialsCount orders the results by credentials count.
func ByCredentialsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newCredentialsStep(), opts...)
	}
}

// ByCredentials orders the results by credentials terms.
func ByCredentials(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCredentialsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByStatusesCount orders the results by statuses count.
func ByStatusesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newStatusesStep(), opts...)
	}
}

// ByStatuses orders the results by statuses terms.
func ByStatuses(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newStatusesStep(), append([]sql.OrderTerm{term}, terms...)...)
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
func newCredentialsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CredentialsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, CredentialsTable, CredentialsColumn),
	)
}
func newStatusesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(StatusesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, StatusesTable, StatusesColumn),
	)
}
func newScoresStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ScoresInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ScoresTable, ScoresColumn),
	)
}