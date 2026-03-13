package sources

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/yoanbernabeu/daybrief/internal/config"
	"google.golang.org/api/option"
	yt "google.golang.org/api/youtube/v3"
)

func FetchYouTube(channels []config.YouTubeSource, since time.Time, apiKey string, logger *slog.Logger) []SourceItem {
	if apiKey == "" {
		logger.Warn("YouTube API key not set, skipping YouTube sources")
		return nil
	}

	ctx := context.Background()
	service, err := yt.NewService(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		logger.Warn("failed to create YouTube service", "error", err)
		return nil
	}

	var items []SourceItem

	for _, ch := range channels {
		logger.Debug("fetching YouTube channel", "name", ch.Name, "channel_id", ch.ChannelID)

		call := service.Search.List([]string{"id", "snippet"}).
			ChannelId(ch.ChannelID).
			Type("video").
			Order("date").
			PublishedAfter(since.Format(time.RFC3339)).
			MaxResults(10)

		resp, err := call.Do()
		if err != nil {
			logger.Warn("failed to fetch YouTube channel", "name", ch.Name, "error", err)
			continue
		}

		for _, item := range resp.Items {
			publishedAt, err := time.Parse(time.RFC3339, item.Snippet.PublishedAt)
			if err != nil {
				continue
			}

			thumbnailURL := ""
			if item.Snippet.Thumbnails != nil && item.Snippet.Thumbnails.High != nil {
				thumbnailURL = item.Snippet.Thumbnails.High.Url
			}

			items = append(items, SourceItem{
				Title:        item.Snippet.Title,
				URL:          fmt.Sprintf("https://www.youtube.com/watch?v=%s", item.Id.VideoId),
				SourceName:   ch.Name,
				SourceType:   "youtube",
				ThumbnailURL: thumbnailURL,
				PublishedAt:  publishedAt,
			})
		}

		logger.Info("fetched YouTube channel", "name", ch.Name, "new_items", len(resp.Items))
	}

	return items
}
