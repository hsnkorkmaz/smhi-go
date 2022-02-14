package main

import (
	"github.com/hsnkorkmaz/smhi-go/services"
	"fmt"
)

func main() {
	response, err := services.GetVersionResult("1.0")
	
	if err != nil {
		fmt.Println(err)
	}

	for _, v := range response.Resources {
		fmt.Println(v.Title)
	}
}
