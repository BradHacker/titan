// Code generated by entc, DO NOT EDIT.

package heartbeat

const (
	// Label holds the string label denoting the heartbeat type in the database.
	Label = "heartbeat"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldSentAt holds the string denoting the sentat field in the database.
	FieldSentAt = "sent_at"
	// FieldReceivedAt holds the string denoting the receivedat field in the database.
	FieldReceivedAt = "received_at"
	// EdgeAgent holds the string denoting the agent edge name in mutations.
	EdgeAgent = "agent"
	// Table holds the table name of the heartbeat in the database.
	Table = "heartbeats"
	// AgentTable is the table the holds the agent relation/edge.
	AgentTable = "heartbeats"
	// AgentInverseTable is the table name for the Agent entity.
	// It exists in this package in order to avoid circular dependency with the "agent" package.
	AgentInverseTable = "agents"
	// AgentColumn is the table column denoting the agent relation/edge.
	AgentColumn = "heartbeat_agent"
)

// Columns holds all SQL columns for heartbeat fields.
var Columns = []string{
	FieldID,
	FieldSentAt,
	FieldReceivedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "heartbeats"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"heartbeat_agent",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}