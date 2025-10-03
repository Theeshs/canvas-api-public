package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Education holds the schema definition for the Education entity.
type Education struct {
	ent.Schema
}

// Fields of the Education.
func (Education) Fields() []ent.Field {
	return []ent.Field{
		field.Uint("id").SchemaType(map[string]string{"postgres": "serial"}),
		field.String("institute_name"),
		field.Time("start_date"),
		field.Time("end_date").Optional(),
		field.Uint("user_id"),
		field.String("mode_of_study"),
		field.String("degree_type"),
		field.String("area_of_study"),
		field.Bool("currently_studying").Optional(),
		field.String("description").Optional(),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

func (Education) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("educations").
			Unique().
			Required().
			Field("user_id"),
	}
}
func (Education) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: "education"}}
}
