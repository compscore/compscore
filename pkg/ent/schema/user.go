package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").
			Comment("ID of the user").
			StructTag(`json:"id"`).
			Positive().
			Immutable().
			Unique(),
		field.String("name").
			Comment("Name of the user").
			StructTag(`json:"name"`).
			Unique().
			NotEmpty(),
		field.String("password").
			Comment("Password of the user").
			Sensitive().
			NotEmpty(),
		field.Int("team_number").
			Comment("Team number of the user").
			StructTag(`json:"team_number"`).
			Optional().
			Unique(),
		field.Enum("role").
			Comment("Role of the user").
			StructTag(`json:"role"`).
			Values(
				"competitor",
				"admin",
			).
			Default("competitor"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("credential", Credential.Type).
			Comment("Credential of the user").
			StructTag(`json:"credential,omitempty"`).
			Annotations(
				entsql.Annotation{
					OnDelete: entsql.Cascade,
				},
			).
			Ref("user"),
		edge.From("status", Status.Type).
			Comment("Status of the user").
			StructTag(`json:"status,omitempty"`).
			Annotations(
				entsql.Annotation{
					OnDelete: entsql.Cascade,
				},
			).
			Ref("user"),
	}
}
