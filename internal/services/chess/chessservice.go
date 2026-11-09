package chess

import (
	"Golang-Otus-Basic/internal/constants/chess"
	"errors"
	"fmt"
	"os"
	"strings"
)

func GenerateChessBoard() (string, error) {
	var chessBoard string
	var chessBoardBuilder strings.Builder

	width, height, firstName, secondName, err := handleInputParameters()

	if err != nil {
		return "", err
	}

	err = validateInputParameters(width, height)

	if err != nil {
		return "", err
	}

	for i := 1; i <= height; i++ {
		chessBoardBuilder.WriteString(fmt.Sprintf("%d ", i))

		for j := 1; j <= width; j++ {
			color := getCellColor(i, j)
			chessPiece := getChessPieceByLineColumnAndHeight(i, j, height)
			chessBoardBuilder.WriteString(color + chessPiece + chess.AnsiResetColorCommand)
		}

		err = tryAddPlayerNamesByLineAndHeight(i, height, &chessBoardBuilder, firstName, secondName)

		if err != nil {
			return "", err
		}

		chessBoardBuilder.WriteString("\n")
	}

	chessBoardBuilder.WriteString(" " + chess.Alphabet[:width])
	chessBoard = chessBoardBuilder.String()

	return chessBoard, nil
}

func handleInputParameters() (width, height int, firstName, lastName string, err error) {
	fmt.Printf("Введите ширину шахматного поля от %d до %d\n", chess.MinWidth, chess.MaxWidth)
	_, err = fmt.Fscan(os.Stdin, &width)

	if err != nil {
		return 0, 0, "", "", errors.New("не удалось обработать входящую ширину шахматаной доски")
	}

	fmt.Printf("Введите высоту шахматного поля от %d до %d\n", chess.MinHeight, chess.MaxHeight)
	_, err = fmt.Fscan(os.Stdin, &height)

	if err != nil {
		return 0, 0, "", "", errors.New("не удалось обработать входящую высоту шахматаной доски")
	}

	fmt.Printf("Введите имя первого игрока\n")
	_, err = fmt.Fscan(os.Stdin, &firstName)

	if err != nil {
		return 0, 0, "", "", errors.New("не удалось обработать имя первого игрока")
	}

	fmt.Printf("Введите имя первого игрока\n")
	_, err = fmt.Fscan(os.Stdin, &lastName)

	if err != nil {
		return 0, 0, "", "", errors.New("не удалось обработать имя второго игрока")
	}

	return width, height, firstName, lastName, nil
}

func validateInputParameters(width, height int) error {
	if width < chess.MinWidth || width > chess.MaxWidth {
		return errors.New("ширина шахматного поля находится за пределами допустимых значений")
	}

	if height < chess.MinHeight || height > chess.MaxHeight {
		return errors.New("высота шахматного поля находится за пределами допустимых значений")
	}

	return nil
}

func getChessPieceByLineColumnAndHeight(line, column, height int) string {
	columnIndex := column % chess.BaseWidth

	if line > 2 && line < height-1 {
		return chess.Empty
	}

	if line == 2 {
		return chess.BlackPawn
	}

	if line == height-1 {
		return chess.WhitePawn
	}

	switch columnIndex {
	case 1, 0:
		if line == 1 {
			return chess.BlackRook
		}

		return chess.WhiteRook
	case 2, 7:
		if line == 1 {
			return chess.BlackKnight
		}

		return chess.WhiteKnight
	case 3, 6:
		if line == 1 {
			return chess.BlackBishop
		}

		return chess.WhiteBishop
	case 4:
		if line == 1 {
			return chess.BlackQueen
		}

		return chess.WhiteQueen
	case 5:
		if line == 1 {
			return chess.BlackKing
		}

		return chess.WhiteKing
	}

	return chess.Empty
}

func getCellColor(line, column int) (color string) {
	if (line+column)%2 != 0 {
		return chess.AnsiBlackColorCommand
	}

	return chess.AnsiWhiteColorCommand
}

func tryAddPlayerNamesByLineAndHeight(line, height int, chessBoardBuilder *strings.Builder, firstName, secondName string) error {
	if &chessBoardBuilder == nil {
		return errors.New("не удалось получить string Builder при добавлении имени игрока")
	}

	if line == 1 {
		chessBoardBuilder.WriteString(chess.Indentation + "Игрок 1: " + firstName)

		return nil
	}

	if line == height {
		chessBoardBuilder.WriteString(chess.Indentation + "Игрок 2: " + secondName)

		return nil
	}

	return nil
}
