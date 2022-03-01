package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

type Description struct {
	Schema     Schema
	Definition Definitions
}

var Descriptions []Description

func main() {
	builderConf := flag.String("builder-conf", "./builder-conf.xml", "Файл с настроками для сборщика документации")
	wsdlXsdPath := flag.String("wsdl-xsd-path", ".", "Путь до директории, которая содержит wsdl и xsd файлы")
	flag.Parse()

	builderData := getFileData(*builderConf)
	builder := getBuilder(builderData)

	files := getWsdlXsdFiles(*wsdlXsdPath)
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

func getBuilder(data []byte) Builder {
	var builder Builder
	err := xml.Unmarshal(data, &builder)
	if err != nil {
		panic(err)
	}
	return builder
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
