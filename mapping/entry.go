// Copyright 2013 Andreas Koch. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mapping

import (
	"fmt"
	"path/filepath"
	"regexp"
	"strings"
)

var (
	PathMapEntrySeparatorPattern = regexp.MustCompile(`(?:\s{2,}|\t+)`)
)

func newPathMapEntry(baseDirectory, dotmanPathMapEntry string) (*PathMapEntry, error) {

	// find source and target path matching in the supplied map entry
	entries := PathMapEntrySeparatorPattern.Split(dotmanPathMapEntry, -1)
	if len(entries) < 2 {
		return nil, fmt.Errorf("%q is not a valid path map entry.", dotmanPathMapEntry)
	}

	// source path
	sourcePath := filepath.Join(baseDirectory, normalizePathSpecification(entries[0]))

	// target path
	targetPath := expandPathVariables(normalizePathSpecification(entries[1]))

	// glob pattern
	globPattern := ""
	if len(entries) == 3 {
		globPattern = strings.TrimSpace(entries[2])
	}

	return &PathMapEntry{
		Source:  sourcePath,
		Target:  targetPath,
		Pattern: globPattern,
	}, nil
}

type PathMapEntry struct {
	Source  string
	Target  string
	Pattern string
}

func (pathMapEntry *PathMapEntry) String() string {
	return fmt.Sprintf("%s → %s (Pattern: %s)", pathMapEntry.Source, pathMapEntry.Target, pathMapEntry.Pattern)
}
