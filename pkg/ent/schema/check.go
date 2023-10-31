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
		field.String("name").
			StructTag(`json:"name"`).
			Comment("Check name").
			NotEmpty().
			Unique(),
		field.Int("id").
			StructTag(`json:"-"`),
		field.Int("weight").
			StructTag(`json:"-"`).
			NonNegative(),
	}
}

// Edges of the Check.
func (Check) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("status", Status.Type).
			StructTag(`json:"status,omitempty"`).
			Comment("Check statuses").
			Annotations(
				entsql.Annotation{
					OnDelete: entsql.Cascade,
				},
			).
			Ref("check"),
		edge.From("credential", Credential.Type).
			StructTag(`json:"credential,omitempty"`).
			Comment("Check credential").
			Annotations(
				entsql.Annotation{
					OnDelete: entsql.Cascade,
				},
			).
			Ref("check"),
	}
}
