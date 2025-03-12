import java.util.Scanner;

public class BankAccount {
    static BankResources bank = new BankResources();

    public static void main(String[] args) {
        // Create scanner instance
        Scanner scanner = new Scanner(System.in);

        int loginAttempts = 0;

        while(loginAttempts < 3){
            // Authenticate user
            if(authenticateUser()) {
                // Welcome user
                System.out.println("Welcome to Kelz Online Banking");
                System.out.println("What would you like to do today?");

                // Create a loop for user actions
                boolean bankingActive = true;
                while (bankingActive) {
                    System.out.println("""
                            1.) Deposit\
                            
                            2.) Withdrawal \
                            
                            3.) Display Current Balance\
                            
                            4.) Exit App\
                            
                            5.) Display Personal Info\
                            
                            Input:\s""");
                    int userInput = scanner.nextInt();

                    switch (userInput) {
                        case 1 -> depositFunds();
                        case 2 -> withdrawalFunds();
                        case 3 -> System.out.println("Current balance: $" + bank.getBalance());
                        case 4 -> bankingActive = false;
                        case 5 -> displayPersonalInfo();
                        default -> System.out.println("Invalid Input");
                    }

                    //System.out.println("\nWould you like to do anything else? "); TODO Get this statement to work
                }
            } else {
                if (loginAttempts < 2){
                    System.out.println("Login information incorrect. Please try again!");
                } else {
                    System.out.println("Max login attempts reached! Try again later!");
                }
                loginAttempts++;
            }
        }
    }



    public static boolean authenticateUser(){
        Scanner authScan = new Scanner(System.in);
        System.out.print("Enter Username: ");
        String username = authScan.nextLine();

        System.out.print("Enter Password: ");
        String password = authScan.nextLine();

        if(username.equals("Kelly") && password.equals("pass")){
            return true;
        } else {
            return false;
        }
    }

    public static void depositFunds(){
        // Create scanner instance
        Scanner scanner1 = new Scanner(System.in);

        // Prompt and receive input
        System.out.println("Enter amount to deposit: ");
        int deposit = scanner1.nextInt();

        // Get balance
        int balance = bank.getBalance();

        // Add deposit
        balance += deposit;

        // Print out total balance after transaction
        System.out.println("New Balance = $" + balance);

        // Set balance
        bank.setBalance(balance);

    }

    public static void withdrawalFunds(){
        // Create scanner instance
        Scanner scanner2 = new Scanner(System.in);

        // Prompt and receive input
        System.out.println("Enter amount to withdrawal: ");
        int withdrawal = scanner2.nextInt();

        // Get balance
        int balance = bank.getBalance();

        if(balance - withdrawal < 0){
            System.out.println("Invalid: Negative Balance");
        } else {
            balance -= withdrawal;
            // Print out total balance after transaction
            System.out.println("New Balance = $" + balance);
        }

        // Set balance
        bank.setBalance(balance);

    }

    public static void displayPersonalInfo(){
        System.out.println("Name: " + bank.getCustomerName());
        System.out.println("Account Number: " + bank.getAccountNumber());
        System.out.println("Phone: " + bank.getPhoneNumber());
    }
}


