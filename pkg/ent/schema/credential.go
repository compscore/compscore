package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Check holds the schema definition for the Credential entity.
type Credential struct {
	ent.Schema
}

// Fields of the Check.
func (Credential) Fields() []ent.Field {
	return []ent.Field{
		field.String("password").
			StructTag(`json:"password"`).
			Comment("Password of Check"),
		field.Int("id").
			StructTag(`json:"-"`),
	}
}

// Edges of the Check.
func (Credential) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("check", Check.Type).
			StructTag(`json:"check,omitempty"`).
			Comment("Check").
			Required().
			Unique(),
		edge.To("team", Team.Type).
			StructTag(`json:"team,omitempty"`).
			Comment("Team").
			Required().
			Unique(),
	}
}
