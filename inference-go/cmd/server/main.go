package main

import (
    "log"
    "net/http"
)

func main() {
    http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte(`{"status":"ok"}`))
    })
    log.Println("inference-go listening :8081")
    log.Fatal(http.ListenAndServe(":8081", nil))
}
