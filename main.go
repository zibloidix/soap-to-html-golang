package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type Schema struct {
	XMLName     xml.Name     `xml:"schema"`
	Imports     []Import     `xml:"import"`
	SimpleTypes []SimpleType `xml:"simpleType"`
}

type SimpleType struct {
	XMLName     xml.Name `xml:"simpleType"`
	Name        string   `xml:"name,attr"`
	Annotation  Annotation
	Restriction Restriction
}

type Annotation struct {
	XMLName       xml.Name      `xml:"annotation"`
	Documentation Documentation `xml:"documentation"`
}

type Documentation struct {
	XMLName xml.Name `xml:"documentation"`
	Value   string   `xml:",chardata"`
}

type Restriction struct {
	XMLName      xml.Name      `xml:"restriction"`
	Base         string        `xml:"base,attr"`
	Enumerations []Enumeration `xml:"enumeration"`
}

type Enumeration struct {
	XMLName xml.Name `xml:"enumeration"`
	Value   string   `xml:"value,attr"`
}

type Import struct {
	XMLName        xml.Name `xml:"import"`
	Namespace      string   `xml:"namespace,attr"`
	SchemaLocation string   `xml:"schemaLocation,attr"`
}

func main() {
	fileName := "./wsdl-xsd/hcs-rvuo-types.xsd"
	data, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	var schema Schema
	err = xml.Unmarshal(data, &schema)
	if err != nil {
		panic(err)
	}
	fmt.Println(schema)
}
