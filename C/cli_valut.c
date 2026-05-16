#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <stdbool.h>

typedef struct {
    char *entry_name;
    char *entry_username;
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

char *validate_entries(const char *prompt) {
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

void free_program_memory(t_entry *entry) {
    free(entry->entry_name);
    free(entry->entry_username);
    free(entry->entry_pw);
    free(entry);
}

int create_vault_entry(char *hostname) {
    // Track return code
    int rc = 1;

    // Setup entry struct to NULL
    t_entry *entry_ptr = NULL;

    // Allocate to the heap
    entry_ptr = calloc(1, sizeof(t_entry));
    if (entry_ptr == NULL) {
        printf("Memory allocation failed\n");
        goto cleanup;
    }

    // Allocate struct fields to NULL
    entry_ptr->entry_name = NULL;
    entry_ptr->entry_username = NULL;
    entry_ptr->entry_pw = NULL;

    // Print toplevel declaration of application on hostname
    printf("Adding entry to vault @%s\n", hostname);

    entry_ptr->entry_name = validate_entries("Entry Name: ");
    entry_ptr->entry_username = validate_entries("USERNAME: ");
    entry_ptr->entry_pw = validate_entries("PASSWORD: ");

    if (entry_ptr->entry_name == NULL || entry_ptr->entry_pw == NULL) {
        printf("Generally process faileure");
        goto cleanup;
    }

    printf("Entry Name: %s\n", entry_ptr->entry_name);
    printf("Enry Username: %s\n", entry_ptr->entry_username);
    printf("Entry Value: %s\n", entry_ptr->entry_pw);

    // If code reaches this point, we have been successful
    rc = 0;

    // Cleanup label
    cleanup:
        free_program_memory(entry_ptr);
        return rc;
}

int main(void) {
    char *hostname = NULL;
    hostname = get_hostname();
    if (hostname == NULL) {
        printf("Failed to get hostname\n");
        // fallback to allocated memory for default name
        char *hostname_string = "default_hostname";
        int len_of_string = strlen(hostname_string) +1;
        char *h_ptr = malloc(len_of_string);
        if (h_ptr == NULL) {
            printf("Failure to allocate memory for hostname");
            free(hostname);
        }
        strcpy(h_ptr, hostname_string);
        hostname = h_ptr;
    }

    bool continue_program = true;
    while (continue_program) {
        char user_input[8];
        printf("Choose an option\n(add entry = A) (exit = X)\n");
        if (fgets(user_input, sizeof(user_input), stdin) == NULL) {
            break;
        }
        switch (user_input[0]) {
            case 'A':
                create_vault_entry(hostname);
                break;
            case 'X':
                exit(0);
            default:
                printf("Unrecognized input\nPlease try again\n");
            }
        }
    free(hostname);
    return 0;
}
