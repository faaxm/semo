package semo

import (
	"errors"
	"path"
	"strings"
)

type Store struct {
	RootPaths []string
}

func NewStore() *Store {
	var s Store
	return &s
}

func (me *Store) AddSearchPath(path string) {
	me.RootPaths = append(me.RootPaths, path)
}

func (me *Store) Template(name string) (tmpl *Template, err error) {
	if len(me.RootPaths) == 0 {
		return nil, errors.New("no template store configured")
	}

	for _, root := range me.RootPaths {
		tmpl, err = templateAtRoot(root, name)
		if err == nil {
			return
		}
	}

	return
}

func templateAtRoot(root, name string) (*Template, error) {
	template_path := root

	// Construct the template path
	for _, component := range strings.Split(name, ".") {
		template_path = path.Join(template_path, component)
	}

	return NewTemplate(template_path)
}
