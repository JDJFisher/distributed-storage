# distributed-storage

This repository implements chain replication, a technique for primary-backup object replication in distributed servers.

The system runs using docker, you can run the system using the command `make serve`

## Components

-   **Client**: The user facing part of the system, creates requests to read and update data in the system
-   **Node**: A member of the chain (a server) which when put together becomes a network of replicated object storage
-   **Master**: The co-ordinate of the chain, handling health checking, service discovery, chain management

## Getting Started

- Build the service images `make build`
- Start the chain services `make serve`
- Execute storage requests `make request`

## Requirements

- GNU Make
- Docker
- Docker Compose
