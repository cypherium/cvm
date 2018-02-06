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
	"math/big"
)

func incrementPc(pc *uint64, op *operation) {
	*pc += (OpLen + op.regNum*RegLen)
}

func decrementGas(gas *uint64, op *operation) {
	if *gas < op.gasCost {
		*gas = 0
	} else {
		*gas -= op.gasCost
	}
}

func opNop(cf *ContractFrame, regIdx []uint64, op *operation) error {
	defer decrementGas(&cf.gas, op)
	cf.pc += OpLen
	return nil
}

func opJmp(cf *ContractFrame, regIdx []uint64, op *operation) error {
	defer decrementGas(&cf.gas, op)
	rf := cf.rf
	idxA := regIdx[0]
	regA := rf.Read(idxA)
	cf.pc = regA.Uint64()
	return nil
}

func opJmpif(cf *ContractFrame, regIdx []uint64, op *operation) error {
	defer decrementGas(&cf.gas, op)
	rf := cf.rf
	idxA := regIdx[0]
	idxB := regIdx[1]
	regB := rf.Read(idxB)
	if regB.Cmp(big.NewInt(0)) != 0 {
		offset := rf.Read(idxA).Uint64()
		cf.pc = offset
	} else {
		cf.pc += (OpLen + op.regNum*RegLen)
	}
	return nil
}

func opJmpifnot(cf *ContractFrame, regIdx []uint64, op *operation) error {
	defer decrementGas(&cf.gas, op)
	rf := cf.rf
	idxA := regIdx[0]
	idxB := regIdx[1]
	regB := rf.Read(idxB)
	if regB.Cmp(big.NewInt(0)) == 0 {
		offset := rf.Read(idxA).Uint64()
		cf.pc = offset
	} else {
		cf.pc += (OpLen + op.regNum*RegLen)
	}
	return nil
}

func opCall(cf *ContractFrame, regIdx []uint64, op *operation) error {
	defer decrementGas(&cf.gas, op)
	defer incrementPc(&cf.pc, op)
	rf := cf.rf
	idxA := regIdx[0]
	//TODO replace with real public address
	ctcAddr := uint64(rf.Read(idxA).Int64())
	msgSzeBig := rf.Read(rf.regSze - 1)
	msgSze := msgSzeBig.Uint64()
	idxC := rf.regSze - (1 + msgSze)
	amount := rf.Read(idxC)
	var msg []byte
	for i := 1; uint64(i) < msgSze; i++ {
		idx := idxC + uint64(i)
		msg = append(msg, rf.Read(idx).Bytes()...)
	}
	newCtc := NewContract(ctcAddr, cf.contract.Self, msg, amount,
		cf.contract.GasLimit, cf.contract.GasPrice)
	err := cf.CallContractFrame(newCtc)
	if err != nil {
		status := OpError{
			OpName: op.name,
			Cause:  err.Error(),
		}
		return status
	}
	return nil
}

func opRet(cf *ContractFrame, regIdx []uint64, op *operation) error {
	defer decrementGas(&cf.gas, op)
	defer incrementPc(&cf.pc, op)
	rf := cf.rf
	rfCaller := cf.caller.rf
	msgSzeBig := rf.Read(rf.regSze - 1)
	msgSze := msgSzeBig.Uint64()
	startIdx := rf.regSze - (1 + msgSze)
	for i := 0; uint64(i) <= msgSze; i++ {
		idx := startIdx + uint64(i)
		rfCaller.Write(idx, rf.Read(idx))
	}
	return nil
}

func opAnd(cf *ContractFrame, regIdx []uint64, op *operation) error {
	defer decrementGas(&cf.gas, op)
	defer incrementPc(&cf.pc, op)
	rf := cf.rf
	idxA := regIdx[0]
	idxB := regIdx[1]
	regA := rf.Read(idxA)
	regB := rf.Read(idxB)
	regA = regA.And(regA, regB)
	rf.Write(idxA, regA)
	return nil
}

func opOr(cf *ContractFrame, regIdx []uint64, op *operation) error {
	defer decrementGas(&cf.gas, op)
	defer incrementPc(&cf.pc, op)
	rf := cf.rf
	idxA := regIdx[0]
	idxB := regIdx[1]
	regA := rf.Read(idxA)
	regB := rf.Read(idxB)
	regA = regA.Or(regA, regB)
	rf.Write(idxA, regA)
	return nil
}

func opXor(cf *ContractFrame, regIdx []uint64, op *operation) error {
	defer decrementGas(&cf.gas, op)
	defer incrementPc(&cf.pc, op)
	rf := cf.rf
	idxA := regIdx[0]
	idxB := regIdx[1]
	regA := rf.Read(idxA)
	regB := rf.Read(idxB)
	regA = regA.Xor(regA, regB)
	rf.Write(idxA, regA)
	return nil
}

func opNot(cf *ContractFrame, regIdx []uint64, op *operation) error {
	defer decrementGas(&cf.gas, op)
	defer incrementPc(&cf.pc, op)
	rf := cf.rf
	idxA := regIdx[0]
	regA := rf.Read(idxA)
	regA = regA.Not(regA)
	rf.Write(idxA, regA)
	return nil
}

func opByte(cf *ContractFrame, regIdx []uint64, op *operation) error {
	defer decrementGas(&cf.gas, op)
	defer incrementPc(&cf.pc, op)
	rf := cf.rf
	idxC := regIdx[0]
	idxA := regIdx[1]
	offset := regIdx[2]
	byteA := rf.Read(idxA).Bytes()
	if offset >= uint64(len(byteA)) {
		status := OpError{
			OpName: op.name,
			Cause:  ErrorAryBoundary.Error(),
		}
		return status
	}
	byteIns := []byte{byteA[offset]}
	rf.Write(idxC, new(big.Int).SetBytes(byteIns))
	return nil
}

func opMov(cf *ContractFrame, regIdx []uint64, op *operation) error {
	defer decrementGas(&cf.gas, op)
	defer incrementPc(&cf.pc, op)
	rf := cf.rf
	idxA := regIdx[0]
	idxB := regIdx[1]
	regB := rf.Read(idxB)
	rf.Write(idxA, regB)
	return nil
}

func opMovi(cf *ContractFrame, regIdx []uint64, op *operation) error {
	defer decrementGas(&cf.gas, op)
	defer incrementPc(&cf.pc, op)
	rf := cf.rf
	idxA := regIdx[0]
	insB := regIdx[1]
	rf.Write(idxA, big.NewInt(int64(insB)))
	return nil
}

func opLdMsg(cf *ContractFrame, regIdx []uint64, op *operation) error {
	defer decrementGas(&cf.gas, op)
	defer incrementPc(&cf.pc, op)
	msg := cf.contract.Input
	rf := cf.rf
	idxA := regIdx[0]
	idxB := regIdx[1]
	regB := rf.Read(idxB).Uint64()
	offset := regIdx[2]
	if (regB + offset) > uint64(len(msg)) {
		status := OpError{
			OpName: op.name,
			Cause:  ErrorMsgBoundary.Error(),
		}
		return status
	}
	msgB := msg[regB : regB+offset]
	rf.Write(idxA, new(big.Int).SetBytes(msgB))
	return nil
}

func opLdMem(cf *ContractFrame, regIdx []uint64, op *operation) error {
	defer decrementGas(&cf.gas, op)
	defer incrementPc(&cf.pc, op)
	mem := cf.mem
	rf := cf.rf
	idxA := regIdx[0]
	idxB := regIdx[1]
	regB := rf.Read(idxB).Uint64()
	offset := regIdx[2]
	dataB, err := mem.Retrieve(regB, offset)
	if err != nil {
		status := OpError{
			OpName: op.name,
			Cause:  err.Error(),
		}
		return status
	}
	rf.Write(idxA, new(big.Int).SetBytes(dataB))
	return nil
}

func opLdStg(cf *ContractFrame, regIdx []uint64, op *operation) error {
	defer decrementGas(&cf.gas, op)
	defer incrementPc(&cf.pc, op)
	stg := cf.contract.Storage
	rf := cf.rf
	idxA := regIdx[0]
	idxB := regIdx[1]
	regB := rf.Read(idxB).Uint64()
	offset := regIdx[2]
	if (regB + offset) > uint64(len(stg)) {
		status := OpError{
			OpName: op.name,
			Cause:  ErrorStgBoundary.Error(),
		}
		return status
	}
	stgB := stg[regB : regB+offset]
	rf.Write(idxA, new(big.Int).SetBytes(stgB))
	return nil
}

func opStMem(cf *ContractFrame, regIdx []uint64, op *operation) error {
	defer decrementGas(&cf.gas, op)
	defer incrementPc(&cf.pc, op)
	mem := cf.mem
	rf := cf.rf
	idxA := regIdx[0]
	idxB := regIdx[1]
	regA := rf.Read(idxA).Uint64()
	dataB := rf.Read(idxB).Bytes()
	err := mem.Store(regA, uint64(len(dataB)), dataB)
	if err != nil {
		status := OpError{
			OpName: op.name,
			Cause:  err.Error(),
		}
		return status
	}
	return nil
}

func opStStg(cf *ContractFrame, regIdx []uint64, op *operation) error {
	defer decrementGas(&cf.gas, op)
	defer incrementPc(&cf.pc, op)
	stg := cf.contract.Storage
	rf := cf.rf
	idxA := regIdx[0]
	idxB := regIdx[1]
	regA := rf.Read(idxA).Uint64()
	stgB := rf.Read(idxB).Bytes()
	if regA+uint64(len(stgB)) > uint64(len(stg)) {
		appendSze := regA + uint64(len(stgB)) - uint64(len(stg))
		stg = append(stg, make([]byte, appendSze)...)
	}
	copy(stg[regA:regA+uint64(len(stgB))], stgB)
	return nil
}

func opInc(cf *ContractFrame, regIdx []uint64, op *operation) error {
	defer decrementGas(&cf.gas, op)
	defer incrementPc(&cf.pc, op)
	rf := cf.rf
	idxA := regIdx[0]
	regA := rf.Read(idxA)
	regA = regA.Add(regA, big.NewInt(1))
	rf.Write(idxA, regA)
	return nil
}

func opDec(cf *ContractFrame, regIdx []uint64, op *operation) error {
	defer decrementGas(&cf.gas, op)
	defer incrementPc(&cf.pc, op)
	rf := cf.rf
	idxA := regIdx[0]
	regA := rf.Read(idxA)
	regA = regA.Sub(regA, big.NewInt(1))
	rf.Write(idxA, regA)
	return nil
}

func opAdd(cf *ContractFrame, regIdx []uint64, op *operation) error {
	defer decrementGas(&cf.gas, op)
	defer incrementPc(&cf.pc, op)
	rf := cf.rf
	idxC := regIdx[0]
	idxA := regIdx[1]
	idxB := regIdx[2]
	regA := rf.Read(idxA)
	regB := rf.Read(idxB)
	regC := regA.Add(regA, regB)
	rf.Write(idxC, regC)
	return nil
}

func opAddi(cf *ContractFrame, regIdx []uint64, op *operation) error {
	defer decrementGas(&cf.gas, op)
	defer incrementPc(&cf.pc, op)
	rf := cf.rf
	idxC := regIdx[0]
	idxA := regIdx[1]
	insB := regIdx[2]
	regA := rf.Read(idxA)
	regC := regA.Add(regA, big.NewInt(int64(insB)))
	rf.Write(idxC, regC)
	return nil
}

func opSub(cf *ContractFrame, regIdx []uint64, op *operation) error {
	defer decrementGas(&cf.gas, op)
	defer incrementPc(&cf.pc, op)
	rf := cf.rf
	idxC := regIdx[0]
	idxA := regIdx[1]
	idxB := regIdx[2]
	regA := rf.Read(idxA)
	regB := rf.Read(idxB)
	regC := regA.Sub(regA, regB)
	rf.Write(idxC, regC)
	return nil
}

func opSubi(cf *ContractFrame, regIdx []uint64, op *operation) error {
	defer decrementGas(&cf.gas, op)
	defer incrementPc(&cf.pc, op)
	rf := cf.rf
	idxC := regIdx[0]
	idxA := regIdx[1]
	insB := regIdx[2]
	regA := rf.Read(idxA)
	regC := regA.Sub(regA, big.NewInt(int64(insB)))
	rf.Write(idxC, regC)
	return nil
}

func opMul(cf *ContractFrame, regIdx []uint64, op *operation) error {
	defer decrementGas(&cf.gas, op)
	defer incrementPc(&cf.pc, op)
	rf := cf.rf
	idxC := regIdx[0]
	idxA := regIdx[1]
	idxB := regIdx[2]
	regA := rf.Read(idxA)
	regB := rf.Read(idxB)
	regC := regA.Mul(regA, regB)
	rf.Write(idxC, regC)
	return nil
}

func opMuli(cf *ContractFrame, regIdx []uint64, op *operation) error {
	defer decrementGas(&cf.gas, op)
	defer incrementPc(&cf.pc, op)
	rf := cf.rf
	idxC := regIdx[0]
	idxA := regIdx[1]
	insB := regIdx[2]
	regA := rf.Read(idxA)
	regC := regA.Mul(regA, big.NewInt(int64(insB)))
	rf.Write(idxC, regC)
	return nil
}

func opDiv(cf *ContractFrame, regIdx []uint64, op *operation) error {
	defer decrementGas(&cf.gas, op)
	defer incrementPc(&cf.pc, op)
	rf := cf.rf
	idxC := regIdx[0]
	idxA := regIdx[1]
	idxB := regIdx[2]
	regA := rf.Read(idxA)
	regB := rf.Read(idxB)
	if regB.Sign() != 0 {
		regC := regA.Div(regA, regB)
		rf.Write(idxC, regC)
	} else {
		rf.Write(idxC, new(big.Int))
	}
	return nil
}

func opDivi(cf *ContractFrame, regIdx []uint64, op *operation) error {
	defer decrementGas(&cf.gas, op)
	defer incrementPc(&cf.pc, op)
	rf := cf.rf
	idxC := regIdx[0]
	idxA := regIdx[1]
	insB := regIdx[2]
	regA := rf.Read(idxA)
	if insB != 0 {
		regC := regA.Div(regA, big.NewInt(int64(insB)))
		rf.Write(idxC, regC)
	} else {
		rf.Write(idxC, new(big.Int))
	}
	return nil
}

func opMod(cf *ContractFrame, regIdx []uint64, op *operation) error {
	defer decrementGas(&cf.gas, op)
	defer incrementPc(&cf.pc, op)
	rf := cf.rf
	idxC := regIdx[0]
	idxA := regIdx[1]
	idxB := regIdx[2]
	regA := rf.Read(idxA)
	regB := rf.Read(idxB)
	if regB.Sign() != 0 {
		regC := regA.Mod(regA, regB)
		rf.Write(idxC, regC)
	} else {
		rf.Write(idxC, new(big.Int))
	}
	return nil
}

func opModi(cf *ContractFrame, regIdx []uint64, op *operation) error {
	defer decrementGas(&cf.gas, op)
	defer incrementPc(&cf.pc, op)
	rf := cf.rf
	idxC := regIdx[0]
	idxA := regIdx[1]
	insB := regIdx[2]
	regA := rf.Read(idxA)
	if insB != 0 {
		regC := regA.Mod(regA, big.NewInt(int64(insB)))
		rf.Write(idxC, regC)
	} else {
		rf.Write(idxC, new(big.Int))
	}
	return nil
}

func opExp(cf *ContractFrame, regIdx []uint64, op *operation) error {
	defer decrementGas(&cf.gas, op)
	defer incrementPc(&cf.pc, op)
	rf := cf.rf
	idxC := regIdx[0]
	idxA := regIdx[1]
	idxB := regIdx[2]
	regA := rf.Read(idxA)
	regB := rf.Read(idxB)
	regC := new(big.Int)
	regC.Exp(regA, regB, big.NewInt(1))
	rf.Write(idxC, regC)
	return nil
}

func opExpi(cf *ContractFrame, regIdx []uint64, op *operation) error {
	defer decrementGas(&cf.gas, op)
	defer incrementPc(&cf.pc, op)
	rf := cf.rf
	idxC := regIdx[0]
	idxA := regIdx[1]
	insB := regIdx[2]
	regA := rf.Read(idxA)
	regC := new(big.Int)
	regC.Exp(regA, big.NewInt(int64(insB)), big.NewInt(1))
	rf.Write(idxC, regC)
	return nil
}

func opEq(cf *ContractFrame, regIdx []uint64, op *operation) error {
	defer decrementGas(&cf.gas, op)
	defer incrementPc(&cf.pc, op)
	rf := cf.rf
	idxC := regIdx[0]
	idxA := regIdx[1]
	idxB := regIdx[2]
	regA := rf.Read(idxA)
	regB := rf.Read(idxB)
	if regA.Cmp(regB) == 0 {
		rf.Write(idxC, big.NewInt(1))
	} else {
		rf.Write(idxC, big.NewInt(0))
	}
	return nil
}

func opLt(cf *ContractFrame, regIdx []uint64, op *operation) error {
	defer decrementGas(&cf.gas, op)
	defer incrementPc(&cf.pc, op)
	rf := cf.rf
	idxC := regIdx[0]
	idxA := regIdx[1]
	idxB := regIdx[2]
	regA := rf.Read(idxA)
	regB := rf.Read(idxB)
	if regA.Cmp(regB) < 0 {
		rf.Write(idxC, big.NewInt(1))
	} else {
		rf.Write(idxC, big.NewInt(0))
	}
	return nil
}

func opGt(cf *ContractFrame, regIdx []uint64, op *operation) error {
	defer decrementGas(&cf.gas, op)
	defer incrementPc(&cf.pc, op)
	rf := cf.rf
	idxC := regIdx[0]
	idxA := regIdx[1]
	idxB := regIdx[2]
	regA := rf.Read(idxA)
	regB := rf.Read(idxB)
	if regA.Cmp(regB) > 0 {
		rf.Write(idxC, big.NewInt(1))
	} else {
		rf.Write(idxC, big.NewInt(0))
	}
	return nil
}

func opLte(cf *ContractFrame, regIdx []uint64, op *operation) error {
	defer decrementGas(&cf.gas, op)
	defer incrementPc(&cf.pc, op)
	rf := cf.rf
	idxC := regIdx[0]
	idxA := regIdx[1]
	idxB := regIdx[2]
	regA := rf.Read(idxA)
	regB := rf.Read(idxB)
	if regA.Cmp(regB) <= 0 {
		rf.Write(idxC, big.NewInt(1))
	} else {
		rf.Write(idxC, big.NewInt(0))
	}
	return nil
}

func opGte(cf *ContractFrame, regIdx []uint64, op *operation) error {
	defer decrementGas(&cf.gas, op)
	defer incrementPc(&cf.pc, op)
	rf := cf.rf
	idxC := regIdx[0]
	idxA := regIdx[1]
	idxB := regIdx[2]
	regA := rf.Read(idxA)
	regB := rf.Read(idxB)
	if regA.Cmp(regB) >= 0 {
		rf.Write(idxC, big.NewInt(1))
	} else {
		rf.Write(idxC, big.NewInt(0))
	}
	return nil
}

func opMax(cf *ContractFrame, regIdx []uint64, op *operation) error {
	defer decrementGas(&cf.gas, op)
	defer incrementPc(&cf.pc, op)
	rf := cf.rf
	idxC := regIdx[0]
	idxA := regIdx[1]
	idxB := regIdx[2]
	regA := rf.Read(idxA)
	regB := rf.Read(idxB)
	if regA.Cmp(regB) >= 0 {
		rf.Write(idxC, regA)
	} else {
		rf.Write(idxC, regB)
	}
	return nil
}

func opMin(cf *ContractFrame, regIdx []uint64, op *operation) error {
	defer decrementGas(&cf.gas, op)
	defer incrementPc(&cf.pc, op)
	rf := cf.rf
	idxC := regIdx[0]
	idxA := regIdx[1]
	idxB := regIdx[2]
	regA := rf.Read(idxA)
	regB := rf.Read(idxB)
	if regA.Cmp(regB) <= 0 {
		rf.Write(idxC, regA)
	} else {
		rf.Write(idxC, regB)
	}
	return nil
}
