package chess

import (
	chessModels "Golang-Otus-Basic/internal/models/chess"
	chessModelsInterfaces "Golang-Otus-Basic/internal/models/chess/interfaces"
)

type ChessModelsRepository struct {
	boards  []chessModels.Board
	games   []chessModels.Game
	moves   []chessModels.Move
	players []chessModels.Player
}

func NewChessModelsRepository() *ChessModelsRepository {
	return &ChessModelsRepository{}
}

func (repository *ChessModelsRepository) AddModel(model chessModelsInterfaces.ISavingModel) {
	switch m := any(model).(type) {
	case chessModels.Board:
		repository.boards = append(repository.boards, m)
	case chessModels.Game:
		repository.games = append(repository.games, m)
	case chessModels.Move:
		repository.moves = append(repository.moves, m)
	case chessModels.Player:
		repository.players = append(repository.players, m)
	}
}
