package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Status holds the schema definition for the Status entity.
type Status struct {
	ent.Schema
}

// Fields of the Status.
func (Status) Fields() []ent.Field {
	return []ent.Field{
		field.String("error").
			StructTag(`json:"error"`).
			Comment("Error message").
			Optional(),
		field.Enum("status").
			StructTag(`json:"status"`).
			Comment("Status").
			Values("up", "down", "unknown").
			Default("unknown"),
		field.Time("time").
			StructTag(`json:"time"`).
			Comment("Time of check").
			Default(time.Now),
		field.Int("id").
			StructTag(`json:"-"`),
	}
}

// Edges of the Status.
func (Status) Edges() []ent.Edge {
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
		edge.To("round", Round.Type).
			StructTag(`json:"round,omitempty"`).
			Comment("Round").
			Required().
			Unique(),
	}
}
