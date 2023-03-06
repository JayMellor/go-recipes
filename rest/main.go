package main

import (
	"encoding/json"
	"fmt"
	"io"
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

// add \?.* to end of routes to allow query params
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
			http.MethodPost: func(request *http.Request, _ map[string]string) (response string, status int) {

				body, error := io.ReadAll(request.Body)
				if error != nil {
					return fmt.Sprintln("Error reading body:", error), http.StatusInternalServerError
				}
				bodyRequest := map[string]any{}
				if error := json.Unmarshal(body, &bodyRequest); error != nil {
					return fmt.Sprintln("Error parsing request", error), http.StatusBadRequest
				}

				description, descriptionExists := bodyRequest["Description"]
				stockLevel, stockLevelExists := bodyRequest["StockLevel"]
				missingFields := []string{}
				if !descriptionExists {
					missingFields = append(missingFields, "Description")
				}
				if !stockLevelExists {
					missingFields = append(missingFields, "Stock Level")
				}

				if !descriptionExists || !stockLevelExists {
					return fmt.Sprintln("The following fields are required:", missingFields), http.StatusBadRequest
				}

				stkLvl, error := strconv.Atoi(fmt.Sprint(stockLevel))
				if error != nil {
					return fmt.Sprintln("Stock Level must be an integer. Received", stockLevel), http.StatusBadRequest
				}

				newProduct := Product{
					id:          stock[len(stock)-1].id + 1,
					description: fmt.Sprint(description),
					stockLevel:  stkLvl,
				}

				stock = append(stock, newProduct)

				return fmt.Sprintf("%+v", newProduct), http.StatusOK
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

	// endpoint, matches, error := find_endpoint(request.URL, request.Method)
	// if error != nil {
	// 	writer.WriteHeader(http.StatusNotFound)
	// 	writer.Write([]byte(fmt.Sprint(error)))
	// 	return
	// }

	// handler, handlerFound := endpoint.handlers[request.Method]
	// if !handlerFound {
	// 	handler = defaultHandler
	// }

	handler, matches := getHandler(request)
	response, status := handler(request, matches)
	if status != http.StatusOK {
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
