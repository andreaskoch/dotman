// Copyright 2013 Andreas Koch. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mapping

import (
	"fmt"
	"github.com/andreaskoch/dotman/ui"
	"github.com/andreaskoch/dotman/util/fs"
	"os"
	"path/filepath"
)

func NewPathMap(sourceFile string) (*PathMap, error) {

	// check if the source file exists
	if !fs.IsFile(sourceFile) {
		return nil, fmt.Errorf("Cannot create a path map because the specified dotfile %q does not exist", sourceFile)
	}

	// open the dotman file
	file, err := os.Open(sourceFile)
	if err != nil {
		return nil, err // unable to read the supplied source file
	}

	// determine the path map directory
	directory := filepath.Dir(sourceFile)

	// read in the lines of the dotman file and create path map entries from it
	pathMapEntries := make([]PathMapEntry, 0)
	lines := fs.GetLines(file)
	for lineNumber, line := range lines {

		// ignore white space and comments
		if isEmptyLine(line) || isComment(line) {
			continue
		}

		// create a path map entry from the line
		pathMapEntry, err := newPathMapEntry(directory, line)
		if err != nil {
			ui.Message("Line %d: %s", lineNumber+1, err)
			continue
		}

		// append the path map entry to the list
		pathMapEntries = append(pathMapEntries, pathMapEntry)
	}

	return &PathMap{
		Directory: directory,
		Entries:   pathMapEntries,
	}, nil
}

type PathMap struct {
	Directory string
	Entries   []PathMapEntry
}

func (pathMap *PathMap) String() string {
	text := ""

	for _, entry := range pathMap.Entries {
		text += fmt.Sprintf("%s", entry.String())
	}

	return text
}
