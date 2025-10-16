package server

import (
	"fmt"
	"log"
	"net/http"
)

const HostPort = ":5000"

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain");

	fmt.Fprintf(w, "Hello from the autocluster!");

	log.Printf("Received request from %s for path: %s", r.RemoteAddr, r.URL.Path);
}

// StartServer sets up and runs the HTTP server.
// It returns an error if the server fails to start.
func StartServer() error {
	mux := http.NewServeMux();

	mux.HandleFunc("/", homeHandler);

	// Start the server
	log.Printf("Starting server on host http://localhost%s", HostPort);

	// ListenAndServe blocks until the server stops, returning an error if it fails
	// We use the custom multiplexer (mux) instead of nil to control routing
	err := http.ListenAndServe(HostPort, mux);

	// The function only reaches here if ListenAndServe returns an error (e.g., port in use)
	if err != nil && err != http.ErrServerClosed {
		return fmt.Errorf("could not start server: %w", err);
	}

	return nil;
}
