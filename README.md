# Primary-Backup

This repository implements chain replication, a technique for primary-backup object replication in distributed servers.

Most of the system is written in Go due to it being a language made for systems like this.

The system runs using docker, you can run the system using the command `make run`

## Components

-   **Client**: The user facing part of the system, creates requests to read and update data in the system
-   **Node**: A member of the chain (a server) which when put together becomes a network of replicated object storage
-   **Master**: The co-ordinate of the chain, handling health checking, service discovery, chain management
