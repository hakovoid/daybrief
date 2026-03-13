package gemini

type SourceSummary struct {
	Title        string   `json:"title"`
	Summary      string   `json:"summary"`
	KeyPoints    []string `json:"key_points"`
	SourceType   string   `json:"source_type"`
	SourceURL    string   `json:"source_url"`
	SourceName   string   `json:"source_name"`
	ThumbnailURL string   `json:"thumbnail_url,omitempty"`
}

type Highlight struct {
	Title        string `json:"title"`
	SourceName   string `json:"source_name"`
	SourceURL    string `json:"source_url"`
	ThumbnailURL string `json:"thumbnail_url,omitempty"`
	Analysis     string `json:"analysis"`
}

type Resource struct {
	Title      string `json:"title"`
	SourceName string `json:"source_name"`
	SourceURL  string `json:"source_url"`
	Summary    string `json:"summary"`
}

type Newsletter struct {
	GeneratedAt string      `json:"generated_at"`
	Subject     string      `json:"subject"`
	Editorial   string      `json:"editorial"`
	Highlights  []Highlight `json:"highlights"`
	Resources   []Resource  `json:"resources"`
}
