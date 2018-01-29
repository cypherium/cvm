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
	"fmt"
	"math/big"
	"testing"
)

type CVM struct {
	GasLimit uint64
	GasPrice *big.Int
	cf       *ContractFrame
}

func TestCvm(t *testing.T) {
	cvm := CVM{
		GasLimit: 1000,
		GasPrice: big.NewInt(10),
	}
	//Create a VM contract instance using an account's public address
	contract0 := NewContract(0, 0, nil, big.NewInt(100), cvm.GasLimit, cvm.GasPrice)
	contract1 := NewContract(1, 0, nil, big.NewInt(100), cvm.GasLimit, cvm.GasPrice)
	//Create a context associated with the given contract
	cvm.cf = NewContractFrame(contract0)
	if err := cvm.cf.Execute(); err != nil {
		t.Errorf("%s\n", err.Error())
	}
	fmt.Printf("End of execution\n")
	cvm.cf = NewContractFrame(contract1)
	if err := cvm.cf.Execute(); err != nil {
		t.Errorf("%s\n", err.Error())
	}
	fmt.Printf("End of execution\n")
}
