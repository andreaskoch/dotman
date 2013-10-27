// Copyright 2013 Andreas Koch. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/andreaskoch/dotman/actions"
	"github.com/andreaskoch/dotman/ui"
	"os"
)

const (
	VERSION = "0.1.0"
)

func main() {

	// determine working directory
	workingDirectory, err := os.Getwd()
	if err != nil {
		ui.Fatal("Cannot determine working directory. %s", err)
	}

	// determine the command name and command arguments
	commandName := ""
	commandArguments := make([]string, 0)
	if len(os.Args) > 1 {
		commandName = os.Args[1]
		commandArguments = os.Args[1:]
	}

	if command := actions.Get(workingDirectory, commandName, commandArguments); command != nil {
		command.Execute()
		os.Exit(0)
	}

	// print the help if no command was recognized
	usage()
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
	for _, action := range actions.GetAll() {
		ui.Message("    %s %s  %s", action.Name(), getActionSpacer(action.Name()), action.Description())
	}
}

func getActionSpacer(action string) string {
	maxLen := 0
	for _, action := range actions.GetAll() {
		if len(action.Name()) > maxLen {
			maxLen = len(action.Name())
		}
	}

	spacer := ""
	for i := 0; i < maxLen-len(action); i++ {
		spacer += " "
	}

	return spacer
}
