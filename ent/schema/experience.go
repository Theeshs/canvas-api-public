package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Experience struct {
	ent.Schema
}

func (Experience) Fields() []ent.Field {
	return []ent.Field{
		field.Uint("id").SchemaType(map[string]string{"postgres": "serial"}),
		field.String("company_name"),
		field.Time("start_date"),
		field.Time("end_date").Optional(),
		field.Bool("current_place").Optional(),
		field.String("position"),
		field.Uint("user_id"),
		field.Time("created_at").Optional().Default(time.Now),
		field.Time("updated_at").Optional().Default(time.Now),
		field.String("description").Optional(),
	}
}

func (Experience) Edges() []ent.Edge {
	return []ent.Edge{
		// Experience belongs to one User
		edge.From("user", User.Type).
			Ref("experiences").
			Unique().
			Required().
			Field("user_id"),
		// Experience can have multiple UserSkillAssociations (many skills)
		edge.To("user_skill_association", UserSkillAssociation.Type),
	}
}

func (Experience) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "experience"},
	}
}
