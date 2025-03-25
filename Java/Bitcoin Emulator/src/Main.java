import java.security.NoSuchAlgorithmException;
import java.util.Scanner;

public class Main {

    public static void main(String[] args) {
        try {

            Resources.introText();

            System.out.println("-".repeat(60));

            //Initialize block ID outside the loop
            int blockID = 0;
            String previousHash = "";

            for(int i = 0; i <= 100000; i++) {
                StringBuilder inputData = new StringBuilder();
                inputData.append(i);
                String coinHash = Hashing.generateSHA256(String.valueOf(inputData));

                if (coinHash.startsWith("000")) {
                    blockID += 1;

                    //String userData = createMessage();

                    // Block header
                    System.out.println("┌" + "─".repeat(98) + "┐");
                    System.out.println("│ BLOCK #" + blockID + String.format("%" + (90 - String.valueOf(blockID).length()) + "s", " ") + "│");
                    System.out.println("├" + "─".repeat(98) + "┤");

                    // Current block data
                    System.out.println("│ Input Data: " + inputData + String.format("%" + (85 - inputData.length()) + "s", " ") + "│");
                    System.out.println("│ Hash: " + coinHash + String.format("%" + (91 - coinHash.length()) + "s", " ") + "│");

                    //Display message
                    //System.out.println("│ Message: " + userData + String.format("%" + (88 - userData.length()) + "s", " ") + "│");

                    // Previous block data (if not the first block)
                    if (blockID > 1) {
                        System.out.println("├" + "─".repeat(98) + "┤");
                        System.out.println("│ Previous Block: #" + (blockID - 1) + String.format("%" + (80 - String.valueOf(blockID - 1).length()) + "s", " ") + "│");
                        System.out.println("│ Previous Hash: " + previousHash + String.format("%" + (82 - previousHash.length()) + "s", " ") + "│");
                    }

                    // Block footer
                    System.out.println("└" + "─".repeat(98) + "┘");
                    System.out.println(); // Empty line between blocks

                    previousHash = coinHash;
                }
            }


        } catch (NoSuchAlgorithmException e) {
            System.err.println("SHA-256 algorithm not available: " + e.getMessage());
        }

    }

    public static String createMessage(){
        String message;
        Scanner sc = new Scanner(System.in);

        System.out.println("Input your message! ");
        message = sc.nextLine();

        return message;
    }

}
