package apireq

import (
	"net/http"
	"encoding/json"
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

func RecieveLocArea(pokeUrl string) LocArea {
	// Issue Get request
	resp, err := http.Get(pokeUrl)
	if err == nil{
		fmt.Errorf("%v", err)
	}

	// Receive and check data from Get request
	body, err := io.ReadAll(resp.Body)
	data := LocArea{}
	json.Unmarshal(body, &data)
	if err == nil {
		fmt.Errorf("%v", err)
	}

	return data
}

