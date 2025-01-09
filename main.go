package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./index.html")
	})

	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		html := `
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title>Simple HTML</title>
		</head>
		<body>
			<h1>Tech Talk AMCC #1</h1>
			<p>This is about Page</p>
		</body>
		</html>
		`
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprint(w, html)
	})

	fmt.Println("Server nya jalan di sini ya beb http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
