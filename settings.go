// Copyright 2013 Andreas Koch. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
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
	Map              *PathMap
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
	if FileExists(dotmanFilePath) {

		if pathMap, err := NewPathMap(dotmanFilePath); err != nil {
			message("Unable to read dotman file. %s", err)
		} else {
			settings.Map = pathMap
		}

		fmt.Printf("%#v", settings.Map)

	}

	// parse command line arguments
	if len(os.Args) > 1 {
		settings.Action = newAction(os.Args[1], os.Args[1:])
	}

}

var usage = func() {
	// description
	message("Backup and bootstrap your dotfiles and system configuration.")
	message("")

	// usage
	message("usage: %s <command> [args]", os.Args[0])
	message("")

	// commands
	message("Available commands are:")
	message("    %s %s  %s", helpAction, getActionSpacer(helpAction), "Prints this help text. Add <command> to get specific help.")
	message("    %s %s  %s", deployAction, getActionSpacer(deployAction), "Deploy all configuration files")
	message("    %s %s  %s", backupAction, getActionSpacer(backupAction), "Backup your current configuration")
	message("    %s %s  %s", importAction, getActionSpacer(importAction), "Import your current configuration files")
	message("    %s %s  %s", updateAction, getActionSpacer(updateAction), "Pull the latest changes from your remote")
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
