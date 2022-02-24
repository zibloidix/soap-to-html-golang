package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

// Для примера разбора wsdl и xsd
// https://github.com/fiorix/wsdl2go/blob/master/wsdl/types.go

type Schema struct {
	XMLName         xml.Name          `xml:"schema"`
	TargetNamespace string            `xml:"targetNamespace,attr"`
	Namespaces      map[string]string `xml:"-"`
	Imports         []ImportSchema    `xml:"import"`
	Includes        []IncludeSchema   `xml:"include"`
	SimpleTypes     []SimpleType      `xml:"simpleType"`
	ComplexTypes    []ComplexType     `xml:"complexType"`
	Elements        []Element         `xml:"element"`
}

type SimpleType struct {
	XMLName         xml.Name    `xml:"simpleType"`
	Name            string      `xml:"name,attr"`
	Union           Union       `xml:"union"`
	Restriction     Restriction `xml:"restriction"`
	TargetNamespace string
}

type Union struct {
	XMLName     xml.Name `xml:"union"`
	MemberTypes string   `xml:"memberTypes,attr"`
}

type Restriction struct {
	XMLName    xml.Name    `xml:"restriction"`
	Base       string      `xml:"base,attr"`
	Enum       []Enum      `xml:"enumeration"`
	Attributes []Attribute `xml:"attribute"`
}

type Enum struct {
	XMLName xml.Name `xml:"enumeration"`
	Value   string   `xml:"value,attr"`
}

type ComplexType struct {
	XMLName         xml.Name       `xml:"complexType"`
	Name            string         `xml:"name,attr"`
	Abstract        bool           `xml:"abstract,attr"`
	Doc             string         `xml:"annotation>documentation"`
	AllElements     []Element      `xml:"all>element"`
	ComplexContent  ComplexContent `xml:"complexContent"`
	SimpleContent   SimpleContent  `xml:"simpleContent"`
	Sequence        Sequence       `xml:"sequence"`
	Choice          Choice         `xml:"choice"`
	Attributes      []Attribute    `xml:"attribute"`
	TargetNamespace string
}

type SimpleContent struct {
	XMLName     xml.Name    `xml:"simpleContent"`
	Extension   Extension   `xml:"extension"`
	Restriction Restriction `xml:"restriction"`
}

type ComplexContent struct {
	XMLName     xml.Name    `xml:"complexContent"`
	Extension   Extension   `xml:"extension"`
	Restriction Restriction `xml:"restriction"`
}

type Extension struct {
	XMLName    xml.Name    `xml:"extension"`
	Base       string      `xml:"base,attr"`
	Sequence   Sequence    `xml:"sequence"`
	Choice     Choice      `xml:"choice"`
	Attributes []Attribute `xml:"attribute"`
}

type Sequence struct {
	XMLName      xml.Name      `xml:"sequence"`
	ComplexTypes []ComplexType `xml:"complexType"`
	Elements     []Element     `xml:"element"`
	Any          []AnyElement  `xml:"any"`
	Choices      []Choice      `xml:"choice"`
}

type Choice struct {
	XMLName      xml.Name      `xml:"choice"`
	ComplexTypes []ComplexType `xml:"complexType"`
	Elements     []Element     `xml:"element"`
	Any          []AnyElement  `xml:"any"`
}

type Attribute struct {
	XMLName   xml.Name `xml:"attribute"`
	Name      string   `xml:"name,attr"`
	Ref       string   `xml:"ref,attr"`
	Type      string   `xml:"type,attr"`
	ArrayType string   `xml:"arrayType,attr"`
	Min       int      `xml:"minOccurs,attr"`
	Max       string   `xml:"maxOccurs,attr"`
	Nillable  bool     `xml:"nillable,attr"`
}

type Element struct {
	XMLName     xml.Name    `xml:"element"`
	Name        string      `xml:"name,attr"`
	Ref         string      `xml:"ref,attr"`
	Type        string      `xml:"type,attr"`
	Min         int         `xml:"minOccurs,attr"`
	Max         string      `xml:"maxOccurs,attr"`
	Nillable    bool        `xml:"nillable,attr"`
	ComplexType ComplexType `xml:"complexType"`
}

type AnyElement struct {
	XMLName xml.Name `xml:"any"`
	Min     int      `xml:"minOccurs,attr"`
	Max     string   `xml:"maxOccurs,attr"`
}

type ImportSchema struct {
	XMLName   xml.Name `xml:"import"`
	Namespace string   `xml:"namespace,attr"`
	Location  string   `xml:"schemaLocation,attr"`
}

type IncludeSchema struct {
	XMLName   xml.Name `xml:"include"`
	Namespace string   `xml:"namespace,attr"`
	Location  string   `xml:"schemaLocation,attr"`
}

func main() {
	fileName := "./wsdl-xsd/hcs-house-management-types.xsd"
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
