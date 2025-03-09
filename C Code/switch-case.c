#include <stdio.h>
#include <stdint.h>

int main()
{
    char grade;
    printf("Enter grade (A-F): ");
    scanf("%c", &grade);
    switch(grade)
    {
        case 'A':
            printf("Grade between 90 to 100 \n");
            break;
        case 'B':
            printf("Grade between 80 to 89 \n");
            break;
        case 'C':
            printf("Grade between 70 to 79 \n");
            break;
        case 'D':
            printf("Grade between 60 to 69 \n");
            break;
        case 'F':
            printf("Grade between 0 to 59 \n");
            break;
        default:
            printf("Input Error...");
            break;
    }
    return 0;
}
