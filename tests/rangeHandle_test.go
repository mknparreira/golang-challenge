package tests

import (
	"testing"
	"net/http"
	"encoding/json"
	"net/http/httptest"
	"challenger/endpoints"
)

func TestDisplayNumbersAfterGetParams(t *testing.T){
	request, _ := http.NewRequest("GET", "/challenger/api/v1/from/1/to/5", nil)

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