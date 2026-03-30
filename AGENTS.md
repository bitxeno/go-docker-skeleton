# Repository Guidelines

## Project Structure & Module Organization
`main.go` is the CLI entrypoint. Command handlers live under `cmd/` (`cmd/server`, `cmd/gen`). Core application code is in `internal/`, grouped by concern: `app/` for bootstrap and config, `db/` for persistence, `http/` for HTTP helpers, `log/` for logging, `model/` for data types, and `utils/` for shared helpers. The Fiber web layer is under `web/`; `web/static/` contains the Vue 3 frontend source, with pages in `src/page/`, routing in `src/router/`, and assets in `src/assets/`. Example config lives in `docs/config.yaml.example`.

## Build, Test, and Development Commands
Run the frontend watcher with `cd web/static && npm install && npm run dev`; this writes development assets for the Go server to serve. Run the backend locally with `go run . server -vv`. For live reload, use `air -c .air.toml`. Create a production frontend bundle with `cd web/static && npm run build`. Run backend checks with `go test ./...`, and use `golangci-lint run` to match CI.

## Coding Style & Naming Conventions
Format Go code with `gofmt` and keep package names lowercase. Exported Go identifiers use `CamelCase`; unexported helpers use `camelCase`. Follow the existing package layout instead of adding broad utility files. Frontend code uses 2-space indentation, double quotes, and Vue single-file components. Keep route/page files under `web/static/src/page/` in lowercase directories such as `home/index.vue`.

## Testing Guidelines
CI runs `go test -coverprofile=coverage.txt -covermode=atomic ./...` plus `golangci-lint`, so new Go packages should include `_test.go` coverage where behavior changes. Place tests next to the code they verify and prefer table-driven Go tests for handlers, config parsing, and utility functions. The frontend currently has no test harness; at minimum, verify `npm run build` before opening a PR.

## Commit & Pull Request Guidelines
Recent history uses concise, scoped subjects such as `build(deps): bump axios...`. Prefer imperative commit messages with an optional scope (`build:`, `feat:`, `fix:`). Keep dependency bumps separate from feature work. PRs should describe the change, note config or Docker impact, link the related issue when applicable, and include screenshots for UI changes in `web/static/`.

## Configuration & Environment
Target the repo toolchain documented in `README.md`: Go 1.21+ and Node.js 20+. Do not commit secrets; start from `docs/config.yaml.example` and mount runtime config into `/data` for Docker deployments.
