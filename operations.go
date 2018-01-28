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

type OpCode byte

//TODO Complete the instruction list
//Flow control operations
const (
	NOP OpCode = iota
	ABORT
	JMP
	JMPI
	JMPIF
	JMPIFNOT
	CALL
	RET
	CREATE
)

//Bitwise logic
const (
	AND OpCode = iota + 0x31
	OR
	XOR
	NOT
	BYTE
	SHL
	SHR
)

//Load/Store
const (
	MOV OpCode = iota + 0x41
	MOVI
	LDMSG
	LDMEM
	LDSTG
	STMEM
	STSTG
)

//Arithmetic oeprations
const (
	INC OpCode = iota + 0x51
	DEC
	NEG
	ABS
	ADD
	ADDI
	SUB
	SUBI
	MUL
	MULI
	DIV
	DIVI
	MOD
	MODI
	EXP
	EXPI
	EQ
	LT
	GT
	LTE
	GTE
	MAX
	MIN
)

type operation struct {
	execute executionFunc
	gasCost uint64
	regNum  uint64
	name    string
}

type executionFunc func(cf *ContractFrame, regIdx []uint64, op *operation) error

type OperationSet [256]operation

func NewOperationSet() OperationSet {
	return OperationSet{
		//Flow Control
		NOP: {
			execute: opNop,
			gasCost: 0,
			regNum:  0,
			name:    "opNop",
		},

		JMP: {
			execute: opJmp,
			gasCost: 2,
			regNum:  1,
			name:    "opJmp",
		},

		JMPI: {
			execute: opJmpi,
			gasCost: 2,
			regNum:  1,
			name:    "opJmpi",
		},

		JMPIF: {
			execute: opJmpif,
			gasCost: 2,
			regNum:  2,
			name:    "opJmpif",
		},

		JMPIFNOT: {
			execute: opJmpifnot,
			gasCost: 2,
			regNum:  2,
			name:    "opJmpifnot",
		},

		CALL: {
			execute: opCall,
			gasCost: 2,
			regNum:  1,
			name:    "opCall",
		},

		RET: {
			execute: opRet,
			gasCost: 2,
			regNum:  0,
			name:    "opRet",
		},

		//Bitwise logic
		AND: {
			execute: opAnd,
			gasCost: 2,
			regNum:  2,
			name:    "opAnd",
		},

		OR: {
			execute: opOr,
			gasCost: 2,
			regNum:  2,
			name:    "opOr",
		},

		XOR: {
			execute: opXor,
			gasCost: 2,
			regNum:  2,
			name:    "opXor",
		},

		NOT: {
			execute: opNot,
			gasCost: 2,
			regNum:  2,
			name:    "opNot",
		},

		BYTE: {
			execute: opByte,
			gasCost: 2,
			regNum:  3,
			name:    "opByte",
		},

		//Load/Store
		MOV: {
			execute: opMov,
			gasCost: 1,
			regNum:  2,
			name:    "opMov",
		},

		MOVI: {
			execute: opMovi,
			gasCost: 1,
			regNum:  2,
			name:    "opMovi",
		},

		LDMSG: {
			execute: opLdMsg,
			gasCost: 1,
			regNum:  3,
			name:    "opLdMsg",
		},

		LDMEM: {
			execute: opLdMem,
			gasCost: 1,
			regNum:  3,
			name:    "opLdMem",
		},

		LDSTG: {
			execute: opLdStg,
			gasCost: 1,
			regNum:  3,
			name:    "opLdStg",
		},

		STMEM: {
			execute: opStMem,
			gasCost: 1,
			regNum:  2,
			name:    "opStMem",
		},

		STSTG: {
			execute: opStStg,
			gasCost: 1,
			regNum:  2,
			name:    "opStStg",
		},

		//Arithmetic
		INC: {
			execute: opInc,
			gasCost: 1,
			regNum:  1,
			name:    "opInc",
		},

		DEC: {
			execute: opDec,
			gasCost: 1,
			regNum:  1,
			name:    "opDec",
		},

		ADD: {
			execute: opAdd,
			gasCost: 1,
			regNum:  3,
			name:    "opAdd",
		},

		ADDI: {
			execute: opAddi,
			gasCost: 1,
			regNum:  3,
			name:    "opAddi",
		},

		SUB: {
			execute: opSub,
			gasCost: 1,
			regNum:  3,
			name:    "opSub",
		},

		SUBI: {
			execute: opSubi,
			gasCost: 1,
			regNum:  3,
			name:    "opSubi",
		},

		MUL: {
			execute: opMul,
			gasCost: 2,
			regNum:  3,
			name:    "opMul",
		},

		MULI: {
			execute: opMuli,
			gasCost: 2,
			regNum:  3,
			name:    "opMuli",
		},

		DIV: {
			execute: opDiv,
			gasCost: 2,
			regNum:  3,
			name:    "opDiv",
		},

		DIVI: {
			execute: opDivi,
			gasCost: 2,
			regNum:  3,
			name:    "opDivi",
		},

		MOD: {
			execute: opMod,
			gasCost: 2,
			regNum:  3,
			name:    "opMod",
		},

		MODI: {
			execute: opModi,
			gasCost: 2,
			regNum:  3,
			name:    "opModi",
		},

		EXP: {
			execute: opExp,
			gasCost: 2,
			regNum:  3,
			name:    "opExp",
		},

		EXPI: {
			execute: opExpi,
			gasCost: 2,
			regNum:  3,
			name:    "opExpi",
		},

		EQ: {
			execute: opEq,
			gasCost: 2,
			regNum:  3,
			name:    "opEq",
		},

		LT: {
			execute: opLt,
			gasCost: 2,
			regNum:  3,
			name:    "opLt",
		},

		GT: {
			execute: opGt,
			gasCost: 2,
			regNum:  3,
			name:    "opGt",
		},

		LTE: {
			execute: opLte,
			gasCost: 2,
			regNum:  3,
			name:    "opLte",
		},

		GTE: {
			execute: opGte,
			gasCost: 2,
			regNum:  3,
			name:    "opGte",
		},

		MAX: {
			execute: opMax,
			gasCost: 2,
			regNum:  3,
			name:    "opMax",
		},

		MIN: {
			execute: opMin,
			gasCost: 2,
			regNum:  3,
			name:    "opMin",
		},
	}
}
