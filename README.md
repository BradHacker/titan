# Project Titan

> **_DISCLAIMER: This tool is for educational, competition, and training purposes only. I am in no way responsible for any abuse of this tool_**

Multi-Platform red team tooling designed for remote monitoring and attacking of target hosts. Titan is designed to be scalable and expandable with a plug and play module based system where new functionality can be easily dropped in for different communication methods or actions.

> _Note: This tool is currently in development._

## Structure

### zeke

Zeke is the C2, the man in charge of this whole operation. This is where all server side code is written and beacon heartbeats are handled. The database is also managed by this code.

Zeke contains its own `Transport` module. This handles all communication to/from beacons as well as serialization/deserialization of data into Golang structs.

#### Adding Functionality to `transport`

Functionality can be expanded upon here by adding more communication protocols, encryption methods, etc. Anything that has to do with communcation can be modified and/or expanded upon here.

### reiner

Reiner is the Beacon, the footsoldier/muscle of the operation. This is where all client side code is written and instructions sent from the C2 are handled. Currently, Remote Code Execution (RCE) is also handled here.

Reiner contains its own `Transport` module. This transport module, while similar to Zeke's, contains client side communication code in order to connect back to the C2 server. It also handles serialization/deserialization as well.

Reiner also has a `Handler` module. This module is designed to interpret the instructions sent by the C2 server. This is where any RCE, Automated Script, Persistence, etc. happens. Each of these actions has it's own file with it's own functionality.

#### Adding Functionality to `transport`

Functionality can be expanded upon here by adding more communication protocols, encryption methods, etc. Anything that has to do with communcation can be modified and/or expanded upon here.

#### Adding Functionality to `handler`

Functionality can be added here via switch cases on the `ActionType` enum. `Handle___(instruction models.Instruction)` functions should be added to adapt to each handler module (ex. `HandleRCE(instruction models.Instruction)`).

## Project Goals

This project was ultimately started to learn more about offensive security tooling, but ultimately I wanted to make something that brought together good code structure/practice as well as top-notch functionality.

Goals:

- Scalability
- Expandability
- Cool factor x100 😎

## Future Plans/Ideas

- [ ] ENT Database integration
- [ ] Communication protocols other than TCP
- [ ] Data formats other than JSON
- [ ] Persistence handlers for different OS's
- [ ] Reconnaissance?
