// Copyright 2013 Andreas Koch. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mapping

import (
	"fmt"
	"github.com/andreaskoch/dotman/util/fs"
	"path/filepath"
	"regexp"
	"strings"
)

var (
	// the white space between the source path, target path and pattern
	pathMapEntrySeparatorPattern = regexp.MustCompile(`(?:\s{2,}|\t+)`)
)

func newPathMapEntry(baseDirectory, dotmanPathMapEntry string) (*pathMapEntry, error) {

	// find source and target path matching in the supplied map entry
	entries := pathMapEntrySeparatorPattern.Split(dotmanPathMapEntry, -1)
	if len(entries) < 2 {
		return nil, fmt.Errorf("%q is not a valid path map entry. A path map entry should conists if a source and target path and optionally a pattern, all separated by some whitespace.", dotmanPathMapEntry)
	}

	// source path
	sourcePath := filepath.Join(baseDirectory, normalizePathSpecification(entries[0]))

	// target path
	targetPath := expandPathVariables(normalizePathSpecification(entries[1]))

	// glob pattern
	var pattern *regexp.Regexp
	if len(entries) == 3 {
		patternText := strings.TrimSpace(entries[2])
		parsedPattern, err := getPattern(patternText)
		if err != nil {
			return nil, fmt.Errorf("%q is not a valid regular expression. Error: %s", patternText, err)
		}

		pattern = parsedPattern
	}

	return &pathMapEntry{
		source:  sourcePath,
		target:  targetPath,
		pattern: pattern,
	}, nil
}

type pathMapEntry struct {
	source  string
	target  string
	pattern *regexp.Regexp

	isReversed bool
}

func (entry *pathMapEntry) String() string {
	return fmt.Sprintf("%s → %s (Pattern: %s)", entry.source, entry.target, entry.pattern)
}

func (entry *pathMapEntry) HasPattern() bool {
	return entry.pattern != nil
}

func (entry *pathMapEntry) IsReversed() bool {
	return entry.isReversed
}

// Reverse the source and target path
func (entry *pathMapEntry) Reverse() *pathMapEntry {
	source := entry.source
	target := entry.target

	entry.source = target
	entry.target = source

	entry.isReversed = !entry.isReversed

	return entry
}

func (entry *pathMapEntry) GetInstructions() []*Instruction {

	// single instruction
	if !entry.HasPattern() {
		return []*Instruction{newInstruction(entry.source, entry.target)}
	}

	// multiple instructions
	instructions := make([]*Instruction, 0)

	// find all file which match the pattern
	sourceEntries := fs.GetMatchingDirectoryEntries(entry.source, entry.pattern)
	for _, sourceEntry := range sourceEntries {
		sourceEntryName := filepath.Base(sourceEntry)
		targetEntry := filepath.Join(entry.target, sourceEntryName)

		// add a new instruction
		instructions = append(instructions, newInstruction(sourceEntry, targetEntry))
	}

	return instructions

}

func getPattern(patternText string) (*regexp.Regexp, error) {

	pattern, err := regexp.Compile(patternText)
	if err != nil {
		return nil, err
	}

	return pattern, nil
}
