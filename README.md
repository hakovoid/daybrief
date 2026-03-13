# DayBrief

An open-source Go CLI tool that aggregates content from RSS feeds, YouTube channels, and podcasts, uses the Gemini API to summarize and analyze each source, and sends an automated HTML newsletter by email.

## Features

- **Multi-source aggregation**: RSS feeds, YouTube channels, podcasts
- **AI-powered analysis**: Two-pass Gemini integration (summarize each source, then synthesize a newsletter)
- **Automated delivery**: HTML email via SMTP
- **Incremental updates**: Only processes new content since last execution
- **CI/CD ready**: Designed to run in GitHub Actions via cron

## Installation

Download the latest binary from [GitHub Releases](https://github.com/yoanbernabeu/daybrief/releases):

```bash
curl -sL https://github.com/yoanbernabeu/daybrief/releases/latest/download/daybrief-linux-amd64 -o daybrief
chmod +x daybrief
```

Or build from source:

```bash
git clone https://github.com/yoanbernabeu/daybrief.git
cd daybrief
make build
```

## Configuration

### `config.yaml`

```yaml
gemini:
  model: "gemini-3-flash-preview"

newsletter:
  language: "fr"
  max_highlights: 5
  editorial_prompt: "A casual, tech-savvy tone with a focus on practical insights."

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

### Environment Variables

Create a `.env` file (see `.env.example`):

```env
GEMINI_API_KEY=your-api-key
YOUTUBE_API_KEY=your-youtube-api-key
SMTP_HOST=smtp.example.com
SMTP_PORT=587
SMTP_USERNAME=user
SMTP_PASSWORD=pass
MAIL_FROM_NAME=DayBrief
MAIL_FROM_EMAIL=newsletter@example.com
DAYBRIEF_RECIPIENTS=user1@example.com,user2@example.com
```

## Usage

### Run the full pipeline

```bash
daybrief run --config config.yaml
```

### Preview in browser

```bash
daybrief preview --config config.yaml
```

### Check source health

```bash
daybrief sources --config config.yaml
```

## GitHub Actions

Copy `.github/workflows/daybrief-example.yml` to your repository and configure the required secrets in your GitHub repository settings.

## License

MIT - see [LICENSE](LICENSE) for details.
