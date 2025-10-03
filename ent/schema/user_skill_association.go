package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type UserSkillAssociation struct {
	ent.Schema
}

func (UserSkillAssociation) Fields() []ent.Field {
	return []ent.Field{
		field.Uint("id").SchemaType(map[string]string{"postgres": "serial"}),
		// field.Uint("user_id"),
		field.Uint("experience_id"),
		field.Uint("skill_id"),
		field.Int32("percentage").Optional(),
		field.Time("created_at").Optional().Default(time.Now),
		field.Time("updated_at").Optional().Default(time.Now),
	}
}

func (UserSkillAssociation) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("experience", Experience.Type).
			Ref("user_skill_association").
			Unique().
			Required().
			Field("experience_id"),
		// Belongs to one Skill
		edge.From("skill", Skill.Type).
			Ref("user_skill_association").
			Unique().
			Required().
			Field("skill_id"),
	}
}

func (UserSkillAssociation) Indexes() []ent.Index {
	return []ent.Index{
		// Ensure unique combinations of user_id, experience_id, and skill_id
		index.Fields("experience_id", "skill_id").Unique(),
	}
}

func (UserSkillAssociation) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "user_skill_association"},
	}
}
