#include <stdio.h>
#include <stdlib.h>
#include <string.h>

char *get_command(char *buffer, size_t size) {
    return fgets(buffer, size, stdin);
}

int main() {
    while (1) {
        printf("c_shell:~$ ");

        char cmd[1024];
        if (get_command(cmd, sizeof(cmd)) == NULL) {
            break;
        } 
        
        cmd[strcspn(cmd, "\n")] = '\0';   // remove newline

        if (strcmp(cmd, "exit") == 0) {
            exit(0);
        }

        // For debugging, will remove later
        printf("Entered command: %s\n", cmd);
    }
    return 0;
}
