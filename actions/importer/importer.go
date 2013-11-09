// Copyright 2013 Andreas Koch. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package importer

import (
	"github.com/andreaskoch/dotman/actions/base"
	"github.com/andreaskoch/dotman/modules"
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

func New(moduleCollectionProvider base.ModulesProviderFunc) *Importer {
	return &Importer{
		base.New(ActionName, ActionDescription, moduleCollectionProvider, func(module *modules.Module, executeADryRunOnly bool) {
			ui.Message("\nImporting %q:", module)
			importModule(module, executeADryRunOnly)
		}),
	}
}

func importModule(module *modules.Module, executeADryRunOnly bool) {

	for _, instruction := range module.Map.Reverse().GetInstructions() {

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
