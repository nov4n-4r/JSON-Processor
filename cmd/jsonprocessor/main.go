package main

import (
	"fmt"
	"jsonprocessor/internal/util"
)

func main() {

	data := []byte(`{
		"name": "Deson",
		"age": 39,
		"city": "Khaby Lake"
	}`)

	fields := []string{"name"}

	util.JsonMap(&data, fields)

	fmt.Printf("Hasil akhir: %s\n", string(data))

	filters := []util.JsonStructure{
		{Key: "name", Value: "Deson"},
		{Key: "age", Value: int(39)},
	}

	isValid, e := util.JsonCheck(&data, filters)

	if e != nil {
		fmt.Printf("Error: %s\n", e.Error())
	}

	fmt.Printf("Hasil checking: %t\n", isValid)

}
