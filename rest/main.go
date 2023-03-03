package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

const port string = ":8080"

type Product struct {
	id          int
	description string
	stockLevel  int
}

var stock = []Product{
	{
		id:          1,
		description: "Jacket",
		stockLevel:  0,
	},
	{
		id:          2,
		description: "Shoes",
		stockLevel:  0,
	},
}

type HandlerFunc func(request *http.Request, args map[string]string) (response string, status int)

type Endpoint struct {
	route    string
	handlers map[string]HandlerFunc
}

var routes = []Endpoint{
	{
		route: `^/$`,
		handlers: map[string]HandlerFunc{
			http.MethodGet: func(request *http.Request, _ map[string]string) (response string, status int) {
				return "hello!", http.StatusOK
			},
		},
	},
	{
		route: `^/stock/?$`,
		handlers: map[string]HandlerFunc{
			http.MethodGet: func(request *http.Request, _ map[string]string) (response string, status int) {
				return fmt.Sprint(stock), http.StatusOK
			},
		},
	},
	{
		route: `^/stock/(?P<productId>\d+)/?$`,
		handlers: map[string]HandlerFunc{
			http.MethodGet: func(request *http.Request, args map[string]string) (response string, status int) {
				id, error := strconv.Atoi(args["productId"])
				if error != nil {
					return fmt.Sprint(error), http.StatusBadRequest
				}

				for _, product := range stock {
					if product.id == id {
						return fmt.Sprint(product), http.StatusOK
					}
				}
				return fmt.Sprintln("No product with the following ID found:", id), http.StatusNotFound
			},
		},
	},
}

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
	for _, endpoint := range routes {
		routeExpr, error := regexp.Compile(endpoint.route)
		if error != nil {
			return nil, nil, error
		}

		if namedMatches, found := findNamedStringSubmatch(routeExpr, URL.String()); found {
			return &endpoint, namedMatches, nil
		}

	}
	return nil, nil, fmt.Errorf("No endpoints matching %v were found", URL)
}

func match_endpoint(writer http.ResponseWriter, request *http.Request) {

	endpoint, matches, error := find_endpoint(request.URL, request.Method)
	if error != nil {
		writer.WriteHeader(http.StatusNotFound)
		writer.Write([]byte(fmt.Sprint(error)))
		return
	}

	if handler := endpoint.handlers[request.Method]; handler == nil {
		writer.WriteHeader(http.StatusMethodNotAllowed)
		writer.Write([]byte(fmt.Sprintf("Method %v is not allowed for endpoint %v", request.Method, request.URL)))
	} else {
		response, status := handler(request, matches)
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
	fmt.Println("Listening on port", port)
	if error := http.ListenAndServe(port, nil); error != nil {
		log.Fatal(error)
	}
}
