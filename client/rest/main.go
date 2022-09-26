package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	log "github.com/sirupsen/logrus"
)

func main() {

	address := flag.String("server", "http://localhost:9000", "HTTP gateway url, e.g. http://localhost:9000")
	flag.Parse()

	login(address)

}

func login(address *string) {
	var body string

	resp, err := http.Post(*address+"/v1/login", "application/json", strings.NewReader(fmt.Sprintf(`
		{
			"apiVersion":"v1",
			"emailID": "sjnjaiswal3@gmail.com",
			"password":"password1"
		}
	`)))

	if err != nil {
		log.Fatalf("failed to call login method: %v", err)
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		body = fmt.Sprintf("failed read Login response body: %v", err)
	} else {
		body = string(bodyBytes)
	}
	log.Printf("Loging response: Code=%d, Body=%s\n\n", resp, body)

}
