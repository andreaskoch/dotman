// Copyright 2013 Andreas Koch. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package actions

import (
	"github.com/andreaskoch/dotman/actions/backup"
	"github.com/andreaskoch/dotman/actions/changes"
	"github.com/andreaskoch/dotman/actions/clone"
	"github.com/andreaskoch/dotman/actions/commit"
	"github.com/andreaskoch/dotman/actions/deploy"
	"github.com/andreaskoch/dotman/actions/importer"
	"github.com/andreaskoch/dotman/actions/list"
	"github.com/andreaskoch/dotman/actions/pull"
	"github.com/andreaskoch/dotman/actions/push"
	"github.com/andreaskoch/dotman/modules"
	"github.com/andreaskoch/dotman/ui"
)

var (
	availableActions = make([]ActionMetaData, 0)
)

func init() {

	// initialize the list of available actions
	availableActions = []ActionMetaData{
		NewActionInfo(clone.ActionName, clone.ActionDescription),
		NewActionInfo(importer.ActionName, importer.ActionDescription),
		NewActionInfo(list.ActionName, list.ActionDescription),
		NewActionInfo(backup.ActionName, backup.ActionDescription),
		NewActionInfo(changes.ActionName, changes.ActionDescription),
		NewActionInfo(deploy.ActionName, deploy.ActionDescription),
		NewActionInfo(commit.ActionName, commit.ActionDescription),
		NewActionInfo(push.ActionName, push.ActionDescription),
		NewActionInfo(pull.ActionName, pull.ActionDescription),
	}
}

func Get(workingDirectory string, actionName string) Action {

	// create a modules provider for the supplied working directory
	modulesProvider := func() *modules.Collection {
		return getModuleCollection(workingDirectory)
	}

	// detect which action is requested
	switch actionName {

	case clone.ActionName:
		return clone.New(workingDirectory)

	case list.ActionName:
		return list.New(modulesProvider)

	case importer.ActionName:
		return importer.New(modulesProvider)

	case backup.ActionName:
		return backup.New(modulesProvider)

	case deploy.ActionName:
		return deploy.New(modulesProvider)

	case changes.ActionName:
		return changes.New(modulesProvider)

	case commit.ActionName:
		return commit.New(workingDirectory, modulesProvider)

	case push.ActionName:
		return push.New(workingDirectory, modulesProvider)

	case pull.ActionName:
		return pull.New(workingDirectory, modulesProvider)

	default:
		return nil // no matching found

	}

	panic("Unreachable")
}

func GetAll() []ActionMetaData {
	return availableActions
}

func getModuleCollection(workingDirectory string) *modules.Collection {
	moduleCollection, err := modules.Load(workingDirectory)
	if err != nil {
		ui.Fatal("Unable to load modules. %s", err)
	}

	return moduleCollection
}
