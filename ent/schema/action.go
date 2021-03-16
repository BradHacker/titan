package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Action holds the schema definition for the Action entity.
type Action struct {
	ent.Schema
}

// Fields of the Action.
func (Action) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("actionType").
			Values("EXEC"),
		field.String("cmd"),
		field.Strings("args"),
		field.String("output").Optional().Nillable(),
	}
}

// Edges of the Action.
func (Action) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("instruction", Instruction.Type).Ref("action").Unique(),
	}
}
