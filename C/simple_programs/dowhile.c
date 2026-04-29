#include <stdio.h>
#include <stdint.h>

int main()
{
    int price, totalPrice = 0;

    do {
        printf("Please enter a price: ");
        scanf("%d", &price);
        totalPrice = totalPrice + price;
    } while(price != 0);

    printf("Total Price = %d", totalPrice);
    return 0;
}
