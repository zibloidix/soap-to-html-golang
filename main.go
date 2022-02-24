package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

// Для примера разбора wsdl и xsd
// https://github.com/fiorix/wsdl2go/blob/master/wsdl/types.go

func main() {
	xsdFile := "./wsdl-xsd/hcs-rvuo-types.xsd"
	xsdData := getFileData(xsdFile)
	xsdSchema := getSchema(xsdData)

	wsdlFile := "./wsdl-xsd/hcs-rvuo-service-async.wsdl"
	wsdlData := getFileData(wsdlFile)
	wsdlDef := getDefinition(wsdlData)
	fmt.Println(xsdSchema, wsdlDef)
}

func getFileData(fileName string) []byte {
	data, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	return data
}

func getSchema(data []byte) Schema {
	var schema Schema
	err := xml.Unmarshal(data, &schema)
	if err != nil {
		panic(err)
	}
	return schema
}

func getDefinition(data []byte) Definitions {
	var def Definitions
	err := xml.Unmarshal(data, &def)
	if err != nil {
		panic(err)
	}
	return def
}
