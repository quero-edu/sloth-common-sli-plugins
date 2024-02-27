# Traefik V2 latency

Latency plugin for [Traefik V2][traefik] routers.

Uses Traefik v2 router metrics to get the latency on the serving routers.

## Options

- `bucket`: (**Required**) The max latency allowed histogram bucket.
- `router_name_regex`: (**required**) Regex to match the traefik router.
- `filter`: (**Optional**) A prometheus filter string using concatenated labels
- `exclude_errors`: (**Optional**) Boolean that will exclude errored requests from valid events when measuring latency requests.

## Metric requirements

- `traefik_router_request_duration_seconds_bucket`: From [traefik].
- `traefik_router_request_duration_seconds_count`: From [traefik].

## Usage examples

### Without filter

```yaml
sli:
  plugin:
    id: "sloth-common/traefik/v2/latency-by-router"
    options:
      router_name_regex: "^default-slok-sloth$"
      bucket: "0.3"
```

### With filters

```yaml
sli:
  plugin:
    id: "sloth-common/traefik/v2/latency-by-router"
    options:
      router_name_regex: "^default-slok-sloth$"
      bucket: "0.3"
      filter: method="GET"
```

### Excluding errors (5xx)

```yaml
sli:
  plugin:
    id: "sloth-common/traefik/v2/latency-by-router"
    options:
      router_name_regex: "^default-slok-sloth$"
      bucket: "0.3"
      filter: method="GET"
      exclude_errors: "true"
```

[traefik]: https://doc.traefik.io/traefik/v2.6/
