package chess

type Board struct {
	boardState [][]string
}

func NewBoard(boardState [][]string) Board {
	return Board{
		boardState: boardState,
	}
}

func (board *Board) GetInfo() Board {
	return *board
}

func (board *Board) GetBoardState() [][]string {
	result := make([][]string, len(board.boardState))

	for i := range board.boardState {
		result[i] = make([]string, len(board.boardState[i]))
		copy(result[i], board.boardState[i])
	}

	return result
}

func (board *Board) UpdateBoard(boardState [][]string) {
	board.boardState = boardState
}

func (board Board) Save() {}
