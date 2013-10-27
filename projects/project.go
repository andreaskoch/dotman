// Copyright 2013 Andreas Koch. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package projects

import (
	"fmt"
	"github.com/andreaskoch/dotman/mapping"
	"path/filepath"
)

type Project struct {
	Directory string
	Map       *mapping.PathMap
}

func newProject(directory, projectFileName string) (*Project, error) {

	projectPathMap, err := mapping.NewPathMap(filepath.Join(directory, projectFileName))
	if err != nil {
		return nil, fmt.Errorf("Unable to read dotman file. %s", err)
	}

	return &Project{
		Directory: directory,
		Map:       projectPathMap,
	}, nil
}
