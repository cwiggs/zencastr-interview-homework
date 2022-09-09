package main

import (
	"net/http"
	//"net/http/httptest"
	"testing"
)

func TestGetPodcasts(t *testing.T) {
	_, err := http.NewRequest(http.MethodGet, "/podcasts", nil)
	if err != nil {
		t.Fatal(err)
	}
	/*
	   w := httptest.NewRecorder()
	   RequestHandler(w, req)
	   result := w.Result()
	   defer result.Body.Close()

	   data, err := ioutil.ReadAll(result.Body)

	   if err != nil {
	       t.Errorf("Error: %v", err)
	   }

	   if string(data) != expected {
	       t.Errorf("Expected Joe Rogan but got %v", string(data))
	   }
	*/
}
