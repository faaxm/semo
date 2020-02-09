package semo

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"text/template"
)

const TemplateFilesDirName = "_template_files"
const ConfigFileName = "config.yml"

type Template struct {
	RootPath string
	Data     map[string]string
	Config   Config
}

// NewTemplate creates a new template with its directory located
// at 'rootPath'.
func NewTemplate(rootPath string) (*Template, error) {
	filesDir := path.Join(rootPath, TemplateFilesDirName)

	if info, err := os.Stat(filesDir); err != nil || !info.IsDir() {
		return nil, fmt.Errorf("not a valid template as '%v' is not a directory", filesDir)
	}
	templ := Template{RootPath: rootPath}

	err := templ.readConfigFile()
	if err != nil {
		return nil, err
	}

	return &templ, nil
}

// ReadFieldsFromArgs will take the list of strings in args
// and use them to configure the Data of its Fields.
func (me *Template) ReadFieldsFromArgs(args []string) error {
	me.Data = dataMapFromArgs(args)

	for index, field := range me.Config.Fields {
		if index < len(args) {
			me.Data[field.Name] = args[index]
		}
	}

	return nil
}

// RequestMissingFields asks the user for the missing params.
func (me *Template) RequestMissingFields() error {
	for _, field := range me.Config.Fields {
		if _, ok := me.Data[field.Name]; !ok {
			me.requestFieldFromUser(field)
		}
	}

	return nil
}

// Instantiate will create an instance of the template
// at the 'dstPath' location.
func (me *Template) Instantiate(dstPath string) error {
	return me.instantiateDir(me.filesDir(), dstPath)
}

// readConfigFile reads the config file at the 'RootPath'
// of the template.
func (me *Template) readConfigFile() error {
	configFilePath := path.Join(me.RootPath, ConfigFileName)
	if info, err := os.Stat(configFilePath); err != nil || info.IsDir() {
		return nil // It is valid that there is no config file
	}

	conf, err := NewConfig(configFilePath)
	if err == nil {
		me.Config = *conf
	} else {
		return err
	}

	return nil
}

// requestFieldsFromUser will show a prompt to the user
// and ask for the value for the given field.
// It then updates the 'Data' of the template.
func (me *Template) requestFieldFromUser(field Field) {
	var data string

	fmt.Println("")
	if len(field.Description) > 0 {
		fmt.Println("Description:", field.Description)
	}
	fmt.Printf("%v [%v]: ", field.Name, field.Default)
	fmt.Scanln(&data)

	if len(data) > 0 {
		me.Data[field.Name] = data
	} else {
		me.Data[field.Name] = field.Default
	}
}

// filesDir returns the directory that contains the raw
// template files.
func (me *Template) filesDir() string {
	return path.Join(me.RootPath, TemplateFilesDirName)
}

// instantiateDir creates a copy of the given raw directory
// of the template.
func (me *Template) instantiateDir(sourceRoot, destRoot string) error {
	files, err := ioutil.ReadDir(sourceRoot)
	if err != nil {
		return err
	}

	for _, file := range files {
		sourcePath := path.Join(sourceRoot, file.Name())
		destName, err := me.renderString(file.Name())
		if err != nil {
			return err
		}
		destPath := path.Join(destRoot, destName)

		if file.IsDir() {
			os.MkdirAll(destPath, file.Mode())
			me.instantiateDir(sourcePath, destPath)
		} else {
			err = me.instantiateFile(sourcePath, destPath)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// instantiateFile creates a copy of the given raw file and will
// place it at 'dest'. 'source' has to be inside of the
// template files directory.
func (me *Template) instantiateFile(source, dest string) error {
	tmpl, err := template.New(path.Base(source)).Funcs(defaultFuncMap).ParseFiles(source)
	if err != nil {
		return err
	}

	destFile, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer destFile.Close()

	me.log(fmt.Sprintf("Copying %s -> %s", path.Base(source), dest))
	err = tmpl.Execute(destFile, me.Data)

	return err
}

// renderString renders the given string using the templating engine
// and the 'Data' available in the template struct.
func (me *Template) renderString(str string) (string, error) {
	tmpl, err := template.New("string_tmpl").Funcs(defaultFuncMap).Parse(str)
	if err != nil {
		return "", err
	}
	var b strings.Builder

	err = tmpl.Execute(&b, me.Data)
	if err != nil {
		return "", err
	}
	return b.String(), nil
}

// log is used when the current status is reported
// to the user.
func (me *Template) log(text string) {
	fmt.Println(text)
}

// dataMapFromArgs converts the given slice of strings
// to a map with 'arg_N' style keys.
func dataMapFromArgs(args []string) map[string]string {
	data := make(map[string]string)

	for idx, arg := range args {
		key_name := fmt.Sprintf("arg_%d", idx)
		data[key_name] = arg
	}

	return data
}
