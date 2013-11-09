// Copyright 2013 Andreas Koch. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package modules

import (
	"fmt"
	"github.com/andreaskoch/dotman/util/fs"
	"strings"
)

const (
	ModuleFileName = "dotman"
)

type Collection struct {
	BaseDirectory string
	Collection    []*Module
}

func Load(directory string) (*Collection, error) {

	// check if the directory exists
	if !fs.DirectoryExists(directory) {
		return nil, fmt.Errorf("The directory %q does not exist.", directory)
	}

	// find all folders with module files in them (non-recursive)
	moduleDirectories, err := getAllModuleDirectories(directory, ModuleFileName)
	if err != nil {
		return nil, fmt.Errorf("Unable scan the directory %q for modules. Error: %s", directory, err)
	}

	// try to create modules from each module directory
	modules := make([]*Module, 0)
	errors := make([]string, 0)
	for _, moduleDirectory := range moduleDirectories {

		module, err := newModule(moduleDirectory, ModuleFileName)
		if err != nil {
			errors = append(errors, err.Error())
			continue
		}

		modules = append(modules, module)
	}

	// create the module collection
	collection := &Collection{
		BaseDirectory: directory,
		Collection:    modules,
	}

	// partial success (not all modules could be read)
	if len(errors) > 0 {
		return collection, fmt.Errorf("%s", strings.Join(errors, ",\n"))
	}

	// success
	return collection, nil
}
