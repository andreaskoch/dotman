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

	// check for wildcard
	if isWildcard, wildcardBaseDirectory := isWildcardSpecification(specification); isWildcard {

		// assemble to absolute wildcard base path
		wildcardBasePath := filepath.Join(baseDirectory, wildcardBaseDirectory)

		return &SourcePath{
			Files: getAllFilesInDirectory(wildcardBasePath),
		}, nil
	}

	fullPath := filepath.Join(baseDirectory, specification)
	return &SourcePath{
		Files: []string{fullPath},
	}, nil
}

type SourcePath struct {
	Files []string
}

func (sourcePath *SourcePath) String() string {
	return strings.Join(sourcePath.Files, ", ")
}
