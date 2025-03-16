public class Main {

    public static void main(String[] args) {
    HourlyEmployee Kelly = new HourlyEmployee("Kelly", "1997", "N/A", 30000,
            "2022", 23);
    System.out.println(Kelly);

    }
}

class Worker {
    //Fields
    protected String name;
    protected String birthDate;
    protected String endDate;

    //Constructors
    public Worker(){
        this("Anon", "Default", "N/A");
    }
    public Worker(String name, String birthDate, String endDate) {
        this.name = name;
        this.birthDate = birthDate;
        this.endDate = endDate;
    }

    // Methods
    public int getAge(){
        return 2025 - Integer.parseInt(birthDate);
    }

    public double collectPay() {

        return 0.0;
    }

    public void terminate(String endDate){
        System.out.println(name + " has been terminated, effective " + endDate);
    }

}

class Employee extends Worker {
    //Fields
    protected long employeeid;
    protected String hireDate;

    //Constructors
    public Employee(){
        this("Anon", "Default", "N/A", 30000, "2025");
    }

    public Employee(String name, String birthDate, String endDate, long employeeid, String hireDate){
        super(name, birthDate, endDate);
        this.employeeid = employeeid;
        this.hireDate = hireDate;
    }

}

class SalariedEmployee extends Employee {
    //Fields
    private double annualSalary;
    private boolean isRetired;

    //Constructors
    public SalariedEmployee(){
        this("Anon", "Default", "N/A", 30000, "2025", 50000, false );
    }

    public SalariedEmployee(String name, String birthDate, String endDate, long employeeid, String hireDate,
                            double annualSalary, boolean isRetired){
        super(name, birthDate, endDate, employeeid, hireDate);
        this.annualSalary = annualSalary;
        this.isRetired = isRetired;
    }

    //Methods
    public void retire(){
        isRetired = true;
    }
}

class HourlyEmployee extends Employee {
    //Fields
    private double hourlyPayRate;

    //Constructors
    public HourlyEmployee(){
        this("Anon", "Default", "N/A", 30000, "2025", 25);
    }

    public HourlyEmployee(String name, String birthDate, String endDate, long employeeid, String hireDate,
                            double hourlyPayRate){
        super(name, birthDate, endDate, employeeid, hireDate);
        this.hourlyPayRate = hourlyPayRate;
    }

    //Methods
    public double getDoublePayRate (){
        return hourlyPayRate * 2;
    }

    @Override
    public String toString() {
        return "HourlyEmployee{" +
                "hourlyPayRate=" + hourlyPayRate +
                ", employeeid=" + employeeid +
                ", hireDate='" + hireDate + '\'' +
                ", name='" + name + '\'' +
                ", birthDate='" + birthDate + '\'' +
                ", endDate='" + endDate + '\'' +
                '}';
    }
}
