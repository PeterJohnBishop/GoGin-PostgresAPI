package routes

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// run with: go test . -v

// === RUN   TestTestHandler
//
//	test_test.go:23: Recieved the expected status code.
//	test_test.go:29: Response body verified as map[string]interface{} (JSON).
//	test_test.go:35: Recieved the expected welcome message.
//
// --- PASS: TestTestHandler (0.00s)
func TestTestHandler(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	TestHandler(w, r)

	resp := w.Result()
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, resp.StatusCode)
	}
	t.Log("Recieved the expected status code.")

	var responseBody map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&responseBody); err != nil {
		t.Fatalf("failed to decode response body: %v", err)
	}
	t.Log("Response body verified as map[string]interface{} (JSON).")

	expectedMessage := "Welcome to my Go/HTTP server!"
	if responseBody["message"] != expectedMessage {
		t.Errorf("expected message %q, got %q", expectedMessage, responseBody["message"])
	}
	t.Log("Recieved the expected welcome message.")

}

// === RUN   TestTestHandler_InvalidMethod
// --- PASS: TestTestHandler_InvalidMethod (0.00s)
func TestTestHandler_InvalidMethod(t *testing.T) {
	r := httptest.NewRequest(http.MethodPost, "/", nil)
	w := httptest.NewRecorder()

	TestHandler(w, r)

	resp := w.Result()
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("expected status %d, got %d", http.StatusMethodNotAllowed, resp.StatusCode)
	}
}
