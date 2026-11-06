package services

import (
	"Golang-Otus-Basic/constants"
	"errors"
	"fmt"
	"os"
)

func GenerateChessBoard() (string, error) {
	var chessBoard string
	width, height, firstName, lastName, err := handleInputParameters()

	if err != nil {
		fmt.Println("При обработке входящих параметров произошла ошибка")

		return "", err
	}

	err = validateInputParameters(width, height)

	if err != nil {
		fmt.Println("Валидация входных параметров прошла неуспешна по причине: ", err.Error())

		return "", err
	}

	for i := 1; i <= height; i++ {
		for j := 1; j <= width; j++ {
			color := getCellColor(i, j)
			chessPiece := getChessPieceByLineColumnAndHeight(i, j, height)
			chessBoard += color + chessPiece + constants.AnsiResetColorCommand
		}

		chessBoard += string("\n")
	}

	chessBoard += firstName
	chessBoard += lastName

	return chessBoard, nil
}

func handleInputParameters() (width, height int, firstName, lastName string, err error) {
	fmt.Printf("Введите ширину шахматного поля от %d до %d\n", constants.MinWidth, constants.MaxWidth)
	_, err = fmt.Fscan(os.Stdin, &width)

	if err != nil {
		return 0, 0, "", "", err
	}

	fmt.Printf("Введите высоту шахматного поля от %d до %d\n", constants.MinHeight, constants.MaxHeight)
	_, err = fmt.Fscan(os.Stdin, &height)

	if err != nil {
		return 0, 0, "", "", err
	}

	fmt.Printf("Введите имя первого игрока")
	_, err = fmt.Fscan(os.Stdin, &firstName)

	if err != nil {
		return 0, 0, "", "", err
	}

	fmt.Printf("Введите имя первого игрока")
	_, err = fmt.Fscan(os.Stdin, &lastName)

	if err != nil {
		return 0, 0, "", "", err
	}

	return width, height, firstName, lastName, nil
}

func validateInputParameters(width, height int) error {
	if width < constants.MinWidth || width > constants.MaxWidth {
		return errors.New("ширина шахматного поля находится за пределами допустимых значений")
	}

	if height < constants.MinHeight || height > constants.MaxHeight {
		return errors.New("высота шахматного поля находится за пределами допустимых значений")
	}

	return nil
}

func getChessPieceByLineColumnAndHeight(line, column, height int) string {
	columnIndex := column % constants.BaseWidth

	if line > 2 && line < height-1 {
		return "⛶"
	}

	if line == 2 {
		return constants.BlackPawn
	}

	if line == height-1 {
		return constants.WhitePawn
	}

	switch columnIndex {
	case 1, 0:
		if line == 1 {
			return constants.BlackRook
		}

		return constants.WhiteRook
	case 2, 7:
		if line == 1 {
			return constants.BlackKnight
		}

		return constants.WhiteKnight
	case 3, 6:
		if line == 1 {
			return constants.BlackBishop
		}

		return constants.WhiteBishop
	case 4:
		if line == 1 {
			return constants.BlackQueen
		}

		return constants.WhiteQueen
	case 5:
		if line == 1 {
			return constants.BlackKing
		}

		return constants.WhiteKing
	}

	return constants.Empty
}

func getCellColor(line, column int) (color string) {
	if (line%2 != 0 && column%2 == 0) || (line%2 == 0 && column%2 != 0) {
		return constants.AnsiBlackColorCommand
	}

	return constants.AnsiWhiteColorCommand
}
