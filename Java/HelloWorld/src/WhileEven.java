public class WhileEven {
    public static void main(String[] args) {
        int i = 5;
        int totalEven = 0;
        int totalOdd = 0;
        while(i <= 20 && totalEven < 5){
            if (isEvenNumber(i)) {
                System.out.println(i + " is even.");
                totalEven++;
            } else {
                System.out.println(i + " is odd.");
                totalOdd++;
            }
            i++;
        }
        System.out.println("Total even numbers: " + totalEven);
        System.out.println("Total odd numbers: " + totalOdd);
    }

    public static boolean isEvenNumber(int number){
        if (number % 2 == 0){
            return true;
        } else {
            return false;
        }
    }
}
