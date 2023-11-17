// Code generated by ent, DO NOT EDIT.

package credential

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the credential type in the database.
	Label = "credential"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// EdgeCheck holds the string denoting the check edge name in mutations.
	EdgeCheck = "check"
	// Table holds the table name of the credential in the database.
	Table = "credentials"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "credentials"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "credential_user"
	// CheckTable is the table that holds the check relation/edge.
	CheckTable = "credentials"
	// CheckInverseTable is the table name for the Check entity.
	// It exists in this package in order to avoid circular dependency with the "check" package.
	CheckInverseTable = "checks"
	// CheckColumn is the table column denoting the check relation/edge.
	CheckColumn = "credential_check"
)

// Columns holds all SQL columns for credential fields.
var Columns = []string{
	FieldID,
	FieldPassword,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "credentials"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"credential_user",
	"credential_check",
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
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(int) error
)

// OrderOption defines the ordering options for the Credential queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByPassword orders the results by the password field.
func ByPassword(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPassword, opts...).ToFunc()
}

// ByUserField orders the results by user field.
func ByUserField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserStep(), sql.OrderByField(field, opts...))
	}
}

// ByCheckField orders the results by check field.
func ByCheckField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCheckStep(), sql.OrderByField(field, opts...))
	}
}
func newUserStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, UserTable, UserColumn),
	)
}
func newCheckStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CheckInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, CheckTable, CheckColumn),
	)
}
