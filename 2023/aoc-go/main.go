package main

import (
	"embed"
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

//go:embed templates/*
var content embed.FS

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s DAY\n", os.Args[0])
		os.Exit(1)
	}

	if err := run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run() error {
	funcMap := template.FuncMap{
		"title": cases.Title(language.Und, cases.NoLower).String,
	}
	tmpl, err := template.New("generate").Funcs(funcMap).ParseFS(content, "templates/*")
	if err != nil {
		return err
	}

	if err := os.MkdirAll(os.Args[1], 0775); err != nil {
		return err
	}

	info := []struct {
		templateName string
		fileName     string
		Day          string
	}{
		{
			templateName: "day.go.tmpl",
			fileName:     fmt.Sprintf("%s.go", os.Args[1]),
			Day:          os.Args[1],
		},
		{
			templateName: "day_test.go.tmpl",
			fileName:     fmt.Sprintf("%s_test.go", os.Args[1]),
			Day:          os.Args[1],
		},
	}

	for _, v := range info {
		f, err := os.Create(filepath.Join(os.Args[1], v.fileName))
		if err != nil {
			return err
		}

		if err := tmpl.ExecuteTemplate(f, v.templateName, v); err != nil {
			f.Close()
			return err
		}

		f.Close()
	}

	return nil
}
