package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"reflect"

	nmap "github.com/lair-framework/go-nmap"
)

var IPs = []string{"192.168.96.184", "192.168.96.59", "192.168.96.205"}

type Payload struct {
	Devices []string `json:"version"`
}

func main() {
	payload := &Payload{}
	err := json.Unmarshal([]byte(os.Args[1]), payload)
	if err != nil {
		log.Fatalf("error unmarshalling payload: %s", err)
	}

	fmt.Printf("-- previous state --\n%v\n\n", payload.Devices)

	currentDevices := getCurrentState()

	fmt.Printf("-- current state --\n%v\n\n", currentDevices)

	if reflect.DeepEqual(payload.Devices, currentDevices) {
		writeVersions([][]string{payload.Devices})
		return
	}

	writeVersions([][]string{payload.Devices, currentDevices})
}

func getCurrentState() []string {
	var hwAddr string
	var hwAddrs []string

	for _, ip := range IPs {
		xml, err := exec.Command("nmap", "-oX", "-", "-sPn", "-PS22", ip).CombinedOutput()
		if err != nil {
			log.Fatalf("error executing nmap: %s -> %s\n", err, string(xml))
		}

		// parse XML result
		nmapRun, err := nmap.Parse(xml)
		if err != nil {
			log.Fatalf("error parsing nmap xml: %s\n", err)
		}

		if len(nmapRun.Hosts) != 0 {
			// pull MAC from result
			for _, addr := range nmapRun.Hosts[0].Addresses {
				if addr.AddrType == "mac" {
					hwAddr = addr.Addr
				}
			}

			hwAddrs = append(hwAddrs, hwAddr)
		}
	}

	return hwAddrs
}

func writeVersions(versions [][]string) {
	newVersions, err := json.Marshal(versions)
	if err != nil {
		log.Fatalf("error marshalling payload: %s", err)
	}

	_, err = os.Stdout.Write(newVersions)
	if err != nil {
		panic(err)
	}
}
