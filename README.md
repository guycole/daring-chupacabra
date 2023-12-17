# daring-chupacabra
[Discrete event simulation](https://en.wikipedia.org/wiki/Discrete-event_simulation) implemented in [golang](https://go.dev/) and deploys to [kubernetes](https://kubernetes.io/).  Clients submit events and poll for state updates via [gRPC](https://en.wikipedia.org/wiki/GRPC).

## Goals
1. Asynchronous or synchronous operation
1. Portable communication via [gRPC](https://en.wikipedia.org/wiki/GRPC)
1. Deploy to [kubernetes](https://kubernetes.io/)
1. Minimal user interface
1. Optional: tate saved to Redis](https://redis.com/) pub/sub
1. Expose [prometheus](https://prometheus.io) scrape target
