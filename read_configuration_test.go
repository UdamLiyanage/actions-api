package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestValidDeviceGet(t *testing.T) {
	//Object ID - 5e1013575699e4fbbc806e3e is always available as a test document
	r := newRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/configurations/5e1013575699e4fbbc806e3e", nil)
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Status should be 200, got %d", w.Code)
	}
}

func TestInvalidDeviceGet(t *testing.T) {
	r := newRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/configurations/000000000000000000000000", nil)
	r.ServeHTTP(w, req)
	if w.Code != http.StatusNotFound {
		t.Errorf("Status should be 404, got %d", w.Code)
	}
}
