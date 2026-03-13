package gemini

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"google.golang.org/genai"
)

type Client struct {
	client          *genai.Client
	model           string
	language        string
	maxHighlights   int
	editorialPrompt string
	logger          *slog.Logger
}

func NewClient(ctx context.Context, apiKey, model, language string, maxHighlights int, editorialPrompt string, logger *slog.Logger) (*Client, error) {
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  apiKey,
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		return nil, fmt.Errorf("creating Gemini client: %w", err)
	}

	return &Client{
		client:          client,
		model:           model,
		language:        language,
		maxHighlights:   maxHighlights,
		editorialPrompt: editorialPrompt,
		logger:          logger,
	}, nil
}

func withRetry[T any](fn func() (T, error), logger *slog.Logger) (T, error) {
	delays := []time.Duration{1 * time.Second, 2 * time.Second, 4 * time.Second}
	var lastErr error

	for i := 0; i <= len(delays); i++ {
		result, err := fn()
		if err == nil {
			return result, nil
		}
		lastErr = err

		if i < len(delays) {
			logger.Warn("retrying after error", "attempt", i+1, "error", err, "delay", delays[i])
			time.Sleep(delays[i])
		}
	}

	var zero T
	return zero, fmt.Errorf("all retries exhausted: %w", lastErr)
}
