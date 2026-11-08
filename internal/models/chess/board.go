package chess

type Board struct {
	boardState string
}

func NewBoard(boardState string) Board {
	return Board{
		boardState: boardState,
	}
}

func (board *Board) GetInfo() Board {
	return *board
}

func (board *Board) UpdateBoard(boardState string) {
	board.boardState = boardState
}
