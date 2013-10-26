// Copyright 2013 Andreas Koch. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

// readLine returns a single line (without the ending \n)
// from the input buffered reader.
// An error is returned iff there is an error with the
// buffered reader.
func readLine(bufferedReader *bufio.Reader) (string, error) {
	var (
		isPrefix bool  = true
		err      error = nil
		line, ln []byte
	)

	for isPrefix && err == nil {
		line, isPrefix, err = bufferedReader.ReadLine()
		ln = append(ln, line...)
	}

	return string(ln), err
}

// Get all lines of a given file
func GetLines(inFile io.Reader) []string {

	lines := make([]string, 0, 10)
	bufferedReader := bufio.NewReader(inFile)
	line, err := readLine(bufferedReader)
	for err == nil {
		lines = append(lines, line)
		line, err = readLine(bufferedReader)
	}

	return lines
}

func CreateDirectory(path string) bool {
	err := os.MkdirAll(path, 0700)
	return err == nil
}

func CreateFile(filePath string) (success bool, err error) {

	// make sure the parent directory exists
	directory := filepath.Dir(filePath)
	if !DirectoryExists(directory) {
		if !CreateDirectory(directory) {
			return false, fmt.Errorf("Cannot create the directory for the given file %q.", filePath)
		}
	}

	// create the file
	if _, err := os.Create(filePath); err != nil {
		return false, fmt.Errorf("Could not create file %q. Error: ", filePath, err)
	}

	return true, nil
}

func PathExists(path string) bool {
	if strings.TrimSpace(path) == "" {
		return false
	}

	if _, err := os.Stat(path); err != nil && os.IsNotExist(err) {
		return false
	}

	return true
}

func FileExists(path string) bool {
	if strings.TrimSpace(path) == "" {
		return false
	}

	file, err := os.Stat(path)
	if err != nil && os.IsNotExist(err) {
		return false
	}

	return !file.IsDir()
}

func DirectoryExists(path string) bool {
	if strings.TrimSpace(path) == "" {
		return false
	}

	file, err := os.Stat(path)
	if err != nil && os.IsNotExist(err) {
		return false
	}

	return file.IsDir()
}

func IsFile(path string) bool {

	fileInfo, err := os.Stat(path)
	if err != nil {
		return false
	}

	return fileInfo.IsDir() == false
}

func IsDirectory(path string) bool {

	fileInfo, err := os.Stat(path)
	if err != nil {
		return false
	}

	return fileInfo.IsDir()
}

func GetAllFiles(path string) []string {

	files := make([]string, 0)

	if !IsDirectory(path) {
		return files
	}

	directoryEntries, err := ioutil.ReadDir(path)
	if err != nil {
		return files
	}

	for _, entry := range directoryEntries {

		entryPath := filepath.Join(path, entry.Name())

		// recurse
		if entry.IsDir() {
			files = append(files, GetAllFiles(entryPath)...)
		}

		files = append(files, entryPath)
	}

	return files
}

func getUserHomeDirectory() (string, error) {

	usr, err := user.Current()
	if err != nil {
		return "", err
	}

	return filepath.Clean(usr.HomeDir), nil
}
