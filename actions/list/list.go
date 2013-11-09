// Copyright 2013 Andreas Koch. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package list

import (
	"github.com/andreaskoch/dotman/actions/base"
	"github.com/andreaskoch/dotman/modules"
	"github.com/andreaskoch/dotman/ui"
)

const (
	ActionName        = "list"
	ActionDescription = "Get a list of all modules in the current dotfile collection."
)

type List struct {
	*base.Action
}

func New(moduleCollectionProvider base.ModulesProviderFunc) *List {
	return &List{
		base.New(ActionName, ActionDescription, moduleCollectionProvider, func(module *modules.Module, executeADryRunOnly bool) {
			ui.Message("%s", module)
		}),
	}
}
