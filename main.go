package main

import (
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	// new => initiates a new verification session, takes email and userID as parameters
	// returns a sessionID
	mux.HandleFunc("/new", func(w http.ResponseWriter, r *http.Request) {})
	// verify => takes the sessionID and a verification code, verifies the code against the session
	mux.HandleFunc("/verify", func(w http.ResponseWriter, r *http.Request) {})
	err := http.ListenAndServe(":8000", mux)
	if err != nil {
		panic(err)
	}
}
