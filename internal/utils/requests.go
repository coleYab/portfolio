package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"portfolio/types"
)

// FetchRepositories fetches public repositories for a user, sorted by last commit time
func FetchRepositories(username string) ([]types.Repository, error) {
	apiURL := fmt.Sprintf("https://api.github.com/users/%s/repos?sort=pushed&direction=desc", username)

	// Create a new HTTP request
	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	// Set headers
	req.Header.Set("Accept", "application/vnd.github+json")

	// Execute the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error fetching repositories: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: status code %d", resp.StatusCode)
	}

	// Decode the response body into a list of repositories
	var repositories []types.Repository
	if err := json.NewDecoder(resp.Body).Decode(&repositories); err != nil {
		return nil, fmt.Errorf("error decoding JSON: %v", err)
	}

	return repositories, nil
}
