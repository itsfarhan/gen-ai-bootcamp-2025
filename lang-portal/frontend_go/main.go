package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"path/filepath"
	"time"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s %s", r.Method, r.RequestURI, time.Since(start))
	})
}

func main() {
	port := "8080"
	
	// Get the current directory
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get current directory: %v", err)
	}
	
	// Create a file server handler for static files
	staticDir := filepath.Join(currentDir)
	fileServer := http.FileServer(http.Dir(staticDir))
	
	// Create a new ServeMux
	mux := http.NewServeMux()
	
	// Register routes
	
	// Serve HTML files for specific routes
	mux.HandleFunc("/words", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filepath.Join(staticDir, "words.html"))
	})
	
	mux.HandleFunc("/groups", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filepath.Join(staticDir, "groups.html"))
	})
	
	mux.HandleFunc("/settings", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filepath.Join(staticDir, "settings.html"))
	})
	
	mux.HandleFunc("/dashboard", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filepath.Join(staticDir, "index.html"))
	})
	
	// Additional page routes
	mux.HandleFunc("/study-sessions", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filepath.Join(staticDir, "study-sessions.html"))
	})
	
	// Activities route
	mux.HandleFunc("/activities", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filepath.Join(staticDir, "activities.html"))
	})
	
	// Set up reverse proxy for API requests to backend server
	backendURL, err := url.Parse("http://localhost:8082")
	if err != nil {
		log.Fatalf("Error parsing backend URL: %v", err)
	}
	
	// Create the reverse proxy
	proxy := httputil.NewSingleHostReverseProxy(backendURL)
	
	// Log the proxied requests
	originalDirector := proxy.Director
	proxy.Director = func(req *http.Request) {
		originalDirector(req)
		log.Printf("Proxying request: %s %s -> %s", req.Method, req.URL.Path, backendURL.String() + req.URL.Path)
	}
	
	// Handle all /api/* requests with the proxy
	mux.HandleFunc("/api/", func(w http.ResponseWriter, r *http.Request) {
		proxy.ServeHTTP(w, r)
	})
	
	// Default handler for static files
	mux.Handle("/", loggingMiddleware(fileServer))
	
	// Log server startup
	serverAddr := fmt.Sprintf(":%s", port)
	log.Printf("Starting server on http://localhost%s", serverAddr)
	log.Printf("Serving static files from: %s", staticDir)
	
	// Start the server
	err = http.ListenAndServe(serverAddr, mux)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

