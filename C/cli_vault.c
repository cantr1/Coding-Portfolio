#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <stdbool.h>
#include <unistd.h>
#include <ctype.h>

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

void create_vault_entry(char *hostname) {
    // Base file path
    const char *base_file_path = "./";

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

    // Determine file name
    char buffer[100];
    snprintf(buffer, sizeof(buffer), "%s%s_vault.txt", base_file_path, entry_ptr->entry_name);
    char *new_file_path = buffer;
    FILE *fptr;
    fptr = fopen(new_file_path, "w");
    fprintf(fptr, "Username: %s\nPassword: %s", entry_ptr->entry_username, entry_ptr->entry_pw);
    fclose(fptr);

    printf("Wrote the follwoing to %s\n", new_file_path);
    printf("Entry Username: %s\n", entry_ptr->entry_username);
    printf("Entry Value: %s\n", entry_ptr->entry_pw);

    // Cleanup label
    cleanup:
        free_program_memory(entry_ptr);
}

void read_vault_entry (void) {
    bool valid = false;
    char entry_name[256];
    printf("Enter entry name: ");
    while (!valid) {
        if (fgets(entry_name, 256, stdin) != NULL) {
            valid = true;
        } else {
            printf("Invalid entry");
        }
    }

    entry_name[strcspn(entry_name, "\n")] = '\0';

    char file_path[300];
    snprintf(file_path, sizeof(file_path), "./%s_vault.txt", entry_name);

    if (access(file_path, F_OK) != 0) {
        printf("File not found --- Unable to update\n");
        return;
    }

    FILE *fptr;
    fptr = fopen(file_path, "r");

    // Store content
    char file_content[100];

    // Read and print
    while(fgets(file_content, 100, fptr)) {
        printf("%s", file_content);
    }

    fclose(fptr);

    printf("\nFile read complete\n");
}

void update_vault_entry(void) {
    bool valid = false;
    char entry_name[256];
    printf("Enter entry name: ");
    while (!valid) {
        if (fgets(entry_name, 256, stdin) != NULL) {
            valid = true;
        } else {
            printf("Invalid entry");
        }
    }

    entry_name[strcspn(entry_name, "\n")] = '\0';

    char file_path[300];
    snprintf(file_path, sizeof(file_path), "./%s_vault.txt", entry_name);

    if (access(file_path, F_OK) != 0) {
        printf("File not found --- Unable to update\n");
        return;
    }

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

    // We don't need to update the entry name in this case
    entry_ptr->entry_username = validate_entries("USERNAME: ");
    entry_ptr->entry_pw = validate_entries("PASSWORD: ");

    // Write new info to file path
    FILE *fptr;
    fptr = fopen(file_path, "w");
    fprintf(fptr, "Username: %s\nPassword: %s", entry_ptr->entry_username, entry_ptr->entry_pw);
    fclose(fptr);

    cleanup:
        free_program_memory(entry_ptr);
}

void delete_vault_entry(void) {
    bool valid = false;
    char entry_name[256];
    printf("Enter entry name: ");
    while (!valid) {
        if (fgets(entry_name, 256, stdin) != NULL) {
            valid = true;
        } else {
            printf("Invalid entry");
        }
    }

    entry_name[strcspn(entry_name, "\n")] = '\0';

    char file_path[300];
    snprintf(file_path, sizeof(file_path), "./%s_vault.txt", entry_name);
    
    // Check file exists
    if (access(file_path, F_OK) == 0) {
        printf("Successfully located vault file\nAttempting deletion\n");
        if (remove(file_path) == 0) {
            printf("File deleted successfully\n");
        } else {
            perror("Error deleting file");
            return;
        }
    } else {
        printf("File not found");
        return;
    }
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
        printf("\nChoose an option\n(Create entry = C) (Read entry = R) (Update entry = U) (Delete entry = D) (exit = X)\n");
        if (fgets(user_input, sizeof(user_input), stdin) == NULL) {
            break;
        }
        switch (toupper(user_input[0])) {
            case 'C':
                create_vault_entry(hostname);
                break;
            case 'R':
                read_vault_entry();
                break;
            case 'U':
                update_vault_entry();
                break;
            case 'D':
                delete_vault_entry();
                break;
            case 'X':
                free(hostname);
                exit(0);
            default:
                printf("Unrecognized input\nPlease try again\n");
            }
        }
}
