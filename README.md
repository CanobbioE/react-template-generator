# React Template Generator
A CLI to automatically generate components and whatnot for the "frankfurt gang" development team.

The idea is to have a tool like the [Angular CLI](https://cli.angular.io/) for ReactTS and ReactNative.

Initially we will only need to generate components like:

```bash
$ ng -g c component
```
should be equivalent to
```bash
$ rtg -g c component
```

## Installation

On Windows, Linux, and macOS, you can use the [pre-built binaries](https://github.com/CanobbioE/template-generator/releases).

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
		Generate a new OBJECT with the given NAME. An object can only be a component (c).

Example:
	rtg -g c MyComponent
	   Created ./components/MyComponent/MyComponent.tsx
	   Created ./components/MyComponent/MyComponentStyle.tsx
```
