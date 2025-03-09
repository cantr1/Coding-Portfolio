import java.util.Scanner;

public class Calculator {

    // Print many newlines to "clear" screen
    public static void clearScreen() {
        for (int i = 0; i < 20; i++) {
            System.out.println();
        }
    }

    // Calculator Feature

    public static void main(String[] args) {
        boolean continueCalculating = true;

        // Intro Statement
        System.out.println("Hello World!\nThis is a simple calculator in Java!");
        System.out.println("""
                |  _________________  |
                | | JO           0. | |
                | |_________________| |
                |  ___ ___ ___   ___  |
                | | 7 | 8 | 9 | | + | |
                | |___|___|___| |___| |
                | | 4 | 5 | 6 | | - | |
                | |___|___|___| |___| |
                | | 1 | 2 | 3 | | x | |
                | |___|___|___| |___| |
                | | . | 0 | = | | / | |
                | |___|___|___| |___| |
                |_____________________|""");

        while (continueCalculating) {
            // Create Scanner Object
            Scanner scanner = new Scanner(System.in);

            // Prompt User
            System.out.print("Enter Num 1: ");
            int num1 = scanner.nextInt();
            System.out.print("Enter Num 2: ");
            int num2 = scanner.nextInt();

            // Get operation
            System.out.println("Enter operation (+, -, *, /, ^): ");
            scanner.nextLine(); // Clear buffer
            String operation = scanner.nextLine();

            // Perform Calculation and display result
            switch (operation) {
                case "+" -> {
                    int result = num1 + num2;
                    System.out.println(num1 + operation + num2 + "=" + result);
                }
                case "-" -> {
                    int result = num1 - num2;
                    System.out.println(num1 + operation + num2 + "=" + result);
                }
                case "*" -> {
                    int result = num1 * num2;
                    System.out.println(num1 + operation + num2 + "=" + result);
                }
                case "/" -> {
                    if (num2 == 0) {
                        System.out.println("ERROR: DIVISION BY ZERO");
                    } else {
                        double result = (double) num1 / num2;
                        System.out.println(num1 + operation + num2 + "=" + result);
                    }
                }
                case "^" -> {
                    int x, result = num1;
                    for (x = 1; x < num2; x++) {
                        result = num1 * result;
                    }
                    System.out.println(num1 + operation + num2 + "=" + result);
                }
                default -> System.out.println("ERROR");
            }

            //Ask if user would like to continue
            System.out.println("Would you like to continue? (1=yes, 0=no): ");
            int answer = scanner.nextInt();
            scanner.nextLine(); // consumes new line buffer


            if (answer == 0) {
                continueCalculating = false;
                // Close scanner at the end
                scanner.close();
                System.out.println("Goodbye!");
            } else if (answer == 1) {
                System.out.println("Would you like to clear the screen? (yes/no): ");
                String clear = scanner.nextLine();
                if (clear.equals("yes")){
                    clearScreen();
                }
            } else {
                System.out.println("ERROR - Continuing Calculations...");
            }

        }
    }
}
