package main

import (
    "flag"
    "log"
    "math"
    "net/http"
    "time"
)
// exponentialBackoff sends a GET request to the specified URL and retries the request with exponential backoff if the response status code is not 200 OK.
func exponentialBackoff(url string, maxRetries int, timeout time.Duration) (*http.Response, error) {
    var resp *http.Response
    var err error
    client := &http.Client{}
    retryInterval := 1.0 // Start with 1 second

    start := time.Now()

    for i := 0; i < maxRetries; i++ {
        resp, err = client.Get(url)
        if err != nil {
            log.Printf("Error when sending request: %v", err)
            return nil, err
        }

        if resp.StatusCode == http.StatusOK {
            return resp, nil
        } else if resp.StatusCode == http.StatusNotFound || resp.StatusCode == http.StatusBadRequest ||
                  resp.StatusCode == http.StatusForbidden || resp.StatusCode == http.StatusConflict {
            log.Printf("Received non-retryable status code: %d", resp.StatusCode)
            return resp, nil
        }

        duration := time.Duration(math.Pow(2, float64(i))) * time.Second * time.Duration(retryInterval)
        if time.Since(start) + duration > timeout {
            log.Println("Timeout reached, stopping retries")
            return resp, err
        }

        log.Printf("Retrying in %v seconds...", duration.Seconds())
        time.Sleep(duration)
    }

    return resp, err
}

func main() {
    var url string
    var maxRetries int
    var timeoutSec int

    flag.StringVar(&url, "url", "https://httpbin.org/get", "URL to request")
    flag.IntVar(&maxRetries, "retries", 5, "Maximum number of retries")
    flag.IntVar(&timeoutSec, "timeout", 60, "Timeout in seconds")
    flag.Parse()

    timeout := time.Duration(timeoutSec) * time.Second

    response, err := exponentialBackoff(url, maxRetries, timeout)
    if err != nil {
        log.Printf("Failed to get a successful response: %v", err)
    } else {
        log.Printf("Success: %d %s", response.StatusCode, response.Status)
    }
}
