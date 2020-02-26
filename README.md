# React Template Generator

A CLI to automatically generate repetitive code files from templates.

The idea is to have a tool like the [Angular CLI](https://cli.angular.io/) for ReactTS and ReactNative.

## Installation

On Windows, Linux, and macOS, you can use the [pre-built binaries](https://github.com/CanobbioE/template-generator/releases).
You can then move the binary into any folder specified in your `PATH` environment variable.
I suggest to rename the executable to `rtg` for an easier usage.

If you have Go 1.13+ you can build from source:

```bash
$ git clone github.com/CanobbioE/react-template-generator
$ cd react-template-generator
$ go install
```

## Usage

```bash
Usage:
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
	   Created ./components/MyComponent/MyComponentStyles.tsx
```
