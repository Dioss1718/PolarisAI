Write-Host "Starting PolarisAI full demo stack..." -ForegroundColor Cyan

$root = Split-Path -Parent $PSScriptRoot

Write-Host "Repo root: $root" -ForegroundColor Yellow

# 1. Synthetic engine
Write-Host "Launching synthetic-engine..." -ForegroundColor Green
Start-Process powershell -ArgumentList "-NoExit", "-Command", "cd '$root\synthetic-engine'; node app.js"

# 2. AI ingest (one-time blocking step)
Write-Host "Skipping ingestion (run manually once if needed)"

# 3. AI uvicorn server
Write-Host "Launching AI engine API..." -ForegroundColor Green
Start-Process powershell -ArgumentList "-NoExit", "-Command", "cd '$root\ai-engine'; uvicorn main:app --reload --port 8000"

# 4. Forecast service
Write-Host "Launching forecast service..." -ForegroundColor Green
Start-Process powershell -ArgumentList "-NoExit", "-Command", "cd '$root\forecast'; python forecast_service.py"

# 5. Go orchestrator
Write-Host "Launching Go orchestrator..." -ForegroundColor Green
Start-Process powershell -ArgumentList "-NoExit", "-Command", "cd '$root\cmd\polarisAI'; go run main.go"

# 6. Frontend
Write-Host "Launching frontend..." -ForegroundColor Green
Start-Process powershell -ArgumentList "-NoExit", "-Command", "cd '$root\frontend'; npm run dev"

Write-Host ""
Write-Host "PolarisAI demo stack launch sequence completed." -ForegroundColor Cyan
Write-Host "Suggested demo scenario: FULL_CHAOS with seed 42" -ForegroundColor Yellow