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
		return nil, fmt.Errorf("Cannot create a path map because the specified dotfile %q does not exist.", sourceFile)
	}

	// open the dotman file
	file, err := os.Open(sourceFile)
	if err != nil {
		return nil, err // unable to read the supplied source file
	}

	// determine the path map directory
	directory := filepath.Dir(sourceFile)

	pathMapEntries := make([]*pathMapEntry, 0)

	// read in the lines of the dotman file and create path map entries from it
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
		directory: directory,
		entries:   pathMapEntries,
	}, nil
}

type PathMap struct {
	directory string
	entries   []*pathMapEntry

	isReversed bool
}

func (pathMap *PathMap) IsReversed() bool {
	return pathMap.isReversed
}

// Reverse the source and target path of all entries in this map
func (pathMap *PathMap) Reverse() *PathMap {
	for _, entry := range pathMap.entries {
		entry.Reverse()
	}

	pathMap.isReversed = !pathMap.isReversed

	return pathMap
}

func (pathMap *PathMap) GetInstructions() []*Instruction {

	instructions := make([]*Instruction, 0)

	// get the instructions for all entries
	for _, entry := range pathMap.entries {
		instructions = append(instructions, entry.GetInstructions()...)
	}

	return instructions
}
