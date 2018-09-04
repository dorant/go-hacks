# Cluster Redis

Connecting a Redis client to a Redis cluster.

## Getting Started

### Start Redis

Start 6 Redis instances, and make a cluster of them:

```
make start
```

Verify cluster is up:

```
make status
```

### Run Go client

```
make run
```

Client will connect, set a key/value and fetch it again.

### Run redis-cli

```
make run-cli
make run-cluster-cli
```

run-cli will connect to a master, but probably the wrong one, getting MOVED back.
run-cluster-cli will connect to a master, and will be forwarded to correct master returning the result.

### Stop cluster

```
make stop
```
