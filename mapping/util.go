// Copyright 2013 Andreas Koch. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mapping

import (
	"github.com/andreaskoch/dotman/util/fs"
	"os"
	"regexp"
	"strings"
)

var (
	DirectorySeparatorPattern = regexp.MustCompile(`[\/]{1,}`)

	HomeDirectoryBashPattern = regexp.MustCompile(`^~`)

	UnixEnvironmentVariablePattern = regexp.MustCompile(`\$(\w+)`)

	WindowsEnvironmentVariablePattern = regexp.MustCompile(`%(\w+)%`)
)

func normalizePathSpecification(path string) string {

	// trim leading and trailing white space
	path = strings.TrimSpace(path)

	// replace all directory separators with the ones for the current platform
	path = DirectorySeparatorPattern.ReplaceAllString(path, string(os.PathSeparator))

	// trim trailing path separators
	path = strings.TrimSuffix(path, string(os.PathSeparator))

	return path
}

func expandPathVariables(path string) string {

	// replace ~/ with the real home directory path
	if homeDirectory, err := fs.GetUserHomeDirectory(); err == nil {
		path = HomeDirectoryBashPattern.ReplaceAllString(path, homeDirectory)
	}

	// replace environment variables
	path = replaceEnvironmentVariables(path, UnixEnvironmentVariablePattern)
	path = replaceEnvironmentVariables(path, WindowsEnvironmentVariablePattern)

	return path
}

func replaceEnvironmentVariables(path string, environmentVariablePattern *regexp.Regexp) string {

	matches := environmentVariablePattern.FindAllStringSubmatch(path, -1)
	for _, submatch := range matches {
		if len(submatch) < 2 {
			continue
		}

		fullmatch := submatch[0]
		variableName := strings.TrimSpace(submatch[1])
		value := os.Getenv(variableName)

		path = strings.Replace(path, fullmatch, value, 1)
	}

	return path
}

func isEmptyLine(line string) bool {
	return strings.TrimSpace(line) == ""
}

func isComment(line string) bool {
	return strings.HasPrefix(strings.TrimSpace(line), "#")
}
