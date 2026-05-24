#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/wait.h>
#include <unistd.h>
#include <stdbool.h>

char *get_command(char *buffer, size_t size) {
    return fgets(buffer, size, stdin);
}

bool is_builtin(char *cmd) {
    const char *built_ins[] = {"cd", "exit", "pwd"};
    const int num_builtins = sizeof(built_ins) / sizeof(built_ins[0]);
    for (int i = 0; i < num_builtins; i++) {
        if (strcmp(built_ins[i], cmd) == 0) {
            return true;
        }
    }
    return false;
}

void run_builtin(char *argv[]) {
    if (strcmp(argv[0], "cd") == 0) {
        char *dir = argv[1];
        if (dir != NULL) {
            if (chdir(dir) != 0) {
            perror("cd");
            }
        } else {
            // Hardcode home dir for now
            if (chdir("/home/kelz") != 0) {
            perror("cd");
            }
        }
    } else if (strcmp(argv[0], "exit") == 0) {
        exit(0);
    } else if (strcmp(argv[0], "pwd") == 0) {
        int size = 100;
        char current_dir[size];
        printf("%s\n", getcwd(current_dir, size));
    }
}

int run_external_cmd(char *argv[]) {
    // Track return code
    int rc = 0;
    // Create fork
    pid_t pid = fork();

    if (pid == 0) {
        if (execvp(argv[0], argv) == -1) {
            perror("execvp");
            exit(EXIT_FAILURE);
        }
    } else if (pid > 0) {
        int status;
        waitpid(pid, &status, 0);
    } else {
        printf("Fork failed!");
        rc = -1;
    }
    return rc;
}

int main() {
    while (1) {
        printf("c_shell:~$ ");

        char cmd[1024];
        if (get_command(cmd, sizeof(cmd)) == NULL) {
            break;
        } 
        
        cmd[strcspn(cmd, "\n")] = '\0';   // remove newline

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

        if (argc == 0) {
            continue;
        }

        // Handle builtin versus external commands
        if (is_builtin(argv[0])) {
            run_builtin(argv);
        } else {
            if (run_external_cmd(argv) != 0) {
                break;
            }
        }
    }
    return 0;
}
