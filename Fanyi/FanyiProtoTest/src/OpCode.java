public class OpCode {
		private int jvm_opcode;
		private int num_operands;
		private int cvm_opcode;
		private int family;
		
		public OpCode(int jvmopcode, int numoperands, int cvmopcode, int family) {
			this.jvm_opcode = jvmopcode;
			this.num_operands = numoperands;
			this.cvm_opcode = cvmopcode;
			this.family = family;
		}
		
		public int getJVMOpcode() {
			return this.jvm_opcode;
		}
		
		public int getNumOperands() {
			return this.num_operands;
		}
		
		public int getCVMOpcode() {
			return this.cvm_opcode;
		}
		
		public int getInstructionFamily() {
			return this.family;
		}
}
