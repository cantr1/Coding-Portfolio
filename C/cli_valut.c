#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <stdbool.h>

typedef struct {
    char *entry_name;
    char *entry_pw;
} t_entry;

char *get_hostname(void) {
    FILE *fp;
    char buffer[256];

    fp = popen("hostname", "r");
    if (fp == NULL) {
        printf("Failed to run command\n");
        return NULL;
    }
    
    while (fgets(buffer, sizeof(buffer), fp) != NULL) {
        // Remove newline character from the end of the buffer
        buffer[strcspn(buffer, "\n")] = '\0';
        // Allocate memory for the hostname string and copy the buffer into it
        char *hostname = malloc((strlen(buffer) + 1) * sizeof(char));
        if (hostname == NULL) {
            printf("Memory allocation failed\n");
            pclose(fp);
            return NULL;
        }
        strcpy(hostname, buffer);
        pclose(fp);
        return hostname;
    }
    pclose(fp);
    return NULL;
}

char *validate_entries(char *prompt) {
    // Returns a pointer to a char array
    bool valid_entry = false;
    char entry[256];
    while (!valid_entry) {
        printf("%s", prompt);
        if (fgets(entry, 256, stdin) != NULL) {
            // Remove newline character from the end of the entry_name
            entry[strcspn(entry, "\n")] = '\0';
            if (strcmp(entry, "") != 0 && strcmp(entry, "\n") != 0) {
                valid_entry = true;
            }
        }
    }
    char *heap_ptr = malloc(strlen(entry) + 1);
    if (heap_ptr == NULL) {
        printf("Pointer allocation failed");
        return NULL;
    }
    strcpy(heap_ptr, entry);
    return heap_ptr;
}

void free_entry_memory(t_entry *entry) {
    free(entry->entry_name);
    free(entry->entry_pw);
}

int main(void) {
    t_entry *entry_ptr = malloc(sizeof(t_entry));
    if (entry_ptr == NULL) {
        printf("Memory allocation failed\n");
        return 1;
    }

    char *hostname = get_hostname();
    if (hostname == NULL) {
        printf("Failed to get hostname\n");
        // fallback to allocated memory for default name
        char *hostname_string = "default_hostname";
        int len_of_string = sizeof(strlen(hostname_string) +1);
        char *h_ptr = malloc(len_of_string);
        if (h_ptr == NULL) {
            printf("Failure to allocate memory for hostname");
            return 1;
        }
        hostname = h_ptr;
    }

    // Print toplevel declaration of application on hostname
    printf("Starting cli_vault@%s\n", hostname);

    entry_ptr->entry_name = validate_entries("Entry Name: ");
    entry_ptr->entry_pw = validate_entries("Entry PW: ");

    if (entry_ptr->entry_name == NULL || entry_ptr->entry_pw == NULL) {
        printf("Generally process faileure");
        return 1;
    }

    printf("\nEntry Name: %s\n", entry_ptr->entry_name);
    printf("Entry Value: %s\n", entry_ptr->entry_pw);

    free_entry_memory(entry_ptr);
    free(hostname);
    return 0;
}
