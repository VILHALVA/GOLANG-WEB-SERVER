package main

import (
	"fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "public"+r.URL.Path)
}

func main() {
    http.HandleFunc("/", handler)

	fmt.Println("Servidor web iniciado. Acesse http://localhost:8080/")
    http.ListenAndServe(":8080", nil)
}
