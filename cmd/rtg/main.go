package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"github.com/fatih/color"
)

const usage = `Usage:
	rtg [OPTION] [OBJECT] NAME

Options:
	-g, --generate OBJECT NAME
		Generate a new OBJECT with the given NAME.
		Any of the following options is a valid OBJECT:
		  c, component   creates a new React Native component
		  x, screen      creates a new React Native screen

Example:
	rtg -g c MyComponent
	   Created ./components/MyComponent/MyComponent.tsx
	   Created ./components/MyComponent/MyComponentStyle.tsx
`

const pathToComponents = "./components/"
const pathToScreens = "./screens/"

type PrintfFunc func(msg string, args ...interface{})
type Logger struct {
	Fatalf, Warnf, Infof, Successf PrintfFunc
}

var log = Logger{
	Fatalf: func(msg string, args ...interface{}) {
		color.Set(color.FgHiRed)
		fmt.Printf(msg, args...)
		color.Unset()
		fmt.Println("\nIf you feel like you need a refresh, try running rtg --help to read the usage notes.")
		fmt.Println("If the error is something unexpected, please open an issue on CanobbioE/react-template-generator/issues.")
		os.Exit(1)
	},
	Infof: func(msg string, args ...interface{}) {
		fmt.Printf(msg, args...)
	},
	Warnf: func(msg string, args ...interface{}) {
		color.Set(color.FgHiYellow)
		fmt.Printf(msg, args...)
		color.Unset()
	},
	Successf: func(msg string, args ...interface{}) {
		color.Set(color.FgHiGreen)
		fmt.Printf(msg, args...)
		color.Unset()
	},
}

func main() {
	flag.Usage = func() { fmt.Printf("%s\n", usage) }
	var generateFlag string

	flag.StringVar(&generateFlag, "g", "", "generate a new object")
	flag.StringVar(&generateFlag, "generate", "", "generate a new object")

	flag.Parse()
	name := flag.Arg(0)

	if flag.NArg() > 1 {
		log.Fatalf("Error: too many arguments.\n" +
			"rtg accepts a single mandatory argument for the object's name.")
	}
	if flag.NArg() < 1 || name == "" || name == " " || name == "-" {
		log.Fatalf("Error: not enough arguments.\n" +
			"rtg accepts a single mandatory argument for the object's name.")
	}

	if generateFlag == "" {
		log.Fatalf("Error: not enough arguments.\n" +
			"Flag -g/--generate expects a single value for the object's type.")
	}

	switch generateFlag {
	case "c":
		fallthrough
	case "component":
		log.Infof("Creating a new React Native component...\n")
		generateTsxFromTemplate(pathToComponents, name)
	case "x":
		fallthrough
	case "screen":
		log.Infof("Creating a new React Native screen...\n")
		generateTsxFromTemplate(pathToScreens, name)
	default:
		log.Fatalf("Error: Unrecognized object's type.\n")
	}
}

// generateTsxFromTemplate creates two files "<name>.tsx"
// and "<name>Style.tsx" inside the "./components/<name>/" folder.
// It works even if name is a path.
func generateTsxFromTemplate(path, name string) {
	generateDir(path + name)
	data := struct{ Name string }{Name: name}
	generateFromTemplate(pathToComponents+"/"+name+".tsx", "componentTemplate", componentTemplate, data)
	generateFromString(pathToComponents+"/"+name+"Style.tsx", styleTemplate)
}

// generateDir creates a directory and all the parents
// directories that do not exist.
func generateDir(path string) {
	path = filepath.FromSlash(path)
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		log.Fatalf("Error: cannot create resource %s.\n"+
			"The error is: %v", path, err)
	}
}

// generateFromTemplate creates a new file at the specified location.
// The fullpath is the complete file path (e.g. ./example/example.txt).
// The created file will be populate using the template defined by
// templateName and actualTemplate populating it with data.
func generateFromTemplate(fullpath, templateName, actualTemplate string, data interface{}) {
	fullpath = filepath.FromSlash(fullpath)
	file, err := os.Create(fullpath)
	if err != nil {
		log.Fatalf("Error: cannot create component file %s.\n"+
			"The error is: %v", fullpath, err)
	}

	t := template.Must(template.New(templateName).Parse(actualTemplate))
	err = t.Execute(file, data)
	if err != nil {
		log.Fatalf("Error: cannot generate from template %s.\n"+
			"The error is: %v", fullpath, err)
	}

	log.Successf("Created %v\n", fullpath)
}

// generateFromString creates a new file at the specified location.
// The fullpath is the complete file path (e.g. ./example/example.txt).
// src is the string that is going to be copied in the file.
func generateFromString(fullpath, src string) {
	fullpath = filepath.FromSlash(fullpath)
	file, err := os.Create(fullpath)
	if err != nil {
		log.Fatalf("Error: cannot generate file from string %s.\n"+
			"The error is: %v", fullpath, err)
	}

	_, err = file.WriteString(src)
	if err != nil {
		log.Fatalf("Error: cannot write from string %s.\n"+
			"The error is: %v", fullpath, err)
	}

	log.Successf("Created %v\n", fullpath)
}
