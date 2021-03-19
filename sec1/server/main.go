package main

import (
	"fmt"
	"io"
	"net/http"

	log "github.com/sirupsen/logrus"
)
const (
	localhost = ":1000"
)
func main(){
	fmt.Printf("Server listening on %v\n", localhost)
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(localhost, nil))
}
func handler (w http.ResponseWriter, r *http.Request){
		const respbo = `<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<meta http-equi v="X-UA-Compatible" content="IE=edge">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title>Document</title>
		</head>
		<body>
			you are listening on port 1000
		</body>
		</html>`
		io.WriteString(w, respbo)
}