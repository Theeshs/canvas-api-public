package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Project struct {
	ent.Schema
}

func (Project) Fields() []ent.Field {
	return []ent.Field{
		field.Uint("id").SchemaType(map[string]string{"postgres": "serial"}),
		field.String("project_name"),
		field.String("description").Optional(),
		field.String("url"),
		field.Uint("user_id"),
		field.Time("created_at").Optional().Default(time.Now),
		field.Time("updated_at").Optional().Default(time.Now),
	}
}

func (Project) Edges() []ent.Edge {
	return []ent.Edge{
		// Experience belongs to one User
		edge.From("user", User.Type).
			Ref("project").
			Unique().
			Required().
			Field("user_id"),
		// Experience can have multiple UserSkillAssociations (many skills)
		edge.From("skill", Skill.Type).Ref("project"),
	}
}

func (Project) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "project"},
	}
}
