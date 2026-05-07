#include <stdio.h>
#include <stdlib.h>

typedef struct {
    char* name;
    int user_id;
} t_user;

typedef struct {
    int process_id;
    t_user user;
    int result;
} t_calculation;

t_user *new_user(char* name) {
    t_user *objPtr = malloc(sizeof(t_user));
    objPtr->name = name;
    objPtr->user_id = 1;
    return objPtr;
}

int main() {
    char name[30];

    printf("What is your name: ");
    fgets(name, sizeof(name), stdin);

    t_user *user_object = new_user(name);

    printf("Hello, %s", user_object->name);

    return 0;
}