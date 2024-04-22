package commands

import "github.com/urfave/cli/v2"

func (commandConfig *CommandContext) ExecuteCommand() *cli.Command {
	return &cli.Command{
		Name:    "exec",
		Aliases: []string{"e"},
		Usage: "Solves a sudoku puzzle provided through argument as base64 representation\n" +
			"of sudoku binary data and outputs similarly encoded solution to the terminal.\n" +
			"You can find more about this format here:\nhttps://github.com/Michu8258/kangaroo/blob/main/documentation/binaryFormat.md",
		Flags: []cli.Flag{},
		Action: func(context *cli.Context) error {
			return commandConfig.executeCommandHandler(context.Args())
		},
	}
}

// executeCommandHandler is an entry point function for exec sudoku command
func (commandConfig *CommandContext) executeCommandHandler(arguments cli.Args) error {
	if arguments.Len() < 1 {
		commandConfig.ServiceCollection.TerminalPrinter.PrintError(
			"Please provide a base64 representation of a sudoku to this command.")
		commandConfig.ServiceCollection.TerminalPrinter.PrintNewLine()
		return nil
	}

	sudokuDto, err := commandConfig.ServiceCollection.SudokuEncoder.ReadFromBase64(
		arguments.First())

	if err != nil {
		commandConfig.ServiceCollection.TerminalPrinter.PrintError(
			"Failed to parse provided data to sudoku object.")
		commandConfig.ServiceCollection.TerminalPrinter.PrintNewLine()
		return nil
	}

	sudoku, ok := commandConfig.executeSudokuInitialization(sudokuDto, false)
	if !ok {
		return nil
	}

	solved, errs := commandConfig.ServiceCollection.Solver.Solve(sudoku)
	if !solved {
		commandConfig.ServiceCollection.TerminalPrinter.
			PrintError("Failed to solve the sudoku.")
		commandConfig.ServiceCollection.TerminalPrinter.PrintNewLine()
		return nil
	}

	if commandConfig.Settings.UseDebugPrints && len(errs) >= 1 {
		commandConfig.ServiceCollection.DataPrinter.PrintErrors(
			"Sudoku solution failure reasons:", err)
		return nil
	}

	solutionBase64, err := commandConfig.ServiceCollection.SudokuEncoder.ToBase64(
		sudoku.ToSudokuDto())

	if err != nil {
		commandConfig.ServiceCollection.TerminalPrinter.PrintError(
			"Failed to encode output sudoku solution.")
		commandConfig.ServiceCollection.TerminalPrinter.PrintNewLine()
		return nil
	}

	commandConfig.ServiceCollection.TerminalPrinter.PrintDefault(solutionBase64)
	commandConfig.ServiceCollection.TerminalPrinter.PrintNewLine()
	return nil
}
