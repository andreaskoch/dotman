// Copyright 2013 Andreas Koch. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package list

import (
	"github.com/andreaskoch/dotman/projects"
	"github.com/andreaskoch/dotman/ui"
)

const (
	ActionName        = "list"
	ActionDescription = "Get a list of all projects."
)

type List struct {
	projectCollectionProvider func() *projects.Collection
}

func New(projectCollectionProvider func() *projects.Collection) *List {
	return &List{
		projectCollectionProvider: projectCollectionProvider,
	}
}

func (list *List) Name() string {
	return ActionName
}

func (list *List) Description() string {
	return ActionDescription
}

func (list *List) Execute() {
	list.execute()
}

func (list *List) DryRun() {
	list.execute()
}

func (list *List) execute() {
	projects := list.projectCollectionProvider()

	for _, project := range projects.Collection {
		ui.Message("%s", project)
	}
}
