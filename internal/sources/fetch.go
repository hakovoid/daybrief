package sources

import (
	"log/slog"
	"time"

	"github.com/yoanbernabeu/daybrief/internal/config"
)

func FetchAll(cfg *config.Config, env *config.EnvConfig, since time.Time, logger *slog.Logger) []SourceItem {
	var all []SourceItem

	rssItems := FetchRSS(cfg.Sources.RSS, since, logger)
	all = append(all, rssItems...)

	ytItems := FetchYouTube(cfg.Sources.YouTube, since, env.YouTubeAPIKey, logger)
	all = append(all, ytItems...)

	podItems := FetchPodcasts(cfg.Sources.Podcasts, since, logger)
	all = append(all, podItems...)

	logger.Info("fetched all sources", "total", len(all))
	return all
}
