#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <stdbool.h>
#include <unistd.h>
#include <ctype.h>

#define MAX_PATH_LEN 300

typedef struct {
    char *entry_name;
    char *entry_username;
    char *entry_pw;
    char action;
} t_entry;

bool validate_entry(const t_entry *entry) {
    if (entry->action != 'C' && entry->action != 'R' && entry->action != 'U' && entry->action != 'D' ) {
        printf("Entry must have a valid action type");
        return false;
    }

    if (entry->action == 'C' || entry->action == 'U') {
        if (entry->entry_name == NULL || entry->entry_username == NULL || entry->entry_pw == NULL) {
            printf("Create / Update Action requires username / password and valid entry name");
            return false;
        }
    } else if (entry->action == 'D' || entry->action == 'R') {
        if (entry->entry_name == NULL) {
            printf("Delete / Read Action requires valid entry name");
            return false;
        }
    } 

    return true;
}

void generate_file_path(const t_entry *entry, char *buffer, size_t size) {
    const char *base_file_path = "./";
    snprintf(buffer, size, "%s%s_vault.txt", base_file_path, entry->entry_name);
}

int create_vault_entry(t_entry *entry) {
    int rc = 1;
    char file_path[MAX_PATH_LEN];

    generate_file_path(entry, file_path, sizeof(file_path));

    FILE *fptr = fopen(file_path, "w");
    if (fptr != NULL) {
        fprintf(fptr, "Username: %s\nPassword: %s", entry->entry_username, entry->entry_pw);
        fclose(fptr);

        printf("Wrote the follwoing to %s\n", file_path);
        printf("Entry Username: %s\n", entry->entry_username);
        printf("Entry Value: %s\n", entry->entry_pw);
        rc = 0;
    } else {
        printf("Unable to create new entry");
    }

    return rc;
}

int read_vault_entry (const t_entry *entry) {
    int rc = 1;
    char file_path[MAX_PATH_LEN];

    generate_file_path(entry, file_path, sizeof(file_path));

    if (access(file_path, F_OK) != 0) {
        printf("File not found --- Unable to update\n");
        return rc;
    }

    FILE *fptr = fopen(file_path, "r");

    if (fptr != NULL) {
        // Store content
        char file_content[100];

        // Read and print
        while(fgets(file_content, 100, fptr)) {
            printf("%s", file_content);
        }

        fclose(fptr);

        printf("\nFile read complete\n");
        rc = 0;
    } else {
        printf("Unable to read from file.");
    }

    return rc;
}

int update_vault_entry(const t_entry *entry) {
    int rc = 1;
    char file_path[MAX_PATH_LEN];

    generate_file_path(entry, file_path, sizeof(file_path));

    if (access(file_path, F_OK) != 0) {
        printf("File not found --- Unable to update\n");
        return rc;
    }

    // Write new info to file path
    FILE *fptr = fopen(file_path, "w");
    if (fptr != NULL) {
        fprintf(fptr, "Username: %s\nPassword: %s", entry->entry_username, entry->entry_pw);
        fclose(fptr);
        rc = 0;
    } else {
        printf("Unable to write data to file.");
    }

    return rc;
}

int delete_vault_entry(const t_entry *entry) {
    int rc = 1;
    char file_path[MAX_PATH_LEN];

    generate_file_path(entry, file_path, sizeof(file_path));
    
    // Check file exists
    if (access(file_path, F_OK) == 0) {
        printf("Successfully located vault file\nAttempting deletion\n");
        if (remove(file_path) == 0) {
            printf("File deleted successfully\n");
            rc = 0;
        } else {
            perror("Error deleting file");
        }
    } else {
        printf("File not found"); 
    }

    return rc;
}

int main(void) {
    int rc = 1;

    t_entry entry = {0}; 

    entry.entry_name = "Udemy";
    entry.entry_username = "kelz";
    entry.entry_pw = "pw";
    entry.action = 'C';

    if (validate_entry(&entry)){
        switch (entry.action) {
        case 'C':
            rc = create_vault_entry(&entry);
            break;
        case 'R':
            rc = read_vault_entry(&entry);
            break;
        case 'U':
            rc = update_vault_entry(&entry);
            break;
        case 'D':
            rc = delete_vault_entry(&entry);
            break;
        default:
            printf("Unrecognized Entry Action\nPlease try again\n");
        }

    } else {
        printf("Invalid entry\n");
    }

    return rc;
}
 