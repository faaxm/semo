# Semo

Semo is a generic scaffolding/templating tool.
Build a template for your favourite project structure and use Semo to instantiate new copies of it with a single command.

You can also place templates at the root of a project to easily extend it and make it easy to have a consistent file structure.
For example, run
```
$ semo new_controller MyControllerName
```
and a new controller class will be created in the right directory of your project.


# How to install

The easiest way is to first install `go` and then run
```
$ go get github.com/faaxm/semo/cmd/semo
$ go install github.com/faaxm/semo/cmd/semo
```
Semo will then be downloaded and installed into your go-workspace, which is usually in your user directory at `~/go`.

Place your template files either in your user directory at `~/.semo/templates` or in your project's root folder
in `semo_templates`.


# First Steps

Semo finds templates by converting their name to a path. For example the template `cpp.class` can be found in
`semo_templates` in the subdirectory `cpp/class`.

To create a new c++ class, consisting of a header and a cpp-file, go to the root directory of the repository (or any subdirectory)
and type
```
$ semo cpp.class MyClass
```
Semo will then create a `MyClass.cpp` and `MyClass.h` file.


# Creating your own templates

The simplest template in Semo is just a directory with the templates name. Inside that root directory, there has to be a
`_template_files` directory containing the files and directories that will be copied when the template is instantiated.

A very simple template for a c header file could look like this:
```
header
└── _template_files
    └── {{ .arg_0 }}.h
```

Notice the name of the template header file contains `{{ .arg_0 }}`. This part will be replaced with the first
argument given to Semo when invoked on the command line. To create a header with the name "TestHeader", you
would call Semo like this:
```
$ semo header TestHeader
```

The same placeholders (`{{ .arg_0 }}`, `{{ .arg_1 }}`, ..., `{{ .arg_N }}`) can be used inside the template header file.

For more information on how these placeholders work, have a look at the `text/template` package from go:
https://golang.org/pkg/text/template/


# More advanced templates

If you don't want to use the standard `.arg_N` placeholders, you can define your own field names together with
default values. To do this, create a `config.yml` file inside the root directory of your template.

This file contains a `fields` entry with a list of the required fields. Each field has to have a `name`, but can
also have a `description` and a `default` value. Have a look at the "cpp.class" template for an example.

Values for each field can still be given on the command line in the order in which they are defined in the
`config.yml` file. However, you can now also call `semo` without providing any arguments other than the
template name, like
```
$ semo cpp.class
```
In that case, Semo will query the user interactively for the data, with the option of using the default
value.


# Functions

Semo provides several functions to alter the input field values:

* `{{ toUpper .arg_0 }}`: Convert all characters of `.arg_0` to upper case.
* `{{ toLower .arg_0 }}`: Convert all characters of `.arg_0` to lower case.
* `{{ noWhitespace .arg_0 }}`: Remove all whitespace from `.arg_0` and replace it with underscores.
* `{{ runID N }}`: Returns a unique ID of maximum length `N`. This ID is generated once when Semo is launched
  and will stay the same for all files in the template.

Additional functions are provided by the `test/template` package of go. You can also concatenate function calls
like this: `{{ toUpper (runID 5) }}`.

# License

This project is licensed under the terms of the GPL version 3.0 license. See LICENSE.txt for more information.
