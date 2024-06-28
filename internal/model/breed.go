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
	id                 string
	name               string
	cfa_url            string
	vetstreet_url      string
	vcahospitals_url   string
	temperament        string
	origin             string
	country_codes      string
	country_code       string
	description        string
	life_span          string
	indoor             int
	lap                int
	alt_names          string
	adaptability       int
	affection_level    int
	child_friendly     int
	dog_friendly       int
	energy_level       int
	grooming           int
	health_issues      int
	intelligence       int
	shedding_level     int
	social_needs       int
	stranger_friendly  int
	vocalisation       int
	experimental       int
	hairless           int
	natural            int
	rare               int
	rex                int
	suppressed_tail    int
	short_legs         int
	wikipedia_url      string
	hypoallergenic     int
	reference_image_id string
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
