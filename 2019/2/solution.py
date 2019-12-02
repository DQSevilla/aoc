import sys

INPUT_FILE = 'input'
ADD = 1
MUL = 2
HALT = 99

def fix_intcode_program(ip):
    """
    Solve AOC 2019 #2-1
    """
    ptr = 0
    while (opcode := ip[ptr]) != HALT:
        [a_ptr, b_ptr, res_ptr] = ip[ptr + 1 : ptr + 4]
        a, b = ip[a_ptr], ip[b_ptr]

        if opcode == ADD:
            ip[res_ptr] = a + b
        elif opcode == MUL:
            ip[res_ptr] = a * b
        else:
            sys.exit('invalid intcode program: invalid opcode %d' % opcode)
        ptr += 4
    return ip

def main():
    with open(INPUT_FILE) as file:
        intcode_program = [int(s) for s in file.read().split(',')]
    intcode_program[1] = 12
    intcode_program[2] = 2
    intcode_program = fix_intcode_program(intcode_program)

    print('Value left at position 0 after intcode program halts:',
          intcode_program[0])

if __name__ == '__main__':
    main()

