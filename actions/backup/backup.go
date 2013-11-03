// Copyright 2013 Andreas Koch. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package backup

import (
	"archive/tar"
	"bytes"
	"fmt"
	"github.com/andreaskoch/dotman/projects"
	"github.com/andreaskoch/dotman/ui"
	"github.com/andreaskoch/dotman/util/fs"
	"io"
	"os"
	"path/filepath"
	"time"
)

const (
	BackupDirectoryName = ".backup"
	ActionName          = "backup"
	ActionDescription   = "Backup your target files."
)

type Backup struct {
	projectCollectionProvider func() *projects.Collection
}

func New(projectCollectionProvider func() *projects.Collection) *Backup {
	return &Backup{
		projectCollectionProvider: projectCollectionProvider,
	}
}

func (backup *Backup) Name() string {
	return ActionName
}

func (backup *Backup) Description() string {
	return ActionDescription
}

func (backup *Backup) Execute(arguments []string) {
	backup.execute(false, arguments)
}

func (backup *Backup) DryRun(arguments []string) {
	backup.execute(true, arguments)
}

func (backup *Backup) execute(executeADryRunOnly bool, arguments []string) {

	projects := backup.projectCollectionProvider()

	// assemble a list of all files to backup
	files := make([]string, 0)
	for _, project := range projects.Collection {

		// add the project file
		files = append(files, project.ProjectFile())

		// add all target files
		for _, entry := range project.Map.Entries {

			targetPath := entry.Target
			if !fs.IsDirectory(targetPath) {
				files = append(files, targetPath)
				continue
			}

			subDirectoryFiles := fs.GetAllFilesRecursively(targetPath)
			files = append(files, subDirectoryFiles...)

		}
	}

	// make sure the archive directory exists
	archiveDirectory := filepath.Join(projects.BaseDirectory, BackupDirectoryName)
	if !fs.DirectoryExists(archiveDirectory) {
		ui.Message("Creating backup directory %q.", archiveDirectory)
		if !executeADryRunOnly && !fs.CreateDirectory(archiveDirectory) {
			ui.Fatal("Unable to create the backup directory %q.", archiveDirectory)
		}
	}

	// assemble a filename for the backup archive
	const dateLayout = "2006-01-02 15:04:05"
	filename := fmt.Sprintf("%s.tar", time.Now().Format(dateLayout))
	archivePath := filepath.Join(archiveDirectory, filename)

	if !executeADryRunOnly {

		// create the archive
		_, err := createTarArchive(archivePath, files)
		if err != nil {
			ui.Fatal("Unable to create a backup %q. %s", archivePath, err)
		}

		ui.Message("The backup has been saved to %q.", archivePath)

	} else {

		ui.Message("Creating archive %s:", archivePath)
		for number, file := range files {
			ui.Message("%d. Adding file %q", (number + 1), file)
		}

	}
}

func createTarArchive(archivePath string, files []string) (success bool, err error) {

	// Create a buffer to write our archive to.
	buffer := new(bytes.Buffer)

	// create the archive writer
	archiveWriter, archiveWriterError := os.OpenFile(archivePath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0600)
	if archiveWriterError != nil {
		return false, fmt.Errorf("%s", archiveWriterError)
	}

	defer func() {
		archiveWriter.Write(buffer.Bytes())
		archiveWriter.Close()
	}()

	// Create a new tar archive.
	archive := tar.NewWriter(buffer)

	// Add some files to the archive.
	for _, file := range files {

		// open the source file
		fileReader, readerError := os.Open(file)
		if readerError != nil {
			return false, fmt.Errorf("%s", readerError) // unable to open target file
		}

		defer fileReader.Close()

		fileInfo, fileInfoError := os.Stat(file)
		if fileInfoError != nil {
			return false, fmt.Errorf("%s", fileInfoError) // unable to get file info
		}

		// create the file header
		fileHeader := &tar.Header{
			Name: file,
			Size: fileInfo.Size(),
		}

		// write the file header
		if fileHeaderError := archive.WriteHeader(fileHeader); fileHeaderError != nil {
			return false, fmt.Errorf("%s", fileHeaderError)
		}

		// write the file content
		if _, copyError := io.Copy(archive, fileReader); copyError != nil {
			return false, fmt.Errorf("%s", copyError)
		}
	}

	// check for errors
	if closeError := archive.Close(); closeError != nil {
		return false, fmt.Errorf("%s", closeError)
	}

	return true, nil
}
