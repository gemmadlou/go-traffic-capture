package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)

		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println(r.Method)
		fmt.Println(string(body))
		fmt.Println(r.Header)
		fmt.Println(r.RequestURI)

		fmt.Fprintf(w, "Hello World")
	})

	fmt.Println("Listening on http://localhost:9990")
	http.ListenAndServe(":9990", nil)
}
