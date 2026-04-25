<p align="center">
  <img src="web/public/favicon.svg" width="60" alt="DayBrief logo" />
</p>

<h1 align="center">DayBrief</h1>

<p align="center">
  <strong>AI-powered newsletter from your favorite sources.</strong><br/>
  Aggregate RSS, YouTube & Podcasts — summarize with Gemini — deliver by email.
</p>

<p align="center">
  <a href="https://github.com/yoanbernabeu/daybrief/actions"><img src="https://github.com/yoanbernabeu/daybrief/actions/workflows/ci.yml/badge.svg" alt="CI" /></a>
  <a href="https://github.com/yoanbernabeu/daybrief/releases"><img src="https://img.shields.io/github/v/release/yoanbernabeu/daybrief" alt="Release" /></a>
  <a href="https://github.com/yoanbernabeu/daybrief/blob/main/LICENSE"><img src="https://img.shields.io/github/license/yoanbernabeu/daybrief" alt="License" /></a>
  <a href="https://yoanbernabeu.github.io/daybrief"><img src="https://img.shields.io/badge/website-daybrief-blue" alt="Website" /></a>
</p>

<p align="center">
  <a href="https://yoanbernabeu.github.io/daybrief">Website</a> &middot;
  <a href="https://yoanbernabeu.github.io/daybrief/guide/">Setup Guide</a> &middot;
  <a href="https://yoanbernabeu.github.io/daybrief/admin/">Dashboard</a> &middot;
  <a href="https://github.com/yoanbernabeu/daybrief/releases">Releases</a>
</p>

---

## Configuration personnelle (instance hakovoid)

> Cette section décrit la configuration concrète de ce fork. Le reste du README est la doc upstream du projet, conservée pour référence.

### Setup actuel

- **Repo** : `hakovoid/daybrief` — workflow autonome, ne dépend plus de `yoanbernabeu/daybrief` à l'exécution
- **Cron** : tous les jours à 05:00 UTC
- **Action** : `uses: hakovoid/daybrief@main` (binaire téléchargé depuis nos propres releases)
- **Sources** (voir `config.yaml`) :
  - RSS Flux FR : `https://hakovoid-flux.netlify.app/rss.xml`
  - RSS Flux World : `https://hakovoid-flux.netlify.app/world/rss.xml`
  - YouTube : chaîne YoanDev (`UCRlsJWh1XwmNGxZPFgJ0Zlw`)
  - Podcast : IFTTD
- **SMTP** : Gmail (avec App Password, 2FA active)
- **Destinataire** : `(removed)`
- **Lookback par défaut** : 48h
- **Modèle Gemini** : `gemini-3-flash-preview`

### Mettre à jour la config

```bash
cd /path/to/daybrief
# éditer config.yaml (sources, lookback, ton éditorial…)
git add config.yaml
git commit -m "config: <description>"
git push
```

### Déclencher manuellement un envoi

```bash
gh workflow run daybrief.yml --repo hakovoid/daybrief
gh run list --repo hakovoid/daybrief --workflow daybrief.yml --limit 3
```

### Mettre à jour un secret

Settings → Secrets and variables → Actions → cliquer sur le secret → Update.
Ou en CLI : `gh secret set <NAME> --repo hakovoid/daybrief`.

### Documentation interne

Notes complètes dans `internal-notes/` :

- `STATUS.md` — état courant + TODO
- `flux-et-daybrief-deploiement.md` — tuto complet (Flux + DayBrief + procédure rename)
- `migration-netlify-vers-cloudflare.md` — plan migration de Flux vers Cloudflare

### Liens utiles

- Workflow : https://github.com/hakovoid/daybrief/actions/workflows/daybrief.yml
- Releases : https://github.com/hakovoid/daybrief/releases
- Site Flux source : https://hakovoid-flux.netlify.app

---

## What is DayBrief?

DayBrief is an open-source GitHub Action that monitors your content sources overnight and delivers a concise, AI-generated newsletter every morning. No server to manage — it runs entirely on GitHub Actions.

**How it works:**

1. **Fetch** — Collects new content from RSS feeds, YouTube channels, and podcasts
2. **Summarize** — Sends each item to Gemini AI for individual analysis
3. **Synthesize** — Generates a cohesive newsletter with editorial, highlights, and resources
4. **Deliver** — Sends the newsletter by email via SMTP and archives it as JSON

## Features

- **Multi-source aggregation** — RSS feeds, YouTube channels, podcasts
- **Two-pass AI analysis** — Individual source summaries, then editorial synthesis via Gemini
- **Incremental processing** — Only processes content published since the last run
- **Zero infrastructure** — Runs entirely on GitHub Actions, no server needed
- **Web dashboard** — Edit config, manage sources, preview newsletters from the browser
- **Shareable archive** — Public web page to browse past newsletter editions

## Quick Start

### 1. Create a repository

Create a new GitHub repository for your newsletter.

### 2. Add `config.yaml`

```yaml
gemini:
  model: "gemini-3-flash-preview"

newsletter:
  language: "fr"
  max_highlights: 5
  default_lookback: "48h"
  editorial_prompt: "A casual, tech-savvy tone with practical insights."

mail:
  subject_prefix: "[DayBrief]"

sources:
  rss:
    - url: "https://blog.golang.org/feed.atom"
      name: "Go Blog"
  youtube:
    - channel_id: "UCxxxx"
      name: "My Channel"
  podcasts:
    - url: "https://example.com/podcast.xml"
      name: "My Podcast"
```

<details>
<summary><strong>Configuration reference</strong></summary>

| Option | Default | Description |
|---|---|---|
| `gemini.model` | `gemini-3-flash-preview` | Gemini model to use |
| `newsletter.language` | `en` | Newsletter language |
| `newsletter.max_highlights` | `5` | Number of highlights |
| `newsletter.default_lookback` | `48h` | Time window for first run |
| `newsletter.editorial_prompt` | — | Tone and style instructions for the AI |
| `mail.subject_prefix` | — | Prefix added to email subjects |

</details>

### 3. Configure secrets

Go to **Settings > Secrets and variables > Actions** and add:

| Secret | Required | Description |
|---|---|---|
| `GEMINI_API_KEY` | Yes | [Google Gemini API key](https://ai.google.dev/) |
| `YOUTUBE_API_KEY` | If YouTube | [YouTube Data API key](https://console.cloud.google.com/) |
| `SMTP_HOST` | Yes | SMTP server host |
| `SMTP_PORT` | No | SMTP port (default: `587`) |
| `SMTP_USERNAME` | Yes | SMTP username |
| `SMTP_PASSWORD` | Yes | SMTP password |
| `MAIL_FROM_NAME` | No | Sender name (default: `DayBrief`) |
| `MAIL_FROM_EMAIL` | Yes | Sender email address |
| `DAYBRIEF_RECIPIENTS` | Yes | Comma-separated recipient emails |

> Need help? See [how to get a Gemini API key](https://yoanbernabeu.github.io/daybrief/guide/gemini-api/) and [free email providers](https://yoanbernabeu.github.io/daybrief/guide/email-providers/).

### 4. Add the workflow

Create `.github/workflows/daybrief.yml`:

```yaml
name: DayBrief Newsletter

on:
  workflow_dispatch:
  schedule:
    - cron: "0 7 * * 1" # Every Monday at 7:00 UTC

permissions:
  contents: write

jobs:
  newsletter:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4

      - uses: yoanbernabeu/daybrief@v0.1.0
        with:
          config: config.yaml
        env:
          GEMINI_API_KEY: ${{ secrets.GEMINI_API_KEY }}
          YOUTUBE_API_KEY: ${{ secrets.YOUTUBE_API_KEY }}
          SMTP_HOST: ${{ secrets.SMTP_HOST }}
          SMTP_PORT: ${{ secrets.SMTP_PORT }}
          SMTP_USERNAME: ${{ secrets.SMTP_USERNAME }}
          SMTP_PASSWORD: ${{ secrets.SMTP_PASSWORD }}
          MAIL_FROM_NAME: ${{ secrets.MAIL_FROM_NAME }}
          MAIL_FROM_EMAIL: ${{ secrets.MAIL_FROM_EMAIL }}
          DAYBRIEF_RECIPIENTS: ${{ secrets.DAYBRIEF_RECIPIENTS }}
```

### 5. Run it

Go to **Actions**, select "DayBrief Newsletter", click **Run workflow**. Once confirmed, the cron schedule handles the rest.

The action automatically commits newsletter output to `output/` for incremental processing.

## Web App

DayBrief includes a web app built with Astro 6, deployed on GitHub Pages:

| Page | Description |
|---|---|
| [Landing](https://yoanbernabeu.github.io/daybrief) | Project presentation |
| [Dashboard](https://yoanbernabeu.github.io/daybrief/admin/) | Edit config.yaml visually, manage sources, preview newsletters |
| [Newsletter](https://yoanbernabeu.github.io/daybrief/newsletter/?owner=yoanbernabeu&repo=newsletter) | Public shareable newsletter archive |
| [Setup Guide](https://yoanbernabeu.github.io/daybrief/guide/) | Step-by-step documentation |

```bash
cd web
npm install
npm run dev      # Dev server
npm run build    # Production build
```

## CLI Usage

DayBrief can also run as a standalone CLI.

```bash
# Download latest binary
curl -sL https://github.com/yoanbernabeu/daybrief/releases/latest/download/daybrief-linux-amd64 -o daybrief
chmod +x daybrief

# Or build from source
git clone https://github.com/yoanbernabeu/daybrief.git && cd daybrief && make build
```

```bash
daybrief run --config config.yaml       # Run full newsletter pipeline
daybrief preview --config config.yaml   # Generate and preview in browser
daybrief sources --config config.yaml   # Check source accessibility
```

When running locally, create a `.env` file with the same variables as the GitHub secrets.

## Architecture

```
RSS / YouTube / Podcasts
        │
        ▼
   ┌─────────┐     ┌───────────┐     ┌──────────┐     ┌───────┐
   │  Fetch   │────▶│ Summarize │────▶│Synthesize│────▶│ Email │
   │ sources  │     │ (Gemini)  │     │(Gemini)  │     │ SMTP  │
   └─────────┘     └───────────┘     └──────────┘     └───────┘
                                           │
                                           ▼
                                     output/*.json
```

## Contributing

Contributions are welcome! Please open an issue first to discuss what you'd like to change.

```bash
make build        # Build binary
make test         # Run tests
make lint         # Run linter
```

## License

[MIT](LICENSE)
