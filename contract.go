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

/*
 * Contract defines the structure of the essential Cypherium smart contract
 * for the VM, the data are extracted from the actual Cypherium blockchain account.
 */
type Contract struct {
	//Caller   PublicAddress //Public address of the caller
	//Self     PublicAddress //Public address of ourself
	Caller   uint64 //Public address of the caller
	Self     uint64 //Public address of ourself
	Code     []byte
	Input    []byte
	Storage  []byte
	GasLimit uint64
	GasPrice *big.Int

	amount *big.Int
}

/*
 * TODO Need to provide a function to create a contract by retriving contract
 * info from the Cypherium UTOX, the following function is a simulation of that
 */
func NewContract(accAddr uint64, caller uint64, input []byte,
	amount *big.Int, gasLimit uint64, gasPrice *big.Int) (ctc *Contract) {

	acc := AccountExamples[accAddr]
	ctc = &Contract{
		Caller:   caller,
		Self:     acc.Address,
		Code:     acc.Code,
		Storage:  acc.Storage,
		Input:    input,
		GasLimit: gasLimit,
		GasPrice: gasPrice,
		amount:   amount,
	}
	return
}

func (ctc *Contract) Print(log *log.Logger) {
	log.Println("Contract info:")
	log.Printf("caller %d, self %d, GasLimit %d, GasPrice %s\n", ctc.Caller,
		ctc.Self, ctc.GasLimit, ctc.GasPrice.String())
}

func (ctc *Contract) PrintCode(log *log.Logger) {
	log.Printf("Code of address %d\n", ctc.Self)
	log.Printf("%X\n", ctc.Code)
}
