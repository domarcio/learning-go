package main

import "fmt"

// Property informations
type Property struct {
	name             string
	location         string
	price            float64
	constructionYear int
}

func main() {
	var house Property
	fmt.Printf("Basic house informations: %+v\n", house)

	house.name = "A beautiful house"
	house.location = "São Paulo"
	house.price = 500000.00
	house.constructionYear = 2009
	fmt.Printf("House informations: %+v\n", house)

	apartament := Property{
		name: "Downtown apartment",
	}
	fmt.Printf("Apartament basic informations: %+v\n", apartament)
}
