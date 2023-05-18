package schema

import (
	"entgo.io/ent"
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
			Comment("Name of the check").
			NotEmpty().
			Unique(),
		field.String("description").
			Comment("Description of the check").
			NotEmpty(),
		field.String("function").
			Comment("Function to run for the check").
			NotEmpty(),
		field.String("host").
			Comment("Host to run the check on").
			NotEmpty(),
	}
}

// Edges of the Check.
func (Check) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("checklogs", CheckLog.Type).Ref("check"),
	}
}
