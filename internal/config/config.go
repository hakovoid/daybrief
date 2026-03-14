package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
)

type RSSSource struct {
	URL  string `yaml:"url"`
	Name string `yaml:"name"`
}

type YouTubeSource struct {
	ChannelID string `yaml:"channel_id"`
	Name      string `yaml:"name"`
}

type PodcastSource struct {
	URL  string `yaml:"url"`
	Name string `yaml:"name"`
}

type Sources struct {
	RSS      []RSSSource     `yaml:"rss"`
	YouTube  []YouTubeSource `yaml:"youtube"`
	Podcasts []PodcastSource `yaml:"podcasts"`
}

type GeminiConfig struct {
	Model string `yaml:"model"`
}

type NewsletterConfig struct {
	Language         string `yaml:"language"`
	MaxHighlights    int    `yaml:"max_highlights"`
	EditorialPrompt  string `yaml:"editorial_prompt"`
	DefaultLookback  string `yaml:"default_lookback"`
}

type MailConfig struct {
	SubjectPrefix string `yaml:"subject_prefix"`
}

type Config struct {
	Gemini     GeminiConfig     `yaml:"gemini"`
	Newsletter NewsletterConfig `yaml:"newsletter"`
	Mail       MailConfig       `yaml:"mail"`
	Sources    Sources          `yaml:"sources"`
}

type EnvConfig struct {
	GeminiAPIKey  string
	YouTubeAPIKey string
	SMTPHost      string
	SMTPPort      string
	SMTPUsername  string
	SMTPPassword  string
	MailFromName  string
	MailFromEmail string
	Recipients    []string
}

func Load(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading config file: %w", err)
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("parsing config file: %w", err)
	}

	// Set defaults
	if cfg.Gemini.Model == "" {
		cfg.Gemini.Model = "gemini-3-flash-preview"
	}
	if cfg.Newsletter.MaxHighlights == 0 {
		cfg.Newsletter.MaxHighlights = 5
	}
	if cfg.Newsletter.Language == "" {
		cfg.Newsletter.Language = "en"
	}
	if cfg.Newsletter.DefaultLookback == "" {
		cfg.Newsletter.DefaultLookback = "48h"
	}

	return &cfg, nil
}

func LoadEnv() (*EnvConfig, error) {
	// Load .env file if it exists, ignore error if not found
	_ = godotenv.Load()

	env := &EnvConfig{
		GeminiAPIKey:  os.Getenv("GEMINI_API_KEY"),
		YouTubeAPIKey: os.Getenv("YOUTUBE_API_KEY"),
		SMTPHost:      os.Getenv("SMTP_HOST"),
		SMTPPort:      os.Getenv("SMTP_PORT"),
		SMTPUsername:  os.Getenv("SMTP_USERNAME"),
		SMTPPassword:  os.Getenv("SMTP_PASSWORD"),
		MailFromName:  os.Getenv("MAIL_FROM_NAME"),
		MailFromEmail: os.Getenv("MAIL_FROM_EMAIL"),
	}

	if env.SMTPPort == "" {
		env.SMTPPort = "587"
	}
	if env.MailFromName == "" {
		env.MailFromName = "DayBrief"
	}

	recipients := os.Getenv("DAYBRIEF_RECIPIENTS")
	if recipients != "" {
		for _, r := range strings.Split(recipients, ",") {
			r = strings.TrimSpace(r)
			if r != "" {
				env.Recipients = append(env.Recipients, r)
			}
		}
	}

	return env, nil
}
