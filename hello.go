package main

import (
    "fmt"
    "html"
    "net/http"
)

func main() {

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "why %q \n", html.EscapeString(r.URL.Path))
        fmt.Fprintf(w, "where %q \n", html.EscapeString(r.URL.Path))
        fmt.Fprintf(w, "what %q \n", html.EscapeString(r.URL.Path))
        fmt.Fprintf(w, "ok  fine %q \n", html.EscapeString(r.URL.Path))
    })
    
    http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request){
        fmt.Fprintf(w, "Hi")
    })

    http.ListenAndServe(":8000", nil)
 
}