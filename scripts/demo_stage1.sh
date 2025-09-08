#!/bin/bash
set -e

echo "▶️ Building V1..."
go build -o v1_server ./cmd/stage1/v1
./v1_server & 
PID=$!
sleep 2

echo "▶️ Curl V1..."
curl -s localhost:8080

echo "🛑 Killing V1 for patch..."
kill $PID
sleep 1

echo "▶️ Building V2..."
go build -o v2_server ./cmd/stage1/v2
./v2_server & 
PID=$!
sleep 2

echo "▶️ Curl V2..."
curl -s localhost:8080

kill $PID
