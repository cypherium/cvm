/*
 * Copyright (C) 2018 The Cypherium VM (CVM) authors
 *
 * This file is part of the CVM library.
 *
 * The CVM library is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The CVM library is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Lesser General Public License for more details.
 *
 * You should have received a copy of the GNU Lesser General Public License
 * along with the CVM library. If not, see <http://www.gnu.org/licenses/>.
 *
 */

package cvm

import "errors"

var (
	ErrorCallDepth      = errors.New("Max call depth reached")
	ErrorOutOfGas       = errors.New("Out of gas")
	ErrorOutOfMemory    = errors.New("Out of memory")
	ErrorMemoryBoundary = errors.New("Memory boundary reached")
	ErrorMsgBoundary    = errors.New("Message boundary reached")
	ErrorStgBoundary    = errors.New("Storage boundary reached")
	ErrorAryBoundary    = errors.New("Array boundary reached")
)

type OpError struct {
	OpName string
	Cause  string
}

func (err OpError) Error() string {
	errStr := "OpName " + err.OpName + "cause " + err.Cause
	return errStr
}
