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

func createProduct(request *http.Request, _ map[string]string) (response string, status int) {

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
		Id:          stock[len(stock)-1].Id + 1,
		Description: fmt.Sprint(description),
		StockLevel:  stkLvl,
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

// add \?.* to end of routes to allow query params
var Routes = []Endpoint{
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
			http.MethodGet:  getProducts,
			http.MethodPost: createProduct,
		},
	},
	{
		route: `^/stock/(?P<productId>\d+)/?$`,
		handlers: map[string]HandlerFunc{
			http.MethodGet: getProduct,
		},
	},
}
