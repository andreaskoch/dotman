// Copyright 2013 Andreas Koch. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package clone

import (
	"fmt"
	"github.com/andreaskoch/dotman/ui"
	"github.com/andreaskoch/dotman/util/command"
	"strings"
)

const (
	ActionName        = "clone"
	ActionDescription = "Clone a dotfile repository."
)

type Clone struct {
	baseDirectory string
}

func New(baseDirectory string) *Clone {
	return &Clone{
		baseDirectory: baseDirectory,
	}
}

func (clone *Clone) Name() string {
	return ActionName
}

func (clone *Clone) Description() string {
	return ActionDescription
}

func (clone *Clone) Execute(arguments []string) {
	clone.execute(false, arguments)
}

func (clone *Clone) DryRun(arguments []string) {
	clone.execute(true, arguments)
}

func (clone *Clone) execute(executeADryRunOnly bool, arguments []string) {

	if len(arguments) == 0 {
		ui.Message("Please specifiy a repository path (e.g. git@bitbucket.org:andreaskoch/dotfiles-public.git).")
		return
	}

	// extract the repository url from the arguments
	repositoryUrl := ""
	if len(arguments) > 0 {
		repositoryUrl = strings.TrimSpace(arguments[0])
	}

	ui.Message("Cloning dotfile repository %q into %q.", repositoryUrl, clone.baseDirectory)
	if !executeADryRunOnly {
		command.Execute(clone.baseDirectory, fmt.Sprintf("git clone --recursive %s", repositoryUrl))
	}
}
