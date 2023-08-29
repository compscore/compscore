package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
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
			Comment("Error message").
			Optional(),
		field.Enum("status").
			Comment("Status").
			Values("up", "down", "unknown").
			Default("unknown"),
		field.Time("time").
			Comment("Time of check").
			Default(time.Now),
	}
}

// Edges of the Status.
func (Status) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("check", Check.Type).
			Comment("Check").
			Annotations(
				entsql.Annotation{
					OnDelete: entsql.Cascade,
				},
			).
			Required().
			Unique(),
		edge.To("team", Team.Type).
			Comment("Team").
			Annotations(
				entsql.Annotation{
					OnDelete: entsql.Cascade,
				},
			).
			Required().
			Unique(),
		edge.To("round", Round.Type).
			Comment("Round").
			Annotations(
				entsql.Annotation{
					OnDelete: entsql.Cascade,
				},
			).
			Required().
			Unique(),
	}
}
