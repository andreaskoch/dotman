// Copyright 2013 Andreas Koch. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package action

import (
	"fmt"
	"strings"
)

type Action struct {
	Name      string
	Arguments []string
}

func New(name string, arguments []string) Action {
	return Action{
		Name:      strings.TrimSpace(strings.ToLower(name)),
		Arguments: arguments,
	}
}

func (action Action) String() string {
	return fmt.Sprintf("%s", action.Name)
}
