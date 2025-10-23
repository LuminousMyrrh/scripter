package api

import (
	"encoding/json"
	"net/http"
)

type SearchResponse struct {
	Items []struct {
		FullName string `json:"full_name"`
		Description string `json:"description"`
		HTMLURL string `json:"html_url"`
	} `json:"items"`
}

func AskGithub(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", err
	}

	var res SearchResponse
	if err = json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return "", err
	}

	if len(res.Items) == 0 {
		return "", err
	}

	return res.Items[0].FullName, nil

}
