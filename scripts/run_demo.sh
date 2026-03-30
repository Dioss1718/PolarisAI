#!/usr/bin/env bash
set -e

echo "Starting PolarisAI full demo stack..."

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
echo "Repo root: $ROOT"

echo "Launching synthetic-engine..."
gnome-terminal -- bash -c "cd '$ROOT/synthetic-engine'; node app.js; exec bash" || true

echo "Running AI ingestion..."
echo "Skipping ingestion (run manually once if needed)"



echo "Launching AI engine API..."
gnome-terminal -- bash -c "cd '$ROOT/ai-engine'; uvicorn main:app --reload --port 8000; exec bash" || true

echo "Launching forecast service..."
gnome-terminal -- bash -c "cd '$ROOT/forecast'; python forecast_service.py; exec bash" || true

echo "Launching Go orchestrator..."
gnome-terminal -- bash -c "cd '$ROOT/cmd/polarisAI'; go run main.go; exec bash" || true

echo "Launching frontend..."
gnome-terminal -- bash -c "cd '$ROOT/frontend'; npm run dev; exec bash" || true

echo ""
echo "PolarisAI demo stack launch sequence completed."
echo "Suggested demo scenario: FULL_CHAOS with seed 42"