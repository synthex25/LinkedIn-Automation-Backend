package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

/*
Global rate-limiting state.

lastRequestTime stores the timestamp of the most recent request
processed by the automation endpoints.

rateLimitMutex protects access to lastRequestTime to ensure
thread-safe behavior across concurrent HTTP requests.
*/
var (
	lastRequestTime time.Time
	rateLimitMutex  sync.Mutex
)

/*
healthHandler provides a simple health check endpoint.

This endpoint is used to verify that the server is running
and capable of handling requests.
*/
func healthHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{
		"status": "OK",
	}
	json.NewEncoder(w).Encode(response)
}

/*
sendMessageHandler simulates sending a LinkedIn message.

Responsibilities:
- Enforces rate limiting to simulate human-like interaction pacing
- Validates incoming JSON payload
- Ensures required fields are present
- Returns appropriate HTTP status codes for success and failure

This endpoint operates in mock mode and does not perform
real LinkedIn message automation.
*/
func sendMessageHandler(w http.ResponseWriter, r *http.Request) {

	// Enforce rate limiting to prevent rapid, bot-like requests
	rateLimitMutex.Lock()
	if time.Since(lastRequestTime) < time.Second {
		rateLimitMutex.Unlock()
		w.WriteHeader(http.StatusTooManyRequests)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "too many requests",
		})
		return
	}
	lastRequestTime = time.Now()
	rateLimitMutex.Unlock()

	// Decode the incoming JSON request body
	var req MessageRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "invalid JSON payload",
		})
		return
	}

	// Validate required input fields
	if req.ProfileURL == "" || req.Message == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "profile_url and message are required",
		})
		return
	}

	// Simulated success response
	response := map[string]string{
		"status":  "message sent",
		"profile": req.ProfileURL,
	}

	json.NewEncoder(w).Encode(response)
}

/*
sendConnectionHandler simulates sending a LinkedIn connection request.

Responsibilities:
- Applies global rate limiting to enforce controlled execution
- Parses and validates JSON input
- Ensures mandatory connection fields are provided
- Returns deterministic and consistent responses

This endpoint represents the connection automation stage
of a LinkedIn workflow.
*/
func sendConnectionHandler(w http.ResponseWriter, r *http.Request) {

	// Enforce rate limiting to avoid burst behavior
	rateLimitMutex.Lock()
	if time.Since(lastRequestTime) < time.Second {
		rateLimitMutex.Unlock()
		w.WriteHeader(http.StatusTooManyRequests)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "too many requests",
		})
		return
	}
	lastRequestTime = time.Now()
	rateLimitMutex.Unlock()

	// Decode the incoming JSON request body
	var req ConnectionRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "invalid JSON payload",
		})
		return
	}

	// Validate required input fields
	if req.ProfileURL == "" || req.Note == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "profile_url and note are required",
		})
		return
	}

	// Simulated success response
	response := map[string]string{
		"status":  "connection request sent",
		"profile": req.ProfileURL,
	}

	json.NewEncoder(w).Encode(response)
}

/*
loggingMiddleware logs incoming HTTP requests.

It records the HTTP method and request path for basic observability
and debugging purposes.
*/
func loggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("[INFO] %s %s\n", r.Method, r.URL.Path)
		next(w, r)
	}
}
