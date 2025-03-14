public class Wall {
    // Instance variables
    private double width;
    private double height;

    // No-args constructor
    public Wall() {
        this(0, 0);  // Call the parameterized constructor with 0,0
    }

    // Parameterized constructor
    public Wall(double width, double height) {
        // Set width, but use 0 if negative
        this.width = (width < 0) ? 0 : width;

        // Set height, but use 0 if negative
        this.height = (height < 0) ? 0 : height;
    }

    // Getters and setters
    public double getWidth() {
        return width;
    }

    public void setWidth(double width) {
        this.width = (width < 0) ? 0 : width;
    }

    public double getHeight() {
        return height;
    }

    public void setHeight(double height) {
        this.height = (height < 0) ? 0 : height;
    }

    // Calculate area method
    public double getArea() {
        return width * height;
    }
}
