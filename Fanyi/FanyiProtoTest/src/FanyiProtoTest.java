import java.util.HashMap;

public class FanyiProtoTest {
	int REGSIZE = 65536;
	int Register[];  // Register of only ints for now
	int RegisterIndex;
	byte[] data; // JVM bytecodes
	int JByteCodeArrayIndex;
	String REG_PREFIX = " R";
	
	// Limits for handling embedded operand instructions
	// iconst, istore, iload
	int ICONST_START = JVMOpCodes.ICONST_M1, ICONST_END = JVMOpCodes.ICONST_5;
	int ISTORE_START = JVMOpCodes.ISTORE_0, ISTORE_END = JVMOpCodes.ISTORE_3;
	int ILOAD_START = JVMOpCodes.ILOAD_0, ILOAD_END = JVMOpCodes.ILOAD_3;
	
	// Num of operands
	public static int ZERO_OPERANDS = 0;
	public static int ONE_OPERAND = 1;
	public static int TWO_OPERANDS = 2;
	
	OpCode[] OpCodes = { 
			new OpCode(JVMOpCodes.ICONST_M1, ZERO_OPERANDS, CVMOpCodes.MOV, InstructionFamily.EMBEDDED_OPERANDS), // iconst_m1, cvmopcode MOV 
			new OpCode(JVMOpCodes.ICONST_0, ZERO_OPERANDS, CVMOpCodes.MOV, InstructionFamily.EMBEDDED_OPERANDS), // iconst_0, cvmopcode MOV
			new OpCode(JVMOpCodes.ICONST_1, ZERO_OPERANDS, CVMOpCodes.MOV, InstructionFamily.EMBEDDED_OPERANDS), // iconst_1, cvmopcode MOV
			new OpCode(JVMOpCodes.ICONST_2, ZERO_OPERANDS, CVMOpCodes.MOV, InstructionFamily.EMBEDDED_OPERANDS), // iconst_2, cvmopcode MOV
			new OpCode(JVMOpCodes.ICONST_3, ZERO_OPERANDS, CVMOpCodes.MOV, InstructionFamily.EMBEDDED_OPERANDS), // iconst_3, cvmopcode MOV
			new OpCode(JVMOpCodes.ICONST_4, ZERO_OPERANDS, CVMOpCodes.MOV, InstructionFamily.EMBEDDED_OPERANDS), // iconst_4, cvmopcode MOV
			new OpCode(JVMOpCodes.ICONST_5, ZERO_OPERANDS, CVMOpCodes.MOV, InstructionFamily.EMBEDDED_OPERANDS), // iconst_5, cvmopcode MOV
			
			new OpCode(JVMOpCodes.ISTORE_0, ZERO_OPERANDS, CVMOpCodes.STMEM, InstructionFamily.EMBEDDED_OPERANDS), // istore_0, cvmopcode STMEM
			new OpCode(JVMOpCodes.ISTORE_1, ZERO_OPERANDS, CVMOpCodes.STMEM, InstructionFamily.EMBEDDED_OPERANDS), // istore_1, cvmopcode STMEM
			new OpCode(JVMOpCodes.ISTORE_2, ZERO_OPERANDS, CVMOpCodes.STMEM, InstructionFamily.EMBEDDED_OPERANDS), // istore_2, cvmopcode STMEM
			new OpCode(JVMOpCodes.ISTORE_3, ZERO_OPERANDS, CVMOpCodes.STMEM, InstructionFamily.EMBEDDED_OPERANDS), // istore_3, cvmopcode STMEM
            
			new OpCode(JVMOpCodes.ILOAD_0, ZERO_OPERANDS, CVMOpCodes.LDMEM, InstructionFamily.EMBEDDED_OPERANDS), // iload_0, cvmopcode LDMEM
			new OpCode(JVMOpCodes.ILOAD_1, ZERO_OPERANDS, CVMOpCodes.LDMEM, InstructionFamily.EMBEDDED_OPERANDS), // iload_1, cvmopcode LDMEM
			new OpCode(JVMOpCodes.ILOAD_2, ZERO_OPERANDS, CVMOpCodes.LDMEM, InstructionFamily.EMBEDDED_OPERANDS), // iload_2, cvmopcode LDMEM
			new OpCode(JVMOpCodes.ILOAD_3, ZERO_OPERANDS, CVMOpCodes.LDMEM, InstructionFamily.EMBEDDED_OPERANDS), // iload_3, cvmopcode LDMEM
			
			new OpCode(JVMOpCodes.IINC, TWO_OPERANDS, CVMOpCodes.INC, InstructionFamily.ARITHMETIC), // iinc, cvmopcode INC
			new OpCode(JVMOpCodes.IADD, ZERO_OPERANDS, CVMOpCodes.ADD, InstructionFamily.ARITHMETIC), // iadd, cvmopcode ADD

			new OpCode(JVMOpCodes.IRETURN, ZERO_OPERANDS, CVMOpCodes.RET, InstructionFamily.RETURN) // ireturn, cvmopcode RET
	}; 
	
	HashMap<Integer, Integer> opcodemap;

	public FanyiProtoTest() {
		RegisterIndex = 0;
		Register = new int[REGSIZE];
		
		JByteCodeArrayIndex = 0;
		
		BuildOpCodeMap();	
	}
	
	public void BuildOpCodeMap() {
		int index = 0;
		opcodemap = new HashMap<Integer, Integer>();
		
		// Just a subset for now
		opcodemap.put(JVMOpCodes.ICONST_M1, index++); // iconst_m1
		opcodemap.put(JVMOpCodes.ICONST_0, index++); // iconst_0
		opcodemap.put(JVMOpCodes.ICONST_1, index++); // iconst_1
		opcodemap.put(JVMOpCodes.ICONST_2, index++); // iconst_2
		opcodemap.put(JVMOpCodes.ICONST_3, index++); // iconst_3
		opcodemap.put(JVMOpCodes.ICONST_4, index++); // iconst_4
		opcodemap.put(JVMOpCodes.ICONST_5, index++); // iconst_5
		
		opcodemap.put(JVMOpCodes.ISTORE_0, index++); // istore_1
		opcodemap.put(JVMOpCodes.ISTORE_1, index++); // istore_1
		opcodemap.put(JVMOpCodes.ISTORE_2, index++); // istore_1
		opcodemap.put(JVMOpCodes.ISTORE_3, index++); // istore_1
		
		opcodemap.put(JVMOpCodes.ILOAD_1, index++); // iload_1
		opcodemap.put(JVMOpCodes.ILOAD_1, index++); // iload_1
		opcodemap.put(JVMOpCodes.ILOAD_1, index++); // iload_1
		opcodemap.put(JVMOpCodes.ILOAD_1, index++); // iload_1
		
		opcodemap.put(JVMOpCodes.IINC, index++); // iinc
		opcodemap.put(JVMOpCodes.IADD, index++); // iadd
		
		opcodemap.put(JVMOpCodes.IRETURN, index); // ireturn
	}
	
	public static byte[] HexStringToByteArray(String s) {
	    int len = s.length();
	    byte[] data = new byte[len/2];
	    for (int i = 0; i<len; i += 2) {
	        data[i/2] = (byte) ((Character.digit(s.charAt(i), 16) << 4)
	                             + Character.digit(s.charAt(i+1), 16));
	    }
	    return data;
	}
	
	public void LookupOpCodeAndTranslate(int jopcode) {
		int opcodeindex = 0;
		
		if (opcodemap.containsKey(jopcode))
			opcodeindex = opcodemap.get(jopcode);
		
		int numoperands = OpCodes[opcodeindex].getNumOperands();
		
		switch (numoperands) {
			case 0: TranslateOpCodeWithZeroOperands(opcodeindex);
					break;
			case 1: TranslateOpCodeWithOneOperand(opcodeindex);
					break;
			case 2: TranslateOpCodeWithTwoOperands(opcodeindex);
					break;
			default: break;
			
		}
	}
	
	// Process Class file, currently just reading bytecodes for inc()
	// and dumping into byte stream
	// TODO: Parse JVM instructions, track other info from class file (e.g. local variables table etc.) 
	public void ProcessClassFile(String jvmbytes) {	
		// Read class into byte array
		data = HexStringToByteArray(jvmbytes);
	}
	
	// Build CFG here
	public void DecodeByteStream() {
		while (JByteCodeArrayIndex < data.length) {
			int jopcode = data[JByteCodeArrayIndex];
			LookupOpCodeAndTranslate(jopcode);
		}
	}
	
	public void HandleEmbeddedOperandInstructions(int opcodeindex) {
		// e.g. shows only iconst_n --> move rx, n
		// need to handle i_store, i_load etc.
		
		// iconst_n
		String instruction = "";
		int offset = 0;
		
		int jvmopcode = OpCodes[opcodeindex].getJVMOpcode();
		if ((jvmopcode >= ICONST_START) && (jvmopcode <= ICONST_END)) {
			offset = jvmopcode - ICONST_START - 1;
			instruction = OpCodes[opcodeindex].getCVMOpcode() + REG_PREFIX + (RegisterIndex++) + " " + offset;
			System.out.println(instruction);
		}
	}
	
	public void HandleArithmeticInstructions(int opcodeindex) {
		// e.g. iadd
		String instruction = "";
	
		instruction = OpCodes[opcodeindex].getCVMOpcode() + REG_PREFIX + (RegisterIndex++) + REG_PREFIX + (RegisterIndex++) + REG_PREFIX + (RegisterIndex++);
		System.out.println(instruction);
	}
	
	public void HandleReturnInstructions(int opcodeindex) {
		// e.g. ireturn
		String instruction = "";
	
		instruction = OpCodes[opcodeindex].getCVMOpcode() + "";
		System.out.println(instruction);
	}
	
	public void TranslateOpCodeWithZeroOperands(int opcodeindex) {
		if (OpCodes[opcodeindex].getInstructionFamily() == InstructionFamily.EMBEDDED_OPERANDS) {
			// Embedded operands
			HandleEmbeddedOperandInstructions(opcodeindex);
		}
		else if (OpCodes[opcodeindex].getInstructionFamily() == InstructionFamily.ARITHMETIC) {
			// Arithmetic instructions
			HandleArithmeticInstructions(opcodeindex);
		}	
		else if (OpCodes[opcodeindex].getInstructionFamily() == InstructionFamily.RETURN) {
			// e.g. ireturn
			HandleReturnInstructions(opcodeindex);
		}
		else {
			// TODO undefined instruction
		}
		
		JByteCodeArrayIndex++;
	}
	
	// TODO 
	public void TranslateOpCodeWithOneOperand(int opcodeindex) {
		JByteCodeArrayIndex+=2;
	}
	
	// TODO
	public void TranslateOpCodeWithTwoOperands(int opcodeindex) {
		JByteCodeArrayIndex+=3;
	}
	
	public static void main(String[] args) {
		String jvmbytes = "050660"; // iconst_2, iconst_3, iadd
		
		FanyiProtoTest fanyi = new FanyiProtoTest();
		fanyi.ProcessClassFile(jvmbytes);
		fanyi.DecodeByteStream();
	}
}
