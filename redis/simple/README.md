# Simple Redis client

Connecting a Redis client to a Dockerized Redis instance.

## Getting Started

### Start Redis

Startup instance:

```
make start
```

Verify Redis is running:

```
make logs
```

### Run client

```
make run
```

Client will connect, set a key/value and fetch it again.

### Stop Redis

```
make stop
```
