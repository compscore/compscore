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
		field.Int("id").
			Comment("ID of the round").
			StructTag(`json:"id"`).
			Positive().
			Immutable().
			Unique(),
		field.Bool("completed").
			Comment("Whether the round is completed").
			StructTag(`json:"completed"`).
			Default(false),
	}
}

// Edges of the Round.
func (Round) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("status", Status.Type).
			Comment("Status of the round").
			StructTag(`json:"status,omitempty"`).
			Annotations(
				entsql.Annotation{
					OnDelete: entsql.Cascade,
				},
			).
			Ref("round"),
		edge.To("scores", Score.Type).
			Comment("Scores for the round").
			StructTag(`json:"scores,omitempty"`).
			Annotations(
				entsql.Annotation{
					OnDelete: entsql.Cascade,
				},
			),
	}
}
