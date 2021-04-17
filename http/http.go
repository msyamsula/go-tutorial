package http

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func HttpRequest() {
	uri := "https://google.com"

	resp, err := http.Get(uri)
	if err != nil {
		log.Println("something's wrong")
		log.Fatal(err)
	}

	// bs := make([]byte, 9999)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))
	io.Copy(os.Stdout, resp.Body)
	fmt.Println(resp.Status)
	fmt.Println(resp.StatusCode)
}
