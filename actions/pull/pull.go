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

	// pull changes in the main repository
	if err := command.Execute(pull.baseDirectory, "git", "pull", "origin", "master"); err != nil {
		ui.Fatal("Error while pull changes for the main repository:\n%s", err)
	}

	// pull changes for all modules
	modules := pull.moduleCollectionProvider()
	for _, module := range modules.Collection {
		// pull changes in sub-module
		if err := gitPull(module.Directory()); err != nil {
			ui.Message("Error while updating module %s:\n%s", module, err)
		}
	}
}

func gitPull(directory string) error {

	// discard changes
	if err := command.Execute(directory, "git", "checkout", "."); err != nil {
		return err
	}

	// switch to master branch
	if err := command.Execute(directory, "git", "checkout", "master"); err != nil {
		return err
	}

	// pull changes
	if err := command.Execute(directory, "git", "pull", "origin", "master"); err != nil {
		return err
	}

	// update submodules
	if err := command.Execute(directory, "git", "submodule", "update", "--init"); err != nil {
		return err
	}

	return nil
}
