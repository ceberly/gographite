package main

import (
	"log"
	"net/http"
)

import (
	"graphite"
	"helper"
)

func main() {
	graphite, err := graphite.NewWithConnection("tcp", ":2003")
	if err != nil {
		log.Fatal(err)
	}

	// This should be handled by your webserver, but is super annoying when testing.
	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Not Found", 404)
	})

	// host:8080/bucket/<optional subbuckets>/unix_timestamp/value
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
