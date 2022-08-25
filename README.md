# daring-chupacabra
Multiuser simulation framework which is implemented in [golang](https://go.dev/) and deploys to [kubernetes](https://kubernetes.io/).

## Goals
1. Enable the creation of multiplayer ("real time strategy")[https://en.wikipedia.org/wiki/Real-time_strategy] games.
1. Maintains the game board, game tokens, scores and resolves conflict, etc.
1. Asynchronous status updates to players about game progress.
1. Implement in [golang](https://go.dev/)
1. Deploy to [kubernetes](https://kubernetes.io/).
    1. Development in [minikube](https://minikube.sigs.k8s.io/docs/start/).
    1. Should support a [LAN party](https://en.wikipedia.org/wiki/LAN_party) scenario.
1. Web based service (HTTP protocol)
    1. [Web Sockets](https://en.wikipedia.org/wiki/WebSocket)
    1. [gorilla websockets](https://github.com/gorilla/websocket)
1. Monitoring
    1. Expose [prometheus](https://prometheus.io) scrape target

## Not Goals
1. User interface.
1. Game board, etc.

## Implementation
1. Three main components: back end (BE), front end (FE), message broker (redis)
    1. Messages between front and back end are exclusively shared using (Redis)[https://redis.com/] pub/sub.
    1. front end manages the web sockets and proxys messages between web client and back end via Redis pub/sub.
    1. back end reads manages actual game play, and communicates results to front end via Redis.

## Deployment
1. Daring Chupacabra deploys as 3 pods: Redis, Front End (FE) and Back End (BE)

## Message Flow
1. FE is exposed to via kubernetes ingress, communicates to BE via Redis pub/sub
    1. JSON formatted traffic arrives to FE via ingress and is converted to an binary format and then written to a Redis pub/sub channel.  
    1. BE reads all messages from a single Redis pub/sub channel and schedules events for execution.  As events occur, status updates are written to a Redis pub/sub channel for relay to remote clients.  Each remote client has a dedicated pub/sub channel for event updates. 
    1. FE consumes status updates, converts to JSON and writes event to remote client.

## Example
1.  Move ships around 2D grid.