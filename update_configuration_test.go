package main

import (
	"encoding/json"
	"net/http"
	"testing"
)

func TestUpdateDeviceValid(t *testing.T) {
	config := Configuration{
		Configuration: nil,
	}
	body, err := json.Marshal(config)
	if err != nil {
		t.Fatal(err)
	}
	w := testRequestStatusCode("PUT", "/configurations/5e1013575699e4fbbc806e3e", body, http.StatusOK, t)
	testRequestBody(w, "ModifiedCount", 0, t)
}

func TestUpdateDeviceInvalid(t *testing.T) {
	config := Configuration{
		Configuration: nil,
	}
	body, err := json.Marshal(config)
	if err != nil {
		t.Fatal(err)
	}
	w := testRequestStatusCode("PUT", "/configurations/5e1013575699e4fbbc806e3e", body, http.StatusOK, t)
	testRequestBody(w, "MatchedCount", 0, t)
}
