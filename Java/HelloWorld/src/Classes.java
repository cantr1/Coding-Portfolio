public class Classes {
    // In this case, image the class were 'Car'

    // Adding fields
    private String make = "Ford";
    private String model = "Mustang";
    private String color = "Red";
    private int doors = 2;
    private boolean convertable = false;

    // Not static - Getters
    public String getMake(){
        return make;
    }

    public String getModel() {
        return model;
    }

    public String getColor() {
        return color;
    }

    public int getDoors() {
        return doors;
    }

    public boolean isConvertable() {
        return convertable;
    }

    // Setters & This
    public void setMake(String make){
        this.make = make;
    }

    public void setModel(String model) {
        this.model = model;
    }

    public void setColor(String color) {
        this.color = color;
    }

    public void setDoors(int doors) {
        this.doors = doors;
    }

    public void setConvertable(boolean convertable) {
        this.convertable = convertable;
    }

    public void describeCar() {
        System.out.println( color + " " +
                make + " " +
                model + " " +
                doors + " Door " +
                (convertable ? "Convertable" : ""));
    }
}
