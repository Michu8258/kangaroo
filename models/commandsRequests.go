package models

type SudokuConfigRequest struct {
	BoxSize      *int8
	LayoutWidth  *int8
	LayoutHeight *int8
}

func (r *SudokuConfigRequest) GetConfigRequest() *SudokuConfigRequest {
	return r
}

type SolveCommandRequest struct {
	SudokuConfigRequest
	InputJsonFile  *string
	OutputJsonFile *string
	OutputTxtFile  *string
}

type CreateCommandRequest struct {
	SudokuConfigRequest
	Overwrite bool
}
