package main

import (
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

import (
	"graphite"
	"helper"
)

func main() {
	// host:8080/bucket/<optional subbuckets>/unix_timestamp/value
	graphite, err := graphite.NewWithConnection(":2003")
	if (err != nil) {
		log.Fatal(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		key, time, value, err := helper.ParseUrl(r.URL)
		if err != nil {
			log.Printf("%v  (%s)", err, r.URL.Path)
			http.Error(w, "Bad Request", 400)
			return
		}

		graphite.Send(key, time, value)
	})

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
