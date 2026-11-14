package main

import (
	"Golang-Otus-Basic/internal/services/chess"
	"fmt"
)

func main() {
	chessService := chess.NewChessService()
	go chessService.StartGenerateAndSaveModels()
	err := chess.PlayChess()

	if err != nil {
		fmt.Println(err)
	}
}
