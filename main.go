// The rootRouteHandler function in the Go code extracts request details, handles optional sleep and
// status parameters, reads the request body, and returns a JSON response with the request details.
package main

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

type HTTPRequestDetails struct {
	Method string              `json:"method"`
	Path   string              `json:"path"`
	Header map[string][]string `json:"header"`
	Body   string              `json:"body"`
}

func main() {
	log.Println("== Start App ==")

	// Create a new HTTP server
	srv := &http.Server{
		Addr: ":8080",
	}

	// Define HTTP handlers
	http.HandleFunc("/", rootRouteHandler)

	// Channel to listen for interrupt signals
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	// Run the server in a goroutine
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			if err != http.ErrServerClosed {
				log.Fatalf("Could not listen on %s: %v\n", srv.Addr, err)
			}
		}
	}()
	log.Printf("Server is ready to handle requests at %s", srv.Addr)

	// Block until a signal is receivede
	<-stop
	log.Printf("Server is shutting down...")

	// Create a deadline to wait for
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Attempt a graceful shutdown
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Printf("Server exiting")

	log.Println("== End App ==")
}

// rootRouteHandler handles the root route and extracts request details, handles optional sleep and
// status parameters, reads the request body, and returns a JSON response with the request details.
func rootRouteHandler(w http.ResponseWriter, r *http.Request) {
	// Extract request details
	details := HTTPRequestDetails{
		Method: r.Method, // HTTP method
		Path:   r.URL.String(),
		Header: r.Header, // Request headers
	}

	log.Printf("Got request for %s with path %s\n", details.Method, details.Path)

	// Sleep for a specified duration if the 'sleep' query parameter is provided
	sleepParam := r.URL.Query().Get("sleep")
	if sleepParam != "" {
		// Sleep for the specified duration
		sleepTime, err := strconv.Atoi(sleepParam)
		if err != nil {
			// Error parsing sleep duration
			log.Println("Error parsing sleep duration:", err)
			http.Error(w, "Error parsing sleep duration", http.StatusBadRequest)
			return
		}
		time.Sleep(time.Duration(sleepTime) * time.Millisecond)
	}

	// Handle the 'status' query parameter if provided
	statusParam := r.URL.Query().Get("status")
	if statusParam != "" {
		// Convert the status code string to an integer
		statusCode, err := strconv.Atoi(statusParam)
		if err != nil {
			// Invalid status code format
			http.Error(w, "Invalid status code format", http.StatusBadRequest)
			return
		}
		// Check if the status code is within the valid range (100 - 599)
		if statusCode < http.StatusContinue || statusCode > http.StatusInternalServerError {
			// Invalid status code range
			http.Error(w, "Invalid status code range", http.StatusNotFound)
			return
		}

		// Write the status code as the response body
		w.WriteHeader(statusCode)
	}

	// Read request body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		// Error reading request body
		log.Println("Error reading request body:", err)
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}
	defer func() {
		if err := r.Body.Close(); err != nil {
			// Error closing request body
			log.Println("Error closing request body:", err)
		}
	}()

	details.Body = string(body)

	// Marshal request details to JSON
	jsonData, err := json.Marshal(details)
	if err != nil {
		// Error marshalling request details to JSON
		log.Println("Error marshalling request details to JSON:", err)
		http.Error(w, "Error marshalling request body", http.StatusInternalServerError)
		return
	}

	// Set content type and write JSON response
	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(jsonData); err != nil {
		// Error writing JSON response
		log.Println("Error writing JSON response:", err)

		// Consider returning a more specific error code depending on the nature of the write error
		http.Error(w, "Error writing response", http.StatusInternalServerError)
		return
	}
}
