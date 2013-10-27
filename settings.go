// Copyright 2013 Andreas Koch. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"github.com/andreaskoch/dotman/projects"
	"github.com/andreaskoch/dotman/ui"
	"os"
	"strings"
)

const (
	dotmanFileName = "dotman"

	helpAction         = "help"
	backupAction       = "backup"
	importAction       = "import"
	deployAction       = "deploy"
	updateAction       = "update"
	listProjectsAction = "list"
)

var (
	availableActions = []string{helpAction, listProjectsAction, backupAction, importAction, deployAction, updateAction}

	settings Settings = Settings{}
)

type Action struct {
	Name      string
	Arguments []string
}

func newAction(name string, arguments []string) Action {
	return Action{
		Name:      strings.TrimSpace(strings.ToLower(name)),
		Arguments: arguments,
	}
}

func (action Action) String() string {
	return fmt.Sprintf("%s", action.Name)
}

type Settings struct {
	WorkingDirectory string
	Action           Action
	Projects         *projects.Collection
}

func init() {

	// determine working directory
	workingDirectory, err := os.Getwd()
	if err != nil {
		ui.Fatal("Cannot determine working directory. %s", err)
	}

	settings.WorkingDirectory = workingDirectory

	// load the projects
	projectCollection, err := projects.Load(workingDirectory, dotmanFileName)
	if err != nil {
		ui.Fatal("Unable to scan for projects. %s", err)
	}

	settings.Projects = projectCollection

	// parse command line arguments
	if len(os.Args) > 1 {
		settings.Action = newAction(os.Args[1], os.Args[1:])
	}

}

var usage = func() {
	// description
	ui.Message("Backup and bootstrap your dotfiles and system configuration.")
	ui.Message("")

	// usage
	ui.Message("usage: %s <command> [args]", os.Args[0])
	ui.Message("")

	// commands
	ui.Message("Available commands are:")
	ui.Message("    %s %s  %s", helpAction, getActionSpacer(helpAction), "Prints this help text. Add <command> to get specific help.")
	ui.Message("    %s %s  %s", listProjectsAction, getActionSpacer(listProjectsAction), "Get a list of all projects.")
	ui.Message("    %s %s  %s", deployAction, getActionSpacer(deployAction), "Deploy all configuration files")
	ui.Message("    %s %s  %s", backupAction, getActionSpacer(backupAction), "Backup your current configuration")
	ui.Message("    %s %s  %s", importAction, getActionSpacer(importAction), "Import your current configuration files")
	ui.Message("    %s %s  %s", updateAction, getActionSpacer(updateAction), "Pull the latest changes from your remote")
}

func getActionSpacer(action string) string {
	maxLen := 0
	for _, action := range availableActions {
		if len(action) > maxLen {
			maxLen = len(action)
		}
	}

	spacer := ""
	for i := 0; i < maxLen-len(action); i++ {
		spacer += " "
	}

	return spacer
}
