package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Beacon holds the schema definition for the Beacon entity.
type Beacon struct {
	ent.Schema
}

// Fields of the Beacon.
func (Beacon) Fields() []ent.Field {
	return []ent.Field{
		field.Time("sentAt"),
		field.Time("receivedAt").Nillable().Optional(),
	}
}

// Edges of the Beacon.
func (Beacon) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("instruction", Instruction.Type).Unique().Required(),
	}
}
