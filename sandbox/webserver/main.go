package main

import (
    "log"

    "net/http"
)

func main() {
    http.HandleFunc("/", home)

    if err := http.ListenAndServe(":6000", nil); err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}

func home(w http.ResponseWriter, req *http.Request) {
    if req.URL.Path != "/" {
        log.Printf("404: %s", req.URL.String())
        http.NotFound(w, req)
    } else {
        w.Write([]byte("CHALLENGE"))
    }
}
