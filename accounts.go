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

type ContractAccount struct {
	Address uint64
	Code    []byte
	Storage []byte
}

//INC %1 INC %2 INC %1 ADD %1 %1 %2
var contractCode0, _ = hex.DecodeString("51000151000251000155000100010002")

//INC %1 INC %2 INC %1 ADD %1 %1 %2 INC %2 CALL %2
var contractCode1, _ = hex.DecodeString("51000151000251000155000100010002510002060002")

//MOVI %16 $2 MOVI %14 $0xFF MOVI %15 $0xFF00 RET
var contractCode2, _ = hex.DecodeString("42000F000242000D00FF42000EFF0007")

var storage0 = []byte{10, 20, 0, 50, 100}
var storage1 = []byte{10, 20, 0}
var storage2 = []byte{10, 20, 0}

var AccountExamples = []ContractAccount{
	{
		Address: 0,
		Code:    contractCode0,
		Storage: storage0,
	},
	{
		Address: 1,
		Code:    contractCode1,
		Storage: storage1,
	},
	{
		Address: 2,
		Code:    contractCode2,
		Storage: storage2,
	},
}
