# Load Testing (JMeter)

This folder contains JMeter scenarios for local load testing.

## Files

- `Summary Report.jmx` - load test plan for `GET /api/v1/satellite/moon`

## Prerequisites

- Docker Desktop is running
- Apache JMeter 5.6.3+ installed

## 1) Start the service locally

From project root (`awesomeProject`):

```powershell
docker compose up -d --build
```

Service endpoint from host:

- `http://localhost:8082/api/v1/satellite/moon`

Quick check:

```powershell
curl.exe -i "http://localhost:8082/health/check"
curl.exe -i "http://localhost:8082/api/v1/satellite/moon"
```

## 2) Run in JMeter GUI (for debugging)

1. Open JMeter GUI (`jmeter.bat`)
2. Open `load_tests/Summary Report.jmx`
3. Run test (`Start`)
4. Inspect metrics in `Aggregate Report` / `Summary Report`

## 3) Run in CLI (recommended for load runs)

Create result directory once:

```powershell
mkdir load_tests\results -Force
```

Run test:

```powershell
jmeter -n -t "load_tests\Summary Report.jmx" -l "load_tests\results\satellite_load.jtl" -e -o "load_tests\results\report"
```

## 4) Useful metrics to track

- `Error %` (target: close to `0%`)
- `Throughput` (req/sec)
- `95% Line` and `99% Line` latency
- `Max` latency spikes

## 5) Stop local stack

```powershell
docker compose down
```
