// Copyright 2013 Andreas Koch. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import ()

const (
	VERSION = "0.1.0"
)

func main() {

	switch settings.Action.String() {
	case listProjectsAction:

	case deployAction:
		// deploy
	case updateAction:
		// update
	case importAction:
		// import
	case backupAction:
		// backup
	default:
		usage()
	}
}
