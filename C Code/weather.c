#include <stdio.h>
#include <stdint.h>
#define SIZE 7

// Function to collect data and fill array
void collectData(int arr[])
{
    //Collect data from user
    printf("Enter the temperature for the last seven days: \n");
    for(int i = 0; i < SIZE; i++)
    {
        printf("Day %d = ", i + 1);
        scanf("%d", &arr[i]);
    }
}

int maxTemp(int arr[])
{
    int max_temp = arr[0];
    for(int i=0; i<SIZE; i++)
    {
        if (max_temp < arr[i])
            max_temp = arr[i];
    }
    return max_temp;
}

int lowTemp(int arr[])
{
    int low_temp = arr[0];
    for(int i=0; i<SIZE; i++)
    {
        if (low_temp > arr[i])
            low_temp = arr[i];
    }
    return low_temp;
}

int main()
{
    int temperatures[SIZE] = {0};
    int max, low;

    collectData(temperatures);
    max = maxTemp(temperatures);
    low = lowTemp(temperatures);

    // Print out Temps
    printf("Temps for the Week: \n");
    for(int i = 0; i < SIZE; i++)
    {
        printf("%d ", temperatures[i]);
        if (i < SIZE - 1)
            printf("-- ");
    }
    printf("\nMax Temp = %d", max);
    printf("\nLow Temp = %d", low);

    return 0;
}
