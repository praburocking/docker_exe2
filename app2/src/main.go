package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8088", nil)
}
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from " + r.RemoteAddr+" to " +r.Host )
}
