package chess

type Move struct {
	initiator   Player
	oldPosition string
	newPosition string
	chessPiece  string
}

func NewMove(initiator Player, oldPosition, newPosition, chessPiece string) Move {
	return Move{
		initiator:   initiator,
		oldPosition: oldPosition,
		newPosition: newPosition,
		chessPiece:  chessPiece,
	}
}

func (move *Move) GetInfo() Move {
	return *move
}
