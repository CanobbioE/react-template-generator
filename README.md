# React Template Generator
A CLI to automatically generate components and whatnot for the "frankfurt gang" development team.

The idea is to have a tool like the [Angular CLI](https://cli.angular.io/) for ReactTS and ReactNative.

We want this:
```bash
// Angular CLI
$ ng -g c component
```

to be equivalent to this:

```bash
// rtg CLI
$ rtg -g c component
```

## Installation

On Windows, Linux, and macOS, you can use the [pre-built binaries](https://github.com/CanobbioE/template-generator/releases).
You can then move the binary into any folder specified in your `PATH` environment variable.
I suggest to rename the executable to `rtg`.

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
