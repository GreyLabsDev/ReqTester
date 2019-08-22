package main

import (
	// "io/ioutil"
	"fmt"
	"log"
	"net/http"
	"strings"
)

var mainControlChannel chan int
var lastReqParams = ""
var lastTokenData = ""

func webServer(controlChannel chan int) {
	mainControlChannel = controlChannel

	http.HandleFunc("/last_req_params", lastPostParamsHandler)
	http.HandleFunc("/save_req_params", recordPostParamsHandler)
	http.HandleFunc("/api/v1/user/token", getTokenHandler)
	http.HandleFunc("/last_token_params", lastTokenParamsHandler)
	http.HandleFunc("/halthalthalt", terminateAppHandler)
	
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func getTokenHandler(w http.ResponseWriter, r *http.Request) {
	lastTokenData = extractRequestParams(r)
	// body, err := r.GetBody()
	// if err == nil {
	// 	bodyString, readError := ioutil.ReadAll(body)
	// 	if readError == nil {
	// 		lastTokenData += "\nBody:\n" + string(bodyString)
	// 	}
	// }
	fmt.Fprintln(w, "\nOk")
}

func recordPostParamsHandler(w http.ResponseWriter, r *http.Request) {
	lastReqParams = extractRequestParams(r)
	fmt.Fprintln(w, "\nOk")
}

func lastPostParamsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "\nLast request params ")
	fmt.Fprintln(w, "\n"+lastReqParams)
}

func lastTokenParamsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "\nLast request params ")
	fmt.Fprintln(w, "\n"+lastTokenData)
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
