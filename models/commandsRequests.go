package models

type SudokuConfigRequest struct {
	BoxSize      *int8
	LayoutWidth  *int8
	LayoutHeight *int8
}

func (r *SudokuConfigRequest) AsConfigRequest() *SudokuConfigRequest {
	return r
}

type SolveCommandRequest struct {
	SudokuConfigRequest
	Overwrite     bool
	InputJsonFile *string
	OutputFile    *string
}

type CreateCommandRequest struct {
	SudokuConfigRequest
	Overwrite bool
}
