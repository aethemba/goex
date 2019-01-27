package github

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const BaseURL = "https://api.github.com/repos"

func ReadIssue(author string, repo string, number string) (*Issue, error) {

	issueUrl := BaseURL + "/" + author + "/" + repo + "/issues/" + number

	client := &http.Client{}

	resp, err := client.Get(issueUrl)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("reading issues failed: %s", resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalf("Error reading response body: %s", err)
	}
	resp.Body.Close()

	var result Issue
	if err := json.Unmarshal(body, &result); err != nil {
		log.Fatalf("JSON unmarshaling failed: %s", err)
	}

	// fmt.Printf("id: %d\n", result.Number)
	// fmt.Printf("url: %s\n", result.Url)
	// fmt.Printf("body: %s\n", result.Body)

	return &result, nil

}
