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

import (
	"log"
)

type Memory struct {
	data []byte
}

func NewMemory() *Memory {
	return &Memory{}
}

//Expand expands the memory to size
func (mem *Memory) Expand(size uint64) error {
	appendSze := size - uint64(len(mem.data))
	mem.data = append(mem.data, make([]byte, appendSze)...)
	if uint64(len(mem.data)) < size {
		return ErrorOutOfMemory
	}
	return nil
}

//Store stores byte array val of size into the memory region starting from idx
func (mem *Memory) Store(idx uint64, size uint64, val []byte) error {
	/*
	 * Handle the corner case where the location in the memory the content to be
	 * stored is not yet allocated
	 */
	if (idx + size) > uint64(len(mem.data)) {
		if err := mem.Expand(idx + size + 128); err != nil {
			return err
		}
	}
	copy(mem.data[idx:idx+size], val)
	return nil
}

//Retrieve gets size byte of content from the memory region starting from idx
func (mem *Memory) Retrieve(idx uint64, size uint64) ([]byte, error) {
	/*
	 * Handle the corner case where the content to be retrieved is out of the
	 * current memory boundary
	 */
	if (idx + size) > uint64(len(mem.data)) {
		return nil, ErrorOutOfMemory
	}
	content := make([]byte, size)
	copy(content, mem.data[idx:idx+size])
	return content, nil
}

func (mem *Memory) Print(log *log.Logger) {
	log.Printf("-----Memory %p of size %d-----\n", mem, len(mem.data))
	log.Println(mem.data)
}
