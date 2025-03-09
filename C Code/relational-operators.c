#include <stdio.h>
#include <stdint.h>

int main()
{
    // Control Flow
    float x, y, z;
    scanf("%f", &x);
    scanf("%f", &y);
    scanf("%f", &z);
    if (x > y && x >z) {
        printf("X is the largest\n");
        if (y>z)
            printf("Z is the smallest\n");
        else
            printf("Y is the smallest\n");
    }
    else if (y > z && y >x) {
        printf("Y is the largest\n");
        if (x>z)
            printf("Z is the smallest\n");
        else
            printf("X is the smallest\n");
        printf("%.2f is greater!\n", y);
    }
    else {
        printf("Z is the largest\n");
        if (x>y)
            printf("y is the smallest\n");
        else
            printf("X is the smallest\n");
    }
    // && || !(int > 100)
    return 0;
}
