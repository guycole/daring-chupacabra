# daring-chupacabra
[Discrete event simulation](https://en.wikipedia.org/wiki/Discrete-event_simulation) implemented in [golang](https://go.dev/) and deploys to [kubernetes](https://kubernetes.io/).  You can craft a starting state and let run, or clients can interact with server via [gRPC](https://en.wikipedia.org/wiki/GRPC).

## Goals
1. Portable communication via [gRPC](https://en.wikipedia.org/wiki/GRPC)
1. Deploy to [kubernetes](https://kubernetes.io/)
1. Minimal user interface (CLI)
1. Asynchronous or synchronous operation
    1. Turns can be incremented via grpc or run continuously

## Data Structure
1. State
    1. Simulation objects are represented by tokens which have a unique ID, a specialization (type) and a life cycle state.
    1. Existence: all tokens hava a [CatalogItemType](https://github.com/guycole/daring-chupacabra/blob/main/server/catalog.go) within the [CatalogMap](https://github.com/guycole/daring-chupacabra/blob/main/server/catalog.go).  
    1. Life Cycle: objects are scheduled, created and optionally deleted.  All tokens (even deleted) remain resident within CatalogMap during the simulation (deleted items are removed from other data structures).  
    1. Location: a 2D grid of [LocationType](https://github.com/guycole/daring-chupacabra/blob/main/server/location.go) represents position within a 2D grid. 
1. Clock (Time/Turn)
    Time always moves forward.  Events can be scheduled up to 99 turns in advance (mod 100).  Events are kept within [event_array](https://github.com/guycole/daring-chupacabra/blob/main/server/event_array.go) which manages a list of [event_node](https://github.com/guycole/daring-chupacabra/blob/main/server/event_node.go).  Each event_node item represents a scheduled task.
1. Events (Tasks)
    Each scheduled event has a token and an action verb.

## Action
1. createAction: create specified object
1. deleteAction: delete specified object
1. moveAction: move specified object
1. nominalAction: turn housekeeping for object
1. nothingAction: do nothing
1. parseAction: parse a command from grpc client
1. scanAction: report on objects around requestor
1. statusAction: report object status