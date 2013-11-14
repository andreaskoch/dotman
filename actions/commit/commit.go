// Copyright 2013 Andreas Koch. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package commit

import (
	"github.com/andreaskoch/dotman/actions/base"
	"github.com/andreaskoch/dotman/ui"
	"github.com/andreaskoch/dotman/util/command"
	"strings"
)

const (
	ActionName        = "commit"
	ActionDescription = "Commit all changes."
)

type Commit struct {
	baseDirectory            string
	moduleCollectionProvider base.ModulesProviderFunc
}

func New(baseDirectory string, moduleCollectionProvider base.ModulesProviderFunc) *Commit {
	return &Commit{
		baseDirectory:            baseDirectory,
		moduleCollectionProvider: moduleCollectionProvider,
	}
}

func (commit *Commit) Name() string {
	return ActionName
}

func (commit *Commit) Description() string {
	return ActionDescription
}

func (commit *Commit) Execute(arguments []string) {
	commit.execute(false, arguments)
}

func (commit *Commit) DryRun(arguments []string) {
	commit.execute(true, arguments)
}

func (commit *Commit) execute(executeADryRunOnly bool, arguments []string) {

	if len(arguments) == 0 {
		ui.Message("Please specifiy a commit message.")
		return
	}

	// extract the repository url from the arguments
	commitMessage := ""
	if len(arguments) > 0 {
		commitMessage = strings.TrimSpace(strings.Join(arguments[0:], ""))
		commitMessage = strings.Trim(commitMessage, `"'`)
	}

	// commit all submodules
	modules := commit.moduleCollectionProvider()
	for _, module := range modules.Collection {

		ui.Message("Commiting changes in sub-module %q.", module)

		// abort if dry
		if executeADryRunOnly {
			continue
		}

		// commit changes in sub-module
		if err := gitCommit(module.Directory(), commitMessage); err != nil {
			ui.Message("%s", err)
		}
	}

	ui.Message("Committing changes to dotfile repository.")
	if executeADryRunOnly {
		return
	}

	// commit changes in sub-module
	if err := gitCommit(commit.baseDirectory, commitMessage); err != nil {
		ui.Message("%s", err)
	}
}

func gitCommit(directory, message string) error {
	// get the current status
	if err := command.Execute(directory, "git", "status"); err != nil {
		return err
	}

	// add all changes to the index
	if err := command.Execute(directory, "git", "add", "-A", "."); err != nil {
		return err
	}

	// commit the changes
	if err := command.Execute(directory, "git", "commit","-m", message); err != nil {
		return err
	}

	return nil
}
