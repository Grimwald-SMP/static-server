package main

import "fmt"
import "net/http"
import "os"

func addCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cross-Origin-Resource-Sharing", "*")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		next.ServeHTTP(w, r)
	})
}

func githubHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from the GitHub handler!")
}

func main() {
	// Load from environment variables
	port := os.Getenv("PORT")
	if port == "" {
		port = ":8082"
	}
	static_dir := os.Getenv("STATIC_DIR")
	if static_dir == "" {
		static_dir = "./files"
	}

	mux := http.NewServeMux()

	// Serve static
	handler := addCORS(http.StripPrefix("/static/", http.FileServer(http.Dir(static_dir))))
	mux.Handle("/static/", handler)
	mux.HandleFunc("POST /github", githubHandler)

	// Start server
	fmt.Printf("Server listening on port %s - http://localhost%s\n", port, port)
	err := http.ListenAndServe(port, mux)
	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
