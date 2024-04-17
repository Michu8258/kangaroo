package models

type SudokuConfigRequest struct {
	BoxSize      *int8
	LayoutWidth  *int8
	LayoutHeight *int8
	Overwrite    bool
}

func (r *SudokuConfigRequest) AsConfigRequest() *SudokuConfigRequest {
	return r
}

type SolveCommandRequest struct {
	SudokuConfigRequest
	InputJsonFile *string
	OutputFile    *string
}

type CreateCommandRequest struct {
	SudokuConfigRequest
}
