package newsletter

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/yoanbernabeu/daybrief/internal/gemini"
)

func SaveJSON(newsletter *gemini.Newsletter, outputDir string) (string, error) {
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return "", fmt.Errorf("creating output dir: %w", err)
	}

	filename := fmt.Sprintf("%s.json", time.Now().Format("2006-01-02"))
	path := filepath.Join(outputDir, filename)

	data, err := json.MarshalIndent(newsletter, "", "  ")
	if err != nil {
		return "", fmt.Errorf("marshaling newsletter: %w", err)
	}

	if err := os.WriteFile(path, data, 0644); err != nil {
		return "", fmt.Errorf("writing newsletter: %w", err)
	}

	return path, nil
}
