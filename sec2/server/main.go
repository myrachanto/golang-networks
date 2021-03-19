package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	log "github.com/sirupsen/logrus"
)
const (
	localhost = ":1000"
)
var (
	Rootrequest int
)
func main(){
	fmt.Printf("Server listening on %v\n", localhost)
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(localhost, nil))
}
func handler (w http.ResponseWriter, r *http.Request){
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("No data from the client\n")
		fmt.Fprintf(w, "no data provided \n")
		return
	}
	r.Body.Close()
	fmt.Printf("client sent: %v \n", string(body))
	caps := strings.Title(string(body))

	fmt.Fprintf(w, "this is remote address %v \n", string(body))
	fmt.Fprint(w, "<hr />")
	fmt.Fprintf(w, "this is the method called %v \n", caps)
}
// func handler (w http.ResponseWriter, r *http.Request){
// 	remotehost := r.RemoteAddr
// 	remoteURl := r.RequestURI
// 	method := r.Method
// 	fmt.Fprintf(w, "this is remote address %v \n", remotehost)
// 	fmt.Fprintf(w, "this is the uri %v \n", remoteURl)
// 	fmt.Fprintf(w, "this is the method called %v \n", method)
// }