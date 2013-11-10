// Copyright 2013 Andreas Koch. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"github.com/andreaskoch/dotman/actions"
	"github.com/andreaskoch/dotman/ui"
	"os"
	"strings"
)

const (
	VERSION       = "0.1.0"
	REPOSITORYURL = "https://github.com/andreaskoch/dotman"
)

var (
	// the what-if flag
	whatIfFlag            = false
	whatIfFlagName        = "whatif"
	whatIfFlagDescription = "Enable the dry-run mode. Only print out what would happen."

	// module filter argument
	moduleFilterExpressionName        = "filter"
	moduleFilterExpressionDescription = "You can add a module filter expression to the import, list, changes and deploy commands."
)

func init() {
	// define flags
	flag.BoolVar(&whatIfFlag, whatIfFlagName, whatIfFlag, whatIfFlagDescription)

}

func main() {

	// parse the command-line options
	flag.Parse()

	// determine working directory
	workingDirectory, err := os.Getwd()
	if err != nil {
		ui.Fatal("Cannot determine working directory. %s", err)
	}

	// determine the command name and command arguments
	commandName := ""
	commandArguments := make([]string, 0)

	commandLineArguments := getCommandLineArguments()
	if len(commandLineArguments) > 0 {
		commandName = commandLineArguments[0]
		commandArguments = commandLineArguments[1:]
	}

	if command := actions.Get(workingDirectory, commandName); command != nil {

		if whatIfFlag {
			ui.Message("Performing a dry-run. No changes will we applied to the system.")
			command.DryRun(commandArguments)
		} else {
			command.Execute(commandArguments)

		}

		os.Exit(0)
	}

	// print the help if no command was recognized
	usage()
}

func getCommandLineArguments() []string {
	args := make([]string, 0)
	for _, arg := range os.Args[1:] {
		whatIfFlagArg := fmt.Sprintf("-%s", whatIfFlagName)
		if strings.HasPrefix(arg, whatIfFlagArg) {
			continue
		}

		args = append(args, arg)
	}

	return args
}

func getApplicationName() string {
	return os.Args[0]
}

var usage = func() {
	// description
	ui.Message("v%s - Backup and bootstrap your dotfiles and system configuration.", VERSION)
	ui.Message("")

	// usage
	ui.Message("usage: %s [-whatif] <command> [<filter>]", getApplicationName())
	ui.Message("")

	// commands
	ui.Message("Available commands are:")
	for _, action := range actions.GetAll() {
		ui.Message("    %s %s  %s", action.Name(), getActionSpacer(action.Name()), action.Description())
	}

	// flags
	ui.Message("")
	ui.Message("Options:")
	ui.Message("    %s %s  %s", whatIfFlagName, getActionSpacer(whatIfFlagName), whatIfFlagDescription)

	// args
	ui.Message("")
	ui.Message("Arguments:")
	ui.Message("    %s %s  %s", moduleFilterExpressionName, getActionSpacer(moduleFilterExpressionName), moduleFilterExpressionDescription)

	// source code
	ui.Message("")
	ui.Message("Contribute: %s", REPOSITORYURL)
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
