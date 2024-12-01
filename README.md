# Proteus
**File Polyglot Generator**

Proteus is a Go-based tool designed to generate polyglot files. A polyglot file is a file that can be interpreted as multiple file formats.

## Features

- Polyglot File Generation: Combine different file types into a single polyglot file.
- Supported Formats: Currently supports JPG and PDF file formats.
- Command-Line Interface: Easy to use CLI for generating polyglot files.

## Usage
To generate a polyglot file, you can use the following command:
```bash
go run main.go -i samples/file.jpg,samples/file.pdf --output file.pdf
```
### Command-Line Options

- `-i, --input: Input files to be combined. Example: -i file1.jpg,file2.pdf`
- `-o, --output: Output filename for the polyglot file.`

### Example
To create a polyglot file from `sample.jpg` and `sample.pdf`, run:
```bash
go run main.go -i samples/sample.jpg,samples/sample.pdf --output output.pdf
```

## How It Works
Proteus works by reading the input files, identifying their formats, and concatenating their byte data to create a polyglot file.

1. Reading Files:
The tool reads the input files byte-by-byte.

2. Identifying File Types:
It uses specific parsers to identify the file types (JPG, PDF).

3. Concatenating Data:
The byte data of the identified files is concatenated to form a single polyglot file.

4. Writing Output:
The concatenated data is written to the specified output file.

## Installation
```bash
git clone https://github.com/Kaiser784/Proteus.git
cd Proteus
go build
```

