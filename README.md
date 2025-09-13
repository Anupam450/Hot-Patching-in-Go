# 🔥 Hot Patching in Go – Demo Repo

## Run the demo

```bash
./scripts/demo_server.sh
```

## Curls for testing

```bash

# Build and start server
go build -o server ./cmd/server
./server

# Build .so for v2.go
go build -buildmode=plugin -o v2.so ./plugin/v2

# Apply hot patch
curl -s localhost:8080/admin/patch

# List available endpoints
curl http://localhost:8080/

# Addition
curl "http://localhost:8080/add?a=10&b=7"

# Subtraction
curl "http://localhost:8080/subtract?a=10&b=7"

# Multiplication
curl "http://localhost:8080/multiply?a=10&b=7"

# Division
curl "http://localhost:8080/divide?a=10&b=2"
```
