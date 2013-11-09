// Copyright 2013 Andreas Koch. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fs

import (
	"bufio"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
	"regexp"
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

func Copy(source, target string) (success bool, err error) {

	// check if the source is a file
	if IsFile(source) {
		return CopyFile(source, target)
	}

	// the source must be a directory
	// read the source directory
	sourceEntries, err := ioutil.ReadDir(source)
	if err != nil {
		return false, err
	}

	for _, sourceEntry := range sourceEntries {
		sourceEntryPath := filepath.Join(source, sourceEntry.Name())

		// recurse into the sub-directory
		if sourceEntry.IsDir() {
			nestedTargetPath := filepath.Join(target, sourceEntry.Name())
			if _, err := Copy(sourceEntryPath, nestedTargetPath); err != nil {
				return false, err // abort if an error occurs
			}

			continue
		}

		// copy the file
		targetFilePath := filepath.Join(target, sourceEntry.Name())
		if _, err := CopyFile(sourceEntryPath, targetFilePath); err != nil {
			return false, err // abort if an error occurs
		}
	}

	// if no error occured everything must be ok
	return true, nil
}

func CopyFile(source, target string) (success bool, err error) {
	if !IsFile(source) {
		return false, fmt.Errorf("%q is not a file.", source)
	}

	// open the source file
	sourceReader, readerErr := os.Open(source)
	if readerErr != nil {
		return false, readerErr
	}

	defer sourceReader.Close()

	// prepare the target file
	var targetFileMode os.FileMode = 0600
	if !FileExists(target) {

		// make sure the path to the target file exists
		if _, createFileErr := CreateFile(target); createFileErr != nil {
			return false, fmt.Errorf("Unable to create the target file %q. Error: %s", target, err)
		}
	} else {

		// use the same file mode if the file already exists
		if targetFileInfo, targetFileInfoErr := os.Stat(target); targetFileInfoErr == nil {
			targetFileMode = targetFileInfo.Mode().Perm()
		}

	}

	// open the target file for writing
	targetWriter, writerErr := os.OpenFile(target, os.O_RDWR|os.O_CREATE|os.O_TRUNC, targetFileMode)
	if writerErr != nil {
		return false, writerErr
	}

	defer targetWriter.Close()

	// copy from source to target
	_, copyErr := io.Copy(targetWriter, sourceReader)
	if copyErr != nil {
		return false, err
	}

	return true, nil
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
	if err != nil {
		return false
	}

	return !file.IsDir()
}

func DirectoryExists(path string) bool {
	if strings.TrimSpace(path) == "" {
		return false
	}

	file, err := os.Stat(path)
	if err != nil {
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

func GetUserHomeDirectory() (string, error) {

	usr, err := user.Current()
	if err != nil {
		return "", err
	}

	return filepath.Clean(usr.HomeDir), nil
}

func GetAllFilesRecursively(path string) []string {
	recurse := true
	return getAllDirectoryEntries(path, recurse, func(file os.FileInfo) bool {
		return !file.IsDir()
	})
}

func GetMatchingDirectoryEntries(path string, pattern *regexp.Regexp) []string {
	recurse := false
	return getAllDirectoryEntries(path, recurse, func(file os.FileInfo) bool {
		return pattern.MatchString(file.Name())
	})
}

func forEachDirectoryEntry(path string, expression func(file os.FileInfo) error) error {

	directoryEntries, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}

	for _, entry := range directoryEntries {
		if err := expression(entry); err != nil {
			return err
		}
	}

	return nil
}

func getAllDirectoryEntries(path string, recurse bool, includeDirectoryEntry func(file os.FileInfo) bool) []string {

	files := make([]string, 0)

	directoryEntries, err := ioutil.ReadDir(path)
	if err != nil {
		return files
	}

	for _, entry := range directoryEntries {

		entryPath := filepath.Join(path, entry.Name())

		// recurse?
		if entry.IsDir() && recurse {
			files = append(files, getAllDirectoryEntries(entryPath, recurse, includeDirectoryEntry)...)
		}

		if includeDirectoryEntry(entry) {
			files = append(files, entryPath)
		}
	}

	return files
}

func GetFileHash(file string) (string, error) {

	if !IsFile(file) {
		return "", fmt.Errorf("%q is not a file.", file)
	}

	itemBytes, readFileErr := ioutil.ReadFile(file)
	if readFileErr != nil {
		return "", fmt.Errorf("Unable to read file %q.", file)
	}

	sha1Hash := sha1.New()
	sha1Hash.Write(itemBytes)
	hashBytes := sha1Hash.Sum(nil)

	return string(hex.EncodeToString(hashBytes)), nil
}

func DirectoriesAreEqual(source, target string) (directoriesAreEqual bool, filesThatAreDifferent []string, err error) {

	filesThatAreDifferent = make([]string, 0)
	err = forEachDirectoryEntry(source, func(file os.FileInfo) error {

		subSource := filepath.Join(source, file.Name())
		subTarget := filepath.Join(source, file.Name())

		// check if the file is are directory
		if file.IsDir() {

			// recurse
			directoriesAreEqual, changedFiles, subDirectoryError := DirectoriesAreEqual(subSource, subTarget)
			if subDirectoryError != nil {
				return subDirectoryError
			}

			// append the changed files
			if !directoriesAreEqual {
				filesThatAreDifferent = append(filesThatAreDifferent, changedFiles...)
			}

			return nil
		}

		// check if the source and target are different
		sourceAndTargetAreEqual, err := FilesAreEqual(subSource, subTarget)
		if err != nil {
			return err
		}

		// if the files are different, add the target file to the list
		if !sourceAndTargetAreEqual {
			filesThatAreDifferent = append(filesThatAreDifferent, subTarget)
		}

		return nil
	})

	// if no files are different and no error occured the directories are alike
	directoriesAreEqual = len(filesThatAreDifferent) == 0 && err == nil

	return directoriesAreEqual, filesThatAreDifferent, err

}

func FilesAreEqual(source, target string) (bool, error) {

	// determine the hash of the source file
	sourceHash, sourceHashErr := GetFileHash(source)
	if sourceHashErr != nil {
		return false, sourceHashErr
	}

	// determine the hash of the target file
	targetHash, targetHashErr := GetFileHash(target)
	if targetHashErr != nil {
		return false, sourceHashErr
	}

	return sourceHash == targetHash, nil
}
