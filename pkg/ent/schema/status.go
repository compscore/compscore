package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Status holds the schema definition for the Status entity.
type Status struct {
	ent.Schema
}

// Fields of the Status.
func (Status) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Comment("ID of the status").
			StructTag(`json:"id"`).
			Immutable().
			Unique().
			Default(uuid.New),
		field.Enum("status").
			Comment("Status of the status").
			StructTag(`json:"status"`).
			Values(
				"success",
				"failure",
				"unknown",
			).
			Default("unknown"),
		field.String("message").
			Comment("Message of the status").
			StructTag(`json:"message"`).
			Optional(),
		field.Time("timestamp").
			Comment("Timestamp of the status").
			StructTag(`json:"timestamp"`).
			Default(time.Now),
		field.Int("points").
			Comment("Points of the status").
			StructTag(`json:"points"`).
			NonNegative(),
	}
}

// Edges of the Status.
func (Status) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("round", Round.Type).
			Comment("Round of the status").
			StructTag(`json:"round,omitempty"`).
			Required().
			Unique(),
		edge.To("check", Check.Type).
			Comment("Check of the status").
			StructTag(`json:"check,omitempty"`).
			Required().
			Unique(),
		edge.To("user", User.Type).
			Comment("User of the status").
			StructTag(`json:"user,omitempty"`).
			Required().
			Unique(),
	}
}
