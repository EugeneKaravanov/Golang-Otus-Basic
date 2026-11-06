package main

import (
	"fmt"
	"os"
	"strings"
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
	var chessBoardBuilder strings.Builder
	var chessBoard string

	fmt.Printf("Введите ширину шахматного поля от %d до %d\n", minWidth, maxWidth)
	_, err := fmt.Fscan(os.Stdin, &inputWidth)

	if err != nil {
		fmt.Println("ОШИБКА: Не удается преобразовать введенную ширину в число")
		os.Exit(1)
	}

	if inputWidth < minWidth || inputWidth > maxWidth {
		fmt.Println("ОШИБКА: Введенная ширина не соответствует ограничениям")
		os.Exit(1)
	}

	fmt.Printf("Введите высоту шахматного поля от %d до %d\n", minHeight, maxHeight)
	_, err = fmt.Fscan(os.Stdin, &inputHeight)

	if err != nil {
		fmt.Println("ОШИБКА: Не удается преобразовать введенную высоту в число")
		os.Exit(1)
	}

	if inputHeight < minHeight || inputHeight > maxHeight {
		fmt.Println("ОШИБКА: Введенная высота не соответствует ограничениям")
		os.Exit(1)
	}

	for i := 1; i <= inputHeight; i++ {
		for j := 1; j <= inputWidth; j++ {
			if (i+j)%2 != 0 {
				chessBoardBuilder.WriteRune(blackCell)
			} else {
				chessBoardBuilder.WriteRune(whiteCell)
			}
		}

		chessBoardBuilder.WriteRune('\n')
	}

	chessBoard = chessBoardBuilder.String()
	fmt.Println("Сгенерированное шахматное поле:")
	fmt.Println(chessBoard)
}
