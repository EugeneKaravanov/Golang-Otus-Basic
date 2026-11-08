package main

import (
	"Golang-Otus-Basic/internal/services/chess"
	"fmt"
	"os"
)

func main() {
	chessboard, err := chess.GenerateChessBoard()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Сгенерированное шахматное поле:")
	fmt.Println(chessboard)
}
