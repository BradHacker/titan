package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Agent holds the schema definition for the Agent entity.
type Agent struct {
	ent.Schema
}

// Fields of the Agent.
func (Agent) Fields() []ent.Field {
	return []ent.Field{
		field.String("uuid"),
		field.String("hostname"),
		field.String("ip"),
		field.String("port"),
		field.Int("pid"),
	}
}

// Edges of the Agent.
func (Agent) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("instruction", Instruction.Type).Ref("agent"),
		edge.From("heartbeat", Heartbeat.Type).Ref("agent"),
	}
}
