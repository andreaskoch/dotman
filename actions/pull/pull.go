// Copyright 2013 Andreas Koch. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pull

import (
	"github.com/andreaskoch/dotman/actions/base"
	"github.com/andreaskoch/dotman/ui"
	"github.com/andreaskoch/dotman/util/command"
)

const (
	ActionName        = "pull"
	ActionDescription = "Pull changes from the remote repository."
)

type Pull struct {
	baseDirectory            string
	moduleCollectionProvider base.ModulesProviderFunc
}

func New(baseDirectory string, moduleCollectionProvider base.ModulesProviderFunc) *Pull {
	return &Pull{
		baseDirectory:            baseDirectory,
		moduleCollectionProvider: moduleCollectionProvider,
	}
}

func (pull *Pull) Name() string {
	return ActionName
}

func (pull *Pull) Description() string {
	return ActionDescription
}

func (pull *Pull) Execute(arguments []string) {
	pull.execute(false, arguments)
}

func (pull *Pull) DryRun(arguments []string) {
	pull.execute(true, arguments)
}

func (pull *Pull) execute(executeADryRunOnly bool, arguments []string) {

	ui.Message("Pulling changes for your dotfile-repository.")
	if executeADryRunOnly {
		return
	}

	// pull changes in sub-module
	if err := gitPull(pull.baseDirectory); err != nil {
		ui.Message("%s", err)
	}
}

func gitPull(directory string) error {
	if err := command.Execute(directory, "git", "pull"); err != nil {
		return err
	}

	if err := command.Execute(directory, "git", "submodule", "foreach", "git", "pull"); err != nil {
		return err
	}

	if err := command.Execute(directory, "git", "submodule", "update", "--init"); err != nil {
		return err
	}

	return nil
}
