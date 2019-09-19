package main

import "fmt"

type computer struct {
	name      string
	createdBy string
	price     float64
}

func main() {
	laptop := new(computer)
	fmt.Printf("Computer located in memory %p - %+v\n", &laptop, laptop)

	tablet := computer{"Apple 4", "Apple", 1000.00}
	fmt.Printf("Table located in memory %p - %+v\n", &tablet, tablet)

	changeComputerInformations(&tablet)
	fmt.Printf("My new table informations are %+v located in memory %p\n", tablet, &tablet)
}

func changeComputerInformations(informations *computer) {
	informations.name = "Apple Inc"
	informations.price = 2000.00
}
