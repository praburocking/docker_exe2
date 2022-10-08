package main

import (
	"fmt"
	"net/http"
	"io"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
func handler(w http.ResponseWriter, r *http.Request) {
	
	requestURL := "http://exe2-app2-1:8088"
	//requestURL ="https://jsonplaceholder.typicode.com/posts/1"
	resp, _ := http.Get(requestURL)
	//defer resp.Body.Close()

	resBody, _ := io.ReadAll(resp.Body)
	//resBody, _ := httputil.DumpResponse(res, true)
	fmt.Fprintf(w, "Hello from " + r.RemoteAddr+" to " +r.Host +string(resBody) )
}