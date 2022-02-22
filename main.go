package main

import (
	"fmt"
	"os"
)

func main() {
	fileName := "./wsdl-xsd/hcs-rvuo-types.xsd"
	data, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}
