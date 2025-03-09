import java.util.Random;
import java.util.Scanner;

public class Main {
    public static void main(String[] args) {
        higherLower();
    }

    // Main game loop
    public static void higherLower(){
        boolean gameIsOn = true;

        // Create a scanner object
        Scanner scanner = new Scanner(System.in);

        while(gameIsOn){
            // Create a variable to store the number
            int randomInt = randomNumber();
            //System.out.println(randomInt);

            System.out.println("Guess a number between 0 and 100");

            // Create a loop for guessing the number
            boolean guessing = true;
            while(guessing){
                // Capture user guess
                System.out.print("Guess: ");
                int userGuess = scanner.nextInt();

                // Compare the guess
                if(compareGuess(userGuess, randomInt)){
                    System.out.println("Correct!");
                    break;
                } else {
                    System.out.println("Incorrect...");
                    if(userGuess > randomInt){
                        System.out.print("Lower ---  ");
                        hotCold(userGuess, randomInt);
                    } else {
                        System.out.print("Higher --- ");
                        hotCold(userGuess, randomInt);
                    }
                }
            }

            // After reading an integer with scanner.nextInt()
            scanner.nextLine(); // consume the leftover newline character

            // Ask user if they would like to continue
            boolean confirming = true;
            while(confirming){
                System.out.println("Would you like to play again? (yes or no)");
                String confirmation = scanner.nextLine();

                // Check for confirmation
                switch (confirmation){
                    case "yes" -> {
                        gameIsOn = true;
                        confirming = false;
                        for(int i = 0; i <= 20; i++){
                            System.out.println();
                        }
                    }
                    case "no" -> {
                        gameIsOn = false;
                        confirming = false;
                    }
                    default -> {
                        System.out.println("Invalid input");
                    }
                }
            }
        }
    }

    // Generate a random number from 0 to 100
    public static int randomNumber(){
        // Create a Random object
        Random random = new Random();

        int randomInt = random.nextInt(101);

        return randomInt;
    }

    // Compare the guess to the stored value
    public static boolean compareGuess(int userGuess, int randomInt){
        if(userGuess == randomInt){
            return true;
        } else {
            return false;
        }
    }

    // Tell the user if they are hot or cold
    public static void hotCold(int userGuess, int randomInt){
        if(Math.abs(randomInt - userGuess) >= 31){
            System.out.println("Very Cold");
        } else if(Math.abs(randomInt - userGuess) >= 21 && Math.abs(randomInt - userGuess) <= 30){
            System.out.println("Cold");
        } else if(Math.abs(randomInt - userGuess) >= 11 && Math.abs(randomInt - userGuess) <= 20 ){
            System.out.println("Warm");
        } else if(Math.abs(randomInt - userGuess) <= 11 && Math.abs(randomInt - userGuess) >= 6 ){
            System.out.println("Hot");
        } else if(Math.abs(randomInt - userGuess) <= 5) {
            System.out.println("Very Hot");
        }
    }
}
