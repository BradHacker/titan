package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Heartbeat holds the schema definition for the Heartbeat entity.
type Heartbeat struct {
	ent.Schema
}

// Fields of the Heartbeat.
func (Heartbeat) Fields() []ent.Field {
	return []ent.Field{
		field.Time("sentAt"),
		field.Time("receivedAt"),
	}
}

// Edges of the Heartbeat.
func (Heartbeat) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("agent", Agent.Type).Unique().Required(),
	}
}
