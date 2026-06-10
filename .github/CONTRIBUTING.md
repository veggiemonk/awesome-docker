# Contributing to awesome-docker

Thanks for taking the time to contribute.

This repository is a curated list of Docker/container resources plus a Go-based maintenance CLI used by CI. Contributions are welcome for both content and tooling.

Please read and follow the [Code of Conduct](./CODE_OF_CONDUCT.md).

## What We Accept

- New high-quality projects that are **for Docker** (see the test below)
- Fixes to descriptions, ordering, or categorization
- Removal of broken, archived, deprecated, or duplicate entries

## The "for Docker" Test (read this before submitting)

This list is for projects whose purpose is to make working with Docker better.
It is **not** a directory of "software you can run in a container" — that's most
software ever written. Before opening a PR, apply this test:

> **If you removed the Docker integration, would the project still have a reason to exist?**
>
> - **Yes, it would** → it's a general tool that happens to use Docker. **Reject.**
> - **No, the project is *about* Docker** → it belongs here.

### Examples

| Project shape                                                                       | Verdict       | Why                                                                                                  |
| ----------------------------------------------------------------------------------- | ------------- | ---------------------------------------------------------------------------------------------------- |
| Tool that monitors **containers**, container resources, or container logs           | ✅ Accept     | Its value disappears if you take Docker away.                                                        |
| Tool that monitors **a host / server / service** and happens to ship as Docker      | ❌ Reject     | Its job is monitoring; Docker is just the delivery vehicle.                                          |
| Reverse proxy that **auto-configures from Docker labels/events**                    | ✅ Accept     | Tightly coupled to the Docker API.                                                                   |
| Reverse proxy that just *runs in* Docker and configures from a static YAML file     | ❌ Reject     | Same proxy works fine without Docker.                                                                |
| Dockerfile linter, image scanner, registry, BuildKit frontend, runtime, SDK         | ✅ Accept     | Object of work is the Docker image, daemon, or socket.                                               |
| General IaC scanner that **happens to read Dockerfiles** among many formats         | ⚠️ Borderline | Allowed if Docker support is a first-class feature, not an afterthought.                             |
| "Awesome 150 web apps deployable with `docker run`"                                 | ❌ Reject     | Belongs in `awesome-selfhosted`; Docker is incidental.                                               |
| Cron / privilege-drop / health-check utility designed for **container init**        | ✅ Accept     | Solves a problem that only exists *because* of how containers are built.                             |
| Generic KV store / service mesh / cluster scheduler that supports Docker            | ❌ Reject     | Existed before Docker, works without it.                                                             |
| Tutorial, video, or book whose **subject** is Docker                                | ✅ Accept     | Under Useful Resources / Videos / Where to start.                                                    |
| Tutorial that uses Docker as a setup step for an unrelated topic                    | ❌ Reject     | Belongs with the unrelated topic.                                                                    |

### One-sentence sanity check

Write the sentence: *"This project exists to ____."*
If the blank doesn't contain **Docker, container, image, registry, Dockerfile, Compose, Swarm, BuildKit, or OCI** — it probably doesn't belong here.

## README Entry Rules

- Use one link per entry.
- Prefer GitHub project/repository URLs over marketing pages.
- Keep entries alphabetically sorted within their section (case-insensitive).
- Keep descriptions concise and concrete: one sentence, lead with the verb.
- **The description should make the Docker connection obvious.** "Monitor
  unhealthy containers" is good; "Monitor your services" is not — the latter
  could be any monitoring tool, and a reviewer can't tell whether the project
  passes the test above.
- Use `:yen:` for paid/commercial services.
- The project has been active in the last 2 years.
- Do not use `:skull:`; archived/deprecated projects should be removed.
- Avoid duplicate links and redirect variants.

## Local Validation

```bash

# Build CLI
make build

# Validate README formatting and content
make lint

# Run code tests (when touching Go code)
make test

# Optional: full external checks (requires GITHUB_TOKEN)
./awesome-docker check
./awesome-docker validate
```

## Pull Request Expectations

- Keep the PR focused to one logical change.
- Explain what changed and why.
- If adding entries, include the target category.
- If removing entries, explain why (archived, broken, duplicate, etc.).

## Maintainer Notes

- Changes should be reviewed before merge.
- Prefer helping contributors improve a PR over silently rejecting it.
- Keep `.github` documentation and workflows aligned with current tooling.
