package main

import (
	"fmt"
	"os"
)

func main() {
	const whiteCell = ' '
	const blackCell = '#'
	const minHeight = 1
	const maxHeight = 20
	const minWidth = 1
	const maxWidth = 20

	var inputHeight int
	var inputWidth int
	var chessboard string

	fmt.Printf("Введите ширину шахматного поля от %d до %d\n", minWidth, maxWidth)
	fmt.Fscan(os.Stdin, &inputWidth)

	if inputWidth < minWidth || inputWidth > maxWidth {
		fmt.Println("ОШИБКА: Введенная ширина не соответствует ограничениям")
		os.Exit(0)
	}

	fmt.Printf("Введите высоту шахматного поля от %d до %d\n", minHeight, maxHeight)
	fmt.Fscan(os.Stdin, &inputHeight)

	if inputHeight < minHeight || inputHeight > maxHeight {
		fmt.Println("ОШИБКА: Введенная высота не соответствует ограничениям")
		os.Exit(0)
	}

	for i := 1; i <= inputHeight; i++ {
		for j := 1; j <= inputWidth; j++ {
			if (i%2 != 0 && j%2 == 0) || (i%2 == 0 && j%2 != 0) {
				chessboard += string(blackCell)
			} else {
				chessboard += string(whiteCell)
			}
		}

		chessboard += string("\n")
	}

	fmt.Println("Сгенерированное шахматное поле:")
	fmt.Println(chessboard)
}
