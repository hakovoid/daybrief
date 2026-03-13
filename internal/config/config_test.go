package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoad(t *testing.T) {
	content := `
gemini:
  model: "gemini-3-flash"
newsletter:
  language: "en"
  max_highlights: 3
  editorial_prompt: "Be concise"
mail:
  subject_prefix: "[Test]"
sources:
  rss:
    - url: "https://example.com/feed.xml"
      name: "Example"
  youtube:
    - channel_id: "UC123"
      name: "Test Channel"
  podcasts:
    - url: "https://example.com/podcast.xml"
      name: "Test Podcast"
`
	dir := t.TempDir()
	path := filepath.Join(dir, "config.yaml")
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}

	cfg, err := Load(path)
	if err != nil {
		t.Fatalf("Load() error: %v", err)
	}

	if cfg.Gemini.Model != "gemini-3-flash" {
		t.Errorf("Model = %q, want %q", cfg.Gemini.Model, "gemini-3-flash")
	}
	if cfg.Newsletter.MaxHighlights != 3 {
		t.Errorf("MaxHighlights = %d, want 3", cfg.Newsletter.MaxHighlights)
	}
	if cfg.Newsletter.Language != "en" {
		t.Errorf("Language = %q, want %q", cfg.Newsletter.Language, "en")
	}
	if len(cfg.Sources.RSS) != 1 {
		t.Errorf("RSS sources = %d, want 1", len(cfg.Sources.RSS))
	}
	if len(cfg.Sources.YouTube) != 1 {
		t.Errorf("YouTube sources = %d, want 1", len(cfg.Sources.YouTube))
	}
	if len(cfg.Sources.Podcasts) != 1 {
		t.Errorf("Podcast sources = %d, want 1", len(cfg.Sources.Podcasts))
	}
}

func TestLoadDefaults(t *testing.T) {
	content := `
sources:
  rss: []
`
	dir := t.TempDir()
	path := filepath.Join(dir, "config.yaml")
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}

	cfg, err := Load(path)
	if err != nil {
		t.Fatalf("Load() error: %v", err)
	}

	if cfg.Gemini.Model != "gemini-3-pro" {
		t.Errorf("Default model = %q, want %q", cfg.Gemini.Model, "gemini-3-pro")
	}
	if cfg.Newsletter.MaxHighlights != 5 {
		t.Errorf("Default MaxHighlights = %d, want 5", cfg.Newsletter.MaxHighlights)
	}
	if cfg.Newsletter.Language != "en" {
		t.Errorf("Default Language = %q, want %q", cfg.Newsletter.Language, "en")
	}
}

func TestLoadEnvRecipients(t *testing.T) {
	t.Setenv("DAYBRIEF_RECIPIENTS", "a@test.com, b@test.com, c@test.com")
	t.Setenv("GEMINI_API_KEY", "test-key")

	env, err := LoadEnv()
	if err != nil {
		t.Fatalf("LoadEnv() error: %v", err)
	}

	if len(env.Recipients) != 3 {
		t.Fatalf("Recipients = %d, want 3", len(env.Recipients))
	}
	if env.Recipients[0] != "a@test.com" {
		t.Errorf("Recipients[0] = %q, want %q", env.Recipients[0], "a@test.com")
	}
	if env.GeminiAPIKey != "test-key" {
		t.Errorf("GeminiAPIKey = %q, want %q", env.GeminiAPIKey, "test-key")
	}
	if env.SMTPPort != "587" {
		t.Errorf("Default SMTPPort = %q, want %q", env.SMTPPort, "587")
	}
}
