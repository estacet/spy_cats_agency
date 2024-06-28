package catapi

import (
	"encoding/json"
	"io"
	"net/http"
	"time"
)

type Client struct {
	httpClient *http.Client
}

func NewClient() *Client {
	return &Client{
		httpClient: &http.Client{
			Timeout: time.Second * 5,
		},
	}
}

func (c *Client) GetBreeds() (BreedResponse, error) {
	const apiEndpoint = "https://api.thecatapi.com/v1/breeds"

	req, err := http.NewRequest(http.MethodGet, apiEndpoint, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var breedResp BreedResponse

	if err := json.Unmarshal(bodyBytes, &breedResp); err != nil {
		return nil, err
	}

	return breedResp, nil
}
