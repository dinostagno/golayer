package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		parametro := "Hola desde el servidor Go"
		html := fmt.Sprintf(`
			<!DOCTYPE html>
			<html>
			<head>
			    <title>Hola Mundo</title>
			</head>
			<body>
			    <h1>%s</h1>
			</body>
			</html>
		`, parametro)
		fmt.Fprintln(w, html)
	})

	http.ListenAndServe(":8080", nil)
}
