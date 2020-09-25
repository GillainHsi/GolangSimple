package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleDataGet(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/data", handleReq)

	writer := httptest.NewRecorder()
	request, _ := http.NewRequest("POST", "/data", nil)
	mux.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Errorf("Response code is %v\n", writer.Code)
	}

}
