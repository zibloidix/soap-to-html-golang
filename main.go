package main

import (
	"encoding/xml"
	"flag"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

type ServiceDescription struct {
	Schemas     []Schema
	Definitions []Definitions
}

func main() {
	builderConf := flag.String("builder-conf", "./builder-conf.xml", "Файл с настроками для сборщика документации")
	wsdlXsdPath := flag.String("wsdl-xsd-path", ".", "Путь до директории, которая содержит wsdl и xsd файлы")
	flag.Parse()

	builderData := getFileData(*builderConf)
	builder := getBuilder(builderData)

	files := getWsdlXsdFiles(*wsdlXsdPath)

	builder.Services = getServices(files)
	builder.run()

}

func getServices(files []string) map[string]ServiceDescription {
	services := map[string]ServiceDescription{}
	for _, file := range files {
		dir := path.Dir(file)
		if _, ok := services[dir]; !ok {
			services[dir] = ServiceDescription{}
		}
		fileData := getFileData(file)
		service := services[dir]
		if strings.Contains(file, ".wsdl") {
			wsdlDef := getDefinition(fileData)
			service.Definitions = append(service.Definitions, wsdlDef)
		}
		if strings.Contains(file, ".xsd") {
			xsdSchema := getSchema(fileData)
			service.Schemas = append(service.Schemas, xsdSchema)
		}
		services[dir] = service
	}
	return services
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
