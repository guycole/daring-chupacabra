# daring-chupacabra
[Discrete event simulation](https://en.wikipedia.org/wiki/Discrete-event_simulation) implemented in [golang](https://go.dev/) and deploys to [kubernetes](https://kubernetes.io/).  You can craft a starting state and let run, or clients cat interact with server via [gRPC](https://en.wikipedia.org/wiki/GRPC).

I needed this to support [Artificial Life](https://en.wikipedia.org/wiki/Artificial_life) experiments and perhaps a game or two.

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
    1. Life Cycle: objects are scheduled, created and optionally deleted.  All created/deleted tokens remain resident within CatalogMap during the simulation (deleted items are removed from other data structures).  
    1. Location: tokens can reside anywhere on a 2D grid of cells.  [LocationType](https://github.com/guycole/daring-chupacabra/blob/main/server/location.go) represents position within the 2D grid [CellArrayType](https://github.com/guycole/daring-chupacabra/blob/main/server/cell_array.go).  Only one token can occupy a [CellType](https://github.com/guycole/daring-chupacabra/blob/main/server/cell.go) on any given turn.  
1. Clock (Time/Turn)
    Time always moves forward.  Events can be scheduled up to 99 turns in advance (mod 100).  Events are kept within [event_array](https://github.com/guycole/daring-chupacabra/blob/main/server/event_array.go) which manages a list of [event_node](https://github.com/guycole/daring-chupacabra/blob/main/server/event_node.go).  Each event_node item represents a scheduled task.
1. Events
    Each scheduled event has at least one token and an action verb.  