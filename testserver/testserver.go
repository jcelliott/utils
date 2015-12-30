package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var separator = "--------------------------------------------------------------------------------"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		printRequest(r)
		defer r.Body.Close()
		_, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println("error reading body:", err)
			w.WriteHeader(http.StatusInternalServerError)
			io.WriteString(w, "error")
			return
		}
		io.WriteString(w, "ok")
	})

	log.Println("Starting server on port 8000\n")
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func printRequest(req *http.Request) {
	fmt.Println(separator)
	fmt.Println(time.Now().Format(time.RFC3339))
	fmt.Printf("Method:\t\t%v\n", req.Method)
	fmt.Printf("URL:\t\t%v\n", req.URL)
	fmt.Printf("Protocol:\t%v\n", req.Proto)
	// fmt.Printf("Headers:\n%# v\n", pretty.Formatter(req.Header))
	fmt.Printf("Headers: {\n")
	printHeaders(req.Header)
	fmt.Printf("}\n")
	fmt.Printf("Content Length:\t%v\n", req.ContentLength)
	fmt.Printf("Remote Address:\t%v\n", req.RemoteAddr)
	fmt.Println()
}

func printHeaders(h http.Header) {
	for k, v := range h {
		fmt.Printf("\t%v: ", k)
		for _, s := range v {
			fmt.Printf("%v, ", s)
		}
		fmt.Println()
	}
}
