package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateDeviceAndDelete(t *testing.T) {
	config := Configuration{
		DeviceToken:       "TestToken",
		DeviceSerial:      "TestSerial",
		ConfigurationType: "TestConfigType",
		Configuration:     nil,
	}
	body, err := json.Marshal(config)
	if err != nil {
		t.Fatal(err)
	}
	r := newRouter()
	createRecorder := httptest.NewRecorder()
	createRequest, _ := http.NewRequest("POST", "/configurations", bytes.NewBuffer(body))
	r.ServeHTTP(createRecorder, createRequest)
	if createRecorder.Code != http.StatusCreated {
		t.Errorf("Status should be 201, got %d", createRecorder.Code)
	}

	_ = json.NewDecoder(createRecorder.Body).Decode(&config)
	deleteRequest, _ := http.NewRequest("DELETE", "/configurations/"+config.ID.Hex(), nil)
	deleteRecorder := httptest.NewRecorder()
	r.ServeHTTP(deleteRecorder, deleteRequest)
	if deleteRecorder.Code != http.StatusNotFound {
		t.Errorf("Status should be 404, got %d", deleteRecorder.Code)
	}
}
