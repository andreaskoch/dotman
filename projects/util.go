// Copyright 2013 Andreas Koch. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package projects

import (
	"fmt"
	"github.com/andreaskoch/dotman/util/fs"
	"io/ioutil"
	"path/filepath"
)

func getAllProjectDirectories(baseDirectory, projectFileName string) ([]string, error) {

	baseDirectoryEntries, err := ioutil.ReadDir(baseDirectory)
	if err != nil {
		return []string{}, fmt.Errorf("Cannot read directory %q.", baseDirectory)
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

	return projectDirectories, nil
}
