package main

import (
	"encoding/json"
	"log"
	"os"

	ditch "github.com/VReactor/discount-itch"
)

type itchioItems struct {
	Games []ditch.Item `json:"games"`
}

func itchioItemsToJSON(itchioItems itchioItems) string {
	b, err := json.Marshal(itchioItems)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func logItems(itchioItems itchioItems) {
	log.Println("all games:")
	for _, item := range itchioItems.Games {
		log.Println(item.Link)
	}
}

func main() {
	// init empty list for games
	result := itchioItems{
		Games: []ditch.Item{},
	}
	
	// get all games that are on 100% discount
	log.Println("Getting items for category:", ditch.Games)
	items, err := ditch.GetCategoryItems(ditch.Games)
	if err != nil {
		log.Println(err)
	} else {
		result.Games = append(result.Games, items...)
	}

	logItems(result)

	// transform items to json
	resultJSON := itchioItemsToJSON(result)

	// resultJSON to a file
	err = os.WriteFile("items.json", []byte(resultJSON), 0644)
	if err != nil {
		panic(err)
	}
}
