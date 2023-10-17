package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		currentTime := time.Now()

		err := godotenv.Load()
		if os.Getenv("ENVIRONMENT") != "production" && err != nil {
			log.Fatal("Error loading .env file")
		}

		body, err := io.ReadAll(r.Body)

		if err != nil {
			log.Fatalln(err)
		}

		hideableHeader := os.Getenv("HIDE_HEADER")
		if hideableHeader != "" {
			r.Header.Set(hideableHeader, "OBFUSCATED")
		}

		var jsonBody map[string]interface{}
		json.Unmarshal([]byte(body), &jsonBody)

		data := map[string]interface{}{
			"app_url":     os.Getenv("APP_URL"),
			"method":      r.Method,
			"headers":     r.Header,
			"body":        jsonBody,
			"request_uri": r.RequestURI,
			"_timestamp":  currentTime.UTC().Format("2006-01-02T15:04:05-0700"),
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
		if os.Getenv("AUTH") == "basic" {
			req.SetBasicAuth(os.Getenv("AUTH_USERNAME"), os.Getenv("AUTH_PASSWORD"))
		}

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
