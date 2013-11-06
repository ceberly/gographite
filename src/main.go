package main

import (
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

import "graphite"

func ParseUrl(url *url.URL) ([]string, int64, float32, error) {
	path := strings.Split(url.Path, "/")
	l := len(path)

	t, err := strconv.ParseInt(path[l-2], 10, 64)
	if err != nil {
		return nil, 0, 0, err
	}

	//TODO: validate time

	v, err2 := strconv.ParseFloat(path[l-1], 32)
	if err2 != nil {
		return nil, 0, 0, err
	}

	return path[:l-3], t, float32(v), nil
}

func main() {
	// host:8080/bucket/<optional subbuckets>/unix_timestamp/value
	graphite, err := graphite.NewWithConnection()
	if (err != nil) {
		log.Fatal(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		key, time, value, err := ParseUrl(r.URL)
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
