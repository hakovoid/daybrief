package sources

import "time"

type SourceItem struct {
	Title        string
	URL          string
	SourceName   string
	SourceType   string
	ThumbnailURL string
	AudioURL     string
	PublishedAt  time.Time
}
