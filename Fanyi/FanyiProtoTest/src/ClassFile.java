public class ClassFile {	
	private static final int MAGIC = 0xCAFEBABE;
	private static final int MAGIC_START_INDEX = 0;
	private static final int MINOR_VERSION_START_INDEX = 4;
	private static final int MAJOR_VERSION_START_INDEX = 6;
	
	private String fileName;

    private byte[] bytes;
    
    private int magic;

    private int minorVersion;
    
    private int majorVersion;

    public ClassFile(String fileName, byte[] bytes) {
    		this.fileName = fileName;
        this.bytes = bytes;
    }

    public String getFileName() {
    		return fileName;
    }

    public byte[] getBytes() {
        return bytes;
    }
    
    public int getMagic() {
    		magic = (bytes[MAGIC_START_INDEX] << 24) | ((bytes[MAGIC_START_INDEX + 1] & 0xff) << 16) | ((bytes[MAGIC_START_INDEX + 2] & 0xff) << 8) | (bytes[MAGIC_START_INDEX + 3] & 0xff);
    		System.out.println("Retrieved magic : " + magic);
    		System.out.println("Expected magic : " + MAGIC);
    		return magic;
    }

    public int getMinorVersion() {
    		minorVersion = ((bytes[MINOR_VERSION_START_INDEX] & 0xff) << 8) | (bytes[MINOR_VERSION_START_INDEX + 1] & 0xff);
    		System.out.println("Minor version : " + minorVersion);
    		return minorVersion;
    }
 
    public int getMajorVersion() {
    		majorVersion = ((bytes[MAJOR_VERSION_START_INDEX] & 0xff) << 8) | (bytes[MAJOR_VERSION_START_INDEX + 1] & 0xff);
    		System.out.println("Major version : " + majorVersion);
		return majorVersion;
    } 
}
