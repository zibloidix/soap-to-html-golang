package main

import (
	"fmt"
	"os"
	"text/template"
)

type Builder struct {
	Title        string
	Version      string
	OutputFile   string
	TemplateFile string
	Services     map[string]ServiceDescription
}

func (r *Builder) run() {
	fmt.Println(r)
	data, err := os.ReadFile(r.TemplateFile)
	if err != nil {
		panic(err)
	}

	out, err := os.Create(r.OutputFile)
	if err != nil {
		panic(err)
	}

	t := template.New("tmp")
	t.Parse(string(data))
	t.Execute(out, r)
}
