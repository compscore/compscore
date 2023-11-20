package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Round holds the schema definition for the Round entity.
type Round struct {
	ent.Schema
}

// Fields of the Round.
func (Round) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Comment("ID of the round").
			StructTag(`json:"id"`).
			Immutable().
			Unique().
			Default(uuid.New),
		field.Int("number").
			Comment("Number of the round").
			StructTag(`json:"number"`).
			Positive().
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
