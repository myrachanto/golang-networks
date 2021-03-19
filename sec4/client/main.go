package main

import (
	"io"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
)
const (
	localhost = "http://localhost:1000"
	host = "https://google.com"
)
func main(){
	resp, err := http.Get(host)
	if err != nil {
		log.Fatalf("cant connect to @v: %v", host, err)
		log.Fatalf("Trying %v ..", localhost)
		resp, err = http.Get(localhost)
	}
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)
}