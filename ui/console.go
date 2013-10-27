// Copyright 2013 Andreas Koch. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"fmt"
	"os"
	"strings"
)

func Message(text string, args ...interface{}) {

	// append newline character
	if !strings.HasSuffix(text, "\n") {
		text += "\n"
	}

	fmt.Printf(text, args...)
}

func Fatal(text string, args ...interface{}) {
	Message(text, args)
	os.Exit(2)
}
