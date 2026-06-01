#include <stdlib.h>
#include <stdio.h>

void update_x(int *num) {
    printf("Address of num = %p\n", (void *)num); /* This is a copy of the pointer, so it has it's own value*/
    printf("Address of num's value: %p\n", (void *)&num);
    *num = *num + 1;
}

void swap_function(int *x, int *y) {
    int temp = *x;
    *x = *y;
    *y = temp;
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

    // Create an array
    /* An array is a collection of pointers to values */
    int numbers[] = {11, 12, 13};

    printf("First number = %d\n", *numbers);
    printf("Second number = %d\n", *(numbers + 1));

    // You can also reference by index
    printf("Third number = %d\n", numbers[2]);

    // See how pointer address changes in an array
    printf("%p\n", (void *)numbers);
    printf("%p\n", (void *)&numbers[0]);
    printf("%p\n", (void *)(numbers + 1));
    printf("%p\n", (void *)&numbers[1]);
    printf("Sizeof Integer = %lu\n", sizeof(int));

    // Find the size of the array
    printf("Number of Elements in `number` = %lu\n", sizeof(numbers) / sizeof(numbers[0]));

    int y = 99;
    int *y_ptr = &y;

    swap_function(x_ptr, y_ptr);
    printf("Swapped value of x = %d\n", x);
    printf("Swapped value of y = %d\n", y);
}