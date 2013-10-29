// Copyright 2013 Andreas Koch. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mapping

import (
	"fmt"
	"path/filepath"
	"strings"
)

func newTargetPath(targetPath string) (*TargetPath, error) {

	// validate the source path specification
	if strings.TrimSpace(targetPath) == "" {
		return nil, fmt.Errorf("Empty specification.")
	}

	// normalize the path specification
	targetPath = normalizePathSpecification(targetPath)

	// expand path variables such as ~/ or $HOME
	targetPath = expandPathVariables(targetPath)

	// abort if the path is not absolute
	if !filepath.IsAbs(targetPath) {
		return nil, fmt.Errorf("Target path is not absolute.")
	}

	return &TargetPath{
		path: targetPath,
	}, nil
}

type TargetPath struct {
	path string
}

func (targetPath *TargetPath) Path() string {
	return targetPath.path
}
