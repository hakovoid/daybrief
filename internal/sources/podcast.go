package sources

import (
	"log/slog"
	"time"

	"github.com/mmcdole/gofeed"
	"github.com/yoanbernabeu/daybrief/internal/config"
)

func FetchPodcasts(podcasts []config.PodcastSource, since time.Time, logger *slog.Logger) []SourceItem {
	var items []SourceItem
	parser := gofeed.NewParser()

	for _, pod := range podcasts {
		logger.Debug("fetching podcast", "name", pod.Name, "url", pod.URL)

		f, err := parser.ParseURL(pod.URL)
		if err != nil {
			logger.Warn("failed to fetch podcast", "name", pod.Name, "error", err)
			continue
		}

		for _, item := range f.Items {
			var publishedAt time.Time
			if item.PublishedParsed != nil {
				publishedAt = *item.PublishedParsed
			} else if item.UpdatedParsed != nil {
				publishedAt = *item.UpdatedParsed
			} else {
				continue
			}

			if !publishedAt.After(since) {
				continue
			}

			audioURL := ""
			if len(item.Enclosures) > 0 {
				audioURL = item.Enclosures[0].URL
			}

			thumbnailURL := ""
			if item.Image != nil {
				thumbnailURL = item.Image.URL
			} else if f.Image != nil {
				thumbnailURL = f.Image.URL
			}

			items = append(items, SourceItem{
				Title:        item.Title,
				URL:          item.Link,
				SourceName:   pod.Name,
				SourceType:   "podcast",
				ThumbnailURL: thumbnailURL,
				AudioURL:     audioURL,
				PublishedAt:  publishedAt,
			})
		}

		logger.Info("fetched podcast", "name", pod.Name, "new_items", len(items))
	}

	return items
}
