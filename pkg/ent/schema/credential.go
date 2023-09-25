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
			Comment("Password of Check"),
	}
}

// Edges of the Check.
func (Credential) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("check", Check.Type).
			Comment("Check").
			Required().
			Unique(),
		edge.To("team", Team.Type).
			Comment("Team").
			Required().
			Unique(),
	}
}
