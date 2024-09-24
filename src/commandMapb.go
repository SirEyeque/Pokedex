package main

import(
	"github.com/SirEyeque/readJSON"
	"encoding/json"
	"log"
	"fmt"
)


func commandMapb(urls * config) error{
	url := ""

	if urls.prev == ""{ 
		fmt.Printf("There is not a previous page to get map areas from\n")
		return nil
	}else{
		url = urls.prev
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
