package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type HTTPRequestDetails struct {
	Method string              `json:"method"`
	URL    string              `json:"url"`
	Header map[string][]string `json:"header"`
	Body   string              `json:"body"`
}

func main() {
	log.Println("== Start App ==")
	http.HandleFunc("/", rootRouteHandler)

	log.Println("Server listening on port 8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("== End App ==")
}

func rootRouteHandler(w http.ResponseWriter, r *http.Request) {
	// Extract request details
	details := HTTPRequestDetails{
		Method: r.Method,
		URL:    r.URL.String(),
		Header: r.Header,
	}

	log.Printf("Got request for %s with path %s\n", details.Method, details.URL)

	// Read request body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Error reading request body:", err)
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}
	defer func() {
		if err := r.Body.Close(); err != nil {
			log.Println("Error closing request body:", err)
		}
	}() // Close the body even in case of other errors

	details.Body = string(body)

	// Marshal request details to JSON
	jsonData, err := json.Marshal(details)
	if err != nil {
		log.Println("Error marshalling request details to JSON:", err)
		http.Error(w, "Error marshalling request body", http.StatusInternalServerError)
		return
	}

	// Set content type and write JSON response
	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(jsonData); err != nil {
		log.Println("Error writing JSON response:", err)

		// Consider returning a more specific error code depending on the nature of the write error
		http.Error(w, "Error writing response", http.StatusInternalServerError)
		return
	}
}
