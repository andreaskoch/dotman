// Copyright 2013 Andreas Koch. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package deploy

import (
	"github.com/andreaskoch/dotman/actions/base"
	"github.com/andreaskoch/dotman/modules"
	"github.com/andreaskoch/dotman/ui"
	"github.com/andreaskoch/dotman/util/fs"
)

const (
	ActionName        = "deploy"
	ActionDescription = "Deploy your modules."
)

type Deploy struct {
	*base.Action
}

func New(moduleCollectionProvider base.ModulesProviderFunc) *Deploy {
	return &Deploy{
		base.New(ActionName, ActionDescription, moduleCollectionProvider, func(module *modules.Module, executeADryRunOnly bool) {
			ui.Message("Deploying %q", module)
			deployModule(module, executeADryRunOnly)
		}),
	}
}

func deployModule(module *modules.Module, executeADryRunOnly bool) {

	for _, instruction := range module.Map.GetInstructions() {

		source := instruction.Source()
		target := instruction.Target()

		ui.Message("Copy %s → %s", source, target)
		if !executeADryRunOnly {
			if _, err := fs.Copy(source, target); err != nil {
				ui.Message("%s", err)
			}
		}
	}
}
