// Code generated by ent, DO NOT EDIT.

package user

import (
	"fmt"
	"io"
	"strconv"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
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
	// EdgeCredential holds the string denoting the credential edge name in mutations.
	EdgeCredential = "credential"
	// EdgeStatus holds the string denoting the status edge name in mutations.
	EdgeStatus = "status"
	// Table holds the table name of the user in the database.
	Table = "users"
	// CredentialTable is the table that holds the credential relation/edge.
	CredentialTable = "credentials"
	// CredentialInverseTable is the table name for the Credential entity.
	// It exists in this package in order to avoid circular dependency with the "credential" package.
	CredentialInverseTable = "credentials"
	// CredentialColumn is the table column denoting the credential relation/edge.
	CredentialColumn = "credential_user"
	// StatusTable is the table that holds the status relation/edge.
	StatusTable = "status"
	// StatusInverseTable is the table name for the Status entity.
	// It exists in this package in order to avoid circular dependency with the "status" package.
	StatusInverseTable = "status"
	// StatusColumn is the table column denoting the status relation/edge.
	StatusColumn = "status_user"
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
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(int) error
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
func newCredentialStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CredentialInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, CredentialTable, CredentialColumn),
	)
}
func newStatusStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(StatusInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, StatusTable, StatusColumn),
	)
}

// MarshalGQL implements graphql.Marshaler interface.
func (e Role) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(e.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (e *Role) UnmarshalGQL(val interface{}) error {
	str, ok := val.(string)
	if !ok {
		return fmt.Errorf("enum %T must be a string", val)
	}
	*e = Role(str)
	if err := RoleValidator(*e); err != nil {
		return fmt.Errorf("%s is not a valid Role", str)
	}
	return nil
}