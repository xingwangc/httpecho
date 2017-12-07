package main

import (
	"flag"
	"fmt"
	"html"
	"log"
	"net/http"
	"strconv"
)

func main() {
	port := 8080
	flag.Parse()
	if len(flag.Args()) > 0 {
		p, err := strconv.Atoi(flag.Arg(0))
		if err == nil && 1024 <= p && p < 49151 {
			port = p
		}

	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q\n", html.EscapeString(r.URL.Path))
	})

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
