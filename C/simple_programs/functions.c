#include <stdio.h>
#include <stdint.h>

// Greeting function
void greeting()
{
    printf("Hello!\n");
    printf("This is the average grade calculator!\n");
}


// User input --- limited to whole numbers for now
int inputGrades(int *total, int *count)
{
    int grade = -1; // Initialize with non-zero value
    *total = 0;
    *count = 0;

    printf("Please input grades (0 to exit): \n");

    while(1) //Infinite loop, break condition if input is 0
    {
        printf("Grade%d = ", *count + 1);
        scanf("%d", &grade);

        if(grade == 0) // Exit condition
            break;

        *total += grade;
        (*count)++; // Increment non-zero grades
    }
}

// Calculate the grade average
double calcAvg(int total, int count)
{
    // Avoid division by zero
    if(count == 0)
        return 0;

    // Using double for more accurate average
    return (double)total / count;
}

// Determine letter grade
void printGrade(double avg_grade)
{
    printf("Average Grade: %.2f\n", avg_grade);

    if(avg_grade >= 90)
        printf("Letter Grade: A\n");
    else if (avg_grade >= 80 && avg_grade <= 89)
        printf("Letter Grade: B\n");
    else if (avg_grade >= 70 && avg_grade <= 79)
        printf("Letter Grade: C\n");
    else if (avg_grade >= 60 && avg_grade <= 69)
        printf("Letter Grade: D\n");
    else if (avg_grade <= 59)
        printf("Letter Grade: F\n");
}

int main()
{
    int total, count;
    double average;

    greeting();
    inputGrades(&total, &count);
    average = calcAvg(total, count);
    printGrade(average);

    return 0;
}
