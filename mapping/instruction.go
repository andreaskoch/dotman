// Copyright 2013 Andreas Koch. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mapping

func newInstruction(source, target string) *Instruction {
	return &Instruction{
		sourcePath: source,
		targetPath: target,
	}
}

type Instruction struct {
	sourcePath string
	targetPath string
}

func (instruction *Instruction) Source() string {
	return instruction.sourcePath
}

func (instruction *Instruction) Target() string {
	return instruction.targetPath
}
