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
./scripts/demo_stage2.sh
