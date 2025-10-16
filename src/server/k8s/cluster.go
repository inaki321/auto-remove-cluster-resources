// File: server/k8s/server.go
package k8s

import (
    "fmt"
    "log"
    "net/http"

)

// HomeHandler serves a basic "Hello" message and logs the request.
// It is exported to be used by the main server package.
func HomeHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/plain")

    // The message is adjusted to reflect the k8s package's context
    fmt.Fprintf(w, "Hello from the k8s cluster controller!")

    log.Printf("Received request from %s for path: %s", r.RemoteAddr, r.URL.Path)
}