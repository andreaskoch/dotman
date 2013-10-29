// Copyright 2013 Andreas Koch. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mapping

import (
	"fmt"
	"strings"
)

func newPathMapEntry(baseDirectory, dotmanPathMapEntry string) (PathMapEntry, error) {

	// find source and target path matching in the supplied map entry
	matches := PathMapEntryPattern.FindStringSubmatch(dotmanPathMapEntry)
	if len(matches) < 3 {
		return PathMapEntry{}, fmt.Errorf("%q is not a valid path map entry.", dotmanPathMapEntry)
	}

	// parse the source path
	sourcePathText := strings.TrimSpace(matches[1])
	sourcePath, err := newSourcePath(baseDirectory, sourcePathText)
	if err != nil {
		return PathMapEntry{}, fmt.Errorf("%s", err)
	}

	// parse the target path
	targetPathText := strings.TrimSpace(matches[2])
	targetPath, err := newTargetPath(targetPathText)
	if err != nil {
		return PathMapEntry{}, fmt.Errorf("%s", err)
	}

	return PathMapEntry{
		Source: sourcePath,
		Target: targetPath,
	}, nil
}

type PathMapEntry struct {
	Source *SourcePath
	Target *TargetPath
}

func (pathMapEntry *PathMapEntry) String() string {
	return fmt.Sprintf("%s → %s\n", pathMapEntry.Source.Path(), pathMapEntry.Target.Path())
}
