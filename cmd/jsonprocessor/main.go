package main

import (
	"fmt"
	"jsonprocessor/internal/util"
)

func main() {
	isValid, e := util.JsonValidation(`{
		"name": "John Doe",
		"age": 30
	}`)

	if e != nil {
		fmt.Println("Error validating JSON:", e)
		return
	}

	fmt.Printf("Is the JSON valid? %v\n", isValid)
}
