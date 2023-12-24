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
  1. Existence: all simulation objects hava a [CatalogItemType](https://github.com/guycole/daring-chupacabra/blob/main/server/catalog.go) within the [CatalogMap](https://github.com/guycole/daring-chupacabra/blob/main/server/catalog.go).  
  1. Life Cycle: objects are created, exist and optionally deleted.  
  1. Location: 
1. Clock (Time)
1. Events