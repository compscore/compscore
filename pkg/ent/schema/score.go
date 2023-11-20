package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Score holds the schema definition for the Score entity.
type Score struct {
	ent.Schema
}

// Fields of the Score.
func (Score) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Comment("ID of the score").
			StructTag(`json:"id"`).
			Immutable().
			Unique().
			Default(uuid.New),
		field.Int("score").
			Comment("Score for team at a given round").
			StructTag(`json:"score"`).
			Positive(),
	}
}

// Edges of the Score.
func (Score) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("round", Round.Type).
			Comment("Round for the score").
			StructTag(`json:"round,omitempty"`).
			Ref("scores").
			Required().
			Unique(),
		edge.From("user", User.Type).
			Comment("Team for the score").
			StructTag(`json:"team,omitempty"`).
			Ref("scores").
			Required().
			Unique(),
	}
}
