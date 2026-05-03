#include <stdio.h>
#include <stdlib.h>

typedef struct {
    int id;
    char *name;
} t_person;

t_person *create_person(int id, char *name) {
    t_person *obj = calloc(1, sizeof(t_person));
    obj->id = id;
    obj->name = name;
    return obj;
}

void free_person(t_person *obj) {
    free(obj);
}

int main() {
    t_person *person = create_person(1, "John Doe");
    printf("ID: %d, Name: %s\n", person->id, person->name);
    free_person(person);
    return 0;
}