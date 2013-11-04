// Copyright 2013 Andreas Koch. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package deploy

import (
	"github.com/andreaskoch/dotman/actions/base"
	"github.com/andreaskoch/dotman/projects"
	"github.com/andreaskoch/dotman/ui"
	"github.com/andreaskoch/dotman/util/fs"
)

const (
	ActionName        = "deploy"
	ActionDescription = "Deploy your projects."
)

type Deploy struct {
	*base.Action
}

func New(projectCollectionProvider base.ProjectsProviderFunc) *Deploy {
	return &Deploy{
		base.New(ActionName, ActionDescription, projectCollectionProvider, func(project *projects.Project, executeADryRunOnly bool) {
			ui.Message("Deploying %q", project)
			deployProject(project, executeADryRunOnly)
		}),
	}
}

func deployProject(project *projects.Project, executeADryRunOnly bool) {

	for _, entry := range project.Map.Entries {
		source := entry.Source
		target := entry.Target

		ui.Message("Copy: %s → %s", source, target)
		if !executeADryRunOnly {
			if _, err := fs.Copy(source, target); err != nil {
				ui.Message("%s", err)
			}
		}
	}
}
