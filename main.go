package main

import (
	"fmt"
	"github.com/charmbracelet/log"
	"net/http"
	"os"
)

func authCallbackHandler(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	state := r.URL.Query().Get("state")
	errorParam := r.URL.Query().Get("error")
	errorDescription := r.URL.Query().Get("error_description")

	if errorParam != "" {
		fmt.Fprintf(w, "Authorization failed: %s - %s\n", errorParam, errorDescription)
		log.Error("Auth failed: ", errorParam, errorDescription)
		return
	}

	if code != "" {
		fmt.Fprintf(w, "Authorization code: %s\n", code)
		fmt.Fprintf(w, "You can now close this page and paste the code in your terminal app.\n")
		log.Info("Authorization received!!", code, state)
		fmt.Printf("\nCode: %s\n%s", code, state)

	} else {
		fmt.Fprintf(w, "No Authorization code or error received.\n")
		fmt.Printf("\nNo auth code or error received\n")
		log.Error("No data received....")
	}
}

func main() {
	//port := os.Getenv("PORT")
	//if port != "" {
	//	port = "8080"
	//}

	port := "8080"

	http.HandleFunc("/callback", authCallbackHandler)

	fmt.Printf("Listening on port %s\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("Error starting callback server: %v", err)
		os.Exit(1)
	}
}
