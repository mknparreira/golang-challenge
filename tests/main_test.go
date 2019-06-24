package tests

import (
	"golang-challenge/endpoints"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRouteIndexHandler(t *testing.T) {
	request, err := http.NewRequest("GET", "/challenger/api/v1", nil)

	if err != nil {
		t.Fatal(err)
	}

	response := httptest.NewRecorder()
	handler := http.HandlerFunc(endpoints.IndexHandler)
	handler.ServeHTTP(response, request)

	if status := response.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestRouteRangeHandler(t *testing.T) {
	request, err := http.NewRequest("GET", "/challenger/api/v1/from/1/to/5", nil)

	if err != nil {
		t.Fatal(err)
	}

	response := httptest.NewRecorder()
	handler := http.HandlerFunc(endpoints.IndexHandler)
	handler.ServeHTTP(response, request)

	if status := response.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
