# distributed-storage

This repository implements chain replication, a technique for primary-backup object replication in distributed servers.

The system runs using docker, you can run the system using the command `make serve`

## Components

- **Client**: The user facing part of the system, creates requests to read and update data in the system

- **Node**: A member of the chain (a server) which when put together becomes a network of replicated object storage

- **Master**: The co-ordinate of the chain, handling health checking, service discovery, chain management

## Getting Started

- Build the service images `make build`

- Start the chain services `make serve`

- Execute storage requests `make request`

### Using the system

Clients can directly interface with the chain by making RPC requests to the master for both read and write operations. To save you the effort of creating these you can instead run Makefile commands to mock these actions, these are:

- Writing a key-value: `make request OP=write KEY=<KEY> VALUE=<value>`
- Reading a value: `make request OP=read KEY=<KEY>`

The system is made to be dependable and therefore supports a many-node chain. Nodes are expected to die and be restarted throughout the chain's lifetime, this can be mimicked through the following commands:

- You can mimic server failures and restarts by manually killing/starting the node containers: `docker-compose (kill/start) node-#` where # is the node you want to change (0, 1 or 2)

## Requirements

- GNU Make

- Docker

- Docker Compose (1.28.5+)

## Tests

To run system tests run the following Go command: `make test` from the project root. This command requires you to install the Go programming language (version 1.15) as the tests run using the Golang testing library. Tests can be found within the /testing/main_test.go file. The tests cover all the system and functional requirements.
