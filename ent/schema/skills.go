package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Skill struct {
	ent.Schema
}

func (Skill) Fields() []ent.Field {
	return []ent.Field{
		field.Uint("id").SchemaType(map[string]string{"postgres": "serial"}),
		field.String("name").Unique(),
		field.Time("created_at").Optional().Default(time.Now),
		field.Time("updated_at").Optional().Default(time.Now),
	}
}

func (Skill) Edges() []ent.Edge {
	return []ent.Edge{
		// Skill can be associated with multiple UserSkillAssociations
		edge.To("user_skill_association", UserSkillAssociation.Type),
	}
}

func (Skill) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "skill"},
	}
}
