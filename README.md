# Building `ccwc` Program

`ccwc` is a CLI tool that reads input files or standard input and provides a summary of the number of bytes, characters, words, and lines. It functions similarly to the Linux `wc` command, commonly used for text processing.


## Getting Started

Follow the steps below to set up, build, and test the application:


---

## Usage

### Command Syntax
```bash
ccwc [flags]
```

### Flags

| Flag                 | Description                            |
|----------------------|----------------------------------------|
| `-c`, `--bytes`      | Count only bytes                      |
| `-m`, `--chars`      | Count only characters                 |
| `-f`, `--file`       | Specify the path to the input file     |
| `-h`, `--help`       | Display help information              |
| `-l`, `--lines`      | Count only lines                      |
| `-w`, `--words`      | Count only words                      |

### Examples

1. Count characters, words, and lines in a file:
   ```bash
   ccwc file.txt
   ```

2. Count only the lines in a file:
   ```bash
   ccwc -l file.txt
   ```

3. Process input from a pipe:
   ```bash
   echo "Hello World" | ccwc
   ```

---

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/HrithikSawant/go-ccwc
   cd go-ccwc
   ```

2. Build the tool:
   ```bash
   go build -o ccwc main.go
   ```

3. Verify installation:
   ```bash
   ./ccwc --help
   ```

---

## Running Tests

Unit tests are provided in the `pkg/utils/utils_test.go` file. Run tests using:
```bash
go test .
```

Sample output:
```bash
go-ccwc/pkg/utils
ok  	github.com/HrithikSawant/go-ccwc/pkg/utils	(cached)
```

---


## Steps to Implement `ccwc`

### Step 1: Byte Count (`-c`)
Add functionality to count bytes in a file. Example usage:
```bash
./ccwc -c test.txt
342190 test.txt
```

### Step 2: Line Count (`-l`)
Extend functionality to count lines in a file. Example usage:
```bash
./ccwc -l test.txt
7145 test.txt
```

### Step 3: Word Count (`-w`)
Extend functionality to count words in a file. Example usage:
```bash
./ccwc -w test.txt
58164 test.txt
```

### Step 4: Character Count (`-m`)
Add functionality to count characters in a file. Example usage:
```bash
./ccwc -m test.txt
339292 test.txt
```

If multibyte characters are not supported, the result will match the output for `-c`.

### Step 5: Default Count
Implement the default behavior to show lines, words, and bytes if no options are provided. Example usage:
```bash
./ccwc test.txt
7145 58164 342190 test.txt
```


**DISCLAIMER**: This project is intended for educational purposes as part of coding challenges and should not be used in production environments or for critical applications. It is inspired by the coding challenges created by John Crickett, available at [codingchallenges.fyi](https://codingchallenges.fyi/challenges/challenge-wc).


## License

This project is licensed under the [MIT License](LICENSE).
