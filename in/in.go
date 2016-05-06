package main

import (
	"encoding/json"
	"log"
	"os"
)

type Payload struct {
	Versions []string   `json:"version"`
	MetaData []MetaData `json:"metadata"`
}

type MetaData struct {
	Name string `json:"name"`
	MAC  string `json:"mac"`
}

func main() {
	payload := &Payload{}
	err := json.Unmarshal([]byte(os.Args[1]), payload)
	if err != nil {
		log.Fatalf("error unmarshalling payload: %s", err)
	}

	payload.MetaData = []MetaData{
		{Name: "Michael", MAC: "SO:ME:MA:C1"},
		{Name: "David", MAC: "SO:ME:MA:C2"},
		{Name: "Gabe", MAC: "SO:ME:MA:C3"},
	}

	jsonResp, err := json.Marshal(payload)
	if err != nil {
		log.Fatalf("error marshalling payload: %s", err)
	}

	_, err = os.Stdout.Write(jsonResp)
	if err != nil {
		panic(err)
	}
}
