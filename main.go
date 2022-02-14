package main

import (
	"fmt"

	"github.com/hsnkorkmaz/smhi-go/services"
)

func testFunc() {
	response, err := services.GetVersionResult("1.0")

	if err != nil {
		fmt.Println(err)
	}

	for _, v := range response.Resources {
		fmt.Println(v.Title)
	}
}
