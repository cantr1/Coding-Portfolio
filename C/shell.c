#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/wait.h>
#include <unistd.h>

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
            break;
        }

        // Tokenize command
        char delimiter[] = " ";
        char *token = strtok(cmd, delimiter);

        // Pass token to argv
        char *argv[64];
        int argc = 0;

        while (token && argc < 63) {
            argv[argc] = token;
            token = strtok(NULL, delimiter);
            argc++;
        }
        argv[argc] = NULL;

        // Create a fork
        pid_t pid = fork();
        
        if (pid < 0) {
            printf("Fork failed!");
            break;
        } else if (pid == 0) {
            int result;
            result = execvp(argv[0], argv);
            if (result != 0) {
                printf("Unrecognized command!\n");
            }
            break;
        } else {
            int status;
            waitpid(pid, &status, 0);
        }
    }
    return 0;
}
