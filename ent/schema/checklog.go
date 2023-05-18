package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// CheckLog holds the schema definition for the CheckLog entity.
type CheckLog struct {
	ent.Schema
}

// Fields of the CheckLog.
func (CheckLog) Fields() []ent.Field {
	return []ent.Field{
		field.String("error").
			Comment("Error message of check").
			Optional(),
		field.Bool("status").
			Comment("Status of check"),
		field.Time("time").
			Comment("Time of check").
			Default(time.Now),
	}
}

// Edges of the CheckLog.
func (CheckLog) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("check", Check.Type).
			Comment("Check that was run").
			Required().
			Unique(),
		edge.To("team", Team.Type).
			Comment("Team that was checked").
			Required().
			Unique(),
		edge.To("round", Round.Type).
			Comment("Round that the check was run for").
			Required().
			Unique(),
	}
}
