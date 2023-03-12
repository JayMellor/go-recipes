package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

type Endpoint struct {
	route    string
	handlers map[string]HandlerFunc
}

func getProducts(request *http.Request, _ map[string]string) (response string, status int) {
	productJson, error := json.Marshal(stock)
	if error != nil {
		return fmt.Sprint(error), http.StatusInternalServerError
	}
	return string(productJson), http.StatusOK
}

func requestBody(request *http.Request) (body map[string]any, error error) {
	jsonBody, error := io.ReadAll(request.Body)
	if error != nil {
		return body, fmt.Errorf("Error reading body: %w", error)
	}
	if error := json.Unmarshal(jsonBody, &body); error != nil {
		return body, fmt.Errorf("Error parsing request body: %w", error)
	}

	return body, nil
}

type Validator func(value any) error

type Field struct {
	validator Validator
	optional  bool
}

func missingFieldError(missingFields []string) error {
	message := ""
	for _, missingField := range missingFields {
		message += fmt.Sprintf("[%v]: field is required\n", missingField)
	}
	return fmt.Errorf("%s", message)
}

func validateBody(body map[string]any, fields map[string]Field) error {

	missingFields := []string{}
	invalidFields := map[string]error{}
	for fieldName, fieldProps := range fields {
		if value, fieldExists := body[fieldName]; fieldExists {
			if error := fieldProps.validator(value); error != nil {
				invalidFields[fieldName] = error
			}
		} else {
			missingFields = append(missingFields, fieldName)
		}
	}

	if len(missingFields) > 0 {
		return missingFieldError(missingFields)
	}
	if len(invalidFields) > 0 {
		return fmt.Errorf("Fields were invalid: %v", invalidFields)
	}

	return nil
}

func createProduct(request *http.Request, _ map[string]string) (response string, status int) {

	body, err := requestBody(request)
	if err != nil {
		return fmt.Sprint(err), http.StatusInternalServerError
	}

	var fields = map[string]Field{
		"Description": Field{
			validator: func(value any) error {
				return nil
			},
			optional: false,
		},
		"StockLevel": Field{
			validator: func(value any) error {
				if stockLevel, err := strconv.Atoi(fmt.Sprint(value)); err != nil {
					return fmt.Errorf("Stock Level must be an integer. Received %v", stockLevel)
				}
				return nil
			},
			optional: false,
		},
	}

	// todo combine validation and deserialization/parsing somehow?
	if err := validateBody(body, fields); err != nil {
		return fmt.Sprint(err), http.StatusBadRequest
	}
	stockLevel, _ := strconv.Atoi(fmt.Sprint(body["StockLevel"]))
	description := body["Description"]

	newProduct := Product{
		Id:          stock[len(stock)-1].Id + 1,
		Description: fmt.Sprint(description),
		StockLevel:  stockLevel,
	}

	stock = append(stock, newProduct)

	productJson, error := json.Marshal(stock)
	if error != nil {
		return fmt.Sprint(error), http.StatusInternalServerError
	}

	return string(productJson), http.StatusOK
}

func getProduct(request *http.Request, args map[string]string) (response string, status int) {
	id, error := strconv.Atoi(args["productId"])
	if error != nil {
		return fmt.Sprint(error), http.StatusBadRequest
	}

	match, found := findProduct(id)
	if !found {
		return fmt.Sprintln("No product with the following ID found:", id),
			http.StatusNotFound
	}

	productJson, error := json.Marshal(match)
	if error != nil {
		return fmt.Sprint(error), http.StatusInternalServerError
	}

	return string(productJson), http.StatusOK

}

func updateProduct(request *http.Request, args map[string]string) (response string, status int) {
	id, error := strconv.Atoi(args["productId"])
	if error != nil {
		return fmt.Sprint(error), http.StatusBadRequest
	}

	match, found := findProduct(id)
	if !found {
		return fmt.Sprintln("No product with the following ID found:", id),
			http.StatusNotFound
	}

	productJson, error := json.Marshal(match)
	if error != nil {
		return fmt.Sprint(error), http.StatusInternalServerError
	}

	return string(productJson), http.StatusOK
}

// // One possibility to control access to routes
// func PopulateRoutes(endpoints []Endpoint) {
// 	routes = endpoints
// }

// add \?.* to end of routes to allow query params
var Routes = []Endpoint{
	{
		route: `^/$`,
		handlers: map[string]HandlerFunc{
			http.MethodGet: func(request *http.Request, _ map[string]string) (response string, status int) {
				// return all endpoints
				return "hello!", http.StatusOK
			},
		},
	},
	{
		route: `^/products/?$`,
		handlers: map[string]HandlerFunc{
			http.MethodGet:  getProducts,
			http.MethodPost: createProduct,
		},
	},
	{
		route: `^/products/(?P<productId>\d+)/?$`,
		handlers: map[string]HandlerFunc{
			http.MethodGet: getProduct,
		},
	},
}
