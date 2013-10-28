// Copyright 2013 Andreas Koch. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package importer

import (
	"github.com/andreaskoch/dotman/projects"
	"github.com/andreaskoch/dotman/ui"
	"github.com/andreaskoch/dotman/util/fs"
)

const (
	ActionName        = "import"
	ActionDescription = "Import your current configuration files."
)

type Importer struct {
	projectCollectionProvider func() *projects.Collection
}

func New(projectCollectionProvider func() *projects.Collection) *Importer {
	return &Importer{
		projectCollectionProvider: projectCollectionProvider,
	}
}

func (importer *Importer) Name() string {
	return ActionName
}

func (importer *Importer) Description() string {
	return ActionDescription
}

func (importer *Importer) Execute() {
	projects := importer.projectCollectionProvider()
	for _, project := range projects.Collection {
		ui.Message("Importing %q", project)
		importProject(project)
	}
}

func importProject(project *projects.Project) {

	for _, entry := range project.Map.Entries {
		sourceFile := entry.Target.String()
		targetFile := entry.Source.String()

		ui.Message("Copy %q to %q.", sourceFile, targetFile)
		if true {
			if _, err := fs.Copy(sourceFile, targetFile); err != nil {
				ui.Message("%s", err)
			}
		}
	}
}
