#include <stdlib.h>
#include <stdio.h>

void update_x(int *num) {
    printf("Address of num = %p\n", (void *)num); /* This is a copy of the pointer, so it has it's own value*/
    printf("Address of num's value: %p\n", (void *)&num);
    *num = *num + 1;
}

int main() {
    // Create an integer
    int x = 47;

    // Create a pointer to that variable
    int *x_ptr = &x; /* x_ptr is now a value representative of where x lives in memory */

    // Dereference to work with value
    printf("The value of x = %d\n", *x_ptr);
    printf("The location in memory of the pointer: %p\n", x_ptr);
    printf("The location in memory of x itself: %p\n", &x);

    // C always passes by value.
    // Here we pass a copy of x's address, allowing the function
    // to modify the original object through the pointer.
    update_x(x_ptr);

    printf("The value of x is now incremented: %d\n", x);
}