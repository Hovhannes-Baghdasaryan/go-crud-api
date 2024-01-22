package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// TagEntity holds the schema definition for the TagEntity entity.
type TagEntity struct {
	ent.Schema
}

// Annotations of the Tag.
func (TagEntity) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "tags"},
	}
}

// Fields of the TagEntity.
func (TagEntity) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("uuid", uuid.UUID{}).
			Default(uuid.New),
		field.String("name"),
	}
}
