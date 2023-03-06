package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"regexp"
)

const port string = ":8080"

type HandlerFunc func(request *http.Request, args map[string]string) (response string, status int)

func findNamedStringSubmatch(regex *regexp.Regexp, url string) (namedMatches map[string]string, found bool) {
	matches := regex.FindStringSubmatch(url)
	if len(matches) == 0 {
		return namedMatches, false
	}

	namedMatches = make(map[string]string)
	for place, name := range regex.SubexpNames() {
		namedMatches[name] = matches[place]
	}
	return namedMatches, true
}

func find_endpoint(URL *url.URL, method string) (endpoint *Endpoint, matches map[string]string, error error) {
	for _, endpoint := range Routes {
		routeExpr, error := regexp.Compile(endpoint.route)
		if error != nil {
			log.Fatal("Error compiling endpoint", endpoint, error)
		}

		if namedMatches, found := findNamedStringSubmatch(routeExpr, URL.String()); found {
			return &endpoint, namedMatches, nil
		}

	}
	return nil, nil, fmt.Errorf("No endpoints matching %v were found", URL)
}

func notFoundHandler(request *http.Request, _ map[string]string) (response string, status int) {
	return fmt.Sprintf("No endpoint matching URL %v found", request.URL), http.StatusNotFound
}

func methodNotAllowedHandler(request *http.Request, _ map[string]string) (response string, status int) {
	return fmt.Sprintf("Method %v is not allowed for endpoint %v", request.Method, request.URL), http.StatusMethodNotAllowed
}

func getHandler(request *http.Request) (HandlerFunc, map[string]string) {
	endpoint, matches, error := find_endpoint(request.URL, request.Method)
	if error != nil {
		return notFoundHandler, matches
	}

	handler, handlerFound := endpoint.handlers[request.Method]
	if !handlerFound {
		return methodNotAllowedHandler, matches
	}

	return handler, matches

}

func match_endpoint(writer http.ResponseWriter, request *http.Request) {

	handler, matches := getHandler(request)
	response, status := handler(request, matches)
	if status == http.StatusOK {
		// todo find a nicer way of doing
		writer.Header().Set("Content-Type", "application/json")
	} else {
		writer.WriteHeader(status)
	}
	writer.Write([]byte(response))

	log.Println(request.Method, request.URL, status)
	if status >= http.StatusBadRequest {
		log.Print(response)
	}
}

func main() {
	http.HandleFunc("/", match_endpoint)
	fmt.Println("Listening on port", port)
	if error := http.ListenAndServe(port, nil); error != nil {
		log.Fatal(error)
	}
}
