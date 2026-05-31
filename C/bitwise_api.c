#include <stdio.h>
#define READ (1 << 0)
#define WRITE (1 << 1)
#define EXECUTE (1 << 2)
#define ADMIN (1 << 3)

int grant_permission(char perm, int permissions);
int remove_permission(char perm, int permissions);
int toggle_permission(char perm, int permissions);
int has_perm(char perm, int permissions);
void print_perm(int permissions);

int main() {
    int perm;
    int action;
    int permissions = 0b0000;
    while (1) {
        printf("ENTER PERM TO MODIFY (RWXA): ");
        perm = getchar();
        if (perm != 'R' && perm != 'W' && perm != 'X' && perm != 'A') {
            continue;
        }

        getchar(); // Consumes the '\n' left in the buffer

        printf("ENTER ACTION (GRTH): ");
        action = getchar();
        switch (action) {
            case 'G':
                permissions = grant_permission(perm, permissions);
                break;
            case 'R':
                permissions = remove_permission(perm, permissions);
                break;
            case 'T':
                permissions = toggle_permission(perm, permissions);
                break;
            case 'H':
                if (has_perm(perm, permissions)) {
                    printf("Has Permissions (%c) Enabled\n", perm);
                } else {
                    printf("Has Permissions (%c) Disabled\n", perm);
                }
                break;
        }

        getchar(); // Consumes the '\n' left in the buffer
        print_perm(permissions);
    }
    return 0;
}

int grant_permission(char perm, int permissions) {
    switch (perm) {
        case 'R':
            return permissions |= READ;
        case 'W':
            return permissions |= WRITE;
        case 'X':
            return permissions |= EXECUTE;
        case 'A':
            return permissions |= ADMIN;
        default:
            return permissions;

    }
}

int remove_permission(char perm, int permissions) {
    switch (perm) {
        case 'R':
            return permissions &= ~READ;
        case 'W':
            return permissions &= ~WRITE;
        case 'X':
            return permissions &= ~EXECUTE;
        case 'A':
            return permissions &= ~ADMIN;
        default:
            return permissions;
    }
}

int toggle_permission(char perm, int permissions) {
    switch (perm) {
        case 'R':
            return permissions ^= READ;
        case 'W':
            return permissions ^= WRITE;
        case 'X':
            return permissions ^= EXECUTE;
        case 'A':
            return permissions ^= ADMIN;
        default:
            return permissions;
    }
}

int has_perm(char perm, int permissions) {
    switch (perm) {
        case 'R':
            return permissions & READ;
        case 'W':
            return permissions & WRITE;
        case 'X':
            return permissions & EXECUTE;
        case 'A':
            return permissions & ADMIN;
        default:
            return permissions;
    }
}

void print_perm(int permissions) {
    if (has_perm('R', permissions)) {
        putchar('R');
    }
    if (has_perm('W', permissions)) {
        putchar('W');
    }
    if (has_perm('X', permissions)) {
        putchar('X');
    }
    if (has_perm('A', permissions)) {
        putchar('A');
    }
    putchar('\n');
}