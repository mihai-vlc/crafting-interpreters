package scanner

type Position struct {
	line   int
	column int
}

func NewPosition(line int, column int) *Position {
	return &Position{
		line:   line,
		column: column,
	}
}
