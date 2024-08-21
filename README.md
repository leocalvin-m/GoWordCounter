
# Word and Line Counter

This Go program counts either words or lines from a file or standard input, depending on the command-line flag provided and the presence of file arguments.

## Table of Contents
- [Word and Line Counter](#word-and-line-counter)
  - [Table of Contents](#table-of-contents)
  - [Overview](#overview)
  - [Dependencies](#dependencies)
  - [Main Function](#main-function)
  - [Command-line Flags](#command-line-flags)
  - [Input Handling](#input-handling)
  - [Scanning and Counting](#scanning-and-counting)
  - [Error Handling](#error-handling)
  - [Output](#output)
  - [Usage](#usage)

## Overview

This program is a utility that reads from either a file or standard input and counts either words or lines. The behavior is controlled by a command-line flag and the presence of file arguments.

## Dependencies

The program uses the following standard Go packages:

- `bufio`: For efficient buffered I/O operations
- `flag`: For parsing command-line flags
- `fmt`: For formatted I/O operations
- `io`: For basic I/O primitives
- `log`: For logging messages
- `os`: For interacting with the operating system

## Main Function

The `main()` function is the entry point of the program. It sets up command-line flags, handles input sources, performs the counting operation, and outputs the result.

## Command-line Flags

- `-l`: A boolean flag that, when set, switches the program to count lines instead of words.

The flag is defined and parsed using:

```go
countLine := flag.Bool("l", false, "Count Lines")
flag.Parse()
```

## Input Handling

The program can read input from either a file or standard input:

```go
args := flag.Args()
var input io.Reader
if(len(args) > 0){
    file, err := os.Open(args[0])
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
        os.Exit(1)
    }
    defer file.Close()
    input = file
} else {
    input = os.Stdin
}
```

If a filename is provided as an argument, the program attempts to open and read from that file. Otherwise, it reads from standard input.

## Scanning and Counting

The program uses a `bufio.Scanner` to read from the input:

```go
scanner := bufio.NewScanner(input)
if !*countLine {
    scanner.Split(bufio.ScanWords)
}
```

If the `-l` flag is not set, the scanner is configured to split on words. Otherwise, it uses the default line splitting.

The counting is done in a simple loop:

```go
count := 0
for scanner.Scan() {
    count++
}
```

## Error Handling

The program includes error handling for file opening and scanning operations:

```go
if err != nil {
    fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
    os.Exit(1)
}

if err := scanner.Err(); err != nil {
    fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
    os.Exit(1)
}
```

If an error occurs, it's printed to standard error and the program exits with a non-zero status code.

## Output

The program uses the `log` package to output results:

```go
log.Printf("Count: %d\n", count)
```

This prints the final count to the standard error stream.

## Usage

To use the program:

1. Compile the Go file:
   ```
   go build main.go
   ```

2. Run the program:
   - To count words from a file:
     ```
     ./wc inputfile.txt
     ```
   - To count lines from a file:
     ```
     ./wc -l inputfile.txt
     ```
   - To count words from standard input:
     ```
     cat inputfile.txt | ./wc
     ```
   - To count lines from standard input:
     ```
     cat inputfile.txt | ./wc -l
     ```

Replace `inputfile.txt` with the path to your input file when reading from a file.
