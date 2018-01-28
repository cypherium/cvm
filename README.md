# Cypherium Virtual Machine Specification 
The Cypherium Virtual Machine (CVM) is a register-based virtual machine inspired 
by the [Dalvik](https://source.android.com/devices/tech/dalvik/)
(now part of the Android Runtime) architecture. The virtual machine from 
[Ethereum VM](https://github.com/ethereum/wiki/wiki/White-Paper) and [NEO 
VM](https://github.com/neo-project/neo-vm) also influenced its design. It is 
used for executing smart contracts on the Cypherium blockchain platform. A 
reference implementation is provided using the [Go language](https://golang.org/).

## Design Principle
CVM favors small and simple smart contracts that preferably only use the 65536 
available registers (registers can be considered as directly addressable memory) 
without allocating extra memory. We envision a blockchain where there is a rich 
set of simple smart contracts that follow the UNIX philosophy "do one thing and 
do it well" at contract developers' disposal. Those contracts would together form 
a common library so that future contract developers will just call them for 
frequently-used functionalities. This, we reckon, promotes the code reusability 
which in turn prevents the blockchain size from bloating by eliminating code 
redundency.

The smart contracts that comprise the common library would preferably be 
developed and added into the blockchain from the team itself and the contract 
developing community, since they are guarenteed to provide high-quality code. 
Nevertheless, once on the blockchain, those
contracts and their codes will be under public scrutiny. 

## Description
* Each contract is allocated with a contract frame with fixed size upon 
  creation. Each contract frame consists of a particular number of registers 
  (16-bit indexing namely, up to 65536) as well as any adjunct data needed to 
  execute the contract. 
* Each contract can access and manipulate data from 5 locations:
  1. Registers from its own contract frame (r/w)
  1. Persistent contract storage (r/w)
  1. On-chain data (ro)
  1. Message from the transaction that triggered the contract (ro)
  1. An **optional** infinitely expending memory (r/w) 
* All memory region but the persistent contract storage will be cleared and 
  deallocated upon exit.
* Use relative offset for jumping/branching
* Gas mechanism similar to that of the EVM
* Use a PC counter to keep track of the next instruction
* Iteration counter (counts conditional jumps to prevent dead loops)

## Instruction Format
* Dest-Source ordering format
* 8-bit opcode 
* 16-bit register index 
* Registers are able to hold arbitrary-precision data

## Call Convention
Upon calling the CALL op, the runtime system checks the content of the register 
65535 for the size N of the message, the actual content is constructed using the 
register (65535-1)-N to 65535-1. The message will be the input message to the 
callee.

Upon calling the the RET op for returning, the runtime will copy the content 
of N last consecutive registers(65535-1-N to 65535-1) from the callee's contract
frame, specified by the register 65535 from the same contract frame, to the 
caller's last N registers.

## Instruction Sets
* 256 (0x100) opcodes, 
* %A, %B, etc. stands for arbitrary registers
* $INS stands for an instant value

### Flow control (0x000 - 0x010)
- NOP : Do nothing
  + NOP
- ABORT : Abort the current execution and reverse all the changes
  + ABORT
- JMP : PC += %A/$INS
  + JMP %A 
  + JMPI $INS
- JMPIF : if %B: PC += %A
  + JMPIF %A %B 
- JMPIFNOT : if !%B: PC += %A
  + JMPIFNOT %A %B
- CALL : Call another contract with the contract address in %A
  + CALL %A 
- RET: Return control to the caller 
  + RET 
- CREATE : Create a smart contract
  + CREATE

### Environment (0x011 - 0x030)
- ADDRESS : Get address of currently executing account to %A
  + ADDRESS %A
- BALANCE : Get balance of the given account to %A
  + BALANCE %A
- ORIGIN : Get the sender's address to %A
  + ORIGIN %A
- CALLER : Get the callers address to %A
  + CALLER %A
- CALLVALUE : Get deposited value
  + CALLVALUE %A
- MSGSIZE :
  + MSGSIZE %A
- GASPRICE : Get gasprice
  + GASPRICE %A
- GAS : Get avaiable gas
  + GAS %A
- BLOCKHASH : Get the hash of one of the complete micro blocks (%A contains the 
  index of the blocks and %A also will contain the result)
  + BLOCKHASH %A 
- COINBASE : Get one of the micro block %B's beneficiary address (%A contains 
  the index of the blocks and %A also will contain the result)
  + COINBASE %A 
- TIMESTAMP : Get one of the key block's timestamp (%A contains the index of the 
  blocks and %A also will contain the result)
  + TIMESTAMP %A 
- DIFFICULTY : Get one of the key blocks' difficulty (%A contains the index of 
  the blocks and %A also will contain the result)
  + DIFFICULTY %A

### Bitwise logic (0x031 - 0x040)
- AND : %A = %A & %B
  + AND %A %B
- OR : %A = %A | %B
  + OR %A %B
- XOR : %A = %A XOR %B
  + XOR %A %B
- NOT : %A = ~%A
  + NOT %A
- BYTE : Retrieve byte %A[$INS] to %C
  + BYTE %C %A $INS

### Load/Store (0x041 - 0x050)
- MOV : Load value %B/$INS to %A
  + MOV %A %B
  + MOVI %A $INS
- LDMSG : Load $INS-byte from message starting from index %B to %A 
  + LDMSG %A %B $INS
- LDMEM : Load $INS-byte from memory starting from index %B to %A 
  + LDMEM %A %B $INS
- LDSTG : Load $INS-byte from storage starting from index %B to %A
  + LDSTG %A %B $INS
- STMEM : Store %B to memory starting from index %A
  + STMEM %A %B
- STSTG : Store %B to storage starting from index %A
  + STSTG %A %B

### Arithmetic  (0x051 - 0x090) 
In this section, the instant values ($INS) are interpreted as 64-bit signed integers
- INC : %A += 1
  + INC %A
- DEC : %A -= 1
  + DEC %A
- NEG : %A = -%A
  + NEG %A
- ABS : %B = |%A|
  + ABS %B %A
- ADD
  + ADD %C %A %B
  + ADDI %C %A $INS 
- SUB
  + SUB %C %A %B
  + SUBI %C %A $INS
- MUL
  + MUL %C %A %B
  + MULI %C %A $INS
- DIV
  + DIV %C %A %B
  + DIVI %C %A $INS
- MOD
  + MOD %C %A %B
  + MODI %C %A $INS
- EXP : %C = %A ^ %B/$INS
  + EXP %C %A %B
  + EXPI %C %A $INS
- EQ : %C = %A == %B/$INS ? 1 : 0
  + EQ %C %A %B
  + EQI %C %A $INS
- LT : %C = %A < %B/$INS ? 1 : 0
  + LT %C %A %B
  + LTI %C %A $INS
- GT : %C = %A > %B/$INS ? 1 : 0
  + GT %C %A %B
  + GTI %C %A $INS
- LTE : %C = %A <= %B/$INS ? 1 : 0
  + LTE %C %A %B
  + LTEI %C %A $INS
- GTE : %C = %A >= %B/$INS ? 1 : 0
  + GTE %C %A %B
  + GTEI %C %A $INS
- MAX : %C = %A > %B/$INS ? %A : %B/$INS
  + MAX %C %A %B
  + MAXI %C %A $INS
- MIN : %C = %A < %B/$INS ? %A : %B/$INS
  + MIN %C %A %B
  + MINI %C %A $INS

### Cryptography (0x091 - 0x100)
- SHA3
- CHKSIG
- CHKMULSIG
