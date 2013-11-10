// Copyright 2013 Andreas Koch. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package command

import (
	"io"
	"os"
	"os/exec"
)

func Execute(directory, commandName string, arguments ...string) error {

	// get the command
	command := getCmd(directory, commandName, arguments...)

	// execute the command
	if err := command.Start(); err != nil {
		return err
	}

	// wait for the command to finish
	return command.Wait()
}

func getCmd(directory, commandName string, arguments ...string) *exec.Cmd {
	if commandName == "" {
		return nil
	}

	// create the command
	command := exec.Command(commandName, arguments...)

	// set the working directory
	command.Dir = directory

	// redirect command io
	redirectCommandIO(command)

	return command
}

func redirectCommandIO(cmd *exec.Cmd) (*os.File, error) {
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		return nil, err
	}

	go io.Copy(os.Stdout, stdout)
	go io.Copy(os.Stderr, stderr)

	//direct. Masked passwords work OK!
	cmd.Stdin = os.Stdin
	return nil, err
}
