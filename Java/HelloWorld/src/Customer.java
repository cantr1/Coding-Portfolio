public class Customer {
    public String name;
    public int creditLimit;
    public String email;

    public String getName() {
        return name;
    }

    public int getCreditLimit() {
        return creditLimit;
    }

    public String getEmail() {
        return email;
    }


    // No args constructor that calls another constructor
    public Customer(){
        this("No one", "No email");
    }

    // Just email and name which calls another constructor
    public Customer(String name, String email){
        this(name, 1000, email);
    }

    public Customer(String name, int creditLimit, String email){
        this.name = name;
        this.creditLimit = creditLimit;
        this.email = email;
    }
}
