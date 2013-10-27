// Copyright 2013 Andreas Koch. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package projects

import (
	"fmt"
	"github.com/andreaskoch/dotman/util/fs"
	"strings"
)

const (
	ProjectFileName = "dotman"
)

type Collection struct {
	BaseDirectory string
	Collection    []*Project
}

func Load(directory string) (*Collection, error) {

	// check if the directory exists
	if !fs.DirectoryExists(directory) {
		return nil, fmt.Errorf("The directory %q does not exist.", directory)
	}

	// find all folders with project files in them (non-recursive)
	projectDirectories, err := getAllProjectDirectories(directory, ProjectFileName)
	if err != nil {
		return nil, fmt.Errorf("Unable scan the directory %q for projects. Error: %s", directory, err)
	}

	// try to create projects from each project directory
	projects := make([]*Project, 0)
	errors := make([]string, 0)
	for _, projectDirectory := range projectDirectories {

		project, err := newProject(projectDirectory, ProjectFileName)
		if err != nil {
			errors = append(errors, err.Error())
			continue
		}

		projects = append(projects, project)
	}

	// create the project collection
	collection := &Collection{
		BaseDirectory: directory,
		Collection:    projects,
	}

	// partial success (not all projects could be read)
	if len(errors) > 0 {
		return collection, fmt.Errorf("%s", strings.Join(errors, ",\n"))
	}

	// success
	return collection, nil
}
