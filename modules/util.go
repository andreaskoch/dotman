// Copyright 2013 Andreas Koch. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package modules

import (
	"fmt"
	"github.com/andreaskoch/dotman/util/fs"
	"io/ioutil"
	"path/filepath"
)

func getAllModuleDirectories(baseDirectory, moduleFileName string) ([]string, error) {

	baseDirectoryEntries, err := ioutil.ReadDir(baseDirectory)
	if err != nil {
		return []string{}, fmt.Errorf("Cannot read directory %q.", baseDirectory)
	}

	// get al list of all directories (including the base directory)
	moduleDirectories := make([]string, 0)
	for _, entry := range baseDirectoryEntries {

		// check the directory
		if entry.IsDir() {
			subDirectoryPath := filepath.Join(baseDirectory, entry.Name())
			moduleFilePath := filepath.Join(subDirectoryPath, moduleFileName)

			// add the directory if it contains a module file
			if fs.FileExists(moduleFilePath) {
				moduleDirectories = append(moduleDirectories, subDirectoryPath)
			}
		}

		// add the base directory if it contains a projct file
		if entry.Name() == moduleFileName {
			moduleDirectories = append(moduleDirectories, baseDirectory)
		}

	}

	return moduleDirectories, nil
}
