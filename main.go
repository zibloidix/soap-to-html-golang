package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

const BASE_PATH = "/home/aleksey/Coding/SOAP/li"

type Description struct {
	Schema     Schema
	Definition Definitions
}

var Descriptions []Description

func main() {
	builder := Builder{
		Title:        "Test title",
		Version:      "13.1.9.2",
		OutputFile:   "./index.html",
		TemplateFile: "./index.tmpl",
	}

	files := getWsdlXsdFiles(BASE_PATH)
	fmt.Println(files)
	xsdFile := "./wsdl-xsd/hcs-appeals-types.xsd"
	xsdData := getFileData(xsdFile)
	xsdSchema := getSchema(xsdData)

	wsdlFile := "./wsdl-xsd/hcs-appeals-service-async.wsdl"
	wsdlData := getFileData(wsdlFile)
	wsdlDef := getDefinition(wsdlData)

	Descriptions = append(Descriptions, Description{xsdSchema, wsdlDef})
	builder.Descriptions = Descriptions

	fmt.Println(xsdSchema, wsdlDef)
	builder.run()
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

func getWsdlXsdFiles(path string) []string {
	files := []string{}
	err := filepath.Walk(path,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				panic(err)
			}
			if isValidFile(path, info) {
				files = append(files, path)
			}
			return nil
		})
	if err != nil {
		log.Println(err)
	}
	return files
}
