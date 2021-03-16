package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Instruction holds the schema definition for the Instruction entity.
type Instruction struct {
	ent.Schema
}

// Fields of the Instruction.
func (Instruction) Fields() []ent.Field {
	return []ent.Field{
		field.Time("sentAt"),
	}
}

// Edges of the Instruction.
func (Instruction) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("agent", Agent.Type).Unique().Required(),
		edge.To("action", Action.Type).Unique().Required(),
		edge.From("beacon", Beacon.Type).Ref("instruction").Unique(),
	}
}
