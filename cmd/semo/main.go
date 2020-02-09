package main

import (
	"fmt"
	"math/rand"
	"os"
	"path"
	"time"

	"github.com/faaxm/semo/semo"
)

const LocalTemplateDirName = "semo_templates"
const HomeTemplateDirName = ".semo/templates"

func main() {
	rand.Seed(time.Now().UnixNano())

	if len(os.Args) < 2 {
		fmt.Printf("Error: Missing template name.\n")
		fmt.Printf("Expected format:\n%v template_name [template_param...]\n", os.Args[0])
		return
	}
	templateId := os.Args[1]
	templateArgs := os.Args[2:]

	store := semo.NewStore()
	addSearchPathsToStore(store)

	tmpl, err := store.Template(templateId)
	if err == nil {
		err = tmpl.ReadFieldsFromArgs(templateArgs)
	}
	if err == nil {
		err = tmpl.RequestMissingFields()
	}
	if err == nil {
		fmt.Printf("Applying template \"%v\"\n", templateId)
		err = instantiateTemplate(tmpl)
	}

	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}

// addSearchPathsToStore adds available template directories
// to the template store's search paths.
func addSearchPathsToStore(store *semo.Store) {
	// Look for a local template directory in the
	// parent directories
	localTmplDir := findNextLocalStore()
	if localTmplDir != nil {
		store.AddSearchPath(*localTmplDir)
	}

	// Add the common template directory from the
	// user home directory
	homedir, err := os.UserHomeDir()
	if err == nil {
		homeTmplDir := path.Join(homedir, HomeTemplateDirName)
		if dirExists(homeTmplDir) {
			store.AddSearchPath(homeTmplDir)
		}
	}
}

// instantiateTemplate instantiates the given template in the
// current working directory.
func instantiateTemplate(tmpl *semo.Template) error {
	return tmpl.Instantiate(".")
}

// findNextLocalStore iterates over the parent directories
// of the current working directory and looks for the closest
// template directory.
func findNextLocalStore() *string {
	workingdir, err := os.Getwd()
	if err != nil {
		return nil
	}

	var result *string
	iterateParentDirs(workingdir, func(dir string) bool {
		tmplDir := path.Join(dir, LocalTemplateDirName)

		if dirExists(tmplDir) {
			result = &tmplDir
			return false
		}
		return true
	})

	return result
}

// iterateParentDirs calls cb for the directory at 'dir'
// as well as all its parent directories. It breaks once
// cb returns false.
func iterateParentDirs(dir string, cb func(string) bool) {
	keep_running := true
	for prevDir := ""; prevDir != dir && keep_running; prevDir, dir = dir, path.Dir(dir) {
		keep_running = cb(dir)
	}
}

// dirExists returns true if a directory
// exists at the given path.
func dirExists(path string) bool {
	if info, err := os.Stat(path); err == nil && info.IsDir() {
		return true
	}
	return false
}
