package model

import "errors"

// ErrorProductPriceTooLow is a error message for low prices
var ErrorProductPriceTooLow = errors.New("The price for the product is too low")

// ErrorProductPriceTooHigh is a error message for high prices
var ErrorProductPriceTooHigh = errors.New("The price for the product is too high")

var minimalPrice float64 = 5.00
var maximumPrice float64 = 30.00

// Product model represents a model of the a product
type Product struct {
	Sku   string `json:"sku"`
	Name  string `json:"name"`
	price float64
}

// SetPrice for a product
func (product *Product) SetPrice(price float64) (err error) {
	err = nil

	if price < minimalPrice {
		err = ErrorProductPriceTooLow
		return
	}
	if price > maximumPrice {
		err = ErrorProductPriceTooHigh
		return
	}

	product.price = price
	return
}

// GetPrice for a product
func (product *Product) GetPrice() float64 {
	return product.price
}
