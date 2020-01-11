package main

import (
	"net/http"
	"testing"
)

func TestValidDeviceGet(t *testing.T) {
	//Object ID - 5e1013575699e4fbbc806e3e is always available as a test document
	testRequestStatusCode("GET", "/configurations/5e1013575699e4fbbc806e3e", nil, http.StatusOK, t)
}

func TestInvalidDeviceGet(t *testing.T) {
	testRequestStatusCode("GET", "/configurations/000000000000000000000000", nil, http.StatusNotFound, t)
}
