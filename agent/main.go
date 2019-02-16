package main

import (
	"log"
	"net/http"
	"net/http/cookiejar"

	"github.com/efkbook/blog-sample/agent/scenario"
)

func init() {
	jar, _ := cookiejar.New(nil)
	http.DefaultClient.Jar = jar
}

func main() {
	errCh := make(chan error)

	go func(ch chan<- error) {
		ch <- scenario.Shakespeare()
	}(errCh)
	go func(ch chan<- error) {
		ch <- scenario.Crawler()
	}(errCh)

	err := <-errCh
	if err != nil {
		log.Printf("error: %v", err)
	}
}
