package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Credential holds the schema definition for the Credential entity.
type Credential struct {
	ent.Schema
}

// Fields of the Credential.
func (Credential) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Comment("ID of the credential").
			StructTag(`json:"id"`).
			Immutable().
			Unique().
			Default(uuid.New),
		field.String("password").
			Comment("Password of the credential").
			StructTag(`json:"password"`),
	}

}

// Edges of the Credential.
func (Credential) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).
			Comment("User of the credential").
			StructTag(`json:"user,omitempty"`).
			Required().
			Unique(),
		edge.To("check", Check.Type).
			Comment("Check of the credential").
			StructTag(`json:"check,omitempty"`).
			Required().
			Unique(),
	}
}
