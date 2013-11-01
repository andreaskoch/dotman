// Copyright 2013 Andreas Koch. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package importer

import (
	"github.com/andreaskoch/dotman/projects"
	"github.com/andreaskoch/dotman/ui"
	"github.com/andreaskoch/dotman/util/fs"
	"regexp"
	"strings"
)

const (
	ActionName        = "import"
	ActionDescription = "Import files based on your current dotman configurations."
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

func (importer *Importer) Execute(arguments []string) {
	importer.execute(false, arguments)
}

func (importer *Importer) DryRun(arguments []string) {
	importer.execute(true, arguments)
}

func (importer *Importer) execute(executeADryRunOnly bool, arguments []string) {

	// extract the project filter from the arguments
	projectFilter := regexp.MustCompile(`.*`)
	if len(arguments) > 0 && strings.TrimSpace(arguments[0]) != "" {

		// get the custom project filter from the command arguments
		projectFilterText := strings.TrimSpace(arguments[0])

		// try to compile the filter
		customProjectFilter, err := regexp.Compile(projectFilterText)
		if err != nil {
			ui.Fatal("%q is not a valid project filter. Error: %s", projectFilterText, err.Error())
		}

		// assign the supplied custom filter
		projectFilter = customProjectFilter
	}

	projects := importer.projectCollectionProvider()
	for _, project := range projects.Collection {

		// skip projects which don't match the filter
		if !projectFilter.MatchString(project.String()) {
			continue
		}

		ui.Message("Importing %q", project)
		importProject(project, executeADryRunOnly)
	}

}

func importProject(project *projects.Project, executeADryRunOnly bool) {

	for _, entry := range project.Map.Entries {
		source := entry.Target.Path()
		target := entry.Source.Path()

		ui.Message("Copy: %s → %s", source, target)
		if !executeADryRunOnly {
			if _, err := fs.Copy(source, target); err != nil {
				ui.Message("%s", err)
			}
		}
	}
}
