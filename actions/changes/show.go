// Copyright 2013 Andreas Koch. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package changes

import (
	"fmt"
	"github.com/andreaskoch/dotman/actions/base"
	"github.com/andreaskoch/dotman/projects"
	"github.com/andreaskoch/dotman/ui"
	"github.com/andreaskoch/dotman/util/fs"
)

const (
	ActionName        = "changes"
	ActionDescription = "Show changed files."
)

type Importer struct {
	*base.Action
}

func New(projectCollectionProvider base.ProjectsProviderFunc) *Importer {
	return &Importer{
		base.New(ActionName, ActionDescription, projectCollectionProvider, func(project *projects.Project, executeADryRunOnly bool) {

			projectTitleHasBeenPrinted := false
			for change := range showChanges(project) {

				// print project title
				if !projectTitleHasBeenPrinted {
					ui.Message("\n%s:", project)
					projectTitleHasBeenPrinted = true
				}

				// report the change
				ui.Message(change)
			}

		}),
	}
}

func showChanges(project *projects.Project) (changes chan string) {

	changes = make(chan string, 10)

	go func() {
		for _, instruction := range project.Map.GetInstructions() {

			source := instruction.Source()
			target := instruction.Target()

			// compare directories
			if fs.IsDirectory(source) {

				directoriesAreEqual, filesThatAreDifferent, err := fs.DirectoriesAreEqual(source, target)
				if err != nil {
					ui.Fatal("Error while comparing the directories %q and %q. Error: %s", source, target, err)
				}

				if !directoriesAreEqual {
					for _, changedFile := range filesThatAreDifferent {
						changes <- fmt.Sprintf("%s", changedFile)
					}
				}

				// check next instruction
				continue
			}

			// compare files
			areEqual, err := fs.FilesAreEqual(source, target)
			if err != nil {
				ui.Fatal("Error while comparing the files %q and %q. Error: %s", err)
			}

			if !areEqual {
				changes <- fmt.Sprintf("%s", target)
			}
		}

		close(changes)
	}()

	return changes
}
