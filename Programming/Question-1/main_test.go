package main 
import (
    "net/http"
    "net/http/httptest"
    "testing"
    "time"
)

// TestSuccessfulRequest checks if the function can handle a 200 OK response properly.
func TestSuccessfulRequest(t *testing.T) {
    server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK) // return 200 OK
    }))
    defer server.Close()

    _, err := exponentialBackoff(server.URL, 1, 10*time.Second)
    if err != nil {
        t.Errorf("Expected no error, got %v", err)
    }
}

// TestRetryOnFailure checks if the function retries the request after an initial failure.
func TestRetryOnFailure(t *testing.T) {
    attempts := 0
    server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        attempts++
        if attempts == 1 {
            w.WriteHeader(http.StatusInternalServerError) // return 500 on first attempt
        } else {
            w.WriteHeader(http.StatusOK) // return 200 on subsequent attempts
        }
    }))
    defer server.Close()

    _, err := exponentialBackoff(server.URL, 5, 10*time.Second)
    if err != nil {
        t.Errorf("Expected no error, got %v", err)
    }
    if attempts != 2 {
        t.Errorf("Expected 2 attempts, got %d", attempts)
    }
}

// TestNonRetryableStatusCode checks if the function stops retrying after a non-retryable status code.
func TestNonRetryableStatusCode(t *testing.T) {
    server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusBadRequest) // return 400 Bad Request
    }))
    defer server.Close()

    resp, _ := exponentialBackoff(server.URL, 5, 10*time.Second)
    if resp.StatusCode != http.StatusBadRequest {
        t.Errorf("Expected 400 Bad Request, got %d", resp.StatusCode)
    }
}

// TestTimeout checks if the function respects the timeout specified.
func TestTimeout(t *testing.T) {
    server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        time.Sleep(6 * time.Second) // Delay the response longer than the timeout
        w.WriteHeader(http.StatusOK)
    }))
    defer server.Close()

    start := time.Now()
    _, _ = exponentialBackoff(server.URL, 5, 5*time.Second)
    if time.Since(start) > 6*time.Second {
        t.Errorf("Expected timeout in about 5 seconds, took longer")
    }
}
