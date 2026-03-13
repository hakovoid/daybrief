package sources

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"time"
)

func GetLastExecutionDate(outputDir string, defaultLookback time.Duration) (time.Time, error) {
	matches, err := filepath.Glob(filepath.Join(outputDir, "*.json"))
	if err != nil {
		return time.Time{}, fmt.Errorf("globbing output dir: %w", err)
	}

	if len(matches) == 0 {
		return time.Now().UTC().Add(-defaultLookback), nil
	}

	sort.Sort(sort.Reverse(sort.StringSlice(matches)))

	data, err := os.ReadFile(matches[0])
	if err != nil {
		return time.Time{}, fmt.Errorf("reading latest output: %w", err)
	}

	var result struct {
		GeneratedAt string `json:"generated_at"`
	}
	if err := json.Unmarshal(data, &result); err != nil {
		return time.Time{}, fmt.Errorf("parsing latest output: %w", err)
	}

	t, err := time.Parse(time.RFC3339, result.GeneratedAt)
	if err != nil {
		return time.Time{}, fmt.Errorf("parsing generated_at: %w", err)
	}

	return t, nil
}
