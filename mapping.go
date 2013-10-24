// Copyright 2013 Andreas Koch. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

var (
	PathMapEntryPattern = regexp.MustCompile(`^(\S.+)\s{2,}(\S.+)$`)
)

type PathMap struct {
	Paths []PathMapEntry
}

type PathMapEntry struct {
	Source SourcePath
	Target TargetPath
}

type SourcePath struct {
}

type TargetPath struct {
}

func NewPathMap(sourceFile string) (PathMap, error) {

	// open the dotman file
	file, err := os.Open(sourceFile)
	if err != nil {
		return PathMap{}, err
	}

	// read in the lines of the dotman file and create path map entries from it
	pathMapEntries := make([]PathMapEntry, 0)
	lines := GetLines(file)
	for _, line := range lines {

		// ignore white space and comments
		if isEmptyLine(line) || isComment(line) {
			continue
		}

		// create a path map entry from the line
		pathMapEntry, err := newPathMapEntry(line)
		if err != nil {
			message("Unable to read path map entry %q. Error: %s", line, err)
			continue
		}

		// append the path map entry to the list
		pathMapEntries = append(pathMapEntries, pathMapEntry)
	}

	return PathMap{}, nil
}

func newPathMapEntry(dotmanPathMapEntry string) (PathMapEntry, error) {

	// find source and target path matching in the supplied map entry
	matches := PathMapEntryPattern.FindStringSubmatch(dotmanPathMapEntry)
	if len(matches) < 3 {
		return PathMapEntry{}, fmt.Errorf("%q is not a valid path map entry", dotmanPathMapEntry)
	}

	// parse the source path
	sourcePathText := strings.TrimSpace(matches[1])
	sourcePath, err := newSourcePath(sourcePathText)
	if err != nil {
		return PathMapEntry{}, fmt.Errorf("%q is not a valid source path. Error: %s", sourcePathText, err)
	}

	// parse the target path
	targetPathText := strings.TrimSpace(matches[2])
	targetPath, err := newTargetPath(targetPathText)
	if err != nil {
		return PathMapEntry{}, fmt.Errorf("%q is not a valid target path. Error: %s", targetPathText, err)
	}

	return PathMapEntry{
		Source: sourcePath,
		Target: targetPath,
	}, nil
}

func newSourcePath(sourcePathText string) (SourcePath, error) {
	return SourcePath{}, nil
}

func newTargetPath(targetPathText string) (TargetPath, error) {
	return TargetPath{}, nil
}

func isEmptyLine(line string) bool {
	return strings.TrimSpace(line) == ""
}

func isComment(line string) bool {
	return strings.HasPrefix(strings.TrimSpace(line), "#")
}
