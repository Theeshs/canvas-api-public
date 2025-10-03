package schema

import (
	"time"

	"entgo.io/ent"
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
		field.Uint("id").SchemaType(map[string]string{"postgres": "serial"}),
		field.String("first_name").Optional(),
		field.String("last_name").Optional(),
		field.Time("dob").Optional(),
		field.String("username"),
		field.String("password"),
		field.String("email"),
		field.String("github_username").Optional(),
		field.String("description").Optional(),
		field.Time("created_at").Optional().Default(time.Now),
		field.Time("updated_at").Optional().Default(time.Now),
		field.Int32("mobile_number").Optional(),
		field.String("address_block").Optional(),
		field.String("address_street").Optional(),
		field.String("recidential_country").Optional(),
		field.String("nationality").Optional(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("educations", Education.Type),
		edge.To("experiences", Experience.Type),
	}
}
