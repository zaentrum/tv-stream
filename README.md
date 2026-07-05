# tv-stream

Streaming origin for the **tv** product: a small Go service that bridges a live
MPEG-TS multicast source to HLS for playback clients. The hot path uses a
stream-copy remux (no re-encode) so it stays cheap and low-latency, with
transcoding reserved as a fallback only.

> **Status: early scaffold.** This is a fresh Go skeleton — a chi router with
> `/healthz`, Prometheus `/metrics`, and a stub `/play/*` handler that returns
> `501 Not Implemented`. The real bridge logic lands in a later phase.

## Layout

```
cmd/server        entrypoint (graceful shutdown)
internal/config   env-driven configuration
internal/http     router wiring (health, metrics, /play mount)
internal/play     placeholder playback surface
k8s               Deployment / Service / ServiceMonitor / dashboard manifests
```

## Local development

```
go run ./cmd/server
curl http://localhost:8080/healthz     # -> ok
curl http://localhost:8080/metrics     # Prometheus metrics
curl http://localhost:8080/play/foo    # -> 501, scaffold only
```

Configuration is via environment variables:

| Variable | Default  | Purpose             |
| -------- | -------- | ------------------- |
| `ADDR`   | `:8080`  | HTTP listen address |

## Container image

The `Dockerfile` builds a static, distroless, non-root image. By default it
pulls its base image from public Docker Hub; a private build can override the
registry prefix:

```
docker build -t tv-stream .
docker build --build-arg BASE=registry.example/library/ -t tv-stream .
```

Published images are available at `ghcr.io/zaentrum/tv-stream`.

## License

[MPL-2.0](LICENSE).
