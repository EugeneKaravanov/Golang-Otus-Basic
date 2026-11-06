package main

import (
	"Golang-Otus-Basic/services"
	"fmt"
	"os"
)

func main() {
	chessboard, err := services.GenerateChessBoard()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Сгенерированное шахматное поле:")
	fmt.Println(chessboard)
}
