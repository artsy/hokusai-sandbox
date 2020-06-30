package main

import (
	"os"
	"fmt"
	"net/http"
	"log"
)

func domain() string {
	if os.Getenv("STAGING") != "" {
		return "hokusai-sandbox-staging.artsy.net"
	} else {
		return "hokusai-sandbox.artsy.net"
	}
}

func main() {
	var domain = domain()
	fmt.Println(fmt.Sprintf("Pinging %s...", domain))
	for true {
		resp, err := http.Get(fmt.Sprintf("https://%s/", domain))
		if err != nil {
			log.Fatal(err)
		}
		if os.Getenv("DEBUG") != "" {
			log.Printf(resp.Status)
		}
	}
}
