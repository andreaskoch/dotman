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

			if fs.IsDirectory(source) {
				continue // skip directories for now
			}

			if !filesAreEqual(source, target) {
				changes <- fmt.Sprintf("%s.", target)
			}
		}

		close(changes)
	}()

	return changes
}

func filesAreEqual(source, target string) bool {

	// determine the hash of the source file
	sourceHash, sourceHashErr := fs.GetFileHash(source)
	if sourceHashErr != nil {
		ui.Fatal("%s", sourceHashErr)
	}

	// determine the hash of the target file
	targetHash, targetHashErr := fs.GetFileHash(target)
	if targetHashErr != nil {
		ui.Fatal("%s", targetHashErr)
	}

	return sourceHash == targetHash
}
