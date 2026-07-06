# Go race detector CI templates

Reusable CI jobs that run `go test -race` with JUnit report output.

Battle-tested pattern for catching data races before they reach production. Supports **GitLab CI** and **GitHub Actions**.

## GitLab CI

### Option A — include from this repository (after publishing)

```yaml
include:
  - project: waqazali/go-race-ci-template
    ref: main
    file: /templates/go/test-race.yml
```

### Option B — copy templates into your project

Copy `templates/` into your repo, then in `.gitlab-ci.yml`:

```yaml
include:
  - local: /templates/defaults.yml
  - local: /templates/go/test-race.yml
```

See [examples/gitlab-ci.yml](examples/gitlab-ci.yml).

### Variables

| Variable | Default | Description |
|----------|---------|-------------|
| `TEST_PATH` | `./...` | Packages to test |
| `TEST_ARG` | `""` | Extra arguments passed to `go test` |
| `CGO_ENABLED` | `1` | Required for `-race` on most platforms |

The job sets `allow_failure: true` by default so race detection can run as an advisory check. Remove that line to make it blocking.

## GitHub Actions

Copy [.github/workflows/go-race.yml](.github/workflows/go-race.yml) into your repository, or use it as a reusable workflow.

## Requirements

- Go 1.20+ (race detector)
- `CGO_ENABLED=1` and a C compiler (Alpine: `gcc build-base`; Ubuntu runners include gcc)
- Tests must be runnable with `go test -race`

## License

MIT — see [LICENSE](LICENSE).
