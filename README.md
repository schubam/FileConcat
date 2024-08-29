# FileConcat
**FileConcat** is a Go tool that concatenates multiple code files, removes comments, and outputs the clean code to STDOUT.

## Features

- Concatenates multiple files into one.
- Removes both single-line and multi-line comments.
- Supports wildcard patterns for file selection.
- Outputs the cleaned, concatenated code to STDOUT.

## Installation

To build FileConcat from source, ensure you have [Go installed](https://golang.org/dl/), then clone this repository and run:

```bash
git clone https://github.com/schubam/FileConcat.git
cd FileConcat
go build -o FileConcat
## Usage

You can use FileConcat by passing in individual filenames or wildcard patterns. The results will be printed to STDOUT.

```bash
./FileConcat file1.swift file2.go "**/*.js"
```

## Contributing

Feel free to submit issues or pull requests if you have suggestions for improvements or new features.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
