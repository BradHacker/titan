# Data Models

This is a layout of the data models being used.

```
# What kind of action do we want to execute

ActionType {
  EXEC
  # Others in the future
}

# The action we want to perform and optional

Action {
  type: ActionType;
  cmd: string;
  output: string;
}

# Represents the agent running on the target

Agent {
  hostname: string;
  ip: string;
  pid: number;
}

# Instructions sent by the C2 to the agent

Instruction {
  agent: Agent;
  action: Action;
  sentAt: DateTime;
}

# Response after completion of an action by the agent, sent to the C2

Beacon {
  agent: Agent;
  action: Action;
  sentAt: DateTime;
  recievedAt: DateTime;
  instruction: Insruction;
}
```

## References

- [Json in Golang](https://gobyexample.com/json)
