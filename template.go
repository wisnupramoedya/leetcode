package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"text/template"
)

type TemplateData struct {
	Number int
	Name   string
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run template.go <Number> <Name>")
		return
	}

	// Insert the argument
	number, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	}
	name := os.Args[2]
	data := TemplateData{Name: name, Number: number}

	// Create directory structures
	dir := fmt.Sprintf("%04d.%s/1", data.Number, spaceToMinus(data.Name))
	err = os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		panic(err)
	}

	// Insert template on each file
	funcMap := template.FuncMap{
		"spaceToMinus":      spaceToMinus,
		"spaceToUnderscore": spaceToUnderscore,
		"toLower":           toLower,
	}

	codeTmpl, err := template.New("code.tmpl").Funcs(funcMap).ParseFiles("./templates/code.tmpl")
	if err != nil {
		panic(err)
	}
	testTmpl, err := template.New("test.tmpl").Funcs(funcMap).ParseFiles("./templates/test.tmpl")
	if err != nil {
		panic(err)
	}

	codeFileName := fmt.Sprintf("%s/%s.go", dir, toLower(spaceToUnderscore(data.Name)))
	testFileName := fmt.Sprintf("%s/%s_test.go", dir, toLower(spaceToUnderscore(data.Name)))

	createFileWithTemplate(testFileName, testTmpl, data)
	createFileWithTemplate(codeFileName, codeTmpl, data)
}

// Helper function to replace spaces with dashes
func spaceToMinus(s string) string {
	return strings.ReplaceAll(s, " ", "-")
}

// Helper function to replace spaces with underscores
func spaceToUnderscore(s string) string {
	return strings.ReplaceAll(s, " ", "_")
}

// Helper function to replace spaces with lower
func toLower(s string) string {
	return strings.ToLower(s)
}

// Helper function to create a file with the specified template and data
func createFileWithTemplate(fileName string, tmpl *template.Template, data TemplateData) {
	f, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			panic(err)
		}
	}(f)

	err = tmpl.Execute(f, data)
	if err != nil {
		panic(err)
	}
}
