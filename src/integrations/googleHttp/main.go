package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct {}

func main() {

	resp, error := http.Get("https://www.google.com")

	if error != nil {
		fmt.Println("Error: ", error)
		os.Exit(1)
	}

	io.Copy(os.Stdout, resp.Body)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	return len(bs), nil
}