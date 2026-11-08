package chess

type Game struct {
	firstPlayer  Player
	secondPlayer Player
	board        Board
	moves        []Move
	isActive     bool
	winner       *Player
}

func NewGame(firstPlayer, secondPlayer Player, board Board, moves []Move) Game {
	return Game{
		firstPlayer:  firstPlayer,
		secondPlayer: secondPlayer,
		board:        board,
		moves:        moves,
		winner:       nil,
		isActive:     true,
	}
}

func (game *Game) GetInfo() Game {
	return *game
}

func (game *Game) UpdateBoard(boardState string) {
	game.board.UpdateBoard(boardState)
}

func (game *Game) AddMove(initiator Player, oldPosition, newPosition, chessPiece string) {
	game.moves = append(game.moves, NewMove(initiator, oldPosition, newPosition, chessPiece))
}

func (game *Game) SetWinner(player *Player) {
	game.winner = player
	game.isActive = false
}
