// Copyright 2013 Andreas Koch. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package base

import (
	"github.com/andreaskoch/dotman/modules"
	"github.com/andreaskoch/dotman/ui"
	"regexp"
	"strings"
)

type ForEachModuleFunc func(module *modules.Module, executeADryRunOnly bool)

type ModulesProviderFunc func() *modules.Collection

type Action struct {
	name                     string
	description              string
	moduleCollectionProvider ModulesProviderFunc
	forEachModule            ForEachModuleFunc
}

func New(name, description string, moduleCollectionProvider ModulesProviderFunc, forEachModule ForEachModuleFunc) *Action {
	return &Action{
		name:                     name,
		description:              description,
		moduleCollectionProvider: moduleCollectionProvider,
		forEachModule:            forEachModule,
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

	// extract the module filter from the arguments
	moduleFilter := regexp.MustCompile(`.*`)
	if len(arguments) > 0 && strings.TrimSpace(arguments[0]) != "" {

		// get the custom module filter from the command arguments
		moduleFilterText := strings.TrimSpace(arguments[0])

		// try to compile the filter
		customModuleFilter, err := regexp.Compile(moduleFilterText)
		if err != nil {
			ui.Fatal("%q is not a valid module filter. Error: %s", moduleFilterText, err.Error())
		}

		// assign the supplied custom filter
		moduleFilter = customModuleFilter
	}

	modules := action.moduleCollectionProvider()
	for _, module := range modules.Collection {

		// skip modules which don't match the filter
		if !moduleFilter.MatchString(module.String()) {
			continue
		}

		action.forEachModule(module, executeADryRunOnly)
	}

}
