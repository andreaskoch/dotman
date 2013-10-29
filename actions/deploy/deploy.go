// Copyright 2013 Andreas Koch. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package deploy

import (
	"github.com/andreaskoch/dotman/projects"
	"github.com/andreaskoch/dotman/ui"
	"github.com/andreaskoch/dotman/util/fs"
)

const (
	ActionName        = "deploy"
	ActionDescription = "Deploy your projects."
)

type Deploy struct {
	projectCollectionProvider func() *projects.Collection
}

func New(projectCollectionProvider func() *projects.Collection) *Deploy {
	return &Deploy{
		projectCollectionProvider: projectCollectionProvider,
	}
}

func (deploy *Deploy) Name() string {
	return ActionName
}

func (deploy *Deploy) Description() string {
	return ActionDescription
}

func (deploy *Deploy) Execute() {
	deploy.execute(false)
}

func (deploy *Deploy) DryRun() {
	deploy.execute(true)
}

func (deploy *Deploy) execute(executeADryRunOnly bool) {

	projects := deploy.projectCollectionProvider()
	for _, project := range projects.Collection {
		ui.Message("Deploying %q", project)
		importProject(project, executeADryRunOnly)
	}

}

func importProject(project *projects.Project, executeADryRunOnly bool) {

	for _, entry := range project.Map.Entries {
		source := entry.Source.Path()
		target := entry.Target.Path()

		ui.Message("Copy: %s → %s", source, target)
		if !executeADryRunOnly {
			if _, err := fs.Copy(source, target); err != nil {
				ui.Message("%s", err)
			}
		}
	}
}
