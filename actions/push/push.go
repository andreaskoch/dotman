// Copyright 2013 Andreas Koch. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package push

import (
	"github.com/andreaskoch/dotman/actions/base"
	"github.com/andreaskoch/dotman/ui"
	"github.com/andreaskoch/dotman/util/command"
)

const (
	ActionName        = "push"
	ActionDescription = "Push all commits to their remote repository."
)

type Push struct {
	baseDirectory            string
	moduleCollectionProvider base.ModulesProviderFunc
}

func New(baseDirectory string, moduleCollectionProvider base.ModulesProviderFunc) *Push {
	return &Push{
		baseDirectory:            baseDirectory,
		moduleCollectionProvider: moduleCollectionProvider,
	}
}

func (push *Push) Name() string {
	return ActionName
}

func (push *Push) Description() string {
	return ActionDescription
}

func (push *Push) Execute(arguments []string) {
	push.execute(false, arguments)
}

func (push *Push) DryRun(arguments []string) {
	push.execute(true, arguments)
}

func (push *Push) execute(executeADryRunOnly bool, arguments []string) {

	// push all submodules
	modules := push.moduleCollectionProvider()
	for _, module := range modules.Collection {

		ui.Message("Pushing changes in sub-module %q.", module)

		// abort if dry
		if executeADryRunOnly {
			continue
		}

		// push changes in sub-module
		if err := gitPush(module.Directory()); err != nil {
			ui.Message("%s", err)
		}
	}

	ui.Message("Pushing dotfile repository.")
	if executeADryRunOnly {
		return
	}

	// push changes in sub-module
	if err := gitPush(push.baseDirectory); err != nil {
		ui.Message("%s", err)
	}
}

func gitPush(directory string) error {
	if err := command.Execute(directory, "git", "push"); err != nil {
		return err
	}

	return nil
}
