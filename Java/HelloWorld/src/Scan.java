import java.util.Scanner;

public class Scan {

    public static void main(String[] args) {
        try {
            System.out.println(getInputFromConsole());
        } catch (NullPointerException e) {
            System.out.println(getInputFromScanner());
        }
    }

    public static String getInputFromScanner(){
        Scanner scanner = new Scanner(System.in);

        System.out.println("What is your name: ");
        String name = scanner.nextLine();
        return "Hello " + name;
    }

    public static String getInputFromConsole(){
        String name = System.console().readLine("What is your name: ");
        return "Hello " + name;
    }
}
