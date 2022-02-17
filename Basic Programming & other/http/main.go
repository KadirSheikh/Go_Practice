package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {

	resp, err := http.Get("http://google.com")

	if err != nil {
		os.Exit(1)
	}

	// bs := make([]byte, 99999)

	// resp.Body.Read(bs)

	// fmt.Println(string(bs))

	io.Copy(os.Stdout, resp.Body)

}

func (logWriter) Write(b []byte) (int, error) {

	fmt.Println(string(b))

	fmt.Println("BYTES", len(b))

	return len(b), nil

}
