// Copyright 2013 Andreas Koch. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package list

import (
	"github.com/andreaskoch/dotman/actions/base"
	"github.com/andreaskoch/dotman/projects"
	"github.com/andreaskoch/dotman/ui"
)

const (
	ActionName        = "list"
	ActionDescription = "Get a list of all projects in the current dotfile collection."
)

type List struct {
	*base.Action
}

func New(projectCollectionProvider base.ProjectsProviderFunc) *List {
	return &List{
		base.New(ActionName, ActionDescription, projectCollectionProvider, func(project *projects.Project, executeADryRunOnly bool) {
			ui.Message("%s", project)
		}),
	}
}
