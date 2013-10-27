// Copyright 2013 Andreas Koch. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package projects

import (
	"fmt"
	"github.com/andreaskoch/dotman/mapping"
	"github.com/andreaskoch/dotman/util/fs"
	"path/filepath"
)

type Project struct {
	Map *mapping.PathMap

	name      string
	directory string
}

func newProject(directory, projectFileName string) (*Project, error) {

	projectFilePath := filepath.Join(directory, projectFileName)

	// check if the project file exists
	if !fs.FileExists(projectFilePath) {
		return nil, fmt.Errorf("Project file file %q does not exist.", projectFilePath)
	}

	// read the project file
	projectPathMap, err := mapping.NewPathMap(projectFilePath)
	if err != nil {
		return nil, fmt.Errorf("Unable to read dotman file. %s", err)
	}

	return &Project{
		Map:       projectPathMap,
		name:      filepath.Base(directory),
		directory: directory,
	}, nil
}

func (project *Project) String() string {
	return project.name
}
