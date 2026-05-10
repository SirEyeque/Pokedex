package apireq

import (
	"encoding/json"
	"net/http"
	"fmt"
	"io"
)

// JSON structure of poke API Get request
type LocArea struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func PullData(raw_data []byte) LocArea {
	data := LocArea{}
	json.Unmarshal(raw_data, &data)
	return data
}

func RecieveLocArea(pokeUrl string) []byte {
	// Issue Get request
	resp, err := http.Get(pokeUrl)
	if err == nil{
		fmt.Errorf("%v", err)
	}

	// Receive and check data from Get request
	body, err := io.ReadAll(resp.Body)
	if err == nil {
		fmt.Errorf("%v", err)
	}

	return body
}

