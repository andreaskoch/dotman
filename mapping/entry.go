// Copyright 2013 Andreas Koch. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mapping

import (
	"fmt"
	"path/filepath"
)

func newPathMapEntries(baseDirectory, dotmanPathMapEntry string) ([]*PathMapEntry, error) {

	entries := make([]*PathMapEntry, 0)

	// find source and target path matching in the supplied map entry
	matches := PathMapEntryPattern.FindStringSubmatch(dotmanPathMapEntry)
	if len(matches) < 3 {
		return entries, fmt.Errorf("%q is not a valid path map entry.", dotmanPathMapEntry)
	}

	sourcePathText := normalizePathSpecification(matches[1])
	targetPathText := normalizePathSpecification(matches[2])

	// prepare the source path
	sourcePath := filepath.Join(baseDirectory, sourcePathText)

	// prepare the target path
	targetPath := expandPathVariables(targetPathText)

	// glob pattern matching
	matches, err := filepath.Glob(sourcePath)
	if err != nil {
		return entries, err
	}

	for _, sourcePath := range matches {

		// determine the target path
		sourceEntryName := filepath.Base(sourcePath)
		targetPath := filepath.Join(targetPath, sourceEntryName)

		// add a new path map entry
		entries = append(entries, &PathMapEntry{
			Source: sourcePath,
			Target: targetPath,
		})
	}

	return entries, nil
}

type PathMapEntry struct {
	Source string
	Target string
}

func (pathMapEntry *PathMapEntry) String() string {
	return fmt.Sprintf("%s → %s\n", pathMapEntry.Source, pathMapEntry.Target)
}
