package main

import (
	"fmt"

	"github.com/domarcio/advanced_structs/model"
)

/*
When the cache doesn't instanceid your default value is null
*/
var propertyCache map[string]model.Property

func main() {
	propertyCache = make(map[string]model.Property, 0)

	house := model.Property{}
	house.Name = "Linda casa na Mooca"
	house.Address = "Rua Padre Raposo, 545"
	house.Zip = "03118-010"
	house.SetPrice(350.135)

	propertyCache["moocaHouse"] = house

	fmt.Println("Is there a house in Mooca?")
	moocaHouse, foundHouse := propertyCache["moocaHouse"]
	if foundHouse {
		fmt.Printf("Yes, there is a house in Mooca. Eu found it: %+v\n", moocaHouse)
	}

	apartament := model.Property{}
	apartament.Name = "Lindo apartamento na Mooca"
	apartament.Address = "Rua Padre Raposo, 172"
	apartament.Zip = "03118-011"
	apartament.SetPrice(250.000)

	propertyCache[apartament.Name] = apartament
	fmt.Println("How many items are in the cache?", len(propertyCache))

	for key, property := range propertyCache {
		fmt.Printf("cache[%s] = %+v\n", key, property)
	}

	if foundHouse {
		delete(propertyCache, "moocaHouse")
	}

	fmt.Printf("If mooca house is found, so deleting from the cache. Now, our cache have %d item\n", len(propertyCache))
}
