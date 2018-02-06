public class JVMOpCodes {
	// iconst_n
	public static final int ICONST_M1 	= 0x02; 
	public static final int ICONST_0 	= 0x03; 
	public static final int ICONST_1 	= 0x04; 
	public static final int ICONST_2 	= 0x05; 
	public static final int ICONST_3 	= 0x06; 
	public static final int ICONST_4 	= 0x07; 
	public static final int ICONST_5 	= 0x08; 
	
	// istore_n
	public static final int ISTORE_0 	= 0x3b; 
	public static final int ISTORE_1 	= 0x3c; 
	public static final int ISTORE_2 	= 0x3d; 
	public static final int ISTORE_3 	= 0x3e; 
	
	// iload_n
	public static final int ILOAD_0 		= 0x1a; 
	public static final int ILOAD_1 		= 0x1b; 
	public static final int ILOAD_2 		= 0x1c; 
	public static final int ILOAD_3 		= 0x1d; 
		
	// Arithmetic
	public static final int IINC 		= 0x84; 
	public static final int IADD 		= 0x60; 
    
	// ireturn
	public static final int IRETURN	 	= 0xAC; 
}
