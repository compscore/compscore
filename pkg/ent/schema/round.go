package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Round holds the schema definition for the Round entity.
type Round struct {
	ent.Schema
}

// Fields of the Round.
func (Round) Fields() []ent.Field {
	return []ent.Field{
		field.Int("number").
			Comment("Round number").
			Positive().
			Unique(),
	}
}

// Edges of the Round.
func (Round) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("status", Status.Type).
			Comment("Check statuses").
			Annotations(
				entsql.Annotation{
					OnDelete: entsql.Cascade,
				},
			).
			Ref("round"),
	}
}
