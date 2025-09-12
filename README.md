# 🔥 Hot Patching in Go – Demo Repo

This repo demonstrates two approaches:

## Stage 1 – Downtime

- V1 running, must stop process to apply V2.
- Shows visible downtime.

## Stage 2 – Hot Patch

- V1 starts (with bug).
- After bug detection, we compile `plugin/v2/v2.go` into `v2.so`.
- Patch applied dynamically at runtime → **zero downtime**.

---

## Run the demos

```bash
./scripts/demo_stage1.sh
./scripts/demo_server.sh
```

## Curls for testing

```bash

# Build .so for v2.go
go build -buildmode=plugin -o v2.so ./plugin/v2

# Apply hot patch
curl -s localhost:8080/admin/patch

# List available endpoints (now add, subtract, multiply, divide)
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
