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
		Generate a new OBJECT with the given NAME. An object can only be a component (c).

Example:
	rtg -g c MyComponent
	   Created ./components/MyComponent/MyComponent.tsx
	   Created ./components/MyComponent/MyComponentStyle.tsx
`

const pathToComponents = "./components/"

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
		fmt.Println("If the error is something unexpected, please open an issue on CanobbioE/template-generator.")
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
		generateReactNativeComponent(name)
	default:
		log.Fatalf("Error: Unrecognized object's type.\n")
	}
}

func generateReactNativeComponent(name string) {
	log.Infof("Creating a new React Native component...\n")
	path := filepath.FromSlash(pathToComponents + name)
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		log.Fatalf("Error: cannot create destination folder.\n"+
			"The error is: %v", err)
	}

	componentFilename := filepath.FromSlash(path + "/" + name + ".tsx")
	componentFile, err := os.Create(componentFilename)
	if err != nil {
		log.Fatalf("Error: cannot create component file.\n"+
			"The error is: %v", err)
	}

	styleFilename := filepath.FromSlash(path + "/" + name + "Style.tsx")
	styleFile, err := os.Create(styleFilename)
	if err != nil {
		log.Fatalf("Error: cannot create style file.\n"+
			"The error is: %v", err)
	}

	data := struct {
		Name string
	}{
		Name: name,
	}
	t := template.Must(template.New("componentTemplate").Parse(componentTemplate))
	err = t.Execute(componentFile, data)
	if err != nil {
		log.Fatalf("Error: cannot write component template.\n"+
			"The error is: %v", err)
	}
	log.Successf("Created %v\n", componentFilename)

	_, err = styleFile.WriteString(styleTemplate)
	if err != nil {
		log.Fatalf("Error: cannot write style template.\n"+
			"The error is: %v", err)
	}
	log.Successf("Created %v\n", styleFilename)
}
