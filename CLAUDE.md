# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

DayBrief is a Go CLI tool that aggregates content from RSS feeds, YouTube channels, and podcasts, summarizes each source using the Gemini API (two-pass: summarize then synthesize), and sends an automated HTML newsletter by email. Designed to run as a GitHub Actions cron job.

## Commands

```bash
make build        # Build binary → ./daybrief
make test         # Run all tests (go test ./...)
make lint         # Run golangci-lint (errcheck, govet, staticcheck, unused, gosimple, ineffassign)
make run          # Run full pipeline with config.yaml
```

Run a single test file:
```bash
go test ./internal/config/
go test ./internal/newsletter/ -run TestRenderHTML
```

## Architecture

The pipeline runs sequentially in `internal/cli/run.go`:

1. **Determine lookback window** — `sources.GetLastExecutionDate()` reads the last JSON output in `output/` to find when the pipeline last ran (falls back to `default_lookback` duration from config)
2. **Fetch sources** — `sources.FetchAll()` collects `SourceItem`s from RSS (gofeed), YouTube (Google API), and podcasts (gofeed) published after the lookback date
3. **Summarize each source** — `gemini.Client.SummarizeSource()` sends each item to Gemini with type-specific strategies: URL context tool for RSS, video FileData for YouTube, audio FileData for podcasts. Returns structured JSON (`SourceSummary`)
4. **Synthesize newsletter** — `gemini.Client.SynthesizeNewsletter()` takes all summaries and produces a `Newsletter` (editorial + highlights + resources)
5. **Save JSON** — `newsletter.SaveJSON()` writes to `output/`
6. **Render HTML** — `newsletter.RenderHTML()` uses an embedded Go template (`internal/newsletter/templates/email.html`)
7. **Send email** — `mail.SendEmail()` via SMTP

### Key packages

- `internal/cli` — Cobra commands (`run`, `preview`, `sources`); holds global config/env/logger
- `internal/config` — YAML config (`config.yaml`) + env vars (`.env` via godotenv). Two structs: `Config` (yaml) and `EnvConfig` (env vars)
- `internal/sources` — Fetchers per source type + `lastrun.go` for incremental processing
- `internal/gemini` — Gemini client with generic retry, JSON schema-constrained responses, type-specific prompts
- `internal/newsletter` — HTML rendering with `//go:embed` template, JSON output
- `internal/mail` — SMTP sender

### Configuration

- `config.yaml` — Sources, Gemini model, newsletter language/tone, mail prefix
- `.env` — API keys (GEMINI_API_KEY, YOUTUBE_API_KEY), SMTP credentials, recipients (comma-separated DAYBRIEF_RECIPIENTS)

## Pre-push checklist

Before pushing, always run all three checks and fix any issues:

```bash
make build && make test && make lint
```

## Conventions

- Go 1.25+ with `log/slog` for structured logging
- golangci-lint v2 with config in `.golangci.yml`
- Gemini responses use `ResponseMIMEType: "application/json"` with explicit JSON schemas
- The `output/` directory stores newsletter JSON files and serves as the incremental state (last run detection)
