package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Check holds the schema definition for the Check entity.
type Check struct {
	ent.Schema
}

// Fields of the Check.
func (Check) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").
			Comment("ID of the check").
			StructTag(`json:"id"`).
			Positive().
			Immutable().
			Unique(),
		field.String("name").
			Comment("Name of the check").
			StructTag(`json:"name"`).
			Unique().
			NotEmpty(),
		field.Int("weight").
			Comment("Weight of the check").
			StructTag(`json:"weight"`).
			NonNegative(),
	}
}

// Edges of the Check.
func (Check) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("credential", Credential.Type).
			Comment("Credential of the check").
			StructTag(`json:"credential,omitempty"`).
			Annotations(
				entsql.Annotation{
					OnDelete: entsql.Cascade,
				},
			).
			Ref("check"),
		edge.From("status", Status.Type).
			Comment("Status of the check").
			StructTag(`json:"status,omitempty"`).
			Annotations(
				entsql.Annotation{
					OnDelete: entsql.Cascade,
				},
			).
			Ref("check"),
	}
}
