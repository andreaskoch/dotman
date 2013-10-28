// Copyright 2013 Andreas Koch. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package actions

type ActionMetaData interface {
	Name() string
	Description() string
}

type Action interface {
	Name() string
	Description() string
	DryRun()
	Execute()
}

type ActionInfo struct {
	name        string
	description string
}

func NewActionInfo(name, description string) ActionInfo {
	return ActionInfo{
		name:        name,
		description: description,
	}
}

func (info ActionInfo) Name() string {
	return info.name
}

func (info ActionInfo) Description() string {
	return info.description
}
