package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

const (
	username    = "<github username>"
	baseURL     = "https://api.github.com"
	accessToken = "<github personal access token>"
)

type Repository struct {
	Name    string `json:"name"`
	Private bool   `json:"private"`
}

func main() {
	// List all of the repositories for the authenticated user
	url := fmt.Sprintf("%s/user/repos", baseURL)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		panic(err)
	}

	// Add the access token as a header
	req.Header.Set("Authorization", fmt.Sprintf("token %s", 
accessToken))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var repositories []Repository
	err = json.NewDecoder(resp.Body).Decode(&repositories)
	if err != nil {
		panic(err)
	}

	// Make all repositories public
	for _, repo := range repositories {
		url = fmt.Sprintf("%s/repos/%s/%s", baseURL, username, 
repo.Name)
		payload := 
strings.NewReader(fmt.Sprintf(`{"name":"%s","private":false}`, repo.Name))

		req, err := http.NewRequest(http.MethodPatch, url, 
payload)
		if err != nil {
			panic(err)
		}

		// Add the access token as a header
		req.Header.Set("Authorization", fmt.Sprintf("token %s", 
accessToken))
		req.Header.Set("Content-Type", "application/json")

		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		fmt.Printf("%s set to public\n", repo.Name)
	}
}
