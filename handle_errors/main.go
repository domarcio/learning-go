package main

import (
	"encoding/json"
	"fmt"

	"github.com/domarcio/handle_errors/model"
)

func main() {
	iphone := model.Product{}
	iphone.Sku = "4rt6a546f1e8dg6d1df"
	iphone.Name = "iPhone XS 240GB"

	fmt.Printf("iPhone informations: %+v\n", iphone)

	iphoneJSON, _ := json.Marshal(iphone)
	fmt.Printf("iPhone JSON: %+v\n", string(iphoneJSON))
}
