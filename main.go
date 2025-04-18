package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Event struct {
	Type   string `json:"type"`
	Repo   struct {
		Name string `json:"name"`
	} `json:"repo"`
	Payload struct {
		Size int `json:"size"`
	} `json:"payload"`
}

func fetchGitHubActivity(username string) ([]Event, error) {
	url := fmt.Sprintf("https://api.github.com/users/%s/events", username)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error fetching data: %s", resp.Status)
	}

	
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var events []Event
	if err := json.Unmarshal(body, &events); err != nil {
		return nil, err
	}

	return events, nil
}

func displayActivity(events []Event) {
	if len(events) == 0 {
		fmt.Println("No activity found or invalid username.")
		return
	}

	for _, event := range events {
		switch event.Type {
		case "PushEvent":
			fmt.Printf("Pushed %d commits to %s\n", event.Payload.Size, event.Repo.Name)
		case "IssuesEvent":
			fmt.Printf("Opened a new issue in %s\n", event.Repo.Name)
		case "WatchEvent":
			fmt.Printf("Starred %s\n", event.Repo.Name)
		}
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <username>")
		os.Exit(1)
	}

	username := os.Args[1]
	events, err := fetchGitHubActivity(username)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	displayActivity(events)
}
