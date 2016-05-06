package main

import (
	"encoding/json"
	"log"
	"os"
)

type HWAddrs struct {
	MACs []string `json:"macs"`
}

type Payload struct {
	Version  HWAddrs `json:"version"`
	MetaData HWAddrs `json:"metadata"`
}

func main() {
	if len(os.Args) < 2 {
		println("usage: " + os.Args[0] + " <destination>")
		os.Exit(1)
	}

	payload := &Payload{}
	err := json.Unmarshal([]byte(os.Args[2]), payload)
	if err != nil {
		log.Fatalf("error unmarshalling payload: %s", err)
	}

	for _, mac := range payload.Version.MACs {
		payload.MetaData.MACs = append(payload.MetaData.MACs, mac)
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
