public class PrimeNumber {
    public static void main(String[] args) {
        int wholeNumber = 11;
        boolean prime = isPrime(wholeNumber);
        if (prime) {
            System.out.println(wholeNumber + " is prime!");
        } else {
            System.out.println(wholeNumber + " is not prime!");
        }
    }

    public static boolean isPrime(int wholeNumber) {
        for (int i = 2; i < wholeNumber; i++) {
            if (wholeNumber % i == 0) {
                return false;
            }
        }
        return true;
    }
}
