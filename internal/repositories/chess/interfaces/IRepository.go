package chess

import chessModelIntefaces "Golang-Otus-Basic/internal/models/chess/interfaces"

type IRepository interface {
	AddModel(model chessModelIntefaces.ISavingModel)
}
