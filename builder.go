package main

import (
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

func (b *Builder) run() {
	data, err := os.ReadFile(b.TemplateFile)
	if err != nil {
		panic(err)
	}

	out, err := os.Create(b.OutputFile)
	if err != nil {
		panic(err)
	}

	t := template.New("tmp")
	t.Parse(string(data))
	t.Execute(out, b)
}
