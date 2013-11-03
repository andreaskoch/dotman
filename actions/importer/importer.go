// Copyright 2013 Andreas Koch. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package importer

import (
	"github.com/andreaskoch/dotman/actions/base"
	"github.com/andreaskoch/dotman/mapping"
	"github.com/andreaskoch/dotman/projects"
	"github.com/andreaskoch/dotman/ui"
	"github.com/andreaskoch/dotman/util/fs"
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

	for _, pathMapEntry := range project.Map.Entries {

		source := pathMapEntry.Source
		target := pathMapEntry.Target

		ui.Message("Copy %s → %s", source, target)
		if !executeADryRunOnly {
			if _, err := fs.Copy(source, target); err != nil {
				ui.Message("%s", err)
			}
		}
	}
}
