package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	type Address struct {
		City    string `json:"City"`
		State   string `json:"State"`
		Country string `json:"Country"`
	}

	type Data struct {
		Kvdb map[string]Address
	}
	empJson := []byte(`
		{
			"Kvdb" : {
				"address1": {
					"City": "Mumbai",
					"State": "Maharashtra",
					"Country": "India"
				},
				"address2": {
					"City": "Mumbai",
					"State": "Maharashtra",
					"Country": "India"
				}	
			}
	}`)

	var result Data

	// Unmarshal or Decode the JSON to the interface.
	json.Unmarshal(empJson, &result)

	fmt.Printf("\nRES: %v", result)
}
