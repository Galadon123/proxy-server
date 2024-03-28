package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello from Backend Server 2")
    })

    fmt.Println("Backend Server 2 is listening on port 5002...")
    http.ListenAndServe(":5002", nil)
}
