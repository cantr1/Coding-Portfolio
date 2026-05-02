#include <stdio.h>

// Simple function declaration and variable casting
float snek_score(int num_files, int num_contributors, int num_commits,
                 float avg_bug_criticality) {
  int size_factor = num_files * num_commits;
  int complexity_factor = size_factor + num_contributors;
  return (float)complexity_factor * avg_bug_criticality;
}

// Flow control with if-else statements
char *get_temperature_status(int temp) {
  if (temp > 90) {
    return "too hot";
  } else if (temp < 70) {
    return "too cold";
  } else {
    return "just right";
  }
}

// Similar to Python, can evaluate vars as truthy or falsy
int can_access_registry(int is_premium, int reputation, int has_2fa) {
  if (is_premium) {
    return 1;
  } else if (reputation >= 100 && has_2fa) {
    return 1;
  }
  return 0;
}

// ternary operator for simple conditionals
// a > b ? a : b

int size() {
  // Use %zu is for printing `sizeof` result
  printf("sizeof(char)   = %zu\n", sizeof(char));
  printf("sizeof(bool)   = %zu\n", sizeof(bool));
  printf("sizeof(int)   = %zu\n", sizeof(int));
  printf("sizeof(float)   = %zu\n", sizeof(float));
  printf("sizeof(double)   = %zu\n", sizeof(double));
  printf("sizeof(size_t)   = %zu\n", sizeof(size_t));
}

// Simple for loop
void print_numbers(int start, int end) {
  for (int i = start; i <= end; i++){
    printf("%d\n", i);
  }
}

// Simple while loop
void print_numbers_reverse(int start, int end) {
  int x = start;
  while (x >= end) {
    printf("%d\n", x);
    x--;
  }
}

// Simple do-while loop
void print_numbers_reverse_do_while(int start, int end) {
    do {
    printf("%d\n", start);
    start--;
  }  while (start >= end);
}

// Structs <3
struct Human {
    int age;
    char *name;
    int is_alive;
};

struct Coordinate {
  int x;
  int y;
  int z;
};

struct Coordinate new_coord(int x, int y, int z) {
  return (struct Coordinate){.x = x, .y = y, .z = z};
}

// Typdefs for cleaner code
typedef struct Pastry {
    char *name;
    float weight;
} pastry_t;

pastry_t muffin = {"Muffin", 0.3};

// Simple pointer
// int *x_ptr = &x;

// Defeferece
// int x_value = *x_ptr;

// Enums
typedef enum DaysOfWeek {
  MONDAY,
  TACO_TUESDAY,
  WEDNESDAY,
  THURSDAY,
  FRIDAY,
  SATURDAY,
  FUNDAY,
} days_of_week_t;

int main() {
    printf("Hello, C!\n");
    // Pointer to struct
    struct Coordinate point = {10, 20, 30};
    struct Coordinate *ptrToPoint = &point;
    printf("X: %d\n", ptrToPoint->x); // X: 10
    // *(ptrToPoint).x also works but is less common
    return 0;
}