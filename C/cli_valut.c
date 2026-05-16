#include <stdio.h>
#include <stdlib.h>

typedef struct {
    char *entry_name;
    char *entry_value;
} t_entry;

int main() {
    t_entry *entry_ptr = malloc(sizeof(t_entry));
    if (entry_ptr == NULL) {
        printf("Memory allocation failed\n");
        return 1;
    }

    entry_ptr->entry_name = "example_name";
    entry_ptr->entry_value = "example_value";

    printf("Entry Name: %s\n", entry_ptr->entry_name);
    printf("Entry Value: %s\n", entry_ptr->entry_value);

    free(entry_ptr);
    return 0;
}