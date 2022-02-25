package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type Description struct {
	Schema     Schema
	Definition Definitions
}

var Descriptions []Description

func main() {
	xsdFile := "./wsdl-xsd/hcs-appeals-types.xsd"
	xsdData := getFileData(xsdFile)
	xsdSchema := getSchema(xsdData)

	wsdlFile := "./wsdl-xsd/hcs-appeals-service-async.wsdl"
	wsdlData := getFileData(wsdlFile)
	wsdlDef := getDefinition(wsdlData)

	Descriptions = append(Descriptions, Description{xsdSchema, wsdlDef})

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
