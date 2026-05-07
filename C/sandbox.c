#include <stdio.h>

typedef struct {
    char* name;
    int user_id;
} t_user;

typedef struct {
    int process_id;
    t_user user;
    int result;
} t_calculation;

t_user new_user(char* name) {
    t_user obj = {name, 1};
    return obj;
}

int main() {
    char name[30];

    printf("What is your name: ");
    getchar();
    fgets(name, sizeof(name), stdin);

    t_user user_object = new_user(name);

    

    return 0;
}