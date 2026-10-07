# wh-suggestions

Order-history based product recommendations for Warehouse USAL: what to add to a
cart, what a user usually orders, and what is trending warehouse-wide.

Go service that consumes `order.completed` events from `wh-backend` over Kafka
(Redpanda), keeps its model in Redis, and serves it over HTTP. Design:
[RFC — wh-general#12](https://github.com/Warehouse-USAL/wh-general/issues/12).

> **Status: scaffold.** Configuration, routing, health check, CI/CD and deploy
> are in place. The Kafka consumer, Redis ingestion, JWT validation and the
> recommendation logic are not implemented yet; the recommendation endpoints
> answer `501`.

## API

| Method | Path                              | Status |
|--------|-----------------------------------|--------|
| `POST` | `/v1/recommendations/cart`        | 501    |
| `GET`  | `/v1/recommendations/user/:userId`| 501    |
| `GET`  | `/v1/recommendations/hot`         | 501    |
| `GET`  | `/healthz`                        | 200 (does not check Redis/Kafka yet) |

In production Caddy exposes the service under `/suggestions/*` (prefix stripped).

## Development

```bash
make help        # list commands
make up-dev      # service + Redis (needs wh-backend's Redpanda, see docker-compose.yml)
make pr-checks   # gofmt + go vet + tests + build — run before opening a PR
```

Configuration is environment-driven; see [`.env.example`](.env.example).

## Branching and releases

Same model as `wh-backend`:

- `develop` is the default branch. Day-to-day work goes through
  `feature/*`, `enhancement/*`, `fix/*`, `refactor/*` PRs (squash merge, 2 approvals).
- `master` is production. Only `beta/<date>` (release candidate) or
  `hotfix/<version>` branches may target it (merge commit, 1 approval).
- Pushing `beta/*` / `hotfix/*` publishes a pre-release; merging into `master`
  publishes the stable release, pushes `ghcr.io/warehouse-usal/wh-suggestions`,
  deploys through `wh-autodeploys`, and opens a backport PR to `develop`.
