package newsletter

import (
	"strings"
	"testing"

	"github.com/yoanbernabeu/daybrief/internal/gemini"
)

func TestRenderHTML(t *testing.T) {
	nl := &gemini.Newsletter{
		GeneratedAt: "2026-03-13T08:00:00Z",
		Subject:     "This Week in Tech",
		Editorial:   "Welcome to this week's roundup of the most interesting tech news.\n\nThis week saw major releases across the ecosystem, from Go to Kubernetes.",
		Highlights: []gemini.Highlight{
			{
				Title:        "Go 1.24 Released",
				SourceName:   "Go Blog",
				SourceURL:    "https://go.dev/blog/go1.24",
				ThumbnailURL: "https://go.dev/images/go-logo-blue.svg",
				Analysis:     "Go 1.24 brings exciting new features including enhanced generics support.",
			},
		},
		Resources: []gemini.Resource{
			{
				Title:      "Kubernetes 1.30",
				SourceName: "K8s Blog",
				SourceURL:  "https://kubernetes.io/blog",
				Summary:    "New features in Kubernetes 1.30.",
			},
		},
	}

	html, err := RenderHTML(nl)
	if err != nil {
		t.Fatalf("RenderHTML() error: %v", err)
	}

	checks := []string{
		"This Week in Tech",
		"Welcome to this week",
		"Go 1.24 Released",
		"go.dev/blog/go1.24",
		"Go Blog",
		"enhanced generics",
		"Kubernetes 1.30",
		"K8s Blog",
		"Powered by",
		"DayBrief",
	}

	for _, check := range checks {
		if !strings.Contains(html, check) {
			t.Errorf("HTML output missing %q", check)
		}
	}
}
