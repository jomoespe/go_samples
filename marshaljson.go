package main 

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	// JSON message structs
	type Brand struct {
		Name    string  `json:"name"` 
		BrandId string  `json:"brandId"` 
	}

	type BrandMessage struct {
		Document Brand  `json:"document"`
	}

	// test data creation
	theBrand := Brand {
		Name: "marca 1", 
		BrandId: "1",
	}

	theBrandMsg := BrandMessage {
		Document:  theBrand,
	}

	// struct to json serialization
	b, err := json.Marshal(theBrandMsg)
	if err != nil {
		fmt.Println("error:", err)
	}
	os.Stdout.Write(b)
}