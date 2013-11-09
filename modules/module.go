// Copyright 2013 Andreas Koch. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package modules

import (
	"fmt"
	"github.com/andreaskoch/dotman/mapping"
	"github.com/andreaskoch/dotman/util/fs"
	"path/filepath"
)

func newModule(directory, moduleFileName string) (*Module, error) {

	moduleFilePath := filepath.Join(directory, moduleFileName)

	// check if the module file exists
	if !fs.FileExists(moduleFilePath) {
		return nil, fmt.Errorf("Module file file %q does not exist.", moduleFilePath)
	}

	// read the module file
	modulePathMap, err := mapping.NewPathMap(moduleFilePath)
	if err != nil {
		return nil, fmt.Errorf("Unable to read dotman file. %s", err)
	}

	return &Module{
		Map:        modulePathMap,
		name:       filepath.Base(directory),
		directory:  directory,
		moduleFile: moduleFilePath,
	}, nil
}

type Module struct {
	Map *mapping.PathMap

	name       string
	directory  string
	moduleFile string
}

func (module *Module) String() string {
	return module.name
}

func (module *Module) Directory() string {
	return module.directory
}

func (module *Module) ModuleFile() string {
	return module.moduleFile
}
