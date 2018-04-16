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
	"encoding/hex"
	"fmt"
	"github.com/syndtr/goleveldb/leveldb"
	"math/big"
	mrand "math/rand"
	"testing"
	"time"
)

func TestCvmFunc(t *testing.T) {
	var db *leveldb.DB = nil
	mrand.Seed(time.Now().Unix())
	code0, _ := hex.DecodeString("51000151000251000155000100010002")
	contract0 := Contract{
		Caller:   nil,
		Self:     []byte("2E98"),
		Code:     code0,
		Input:    nil,
		Storage:  nil,
		GasLimit: mrand.Uint64() % 1000,
		GasPrice: new(big.Int).SetUint64(mrand.Uint64() % 90000),
		Balance:  new(big.Int).SetUint64(mrand.Uint64() % 9000),
	}

	//Create a context associated with the given contract
	cf := NewContractFrame(&contract0, db)
	if err := cf.Execute(); err != nil {
		t.Errorf("%s\n", err.Error())
	}
	fmt.Printf("End of execution\n")
	//cf = NewContractFrame(contract1)
	//if err := cvm.cf.Execute(); err != nil {
	//	t.Errorf("%s\n", err.Error())
	//}
	//fmt.Printf("End of execution\n")
}

/*
func TestCvmGas(t *testing.T) {
	cvm := CVM{
		GasLimit: 5,
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
*/
