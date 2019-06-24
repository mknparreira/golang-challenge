package tests

import (
	"testing"
	"net/http"
	"encoding/json"
	"net/http/httptest"
	"challenger/endpoints"
)

func TestDisplayNumbers1To100(t *testing.T){
	request, _ := http.NewRequest("GET", "/challenger/api/v1", nil)

	response := httptest.NewRecorder()
	handler := http.HandlerFunc(endpoints.IndexHandler)
	handler.ServeHTTP(response, request)

	var m interface{}
	decoder := json.NewDecoder(response.Body)
	decoder.Decode(&m)
	
	if m == nil {
		t.Errorf("handler returned unexpected body: got %v",
		response.Body.String())
	}
}