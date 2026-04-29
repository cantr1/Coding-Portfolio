#include <stdio.h>
#include <stdint.h>

int main()
{
    // Would you rather a million dollars today or...
    // 1 cent that doubles every day for the next thirty days?
    int i, choice;
    float x = 0.01;
    printf("Would you rather a million dollars today or...\n");
    printf("1 cent that doubles every day for the next thirty days?\n");
    printf("1 for the million, 2 for the penny: \n");
    scanf("%d", &choice);

    for(i = 1; i < 30; i++)
    {
        x = x * 2;
    }


    printf("Final Value after 30 days = $%.0lf\n", x);
    if(choice == 1){
        printf("Well would you look at that!\n");
        printf("Lost profit = $%d", (int)x - 1000000);
    }
    else
        printf("Wise choice!\n");
    return 0;
}
