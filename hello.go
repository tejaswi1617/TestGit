package main

import (
    "fmt"
    "html"
    "net/http"
)

func main() {

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "my %q \n", html.EscapeString(r.URL.Path))
        fmt.Fprintf(w, "your %q \n", html.EscapeString(r.URL.Path))
        fmt.Fprintf(w, "who's %q \n", html.EscapeString(r.URL.Path))
        fmt.Fprintf(w, "ok you %q \n", html.EscapeString(r.URL.Path))
    })
    
    http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request){
        fmt.Fprintf(w, "Hi")
    })

    http.ListenAndServe(":8000", nil)

}