#   🔀 PUSH-SWAP

## OBJECTIVES

Push-Swap is a simple project that utilizes a Non-Comparative Sorting Algorithm. You have a list of integer values, two stacks (a and b), and a set of instructions at your disposal. The goal is to sort the stack a, which contains the integer values, in ascending order using both stacks and a set of instructions.

## PROGRAMS

There are two programs in this project:

### 1. `PUSH-SWAP`

- `push-swap` calculates and displays on the standard output the smallest program using the push-swap instruction language that sorts integer arguments received.

#### USAGE:

```shell
./push-swap [integer arguments]
```

- The integer arguments should be formatted as a list, with the first integer at the top of stack a.

#### OUTPUT:

- The program displays the smallest list of instructions possible to sort stack a, with the smallest number at the top. Instructions are separated by newline characters.

#### ERRORS:

- If there are errors such as non-integer arguments or duplicates, the program displays "Error" followed by a newline on the standard error.

- If there are no arguments, the program does not display anything (0 instructions).

### 2. `CHECKER`

- `checker` takes integer arguments and reads instructions from the standard input. It then executes the instructions and displays "OK" if the integers are sorted. Otherwise, it displays "KO."

#### USAGE:

```shell
[instructions] | ./checker [integer arguments]
```

- The integer arguments should be formatted as a list, with the first argument at the top of stack a. Instructions are provided via standard input.

#### OUTPUT:

- If after executing the instructions, stack a is sorted, and stack b is empty, the program displays "OK" followed by a newline on the standard output. In all other cases, it displays "KO" followed by a newline.

#### ERRORS:

- Errors, such as non-integer arguments, duplicates, or incorrectly formatted instructions, result in the program displaying "Error" followed by a newline on the standard error.

- If there are no arguments, the program displays nothing.

## USAGE

To build the `push-swap` and `checker` programs, you can use the provided scripts:

### 1. `build.sh`

- This script compiles the `push-swap` and `checker` programs and moves them to the appropriate location.

#### USAGE:

```shell
./build.sh
```

### 2. `audit.sh`

- This script checks your code for potential issues or violations of coding standards using various static analysis tools.

#### USAGE:

```shell
./audit.sh
```

These scripts simplify the process of building and auditing your code, ensuring that your programs are correctly compiled.

Please run `build.sh` before using the `push-swap` and `checker` programs, and use `audit.sh` to maintain code quality and style compliance.

##  AUTHORS
+   @abdouksow
+   @iyossang
+   @serignmbaye