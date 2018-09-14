package main

import (
	"net/http"
	"testing"
)

func TestPutDevice(t *testing.T) {

}

func TestPutData(t *testing.T) {

}

func TestGetDevicesByUser(t *testing.T) {

}

func TestCreateDevice(t *testing.T) {

}

func TestGetDeviceByCreds(t *testing.T) {
	_, err := http.NewRequest("GET", "/v1/users/list", nil)
	if err != nil {
		t.Fatal(err)
	}

	// c := connections

	// rr := httptest.NewRecorder()
	// handler := http.HandlerFunc(GetDeviceByCreds)

}
