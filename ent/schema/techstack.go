package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type TechSctack struct {
	ent.Schema
}

func (TechSctack) Fields() []ent.Field {
	return []ent.Field{
		field.Uint("id").SchemaType(map[string]string{"postgres": "serial"}),
		field.String("name"),
		field.Uint("skill_id"),
		field.Uint("user_id"),
		field.Time("created_at").Optional().Default(time.Now),
		field.Time("updated_at").Optional().Default(time.Now),
	}
}

func (TechSctack) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("skill", Skill.Type).
			Required().
			Unique().
			Field("skill_id"),

		edge.To("user", User.Type).
			Required().
			Unique().
			Field("user_id"),
	}
}

func (TechSctack) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "techstack"},
	}
}
