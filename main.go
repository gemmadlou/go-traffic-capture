package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := godotenv.Load()
		if os.Getenv("ENVIRONMENT") != "production" && err != nil {
			log.Fatal("Error loading .env file")
		}

		body, err := io.ReadAll(r.Body)

		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println(os.Getenv("HIDE_HEADER"))
		hideableHeader := os.Getenv("HIDE_HEADER")
		if hideableHeader != "" {
			r.Header.Set(hideableHeader, "OBFUSCATED")
		}

		var x map[string]interface{}

		json.Unmarshal([]byte(body), &x)

		data := map[string]interface{}{
			"method":      r.Method,
			"headers":     r.Header,
			"body":        x,
			"request_uri": r.RequestURI,
		}

		jsonData, err := json.Marshal(data)
		if err != nil {
			log.Fatal("Error:", err)
			return
		}

		send := []byte(jsonData)
		req, err := http.NewRequest(http.MethodPost, os.Getenv("ELASTICSEARCH_URL")+"/traffic/_doc", bytes.NewReader(send))
		if err != nil {
			log.Fatal(err)
		}

		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}

		resp, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		}

		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		sb := string(respBody)
		fmt.Println(sb)

		fmt.Println(string(jsonData))
		fmt.Fprintf(w, "Request logged")
	})

	fmt.Println("Listening on http://localhost:9990")
	http.ListenAndServe(":9990", nil)
}
