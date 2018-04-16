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
	"encoding/binary"
	"fmt"
	"github.com/syndtr/goleveldb/leveldb"
	"log"
	"os"
)

//ContractFrame defines the runtime info associated with a contract
type ContractFrame struct {
	rf         *RegisterFile
	mem        *Memory
	contract   *Contract
	depth      uint64 //Current call depth
	pc         uint64 //Program counter
	jmpc       uint64 //Conditional jump counter
	gas        uint64 //Remaining gas counter
	Operations OperationSet

	call    bool           //If this contract is called by another
	caller  *ContractFrame //Previous caller
	sender  *ContractFrame //Original caller
	log     *log.Logger
	StateDB *leveldb.DB
}

func NewContractFrame(ctc *Contract, db *leveldb.DB) (cf *ContractFrame) {
	cf = &ContractFrame{
		rf:         NewRegisterFile(),
		mem:        NewMemory(),
		contract:   ctc,
		pc:         0,
		depth:      0,
		gas:        ctc.GasLimit,
		Operations: NewOperationSet(),
		call:       false,
		caller:     nil,
		sender:     nil,
		log:        log.New(os.Stderr, "", 0),
		StateDB:    db,
	}
	return
}

/*
 * CallContractFrame sets up a ContractFrame for a contract called by another
 * one and sets up the correct calling relationships (sender, caller etc.).
 */
func (cf *ContractFrame) CallContractFrame(ctc *Contract) error {
	if cf.depth == MaxCallDepth {
		return ErrorCallDepth
	}
	newCf := &ContractFrame{
		rf:         NewRegisterFile(),
		mem:        NewMemory(),
		contract:   ctc,
		pc:         0,
		depth:      cf.depth + 1,
		gas:        cf.gas,
		Operations: NewOperationSet(),
		call:       true,
		caller:     cf,
		sender:     cf.sender,
		log:        log.New(os.Stderr, "", 0),
	}
	//If the sender is not set it means that the caller is the original sender
	if newCf.sender == nil {
		newCf.sender = cf
	}
	fmt.Println(newCf.pc, newCf.depth, newCf.call)
	newCf.Execute()
	return nil
}

/*
 * Execute reads the contract associated with the current contract frame and
 * execute it
 */
func (cf *ContractFrame) Execute() error {
	cf.contract.PrintCode(cf.log)
	/*
	 * A jump out of the code would result an immediate return, the compiler
	 * should ensure a valid jump
	 */
	for cf.pc < uint64(len(cf.contract.Code)) && cf.pc >= 0 {
		if cf.gas == 0 {
			return ErrorOutOfGas
		}
		opIdx := cf.contract.Code[cf.pc]
		op := cf.Operations[opIdx]
		regIdx := make([]uint64, op.regNum)
		for i := 0; uint64(i) < op.regNum; i++ {
			offset := cf.pc + 1 + RegLen*uint64(i)
			regIdx[i] = uint64(binary.BigEndian.Uint16(cf.contract.Code[offset : offset+2]))
		}
		cf.log.Println("pc:", cf.pc, "opname:", op.name, "regs:", regIdx, "gas", cf.gas)
		if exeErr := op.execute(cf, regIdx, &op); exeErr != nil {
			//log.Println(err.Error())
			return exeErr
		}
		cf.rf.Print(cf.log)
	}
	return nil
}
