# daring-chupacabra
[Discrete event simulation](https://en.wikipedia.org/wiki/Discrete-event_simulation) implemented in [golang](https://go.dev/) and deploys to [kubernetes](https://kubernetes.io/).  Clients submit events and poll for state updates via [gRPC](https://en.wikipedia.org/wiki/GRPC).

## Introduction
1.  Stuff

## Goals
1. Asynchronous or synchronous operation
1. Portable communication via [gRPC](https://en.wikipedia.org/wiki/GRPC)
1. Deploy to [kubernetes](https://kubernetes.io/)
1. Minimal user interface
1. Optional: tate saved to Redis](https://redis.com/) pub/sub
1. Expose [prometheus](https://prometheus.io) scrape target

## Data Structure
1. State
    1. Simulation objects are represented by tokens which have a unique ID, a specialization (type) and a life cycle state.
    1. Existence: all tokens hava a [CatalogItemType](https://github.com/guycole/daring-chupacabra/blob/main/server/catalog.go) within the [CatalogMap](https://github.com/guycole/daring-chupacabra/blob/main/server/catalog.go).  
    1. Life Cycle: objects are scheduled, created and optionally deleted.  All tokens remain resident within CatalogMap during the simulation.  
    1. Location: tokens can reside anywhere on a 2D grid of cells.  [LocationType](https://github.com/guycole/daring-chupacabra/blob/main/server/location.go) represents position within the 2D grid [CellArrayType](https://github.com/guycole/daring-chupacabra/blob/main/server/cell_array.go).  Only one token can occupy a [CellType](https://github.com/guycole/daring-chupacabra/blob/main/server/cell.go).  
1. Clock (Time)
    Time always moves forward.  Events can be scheduled up to 99 turns in advance (mod 100).  Events are kept within [event_array](https://github.com/guycole/daring-chupacabra/blob/main/server/event_array.go) which manages a list of [event_node](https://github.com/guycole/daring-chupacabra/blob/main/server/event_node.go).  Each event_node item represents a scheduled event process.
1. Events
    Each scheduled event has at least one token and an action verb.  Tokens can interract.