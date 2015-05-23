/* This is the file rest-server 
 */
 package main

import (
	"io"
	"log"
	"net/http"
)

/* Get a JSON
 */
func get(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	io.WriteString(res,
		`{
    "infrastructure": "digitalocean.com",
    "platform": "docker",
    "language": "go",
    "result": "bazinga!"
    }`,
	)
}

func main() {
	http.HandleFunc("/", get)                // declare handler for path
	err := http.ListenAndServe(":8080", nil) // start the http listener
	if err != nil {
		log.Fatal(err)
	}
}
