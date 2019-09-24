package model

// Property informations
type Property struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Zip     string `json:"zip"`
	price   float64
}

// SetPrice for a property
func (property *Property) SetPrice(price float64) {
	property.price = price
}

// GetPrice property
func (property *Property) GetPrice() float64 {
	return property.price
}
