// Code generated by ent, DO NOT EDIT.

package check

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the check type in the database.
	Label = "check"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldWeight holds the string denoting the weight field in the database.
	FieldWeight = "weight"
	// EdgeStatus holds the string denoting the status edge name in mutations.
	EdgeStatus = "status"
	// EdgeCredential holds the string denoting the credential edge name in mutations.
	EdgeCredential = "credential"
	// Table holds the table name of the check in the database.
	Table = "checks"
	// StatusTable is the table that holds the status relation/edge.
	StatusTable = "status"
	// StatusInverseTable is the table name for the Status entity.
	// It exists in this package in order to avoid circular dependency with the "status" package.
	StatusInverseTable = "status"
	// StatusColumn is the table column denoting the status relation/edge.
	StatusColumn = "status_check"
	// CredentialTable is the table that holds the credential relation/edge.
	CredentialTable = "credentials"
	// CredentialInverseTable is the table name for the Credential entity.
	// It exists in this package in order to avoid circular dependency with the "credential" package.
	CredentialInverseTable = "credentials"
	// CredentialColumn is the table column denoting the credential relation/edge.
	CredentialColumn = "credential_check"
)

// Columns holds all SQL columns for check fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldWeight,
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
	// WeightValidator is a validator for the "weight" field. It is called by the builders before save.
	WeightValidator func(int) error
)

// OrderOption defines the ordering options for the Check queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByWeight orders the results by the weight field.
func ByWeight(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldWeight, opts...).ToFunc()
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

// ByCredentialCount orders the results by credential count.
func ByCredentialCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newCredentialStep(), opts...)
	}
}

// ByCredential orders the results by credential terms.
func ByCredential(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCredentialStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newStatusStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(StatusInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, StatusTable, StatusColumn),
	)
}
func newCredentialStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CredentialInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, CredentialTable, CredentialColumn),
	)
}