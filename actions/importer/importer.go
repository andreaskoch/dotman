// Copyright 2013 Andreas Koch. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package importer

import (
	"github.com/andreaskoch/dotman/actions/base"
	"github.com/andreaskoch/dotman/projects"
	"github.com/andreaskoch/dotman/ui"
	"github.com/andreaskoch/dotman/util/fs"
	"path/filepath"
	"regexp"
)

const (
	ActionName        = "import"
	ActionDescription = "Import files based on your current dotman configurations."
)

type Importer struct {
	*base.Action
}

func New(projectCollectionProvider func() *projects.Collection) *Importer {
	return &Importer{
		base.New(ActionName, ActionDescription, projectCollectionProvider, func(project *projects.Project, executeADryRunOnly bool) {
			ui.Message("\nImporting %q:", project)
			importProject(project, executeADryRunOnly)
		}),
	}
}

func importProject(project *projects.Project, executeADryRunOnly bool) {

	// build the copy-map
	for _, pathMapEntry := range project.Map.Entries {

		source := pathMapEntry.Target
		target := pathMapEntry.Source
		patternText := pathMapEntry.Pattern

		// copy one or more files
		if patternText != "" {

			pattern, err := regexp.Compile(patternText)
			if err != nil {
				ui.Fatal("%s", err)
			}

			sourceEntries := fs.GetMatchingDirectoryEntries(source, pattern)
			for _, sourceEntry := range sourceEntries {
				sourceEntryName := filepath.Base(sourceEntry)
				targetEntry := filepath.Join(target, sourceEntryName)
				copy(sourceEntry, targetEntry, executeADryRunOnly)
			}

			continue

		}

		// copy a single file or folder
		copy(source, target, executeADryRunOnly)
	}
}

func copy(source, target string, executeADryRunOnly bool) {
	ui.Message("Copy %s → %s", source, target)
	if !executeADryRunOnly {
		if _, err := fs.Copy(source, target); err != nil {
			ui.Message("%s", err)
		}
	}
}
