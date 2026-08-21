# Distributed Storage Control-Plane Prototype

A Go prototype for the control plane of a distributed file-storage system. It explores node registration, heartbeats, capacity tracking, chunk placement, replica metadata, and dead-node detection with a PostgreSQL-backed master service.

## What this project demonstrates

- Separate master and storage-node services
- Node registration, heartbeat timestamps, and liveness status
- Capacity and used-space metadata for storage nodes
- File metadata split into ordered chunk records
- Round-robin placement decisions across active nodes
- Replica acknowledgement counts and a background dead-node checker
- REST handlers built with Fiber and PostgreSQL repositories built with `sqlx`

## Architecture

```text
client -> master API -> PostgreSQL metadata
                    -> choose active nodes for each chunk

storage node -> register/heartbeat/acknowledge -> master API
```

The `master` service owns the metadata and placement decisions. Its client module creates file and chunk records, while its node module registers storage nodes, records heartbeats, updates capacity, and marks stale nodes as dead.

The `node` directory contains a separate service skeleton intended to register with the master and send heartbeats. The actual chunk upload/download data plane is not complete, so this repository should be read as a control-plane and metadata experiment rather than a working distributed filesystem.

## Run locally

Requirements: Go 1.26.3+ and PostgreSQL. Configure each service through its environment file before running.

Start the master:

```bash
cd master
go mod download
go run ./src
```

Start the node in another terminal:

```bash
cd node
go mod download
go run ./src
```

The master exposes generated Swagger documentation under `/swagger/`.

## Limitations

- Chunk bytes are not stored or transferred by the current node implementation.
- Replication metadata exists, but automatic replica copying and recovery are incomplete.
- Node-to-master registration, heartbeat, and acknowledgement paths are still partial in the node service.
- There is no consensus, leader election, checksum verification, authentication, or transport security.
- Failure, concurrency, and integration behavior is not covered by a comprehensive automated test suite.
- The committed environment files are development configuration and should be converted to non-secret examples before reuse.
