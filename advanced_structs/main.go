package main

import (
	"encoding/json"
	"fmt"

	"github.com/domarcio/advanced_structs/model"
)

func main() {
	house := model.Property{}
	house.Name = "Lindo apartamento na Mooca"
	house.Address = "Rua Padre Raposo, 545"
	house.Zip = "03118-010"
	house.SetPrice(350.135)

	fmt.Printf("About house: %+v\n", house)
	fmt.Printf("Price: %v\n", house.GetPrice())

	houseJSON, _ := json.Marshal(house)

	fmt.Println(string(houseJSON))
}
