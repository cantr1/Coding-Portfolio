public class Constructors {
    private String name;
    private String phone;

    public Constructors(){
        this("Kelly", "615-753-1126");
    }
    public Constructors(String userName, String phone){
        name = userName;
        this.phone = phone;
    }

    public String getName(){
        return name;
    }

    public String getPhone() {
        return phone;
    }
}
