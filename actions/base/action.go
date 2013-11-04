// Copyright 2013 Andreas Koch. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package base

import (
	"github.com/andreaskoch/dotman/projects"
	"github.com/andreaskoch/dotman/ui"
	"regexp"
	"strings"
)

type ForEachProjectFunc func(project *projects.Project, executeADryRunOnly bool)

type ProjectsProviderFunc func() *projects.Collection

type Action struct {
	name                      string
	description               string
	projectCollectionProvider ProjectsProviderFunc
	forEachProject            ForEachProjectFunc
}

func New(name, description string, projectCollectionProvider ProjectsProviderFunc, forEachProject ForEachProjectFunc) *Action {
	return &Action{
		name:                      name,
		description:               description,
		projectCollectionProvider: projectCollectionProvider,
		forEachProject:            forEachProject,
	}
}

func (action *Action) Name() string {
	return action.name
}

func (action *Action) Description() string {
	return action.description
}

func (action *Action) Execute(arguments []string) {
	action.execute(false, arguments)
}

func (action *Action) DryRun(arguments []string) {
	action.execute(true, arguments)
}

func (action *Action) execute(executeADryRunOnly bool, arguments []string) {

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

	projects := action.projectCollectionProvider()
	for _, project := range projects.Collection {

		// skip projects which don't match the filter
		if !projectFilter.MatchString(project.String()) {
			continue
		}

		action.forEachProject(project, executeADryRunOnly)
	}

}
