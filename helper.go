package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

func writeToFile(titles map[string]bool) error {
	var titlesList []string
	for title := range titles {
		titlesList = append(titlesList, title)
	}

	data := struct {
		Timestamp string   `json:"timestamp"`
		Source    string   `json:"source"`
		Count     int      `json:"count"`
		Titles    []string `json:"titles"`
	}{
		Timestamp: time.Now().Format(time.RFC3339),
		Source:    URL,
		Count:     len(titlesList),
		Titles:    titlesList,
	}

	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshaling JSON: %w", err)
	}

	filename := "news_titles_" + time.Now().Format("2006-01-02") + ".json"
	err = os.WriteFile(filename, jsonData, 0644)
	if err != nil {
		return fmt.Errorf("error writing to file: %w", err)
	}
	return nil
}
