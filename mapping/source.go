// Copyright 2013 Andreas Koch. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mapping

import (
	"fmt"
	"path/filepath"
	"strings"
)

func newSourcePath(baseDirectory, specification string) (*SourcePath, error) {

	// validate the source path specification
	if strings.TrimSpace(specification) == "" {
		return nil, fmt.Errorf("Empty specification.")
	}

	// normalize the path specification
	specification = normalizePathSpecification(specification)

	fullPath := filepath.Join(baseDirectory, specification)
	return &SourcePath{
		path: fullPath,
	}, nil
}

type SourcePath struct {
	path string
}

func (sourcePath *SourcePath) Path() string {
	return sourcePath.path
}
