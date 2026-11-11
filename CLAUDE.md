# Company Service

## How to work in this repo

- Plan before editing: outline the approach and surface tradeoffs first. Do not hide confusion behind a confident edit.
- Simplicity first: smallest change that works, no premature abstraction (YAGNI).
- Surgical changes only: touch only what the task names; don't refactor unrelated code.
- State the goal and the definition of done up front; loop until verified against it.
- Ask when the spec is ambiguous; do not guess and ship the guess as fact.
- Stay in scope: no drive-by edits to neighboring code just because you're in the file.
- Tests are truth: run tests and they must pass before you claim done.
- Be concise: short answers, no recap of work you just did unless asked.

## Stack

- Go, REST/JSON API, RabbitMQ for events, packaged as a Docker image and deployed with Helm to Kubernetes (AKS). The shared stack is described in `plan.md` section 1.2 in the `general` repository.
- Database: PostgreSQL, owned by this service only.
- Prefer the repo's existing libraries over reaching for a new dependency.
- Don't infer the build tooling or the module layout. Read the repo.

## What burned us before

<!-- Guide for the agent: nothing is recorded yet. Do not add anything here on your own and do not invent incidents.
     When the user tells you to add an entry, write one short line: what went wrong, and the check that would have caught it.
     Example of the shape (not a real incident): "A job reported success while some records were skipped. Verify counts after bulk operations." -->

## Keep this file honest

- When you get corrected, add the lesson to this file (update CLAUDE.md) so the next session starts ahead.

## Purpose and place

Repository `company-service` contains the Company & Membership service (short name `company-membership`) of Mi (Butterfly Give).

It manages companies, company memberships and roles (`employee`, `representative`), and invite codes. It also answers "who is this user's company" for other services through REST. See `plan.md` section 2.1, "Company & Membership".

It does not handle challenges, participation, donations, or impact. Those belong to other services.

## Shared documentation

The repository `rso-project-team-14/general` (cloned locally at `../general/`) holds the project plan. `plan.md` is the source of truth for infrastructure, service responsibilities, and delivery stages. Planned API contracts are in `general/api/`.

If `general` is not accessible, say so instead of assuming its contents.

## Sync with general

Before implementing a feature, list the `changelog/` directory in the `general` repository (`rso-project-team-14/general`) and compare the newest file name with the one recorded below under "Last checked general changelog entry".
- If there are files newer than that one (file names start with a UTC timestamp, so alphabetical order is chronological), read them, then re-read the parts of `plan.md`, `docs/`, and `api/` they affect. Apply any changes that concern this repository, then update the "Last checked" line to the newest file name.
- If there are no newer files, do not re-read `plan.md`.
- Every document in `general` has its content hash on line 3 (`<!-- sha256: ... -->`, covering line 4 to the end). To check whether a specific document changed since you last read it, compare that line with the hash you recorded before reading the whole file.

Last checked general changelog entry: `2026-11-11_19-01-19_gateway-api-with-envoy.md`

## Current development stage

Follow `general/docs/stages/stage-1-demo.md`: deploy and connect the services, no business logic.

Required endpoints:
- `GET /health/started`
- `GET /health/ready`
- `GET /health/live`
- `GET /api/v1/ping`

## Contracts

- Own planned contract: `general/api/company-membership.md`. It does not exist yet. Do not invent endpoints; ask or propose the contract in `general` first.
- Contracts of the services that call this one are not defined yet.
- Until the service is implemented, the planned contract is the guide. After that, this repository's OpenAPI spec and code are the truth, and any difference from the planned contract is reported through a change in `general`.

## Conventions

- Go version: as set in `go.mod`.
- Target layout: `cmd/`, `internal/http`, `internal/domain`, `internal/store`, `internal/messaging`, `migrations/` (`plan.md` section 1.4). The code currently lives in the repository root. Moving it is a separate, deliberate change.
- Runtime conventions (port, health endpoints, configuration variables, JSON logging, `X-Request-ID`, error format, UUID v4, UTC timestamps): `plan.md` section 1.6. No service-specific deviations so far.
- Write code, comments, and documentation in English.
- Follow standard Go conventions and format code with `gofmt`.
- Prefer the standard library when practical.
- Add tests for new functionality.

## Data

- Owns `company_db` (PostgreSQL). Migrations go in `migrations/` and are applied at startup (`plan.md` section 1.8). No tables exist yet; Stage 1 has none.
- Never read or write another service's database. References to other services' entities are plain UUIDs.
- Money, if ever stored here, is in integer cents with an ISO 4217 currency code.

## Events

- The event catalog in `plan.md` section 1.7 lists no events that this service produces or consumes. Section 1.7 is provisional.
- Do not add events until they are defined in `general`. If this service starts producing or consuming events, the outbox and idempotency rules in section 1.7 apply.

## Commands

TBD until local development (`plan.md` section 1.5) is decided. For now:
- Test: `go test ./...` (run before committing)
- Format: `gofmt`
- Lint and migrations: not set up yet.

## Git and pull requests

- GitHub Flow (`plan.md` section 1.13): short-lived branches from `main`, merged through pull requests.
- Pending in issue #5: branch naming, commit format, pull request and review rules, and merge method. Conventional Commits is assumed in #5 but not final. Do not treat any of these as decided.

## Rules for Claude

- Stay within this service's responsibility.
- Do not resolve an open decision from `plan.md` section 1.12 silently. Raise it as an issue or a change in `general`.
- Do not change another service's code or database.
- If the plan looks wrong or contradicts the code, say so instead of quietly diverging from it.
- Never commit secrets or credentials.
- Explain important technical decisions, and ask before significant architectural changes.
- Keep detailed documentation in separate files and read it only when needed.
