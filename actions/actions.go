// Copyright 2013 Andreas Koch. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package actions

import (
	"github.com/andreaskoch/dotman/actions/importer"
	"github.com/andreaskoch/dotman/actions/list"
	"github.com/andreaskoch/dotman/projects"
	"github.com/andreaskoch/dotman/ui"
)

var (
	availableActions = make([]ActionMetaData, 0)
)

func init() {

	// initialize the list of available actions
	availableActions = []ActionMetaData{
		NewActionInfo(list.ActionName, list.ActionDescription),
		NewActionInfo(importer.ActionName, importer.ActionDescription),
	}
}

func Get(workingDirectory string, actionName string, arguments []string) Action {

	// create a projects provider for the supplied working directory
	projectsProvider := func() *projects.Collection {
		return getProjectCollection(workingDirectory)
	}

	// detect which action is requested
	switch actionName {

	case list.ActionName:
		return list.New(projectsProvider)

	case importer.ActionName:
		return importer.New(projectsProvider)

	default:
		return nil // no matching found

	}

	panic("Unreachable")
}

func GetAll() []ActionMetaData {
	return availableActions
}

func getProjectCollection(workingDirectory string) *projects.Collection {
	projectCollection, err := projects.Load(workingDirectory)
	if err != nil {
		ui.Fatal("Unable to load projects. %s", err)
	}

	return projectCollection
}
