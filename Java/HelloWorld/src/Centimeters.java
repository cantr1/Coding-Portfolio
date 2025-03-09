public class Centimeters {
    public static void main(String[] args){
        convertToCentimeters(56);
        convertToCentimeters(6, 0);
    }
    public static double convertToCentimeters(int inches){
        return inches * 2.54;
    }

    public static double convertToCentimeters(int feet, int inches){
        int total = (feet * 12) + inches;
        double heightInCm = convertToCentimeters(total);
        System.out.println(heightInCm);
        return heightInCm;
    }

}
