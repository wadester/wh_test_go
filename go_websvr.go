/*
 * Module:   go_websvr.go
 * Purpose:  go simple web server
 * Date:     ? (examples), N/A WWH
 * Notes:
 * 1) To build:
 *      go build go_websvr.go
 * 2) Ref:  GO tour
 *
*/

package main

import (
	"fmt"
	"log"
	"net/http"
)

// define a struc and a Hello method on the struct 
type Hello struct{}

func (h Hello) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request) {
	fmt.Fprint(w, "<HTML><H1>Hello</H1><BODY>Test 1, 2, 3!</BODY></HTML>")
}

// main function, define host and port
func main() {
	port := 8082
	host := "localhost"
	hostinfo := fmt.Sprintf("%s:%d", host, port)
	fmt.Printf("go_websvr listening on http://%s\n", hostinfo)

	// define struct and serve with callback to that struct's
	// ServeHTTP method defined above
	var h Hello
	err := http.ListenAndServe(hostinfo, h)
	if err != nil {
		log.Fatal(err)
	}
}
