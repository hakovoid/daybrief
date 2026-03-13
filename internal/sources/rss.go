package sources

import (
	"log/slog"
	"time"

	"github.com/mmcdole/gofeed"
	"github.com/yoanbernabeu/daybrief/internal/config"
)

func FetchRSS(feeds []config.RSSSource, since time.Time, logger *slog.Logger) []SourceItem {
	var items []SourceItem
	parser := gofeed.NewParser()

	for _, feed := range feeds {
		logger.Debug("fetching RSS feed", "name", feed.Name, "url", feed.URL)

		f, err := parser.ParseURL(feed.URL)
		if err != nil {
			logger.Warn("failed to fetch RSS feed", "name", feed.Name, "error", err)
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

			items = append(items, SourceItem{
				Title:       item.Title,
				URL:         item.Link,
				SourceName:  feed.Name,
				SourceType:  "rss",
				PublishedAt: publishedAt,
			})
		}

		logger.Info("fetched RSS feed", "name", feed.Name, "new_items", len(items))
	}

	return items
}
