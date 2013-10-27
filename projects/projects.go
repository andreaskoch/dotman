// Copyright 2013 Andreas Koch. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package projects

import (
	"fmt"
	"github.com/andreaskoch/dotman/ui"
	"github.com/andreaskoch/dotman/util/fs"
	"io/ioutil"
	"path/filepath"
	"strings"
)

type Collection struct {
	BaseDirectory string
	Collection    []*Project
}

func Load(directory, projectFileName string) (*Collection, error) {

	// check if the directory exists
	if !fs.DirectoryExists(directory) {
		return nil, fmt.Errorf("The directory %q does not exist.", directory)
	}

	// validate the mapping file name
	if strings.TrimSpace(projectFileName) == "" {
		return nil, fmt.Errorf("The project file name cannot be empty.")
	}

	// scan for projects in the specified directory
	projects := make([]*Project, 0)
	errors := make([]string, 0)
	projectDirectories := getAllProjectDirectories(directory, projectFileName)
	for _, projectDirectory := range projectDirectories {

		project, err := newProject(projectDirectory, projectFileName)
		if err != nil {
			errors = append(errors, err.Error())
			continue
		}

		projects = append(projects, project)
	}

	// create the collection
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

func getAllProjectDirectories(baseDirectory, projectFileName string) []string {

	// abort if the supplied directory is not a directory or does not exist
	if !fs.IsDirectory(baseDirectory) {
		ui.Fatal("%q is not a directory.", baseDirectory)
	}

	baseDirectoryEntries, err := ioutil.ReadDir(baseDirectory)
	if err != nil {
		ui.Fatal("Cannot read directory %q.", baseDirectory)
	}

	// get al list of all directories (including the base directory)
	projectDirectories := make([]string, 0)
	for _, entry := range baseDirectoryEntries {

		// check the directory
		if entry.IsDir() {
			subDirectoryPath := filepath.Join(baseDirectory, entry.Name())
			projectFilePath := filepath.Join(subDirectoryPath, projectFileName)

			// add the directory if it contains a project file
			if fs.FileExists(projectFilePath) {
				projectDirectories = append(projectDirectories, subDirectoryPath)
			}
		}

		// add the base directory if it contains a projct file
		if entry.Name() == projectFileName {
			projectDirectories = append(projectDirectories, baseDirectory)
		}

	}

	return projectDirectories
}
