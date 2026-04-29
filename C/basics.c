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

int main() {
    printf("Hello, C!\n");
    return 0;
}