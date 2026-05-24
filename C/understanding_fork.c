#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/wait.h>
#include <unistd.h>

char *get_command(char *buffer, size_t size) {
    return fgets(buffer, size, stdin);
}

int main() {
    pid_t pid = fork();

    if (pid < 0) {
        printf("Fork failed!");
        exit(1);
    } else if (pid > 0) {
        printf("Hello from parent! (My PID: %d) (Child PID: %d)\n",getpid(), pid);
        // Track status
        int status;
        waitpid(pid, &status, 0);
        if (WIFEXITED(status)) {
        printf("Parent: Child finished with exit status: %d\n", WEXITSTATUS(status));
        }
    } else {
        printf("Hello from child! (PID: %d)\n", pid);
        printf("My parent is (PID: %d)\n", getppid());
        // Simulate work
        sleep(2);
        // Exit with specific code
        exit(42);
    }
    return 0;
}
