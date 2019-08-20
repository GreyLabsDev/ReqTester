package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

var mainControlChannel chan int
var lastReqParams = ""

func webServer(controlChannel chan int) {
	mainControlChannel = controlChannel

	http.HandleFunc("/last_req_params", lastPostParamsHandler)
	http.HandleFunc("/save_req_params", recordPostParamsHandler)
	http.HandleFunc("/halthalthalt", terminateAppHandler)

	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func recordPostParamsHandler(w http.ResponseWriter, r *http.Request) {
	lastReqParams = extractRequestParams(r)
	fmt.Fprintln(w, "\nOk")
}

func lastPostParamsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "\nLast request params ")
	fmt.Fprintln(w, "\n"+lastReqParams)
}

func extractRequestParams(r *http.Request) string {
	var requestData = ""
	requestData += "\nMethod " + r.Method
	requestData += "\nUserAgent " + r.UserAgent()
	requestData += "\n"
	for k, v := range r.Header {
		requestData += "\nHeader field " + k + ", Value " + strings.Join(v[:], ",")
	}
	return requestData
}

func terminateAppHandler(w http.ResponseWriter, r *http.Request) {
	mainControlChannel <- 0
}
