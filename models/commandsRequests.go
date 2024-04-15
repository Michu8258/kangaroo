package models

type SolveCommandRequest struct {
	InputJsonFile  *string
	OutputJsonFile *string
	OutputTxtFile  *string
	BoxSize        *int8
	LayoutWidth    *int8
	LayoutHeight   *int8
}
