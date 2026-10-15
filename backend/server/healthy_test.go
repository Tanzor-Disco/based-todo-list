package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleHealthReturnsOK(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/health", nil)
	response := httptest.NewRecorder()

	handleHealth(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf(
			"expected status %d, got %d",
			http.StatusOK,
			response.Code,
		)
	}
}
