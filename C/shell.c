#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/wait.h>
#include <unistd.h>
#include <stdbool.h>
#include <limits.h>
#include <fcntl.h>

char *get_command(char *buffer, size_t size) {
    return fgets(buffer, size, stdin);
}

void remove_character(char *str, char c) {
    if (str == NULL) {
        return;
    }
    char *src, *dst;
    for (src = dst = str; *src != '\0'; src++) {
        *dst = *src;
        if (*dst != c) {
            dst++;
        }
    }
    *dst = '\0'; // New null terminator
}

void free_history(char ***history, int count) {
    for (int i = 0; i < count; i++) {
        free((*history)[i]);
    }
    free(*history);
}

int write_redirect(char *argv[], int redirect_index) {
    char *outfile = argv[redirect_index + 1];

    if (outfile == NULL) {
        fprintf(stderr, "Must specify a file path\n");
        return 1;
    }

    // Remove redirect token
    argv[redirect_index] = NULL;

    pid_t pid = fork();

    if (pid < 0) {
        perror("fork");
        return 1;
    }

    if (pid == 0) {
        int fd = open(outfile,
                      O_WRONLY | O_CREAT | O_TRUNC,
                      0644);

        if (fd < 0) {
            perror("open");
            exit(EXIT_FAILURE);
        }

        if (dup2(fd, STDOUT_FILENO) < 0) {
            perror("dup2");
            close(fd);
            exit(EXIT_FAILURE);
        }

        close(fd);
        execvp(argv[0], argv);

        perror("execvp");
        exit(EXIT_FAILURE);
    }
    waitpid(pid, NULL, 0);
    return 0;
}

int append_redirect(char *argv[], int redirect_index) {
    char *outfile = argv[redirect_index + 1];

    if (outfile == NULL) {
        fprintf(stderr, "Must specify a file path\n");
        return 1;
    }

    // Remove redirect token
    argv[redirect_index] = NULL;

    pid_t pid = fork();

    if (pid < 0) {
        perror("fork");
        return 1;
    }

    if (pid == 0) {
        int fd = open(outfile,
                      O_WRONLY | O_CREAT | O_APPEND,
                      0644);

        if (fd < 0) {
            perror("open");
            exit(EXIT_FAILURE);
        }

        if (dup2(fd, STDOUT_FILENO) < 0) {
            perror("dup2");
            close(fd);
            exit(EXIT_FAILURE);
        }

        close(fd);
        execvp(argv[0], argv);

        perror("execvp");
        exit(EXIT_FAILURE);
    }
    waitpid(pid, NULL, 0);
    return 0;
}

int pipe_redirect(char *argv[], int redirect_index) {
     int pipefd[2];

    if (pipe(pipefd) < 0) {
        perror("pipe");
        return 1;
    }

    argv[redirect_index] = NULL;

    char **right_argv = &argv[redirect_index + 1];

    pid_t pid1 = fork();

    if (pid1 == 0) {
        dup2(pipefd[1], STDOUT_FILENO);
        close(pipefd[0]);
        close(pipefd[1]);
        execvp(argv[0], argv);
        perror("execvp left");
        exit(EXIT_FAILURE);

    }

    pid_t pid2 = fork();

    if (pid2 == 0) {
        dup2(pipefd[0], STDIN_FILENO);
        close(pipefd[1]);
        close(pipefd[0]);
        execvp(right_argv[0], right_argv);
        perror("execvp right");
        exit(EXIT_FAILURE);
    }

    close(pipefd[0]);
    close(pipefd[1]);

    waitpid(pid1, NULL, 0);
    waitpid(pid2, NULL, 0);
    return 0;
}

int process_redirects(char *argv[], int argc) {
    for (int i = 0; i <  argc; i++) {
        if (strcmp(argv[i], ">") == 0) {
            write_redirect(argv, i);
            return true;
        } else if (strcmp(argv[i], ">>") == 0) {
            append_redirect(argv, i);
            return true;
        } else if (strcmp(argv[i], "|") == 0) {
            pipe_redirect(argv, i);
            return true;
        }
    }
    return false;
}

int add_command_to_history(char ***history, int count, char *cmd) {
    // Expand the array to hold one more element
    char **temp = realloc(*history, (count + 1) * sizeof(char *));
    if (temp == NULL) {
        perror("Allocation failed");
        return 1;
    }

    *history = temp;

    // Allocate the memory for the string
    (*history)[count] = malloc(strlen(cmd) + 1);
    if ((*history)[count] == NULL) {
        perror("Allocation failed");
        free_history(history, count);
        return 1;
    }

    // Copy the string
    strcpy((*history)[count], cmd);
    return 0;
}

int print_history(char **history, int count) {
    for (int i = 0; i < count; i++) {
        printf("%d - %s\n", i, history[i]);
    }
    return 0;
}

bool is_builtin(char *cmd) {
    const char *built_ins[] = {"cd", "exit", "pwd", "export", "history"};
    const int num_builtins = sizeof(built_ins) / sizeof(built_ins[0]);
    for (int i = 0; i < num_builtins; i++) {
        if (strcmp(built_ins[i], cmd) == 0) {
            return true;
        }
    }
    return false;
}

void expand_variables(char *argv[], int num_args) {
    for (int i = 0; i < num_args; i++) {
        if (argv[i] != NULL) {
            if (strchr(argv[i], '$')) {
                remove_character(argv[i], '$'); /* Remove '$'*/
                if (strchr(argv[i], '"')) {
                    remove_character(argv[i], '"'); /* Remove double quotes*/
                }

                if (getenv(argv[i]) != NULL) {
                    argv[i] = getenv(argv[i]);
                } else {
                    argv[i] = ""; /* Handle env vars that do not exist */
                }
            }
        }
    } 
}

int export_env_vars(char *argv[]) {
    if (argv[1] == NULL || argv[2] == NULL) {
        return -1;
    }

    char *env_name = argv[1];
    char *env_value = argv[2];

    if (setenv(env_name, env_value, 1) != 0) {
        return -1;
    } else {
        printf("env (%s) set to (%s)\n", argv[1], getenv(argv[1]));
        return 0;
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

bool run_builtin(char *argv[], char **history, int count) {
    if (strcmp(argv[0], "cd") == 0) {
        char *dir = argv[1];
        if (dir != NULL) {
            if (strcmp(dir, "~") == 0) {
                dir = getenv("HOME");
            }
            if (chdir(dir) != 0) {
            perror("cd");
            }
        } else {
            char *home_dir = getenv("HOME");
            if (home_dir != NULL) {
                if (chdir(home_dir) != 0) {
                    perror("cd");
                }
            } else {
                printf("Error finding home dir env variable");
            }
        }
    } else if (strcmp(argv[0], "exit") == 0) {
        return true;
    } else if (strcmp(argv[0], "pwd") == 0) {
        char current_dir[PATH_MAX];
        if (getcwd(current_dir, PATH_MAX) != NULL) {
            printf("%s\n", current_dir);
        } else {
            perror("pwd");
        }
    } else if (strcmp(argv[0], "export") == 0) {
        if (export_env_vars(argv) != 0) {
            perror("export");
        }
    } else if (strcmp(argv[0], "history") == 0) {
        if (print_history(history, count) != 0) {
            perror("history");
        }
    }
    return false;
}

int main() {
    // Create a variable to track history of commands and count
    char **history = NULL;
    int count = 0;

    while (1) {
        printf("c_shell:~$ ");

        char cmd[1024];
        if (get_command(cmd, sizeof(cmd)) == NULL) {
            break;
        } 
        
        cmd[strcspn(cmd, "\n")] = '\0';   // remove newline

        if (add_command_to_history(&history, count, cmd) != 0) {
            perror("Failed to add command to history");
            break;
        }
        count++;

        // Tokenize command + track raw command
        char raw_cmd[1024];
        strcpy(raw_cmd, cmd);
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

        // Expand any variables
        expand_variables(argv, argc);

        // Handle any redirects, if not process builtin / versus external command
        if (process_redirects(argv, argc)) {
            continue;
        }

        // Handle builtin versus external commands
        if (is_builtin(argv[0])) {
            if (run_builtin(argv, history, count) == true) {
                free_history(&history, count);
                return 0;
            }
        } else {
            if (run_external_cmd(argv) != 0) {
                break;
            }
        }
    }
    free_history(&history, count);
    return 0;
}
