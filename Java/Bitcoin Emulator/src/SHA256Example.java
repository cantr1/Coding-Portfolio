import java.nio.charset.StandardCharsets;
import java.security.MessageDigest;
import java.security.NoSuchAlgorithmException;

public class SHA256Example {

    /**
     * Converts a byte array to a hexadecimal string
     */
    private static String bytesToHex(byte[] hash) {
        StringBuilder hexString = new StringBuilder();
        for (byte b : hash) {
            String hex = Integer.toHexString(0xff & b);
            if (hex.length() == 1) {
                hexString.append('0');
            }
            hexString.append(hex);
        }
        return hexString.toString();
    }

    /**
     * Generate SHA-256 hash for a given input string
     */
    public static String generateSHA256(String input) throws NoSuchAlgorithmException {
        MessageDigest digest = MessageDigest.getInstance("SHA-256");
        byte[] hashBytes = digest.digest(input.getBytes(StandardCharsets.UTF_8));
        return bytesToHex(hashBytes);
    }

    public static void main(String[] args) {
        try {
            String input = "Hello, World!";
            String sha256Hash = generateSHA256(input);

            System.out.println("Input: " + input);
            System.out.println("SHA-256 Hash: " + sha256Hash);

            // Verify that the same input always produces the same hash
            String anotherHash = generateSHA256(input);
            System.out.println("Hash matches: " + sha256Hash.equals(anotherHash));

            // Demonstrate that even a small change produces a completely different hash
            String modifiedInput = "Hello, World";
            String modifiedHash = generateSHA256(modifiedInput);
            System.out.println("Modified input: " + modifiedInput);
            System.out.println("Modified hash: " + modifiedHash);
            System.out.println("Hashes are different: " + !sha256Hash.equals(modifiedHash));

        } catch (NoSuchAlgorithmException e) {
            System.err.println("SHA-256 algorithm not available: " + e.getMessage());
        }
    }
}
