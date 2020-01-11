package main

import (
	"encoding/json"
	"net/http"
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
	createRecorder := testRequestStatusCode("POST", "/configurations", body, http.StatusCreated, t)

	_ = json.NewDecoder(createRecorder.Body).Decode(&config)
	testRequestStatusCode("DELETE", "/configurations/"+config.ID.Hex(), nil, http.StatusNotFound, t)
}
