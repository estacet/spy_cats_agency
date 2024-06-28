package model

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Weight struct {
	imperial string
	metric   string
}

type Breed struct {
	Weight
	Id string `json:"id"`
}

func GetBreeds() (*[]Breed, error) {
	const apiEndpoint = "https://api.thecatapi.com/v1/breeds"

	client := &http.Client{
		Timeout: time.Second * 5,
	}

	req, err := http.NewRequest(http.MethodGet, apiEndpoint, nil)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	fmt.Println(resp.StatusCode)

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer resp.Body.Close()

	breeds := &[]Breed{}

	if err := json.Unmarshal(bodyBytes, breeds); err != nil {
		fmt.Println(err)
		return nil, err
	}

	return breeds, nil
}
