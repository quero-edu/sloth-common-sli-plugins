# Traefik V2 availability by router

Availability plugin for [Traefik V2][traefik] routers.

Uses Traefik v2 router metrics to get the correct and invalid availability on the serving routers.

## Options

- `filter`: (**Optional**) A prometheus filter string using concatenated labels
- `router_name_regex`: (**required**) Regex to match the traefik routers.

## Metric requirements

- `traefik_router_requests_total`: From [traefik].

## Usage examples

### Without filter

```yaml
sli:
  plugin:
    id: "sloth-common/traefik/v2/availability-by-router"
    options:
      router_name_regex: "^default-slok-sloth$"
```

### With filters

```yaml
sli:
  plugin:
    id: "sloth-common/traefik/v2/availability-by-router"
    options:
      router_name_regex: "^default-slok-sloth$"
      filter: method="GET"
```

[traefik]: https://doc.traefik.io/traefik/v2.6/
