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
	errorSetPrice := iphone.SetPrice(20.99)
	if errorSetPrice != nil {
		fmt.Printf("[main] There was an error pricing product: %s\n", errorSetPrice.Error())
		return
	}

	fmt.Printf("iPhone informations: %+v\n", iphone)

	iphoneJSON, errorJSONMarshal := json.Marshal(iphone)
	if errorJSONMarshal != nil {
		fmt.Printf("[main] An error occurred on JSON Marshal: %s\n", errorJSONMarshal.Error())
		return
	}
	fmt.Printf("iPhone JSON: %+v\n", string(iphoneJSON))
	fmt.Printf("My new price is %g\n", iphone.GetPrice())
}
