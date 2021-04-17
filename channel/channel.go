package channel

import (
	"log"
	"net/http"
	"time"
)

// global var
var arr = []string{
	"https://google.com",
	"https://golang.org",
	"https://facebook.com",
	"https://stackoverflow.com",
	"https://amazon.com",
}

// checkUrl primitive
func checkUrl(url string) {
	_, err := http.Get(url)
	if err != nil {
		log.Printf("%v currently error with error: %v\n", url, err)
		return
	}
	log.Printf("%v is up\n", url)
}

// checkUrl with channel to leverage go routine
// send report through channel
func checkUrlwithChannel(url string, c chan string) {
	_, err := http.Get(url)
	if err != nil {
		log.Printf("%v currently error with error: %v\n", url, err)
		c <- url
		return
	}
	c <- url
	log.Printf("%v is up\n", url)
}

// check url with http get, step by step
func StatusCheck() {
	for _, link := range arr {
		checkUrl(link)
	}
}

// check concurenly, with go routine and channel
func StatusCheckConcurent() {
	c := make(chan string)
	for _, link := range arr {
		go checkUrlwithChannel(link, c)
	}

	// wait from channel from each go routine
	// for i := 0; i < len(arr); i++ {
	// 	fmt.Println(<-c)
	// }

	for {
		go checkUrlwithChannel(<-c, c)
		time.Sleep(1 * time.Second)
	}
}
