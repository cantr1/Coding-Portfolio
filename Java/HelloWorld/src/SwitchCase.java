public class SwitchCase {
    public static void main(String[] args) {
        System.out.println(printDayOfWeek(0));
        }

    public static String printDayOfWeek(int day) {
        if (day >= 0 && day <= 6) {
            return switch (day) {
                case 0 -> "Sunday";
                case 1 -> "Monday";
                case 2 -> "Tuesday";
                case 3 -> "Wednesday";
                case 4 -> "Thursday";
                case 5 -> "Friday";
                case 6 -> "Saturday";
                default -> "Invalid day";
            };
        } else {
            return "Invalid day";
        }
    }
}
