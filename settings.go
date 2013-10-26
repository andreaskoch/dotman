// Copyright 2013 Andreas Koch. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"github.com/andreaskoch/dotman/mapping"
	"github.com/andreaskoch/dotman/ui"
	"github.com/andreaskoch/dotman/util/fs"
	"os"
	"path/filepath"
	"strings"
)

const (
	dotmanFileName = "dotman"

	helpAction   = "help"
	backupAction = "backup"
	importAction = "import"
	deployAction = "deploy"
	updateAction = "update"
)

var (
	availableActions = []string{helpAction, backupAction, importAction, deployAction, updateAction}

	settings CommandlineArguments = CommandlineArguments{}
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

type CommandlineArguments struct {
	WorkingDirectory string
	Action           Action
	Map              *mapping.PathMap
}

func init() {

	// determine working directory
	workingDirectory, err := os.Getwd()
	if err != nil {
		panic(fmt.Sprintf("Cannot determine working directory. %s", err))
	}
	settings.WorkingDirectory = workingDirectory

	// load path map
	dotmanFilePath := filepath.Join(workingDirectory, dotmanFileName)
	if fs.FileExists(dotmanFilePath) {

		if pathMap, err := mapping.NewPathMap(dotmanFilePath); err != nil {
			ui.Message("Unable to read dotman file. %s", err)
		} else {
			settings.Map = pathMap
		}

		fmt.Printf("%s", settings.Map)

	}

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
