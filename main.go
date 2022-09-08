package main

import (
    "encoding/json"
	"net/http"
)

type Podcast struct {
	Name string `json:"Name"`
	Author string `json:"Author"`
}

func PodcastHandler(w http.ResponseWriter, r *http.Request) {
    podcasts := []Podcast{
        Podcast{"JRE", "Joe Rogan"},
        Podcast{"CWE", "Chris Wiggins"},
    }

    json.NewEncoder(w).Encode(podcasts)
}

func main() {
    http.HandleFunc("/podcasts", PodcastHandler)
    http.ListenAndServe(":1234", nil)
}
