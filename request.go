package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func request() int {
	resp, err := http.Get("https://google.com")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	fmt.Println(resp.StatusCode)

	file, err := os.Create("google.html")
	if err != nil {
		log.Fatal("Cannot create a file")
	}

	_, err = file.Write(body)
	if err != nil {
		log.Fatal("The data cannot be written to the file")
	}

	return resp.StatusCode
}
