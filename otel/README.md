# Open-telemetry observability

Sample configuration for Kbot that send logs to [OpenTelemetry Collector] and metrics to [OpenTelemetry Collector] or [Prometheus].

## Prerequisites

- [Docker]
- [Docker Compose]

## How to run

```bash
 read -s TELE_TOKEN
 export TELE_TOKEN

 docker-compose -f otel/docker-compose.yaml up

 docker-compose -f otel/docker-compose.yaml down
```
Open Grafana in the browser on port 3002