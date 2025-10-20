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
type Document struct {
	ent.Schema
}

type DocumentType string

const (
	DocumentTypeResume      DocumentType = "resume"
	DocumentTypePassport    DocumentType = "passport"
	DocumentTypeIDCard      DocumentType = "id_card"
	DocumentTypeCertificate DocumentType = "certificate"
	DocumentTypeOther       DocumentType = "other"
)

// Fields of the Education.
func (Document) Fields() []ent.Field {
	return []ent.Field{
		field.Uint("id").SchemaType(map[string]string{"postgres": "serial"}),
		field.String("document_name"),
		field.Uint("user_id"),
		field.Enum("document_type").
			Values(
				string(DocumentTypeResume),
				string(DocumentTypeIDCard),
				string(DocumentTypeCertificate),
				string(DocumentTypeOther),
			).
			Default(string(DocumentTypeOther)),
		field.String("google_id"),
		field.String("document_web_link").Optional(),
		field.String("document_thumnail_link").Optional(),
		field.String("document_export_link").Optional(),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

func (Document) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("documents").
			Unique().
			Required().
			Field("user_id"),
	}
}
func (Document) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: "document"}}
}
