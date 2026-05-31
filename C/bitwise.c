#include <stdio.h>
#define READ    (1 << 0)
#define WRITE   (1 << 1)
#define EXECUTE (1 << 2)
#define ADMIN   (1 << 3)

int main() {
    // AND
    printf("With the AND (&) operator, a bit is only 1 if both bits are 1.\n");
    printf("5 in binary = 101\n");
    printf("3 in binary = 011\n");
    printf("Thus what remains = 001\n");
    printf("(5 & 3) = %d\n", 5 & 3);
    printf("(5 & 5) = %d *All bits match\n\n", 5 & 5);

    // OR
    printf("With the OR (|) operator, a bit becomes 1 if either bit is 1\n");
    printf("8 in binary = 1000\n");
    printf("5 in binary = 0101\n");
    printf("Thus what remains: 1101\n");
    printf("(8 | 5) = %d\n", 8 | 5);
    printf("(8 | 8) = %d\n\n", 8 | 8);

    // XOR
    printf("With XOR (^) operations, a bit only becomes 1 when they differ\n");
    printf("13 in binary = 1101\n");
    printf("10 in binary = 1010\n");
    printf("What differs = 0111\n");
    printf("(13 ^ 10) = %d\n", 13 ^ 10);
    printf("(13 ^ 13) = %d\n\n", 13 ^ 13);

    // Real example
    printf("Lets say we have an integer to track permissions\nBit 0 = Read\nBit 1 = Write\nBit 2 = Execute\nBit 3 = Admin\nKeep in mind, 1 is ON and binary goes right to left\n\n");
    int permissions = 0b1011;
    printf("Our permissions: 1011 (Admin, RW)\n");
    printf("We can write a very simple check with these operations to check our bit positions\n");
    if (permissions & 0b1000) {
        printf("This code works because we have admin perms\n");
    }
    if ((permissions & 0b0011) == 0b0011) {
        printf("This code works because we have RW\n");
    }
    if ((permissions & 0b0111) == 0b0111) {
        printf("This code will never hit because we do not have RWX\n");
    }

    // Set with OR
    printf("If we want to set bits, we can actually use the OR operator");
    printf("Say we want to add execute to our perms, we would do the followign:\npermissions |= 0b0100\n");
    printf("This will insert the bit into that position\n");

    permissions |= 0b0100;
    if (permissions & 0b0100) {
        printf("This code executing confirms that we have X perms\n\n");
    }

    // Toggle with XOR
    printf("Similarly we can toggle the setting with XOR (^)\n");
    printf("We run permissions ^= 0b0100\n");
    printf("this will toggle the bit\n");

    permissions ^= 0b0100;
    if (!(permissions & 0b0100)) {
        printf("This code executing confirms that X perms have toggled OFF\n");
    }

    permissions ^= 0b0100;
    if (permissions & 0b0100) {
        printf("This code executing confirms that X perms have toggled ON\n");
    }

    // Shifts
    printf("Now we approach the concept of shifts\n");
    printf("In essence, it is multiplication and division\n");
    printf("\nSay we have 5 (0b0101) and we left shift by 1\n0b0101 << 1\n");
    printf("This means that each bit moves one to the left, making 0b1010 (10)\nA left shift of one is the same as *2\n");
    printf("The formula behind shifts becomes (x << n == x * 2^n)\n\n");

    printf("A common way to check bits aside from the AND operator follows this pattern:\n");
    printf("if ((permissions >> 3) & 1)\n");
    printf("This moves our third bit over to the first bit position, thus the comparison becomes:\n0b000X = 0b0001\n");

    if ((permissions >> 3) & 1) {
        printf("This code shows that we have X perms\n");
    }

    return 0;
}