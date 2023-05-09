package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type People struct {
	FirstName string
	LastName  string
	Details   struct {
		Height int
		Weight float32
	}
}

type Name struct {
	FirstName string
	LastName  string
}

type Address struct {
	Line1 string
	Line2 string
	Line3 string
}

type Customer struct {
	Name    Name
	Email   string
	Address Address
	DOB     time.Time
}

type AnotherCustomer map[string]interface{}
type AnotherName map[string]interface{}
type AnotherAddress map[string]interface{}

type Rates struct {
	Base        string `json:"base currency"`
	Destination string `json:"destination currency"`
}

func main() {
	var persons []People
	var rates Rates
	var result map[string]interface{}

	layoutISO := "2006-01-01"
	dob, _ := time.Parse(layoutISO, "2010-01-18")
	johnPhiri := AnotherCustomer{
		"Name": AnotherName{
			"FirstName": "John",
			"LastName":  "Phiri",
		},
		"Email": "johnphiri@example.com",
		"Address": AnotherAddress{
			"Line1": "The White House",
			"Line2": "1600 Pennsylvania Avenue NW",
			"Line3": "Washington, DC 20500",
		},
		"DOB": dob,
	}

	john := Customer{
		Name: Name{
			FirstName: "John",
			LastName:  "Smith",
		},
		Email: "johnsmith@example.com",
		Address: Address{
			Line1: "The White House",
			Line2: "1600 Pennysylvania Avenue NW",
			Line3: "Washington, DC 20500",
		},
		DOB: dob,
	}

	ratesJsonString := `{
								"base currency": "EUR",
								"destination currency": "USD"
							}`

	jsonString := `[
						{
							"firstname": "Gregory",
							"lastname": "Phiri",
							"details": {
								"height": 175,
								"weight": 70.0
							}
						},
						{
							"firstname": "Mickey",						
							"lastname": "Mouse",
							"details": {
								"height": 105,
								"weight": 85.5
							}
						}

				   ]`

	currencyJsonString :=
		`{
			"success": true,
			"timestamp": 1588779306,
			"base": "EUR",
			"date": "2020-05-06",
			"rates": {
				"AUD": 1.683349,
				"CAD": 1.528643,
				"SGD": 1.534513,
				"USD": 1.080054
			}
		}`

	err := json.Unmarshal([]byte(jsonString), &persons)
	if err == nil {
		for _, person := range persons {
			fmt.Println(person.FirstName, person.LastName, "weighs", person.Details.Weight, "and he is", person.Details.Height, "cm tall.")
		}
	} else {
		fmt.Println(err)
	}

	ratesErr := json.Unmarshal([]byte(ratesJsonString), &rates)
	if ratesErr == nil {
		fmt.Println("Exchanging the currency from", rates.Base, "to", rates.Destination)
	} else {
		fmt.Println(ratesErr)
	}

	currencyErr := json.Unmarshal([]byte(currencyJsonString), &result)
	if currencyErr == nil {
		rates := result["rates"]
		currencies := rates.(map[string]interface{})
		SGD := currencies["SGD"]
		fmt.Println(result["success"])
		fmt.Println(rates)
		fmt.Println(currencies)
		fmt.Println(SGD)
	} else {
		fmt.Println(currencyErr)
	}

	johnJSON, encodeErr := json.MarshalIndent(john, "", "	")
	if encodeErr == nil {
		fmt.Println(string(johnJSON))
	} else {
		fmt.Println(encodeErr)
	}
	johnPhirJSON, encodeJohnPhiriErr := json.MarshalIndent(johnPhiri, "", "	")
	if encodeJohnPhiriErr == nil {
		fmt.Println(string(johnPhirJSON))
	} else {
		fmt.Println(encodeJohnPhiriErr)
	}
}
