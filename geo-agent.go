package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"

	"github.com/lair-framework/go-nmap"
)

var IPs = []string{"192.168.96.184", "192.168.96.59", "192.168.96.205"}

func main() {
	http.HandleFunc("/", viewHandler)
	http.ListenAndServe(":8080", nil)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	p := getCurrentState()
	fmt.Fprintf(w, "%+v", p)
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

		fmt.Printf("result from nmap: %+v\n", nmapRun)

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
