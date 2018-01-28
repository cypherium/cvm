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
	"math/big"
)

type RegisterFile struct {
	data   []*big.Int
	regSze uint64
}

func NewRegisterFile() *RegisterFile {
	rf := &RegisterFile{data: make([]*big.Int, RegSze, RegSze)}
	for i, _ := range rf.data {
		rf.data[i] = big.NewInt(0)
	}
	rf.regSze = RegSze
	return rf
}

//Read reads the idx-th register
func (rf *RegisterFile) Read(idx uint64) *big.Int {
	return rf.data[idx]
}

//Write writes the value val to the idx-th register
func (rf *RegisterFile) Write(idx uint64, val *big.Int) {
	rf.data[idx] = val
}

//Print prints the content of the entire register file
func (rf *RegisterFile) Print(log *log.Logger) {
	log.Printf("-----Register File %p-----\n", rf)
	log.Println(rf.data)
}
