// Copyright 2013 Andreas Koch. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var (
	PathMapEntryPattern = regexp.MustCompile(`^(\S.+)\s{2,}(\S.+)$`)

	DirectorySeparatorPattern = regexp.MustCompile(`[\/]{1,}`)
)

type PathMap struct {
	Directory string
	Paths     []PathMapEntry
}

type PathMapEntry struct {
	Source *SourcePath
	Target *TargetPath
}

type SourcePath struct {
	Files []string
}

type TargetPath struct {
	Files []string
}

func NewPathMap(sourceFile string) (*PathMap, error) {

	// check if the source file exists
	if !IsFile(sourceFile) {
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
	lines := GetLines(file)
	for lineNumber, line := range lines {

		// ignore white space and comments
		if isEmptyLine(line) || isComment(line) {
			continue
		}

		// create a path map entry from the line
		pathMapEntry, err := newPathMapEntry(directory, line)
		if err != nil {
			message("Line %d: %s", lineNumber+1, err)
			continue
		}

		// append the path map entry to the list
		pathMapEntries = append(pathMapEntries, pathMapEntry)
	}

	return &PathMap{
		Directory: directory,
		Paths:     pathMapEntries,
	}, nil
}

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

func newSourcePath(baseDirectory, specification string) (*SourcePath, error) {

	// check if the base directory exists
	if !IsDirectory(baseDirectory) {
		return nil, fmt.Errorf("The specified base directory %q does not exists.", baseDirectory)
	}

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

		// check if the base path is an existing directory
		if !IsDirectory(wildcardBasePath) {
			return nil, fmt.Errorf("%q does not exist.", wildcardBasePath)
		}

		return &SourcePath{
			Files: GetAllFiles(wildcardBaseDirectory),
		}, nil
	}

	// check if the specified file or directory exists
	fullPath := filepath.Join(baseDirectory, specification)
	if PathExists(fullPath) {
		return &SourcePath{
			Files: []string{fullPath},
		}, nil
	}

	// the specification is invalid
	return nil, fmt.Errorf("%q does not exist.", specification)
}

func isWildcardSpecification(path string) (isWildcard bool, wildcardBaseDirectory string) {

	// split the path into its components
	components := strings.Split(path, string(os.PathSeparator))

	// check if the last path component is a star
	lastComponent := components[len(components)-1]
	isWildcard = lastComponent == "*"

	// determine the wild base directory
	wildcardBaseDirectory = ""
	if isWildcard && len(components) > 1 {
		indexOfLastComponentBeforeWildcard := len(components) - 2
		wildcardBaseDirectory = strings.Join(components[0:indexOfLastComponentBeforeWildcard], string(os.PathSeparator))
	}

	return isWildcard, wildcardBaseDirectory
}

func normalizePathSpecification(path string) string {

	// trim leading and trailing white space
	path = strings.TrimSpace(path)

	// replace all directory separators with the ones for the current platform
	path = DirectorySeparatorPattern.ReplaceAllString(path, string(os.PathSeparator))

	return path
}

func newTargetPath(targetPathText string) (*TargetPath, error) {
	return &TargetPath{}, nil
}

func isEmptyLine(line string) bool {
	return strings.TrimSpace(line) == ""
}

func isComment(line string) bool {
	return strings.HasPrefix(strings.TrimSpace(line), "#")
}
