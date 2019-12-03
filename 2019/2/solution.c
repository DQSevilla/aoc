
#include <sys/types.h>

#include <err.h>
#include <errno.h>
#include <fcntl.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>

#define INPUT_FILE "input"

enum opcode {ADD = 1, MUL = 2, HALT = 99};

static void fix_intcode_program(int *, int);

// precondition: program points to a sequence of at least len integers
static void fix_intcode_program(int *program, int len) {
    int ip, a, b;

    for (ip = 0; program[ip] != HALT && ip < len; ip += 4) {
        a = program[program[ip + 1]];
        b = program[program[ip + 2]];

        if (program[ip] == ADD) {
            program[program[ip + 3]] = a + b;
        } else if (program[ip] == MUL) {
            program[program[ip + 3]] = a * b;
        } else {
            errx(EXIT_FAILURE, "invalid intcode program: invalid opcode "
                "%d", program[ip]);
        }
    }

    if (ip >= len) {
        errx(EXIT_FAILURE, "invalid intcode program: no halting opcode");
    }
}

int main(void) {
    char buffer[BUFSIZ];
    int program[BUFSIZ];
    char *token;
    int fd, nr, len;

    if ((fd = open(INPUT_FILE, O_RDONLY)) == -1) {
        err(EXIT_FAILURE, "open");
    }

    memset(buffer, 0, sizeof(buffer));
    if ((nr = read(fd, buffer, sizeof(buffer))) == -1) {
        err(EXIT_FAILURE, "read");
    }

    (void)close(fd);

    token = strtok(buffer, ",");
    for (len = 0; token != NULL; ++len) {
        if (sscanf(token, "%d", &program[len]) == EOF) {
            if (errno != 0) {
                err(EXIT_FAILURE, "sscanf");
            } else {
                errx(EXIT_FAILURE, "token \"%s\" not an integer", token);
            }
        }
        token = strtok(NULL, ",");
    }

    if (len < 3) {
        errx(EXIT_FAILURE, "invalid intcode program scanned from file");
    }

    program[1] = 12;
    program[2] = 2;
    fix_intcode_program(program, len);

    (void)printf("Value left at position 0 after intcode program halts: %d\n",
        program[0]);
}
