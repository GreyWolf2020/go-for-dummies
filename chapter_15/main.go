package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	url := "https://data.fixer.io/api/latest?access_key=naLDh6Zdamkaa6vF8NJk9HRdbAzXc1xc"

	if resp, err := http.Get(url); err == nil {
		defer resp.Body.Close()
		if body, err := ioutil.ReadAll(resp.Body); err == nil {
			fmt.Println(string(body))
		} else {
			log.Fatal(err)
		}
	}
	fmt.Println("Done")
}
