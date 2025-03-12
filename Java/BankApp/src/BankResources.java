import java.util.Scanner;

public class BankResources {
    //Adding fields
    private int balance = 1000;
    private int accountNumber = 12345;
    private String customerName = "Kelly";
    private String phoneNumber = "615-753-1126";
    private String password = "pass";

    // Getters
    public int getBalance(){
        return balance;
    }
    public int getAccountNumber() {
        return accountNumber;
    }
    public String getCustomerName() {
        return customerName;
    }
    public String getPhoneNumber() {
        return phoneNumber;
    }
    public String accountPassword(){
        return password;
    }

    // Setters
    public void setBalance(int balance){
        this.balance = balance;
    }
    public void setAccountNumber(int accountNumber) {
        this.accountNumber = accountNumber;
    }
    public void setCustomerName(String customerName) {
        this.customerName = customerName;
    }
    public void setPhoneNumber(String phoneNumber) {
        this.phoneNumber = phoneNumber;
    }

    public static int depositFunds(int balance){
        // Create scanner instance
        Scanner scanner1 = new Scanner(System.in);

        // Prompt and receive input
        System.out.println("Enter amount to deposit: ");
        int deposit = scanner1.nextInt();
        balance += deposit;

        return balance;
    }

    public static int withdrawalFunds(int balance){
        // Create scanner instance
        Scanner scanner2 = new Scanner(System.in);

        // Prompt and receive input
        System.out.println("Enter amount to withdrawal: ");
        int withdrawal = scanner2.nextInt();
        balance -= withdrawal;

        return balance;
    }
}
