package main

import(
	"github.com/SirEyeque/readJSON"
	"encoding/json"
	"log"
	"fmt"
)

type locArea struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func commandMap(urls * config) error{
	url := ""

	if urls.next == ""{ 
		url = "https://pokeapi.co/api/v2/location-area/"
	}else{
		url = urls.next
	}

	body, err := readJSON.Read(url)

	if err != nil{
		fmt.Printf("%w", err)
		log.Fatal(err)
	}

	dat := locArea{}

	if err = json.Unmarshal(body, &dat); err != nil{
		log.Fatalf("Error: failed with error: %w", err)
	}

	urls.next = dat.Next
	urls.prev = dat.Previous

	for _, item := range dat.Results{
		fmt.Printf("%s\n", item.Name)
	}

	return nil
}
