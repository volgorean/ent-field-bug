package schema

import (
	bugfield "entgo.io/bug/field"
	"entgo.io/bug/other"
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int("age"),
		field.String("name"),
		field.JSON("metafield", bugfield.MetaField{}),
		field.JSON("otherfield", other.OtherField{}),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
