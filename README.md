# Kangaroo

A simple CLI that will help you to solve a Sudoku puzzle. It implements Crook's method - [a preemptive sets](https://www.sudokuwiki.org/Crooks_Algorithm) alrogithm to solve the sudoku.

### Quick usage

Use `kangaroo create -o <path to file>` command to use terminal's editor to configure the sudoku and save sudoku configuration to a json file.

If you have sudoku file, you can use `kangaroo solve -i <path to file>` to solve sudoku from the file.

Or just use `kangaroo solve` to provide sudoku configuration through the terminal and solve it.

<img src="./documentation/images/SudokuValuesInput.png" alt="Terminal input" width="500"/>

### Commands

**There are just 2 commands in the CLI:**

**create**

```
NAME:
   Kangaroo create - Creates sudoku puzzle data and saves to provided file paths (JSON and TXT files supported, default is JSON)

USAGE:
   Kangaroo create [command options] [arguments...]

OPTIONS:
   --box-size value, -s value         How many rows and columns single sudoku box has - in case of classic sudoku it is 3 (default: 0)
   --layout-width value, --lw value   How many boxes there are in the row - in case of classic sudoku it is 3 (default: 0)
   --layout-height value, --lh value  How many boxes there are in the column - in case of classic sudoku it is 3 (default: 0)
   --overwrite, -r                    Overwrite provided file(s) paths if exist (default: false)
   --help, -h   
```

**solve**

```
NAME:
   Kangaroo solve - Solves a provided sudoku puzzle

USAGE:
   Kangaroo solve [command options] [arguments...]

OPTIONS:
   --box-size value, -s value         How many rows and columns single sudoku box has - in case of classic sudoku it is 3 (default: 0)
   --layout-width value, --lw value   How many boxes there are in the row - in case of classic sudoku it is 3 (default: 0)
   --layout-height value, --lh value  How many boxes there are in the column - in case of classic sudoku it is 3 (default: 0)
   --overwrite, -r                    Overwrite provided file(s) paths if exist (default: false)
   --input-file value, -i value       Specify path to sudoku JSON configuration file
   --output-file value, -o value      Specify path to file where you want to save solution of the sudoku (JSON or TXT, JSON is default)
   --help, -h                         show help
```

### Documentation

Fore more information, please navigate to [./documentation](./documentation/nomenclature.md) directory of this repository.

### Stack

- Go
- [bubbletea](https://github.com/charmbracelet/bubbletea)
- [gouuid](https://github.com/nu7hatch/gouuid)
- [urfave cli](https://github.com/urfave/cli)