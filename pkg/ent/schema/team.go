package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Team holds the schema definition for the Team entity.
type Team struct {
	ent.Schema
}

// Fields of the Team.
func (Team) Fields() []ent.Field {
	return []ent.Field{
		field.Int8("number").
			StructTag(`json:"number"`).
			Comment("Team number").
			Optional().
			Unique().
			Positive(),
		field.String("name").
			StructTag(`json:"name"`).
			Comment("Team name").
			NotEmpty().
			Unique(),
		field.String("password").
			Comment("Team password").
			Sensitive().
			NotEmpty(),
		field.Int("id").
			StructTag(`json:"-"`),
		field.Enum("role").
			StructTag(`json:"role"`).
			Comment("User Permissions").
			Values("admin", "competitor").
			Default("competitor"),
	}
}

// Edges of the Team.
func (Team) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("status", Status.Type).
			StructTag(`json:"status,omitempty"`).
			Comment("Check statuses").
			Annotations(
				entsql.Annotation{
					OnDelete: entsql.Cascade,
				},
			).
			Ref("team"),
		edge.From("credential", Credential.Type).
			StructTag(`json:"credential,omitempty"`).
			Comment("Check credential").
			Annotations(
				entsql.Annotation{
					OnDelete: entsql.Cascade,
				},
			).
			Ref("team"),
	}
}
