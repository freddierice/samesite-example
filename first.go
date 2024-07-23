package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
)

func serveFirstPage(w http.ResponseWriter, r *http.Request) {
	html := `
	<!DOCTYPE html>
	<html>
	<head>
		<title>First Page</title>
	</head>
	<body>
		<h1>First Page</h1>
	</body>
	</html>
	`
	fmt.Fprint(w, html)
}

func getNumber(w http.ResponseWriter, r *http.Request) {
	// the request must be authenticated with a session cookie
	cookie, err := r.Cookie("trends_session")
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	if cookie.Value != "loggedin" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// return json
	w.Header().Set("Content-Type", "application/json")
	number := rand.Intn(100)
	response := map[string]int{"number": number}
	json.NewEncoder(w).Encode(response)
}

func setCookie(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:     "trends_session",
		Value:    "loggedin",
		Path:     "/",
		Domain:   "first.trends.stream",
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	}
	http.SetCookie(w, &cookie)
	fmt.Fprintf(w, "Cookie set")
}
