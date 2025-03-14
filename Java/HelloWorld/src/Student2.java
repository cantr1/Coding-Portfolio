public class Student2 {
    public static void main(String[] args) {

        for (int i = 1; i <= 5; i++) {
            LPAStudent s = new LPAStudent("S92300" + i,
                    switch (i) {
                        case 1 -> "Mary";
                        case 2 -> "Tom";
                        case 3 -> "Henry";
                        case 4 -> "Liz";
                        case 5 -> "Jerry";
                        default -> "Anon";
                    },
                    "07/23/2001",
                    "Java Class");
            System.out.println(s);
        }

        Student pojoStudent = new Student();
        LPAStudent recordStudent = new LPAStudent("Kelly",
                "2003",
                "07/25/1997",
                "Java");

        System.out.println(pojoStudent);
        System.out.println(recordStudent);

        System.out.println(pojoStudent.getName());
        System.out.println(recordStudent.name());
    }
}
