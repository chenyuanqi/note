package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

func main() {
	type FruitBasket struct {
		Name    string
		Fruit   []string
		Id      int64  `json:"ref"`
		private string // An unexported field is not encoded.
		Created time.Time
	}

	basket := FruitBasket{
		Name:    "Standard",
		Fruit:   []string{"Apple", "Banana", "Orange"},
		Id:      999,
		private: "Second-rate",
		Created: time.Now(),
	}

	var jsonData []byte
	//  json.Marshal function in package encoding/json generates JSON data.
	jsonData, err := json.Marshal(basket)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(jsonData)) // {"Name":"Standard","Fruit":["Apple","Banana","Orange"],"ref":999,"Created":"2022-04-29T23:37:17.087307+08:00"}

	// Replace json.Marshal with json.MarshalIndent in the example above to indent the JSON output
	jsonData, err = json.MarshalIndent(basket, "", "    ")
	fmt.Println(string(jsonData))
	/*
		{
			"Name": "Standard",
			"Fruit": [
				"Apple",
				"Banana",
				"Orange"
			],
			"ref": 999,
			"Created": "2022-04-29T23:39:28.038571+08:00"
		}
	*/

	jsonData = []byte(`
	{
		"Name": "Standard",
		"Fruit": [
			"Apple",
			"Banana",
			"Orange"
		],
		"ref": 999,
		"Created": "2018-04-09T23:00:00Z"
	}`)
	// The json.Unmarshal function in package encoding/json parses JSON data
	err = json.Unmarshal(jsonData, &basket)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(basket.Name, basket.Fruit, basket.Id) // Standard [Apple Banana Orange] 999
	fmt.Println(basket.Created)                       // 2018-04-09 23:00:00 +0000 UTC
}
