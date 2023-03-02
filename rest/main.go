package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"regexp"
)

const port string = ":8080"

var stock = map[string]int{
	"Trousers": 0,
	"Jackets":  0,
}

type Endpoint struct {
	route    string
	handlers map[string]func(request *http.Request) (response string, status int)
}

var routes = []Endpoint{{
	route: `^/$`,
	handlers: map[string]func(request *http.Request) (response string, status int){
		http.MethodGet: func(request *http.Request) (response string, status int) {
			return "hello!", http.StatusOK
		},
	},
}}

func find_endpoint(URL *url.URL, method string) (endpoint *Endpoint, error error) {
	for _, endpoint := range routes {
		matched, error := regexp.Match(endpoint.route, []byte(fmt.Sprint(URL)))
		if error != nil {
			return nil, error
		}

		if matched {
			return &endpoint, nil
		}

	}
	return nil, fmt.Errorf("No endpoints matching %v were found", URL)
}

func match_endpoint(writer http.ResponseWriter, request *http.Request) {

	endpoint, error := find_endpoint(request.URL, request.Method)
	if error != nil {
		writer.WriteHeader(http.StatusNotFound)
		writer.Write([]byte(fmt.Sprint(error)))
		return
	}

	if handler := endpoint.handlers[request.Method]; handler == nil {
		writer.WriteHeader(http.StatusMethodNotAllowed)
		writer.Write([]byte(fmt.Sprintf("Method %v is not allowed for endpoint %v", request.Method, request.URL)))
	} else {
		response, status := handler(request)
		if status != http.StatusOK {
			writer.WriteHeader(status)
		}
		writer.Write([]byte(response))
	}

	// check status and message in writer
	// log.Println(request.Method, status, request.URL)
}

func main() {

	http.HandleFunc("/", match_endpoint)

	// http.HandleFunc("/stock", func(writer http.ResponseWriter, request *http.Request) {

	// 	switch method := request.Method; method {
	// 	case http.MethodGet:
	// 		for product, stockLevel := range stock {
	// 			writer.Write([]byte(fmt.Sprint(product, ":", stockLevel, "<br/>")))
	// 		}

	// 	case http.MethodPost:

	// 	default:
	// writer.WriteHeader(http.StatusMethodNotAllowed)
	// writer.Write([]byte(fmt.Sprintf("Method %v is not allowed for endpoint %v", method, request.URL)))
	// 	}
	// })

	// http.HandleFunc("/stock/:id",
	// 	func(writer http.ResponseWriter, request *http.Request) {
	// 		writer.Write([]byte(fmt.Sprint("Url is", request.URL)))
	// 	})

	fmt.Println("Listening on port", port)
	if error := http.ListenAndServe(port, nil); error != nil {
		log.Fatal(error)
	}
}
