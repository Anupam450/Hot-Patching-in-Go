#!/bin/bash
set -e

echo "▶️ Building Stage2 server..."
go build -o server ./cmd/stage2

echo "▶️ Starting server..."
./server &
PID=$!
sleep 3

echo "▶️ Curl before patch..."
curl -s localhost:8080

echo "⏳ Simulating bug detection..."
sleep 12

echo "▶️ Building v2 plugin..."
go build -buildmode=plugin -o v2.so ./plugin/v2

echo "▶️ Applying patch..."
curl -s localhost:8080/admin/patch

echo "▶️ Curl after patch..."
curl -s localhost:8080

kill $PID
