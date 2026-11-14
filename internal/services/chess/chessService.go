package chess

import (
	chessConstants "Golang-Otus-Basic/internal/constants/chess"
	chessModels "Golang-Otus-Basic/internal/models/chess"
	chessRepositories "Golang-Otus-Basic/internal/repositories/chess"
	chessInterfaces "Golang-Otus-Basic/internal/repositories/chess/interfaces"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type ChessService struct {
	repository chessInterfaces.IRepository
}

func NewChessService() ChessService {
	return ChessService{
		repository: chessRepositories.NewChessModelsRepository(),
	}
}

func (service *ChessService) StartGenerateAndSaveModels() {

	for {
		modelType := rand.Intn(chessConstants.GeneratedModelTypes)

		switch modelType {
		case 0:
			{
				model, err := generateChessBoard(8, 8)

				if err == nil {
					service.repository.AddModel(model)
				}
			}
		case 1:
			{
				var moves []chessModels.Move
				player1 := chessModels.NewPlayer("Иван")
				player2 := chessModels.NewPlayer("Антон")
				board, _ := generateChessBoard(8, 8)
				moves = append(moves, chessModels.NewMove(chessModels.NewPlayer("Антон"), "E4", "E4", chessConstants.BlackKnight))
				moves = append(moves, chessModels.NewMove(chessModels.NewPlayer("Иван"), "E4", "E4", chessConstants.WhiteKing))
				model := chessModels.NewGame(player1, player2, board, moves)
				service.repository.AddModel(model)
			}
		case 2:
			{
				model := chessModels.NewMove(chessModels.NewPlayer("Антон"), "E4", "E4", chessConstants.BlackKnight)
				service.repository.AddModel(model)
			}
		case 3:
			{
				model := chessModels.NewPlayer("Иван")
				service.repository.AddModel(model)
			}
		}

		delay := rand.Intn(chessConstants.MaxGenerationPause-chessConstants.MinGenerationPause+1) + chessConstants.MinGenerationPause

		time.Sleep(time.Duration(delay) * time.Second)
	}
}

func PlayChess() error {
	width, height, firstName, secondName, err := handleInputParameters()
	if err != nil {
		return err
	}

	board, err := generateChessBoard(width, height)
	if err != nil {
		return err
	}

	var moves []chessModels.Move

	player1 := chessModels.NewPlayer(firstName)
	player2 := chessModels.NewPlayer(secondName)

	game := chessModels.NewGame(player1, player2, board, moves)

	isFirstPlayerTurn := true
	isActiveGame := true

	for isActiveGame {
		clearConsole()
		printChessBoard(board, player1, player2)

		currentPlayer := player2
		if isFirstPlayerTurn {
			currentPlayer = player1
		}

		fmt.Printf("\nХод игрока: %s\n", currentPlayer.GetName())

		fmt.Println("\nВыберите действие:")
		fmt.Println("1. Сделать ход")
		fmt.Println("2. Автоход")
		fmt.Println("3. Сдаться")

		var action int
		fmt.Scan(&action)

		switch action {
		case 1:
			oldPosition, newPosition, err := handleMoveInput()
			if err != nil {
				fmt.Println(err)
				continue
			}

			err = executeMove(&game, &board, currentPlayer, oldPosition, newPosition)
			if err != nil {
				fmt.Println(err)
				continue
			}

			isFirstPlayerTurn = !isFirstPlayerTurn

		case 2:
			fmt.Println("Введите количество автоходов:")

			var count int
			_, err := fmt.Scan(&count)

			if err != nil || count <= 0 {
				fmt.Println("Количество автоходов должно быть положительным числом")
				continue
			}

			for i := 0; i < count; i++ {
				currentPlayer = player2
				if isFirstPlayerTurn {
					currentPlayer = player1
				}

				oldPosition, newPosition := generateRandomPositions(width, height)

				err := executeMove(
					&game,
					&board,
					currentPlayer,
					oldPosition,
					newPosition,
				)

				if err != nil {
					return err
				}

				isFirstPlayerTurn = !isFirstPlayerTurn
			}

		case 3:
			fmt.Printf("\nИгрок %s сдался. Игра завершена!\n", currentPlayer.GetName())
			isActiveGame = false

		default:
			fmt.Println("Неизвестное действие")
		}
	}

	return nil
}

func generateChessBoard(width, height int) (chessModels.Board, error) {
	err := validateInputParameters(width, height)
	chessBoard := make([][]string, height)

	for i := range chessBoard {
		chessBoard[i] = make([]string, width)
	}

	if err != nil {
		return chessModels.Board{}, err
	}

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			chessPiece := getChessPieceByLineColumnAndHeight(i, j, height)
			chessBoard[i][j] = chessPiece
		}
	}

	return chessModels.NewBoard(chessBoard), nil
}

func handleInputParameters() (width, height int, firstName, lastName string, err error) {
	fmt.Printf("Введите ширину шахматного поля от %d до %d\n", chessConstants.MinWidth, chessConstants.MaxWidth)
	_, err = fmt.Fscan(os.Stdin, &width)

	if err != nil {
		return 0, 0, "", "", errors.New("не удалось обработать входящую ширину шахматаной доски")
	}

	fmt.Printf("Введите высоту шахматного поля от %d до %d\n", chessConstants.MinHeight, chessConstants.MaxHeight)
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
	if width < chessConstants.MinWidth || width > chessConstants.MaxWidth {
		return errors.New("ширина шахматного поля находится за пределами допустимых значений")
	}

	if height < chessConstants.MinHeight || height > chessConstants.MaxHeight {
		return errors.New("высота шахматного поля находится за пределами допустимых значений")
	}

	return nil
}

func getChessPieceByLineColumnAndHeight(line, column, height int) string {
	columnIndex := column % chessConstants.BaseWidth

	if line > 1 && line < height-2 {
		return chessConstants.Empty
	}

	if line == 1 {
		return chessConstants.BlackPawn
	}

	if line == height-2 {
		return chessConstants.WhitePawn
	}

	switch columnIndex {
	case 0, 7:
		if line == 0 {
			return chessConstants.BlackRook
		}

		return chessConstants.WhiteRook
	case 1, 6:
		if line == 0 {
			return chessConstants.BlackKnight
		}

		return chessConstants.WhiteKnight
	case 2, 5:
		if line == 0 {
			return chessConstants.BlackBishop
		}

		return chessConstants.WhiteBishop
	case 3:
		if line == 0 {
			return chessConstants.BlackQueen
		}

		return chessConstants.WhiteQueen
	case 4:
		if line == 0 {
			return chessConstants.BlackKing
		}

		return chessConstants.WhiteKing
	}

	return chessConstants.Empty
}

func getCellColor(line, column int) (color string) {
	if (line+column)%2 != 0 {
		return chessConstants.AnsiBlackColorCommand
	}

	return chessConstants.AnsiWhiteColorCommand
}

func printChessBoard(board chessModels.Board, player1, player2 chessModels.Player) {
	var chessBoard string
	var chessBoardBuilder strings.Builder
	var err error
	boardState := board.GetBoardState()
	height := len(boardState)
	width := len(boardState[0])

	for i := 0; i < height; i++ {
		chessBoardBuilder.WriteString(fmt.Sprintf("%d ", height-i))

		for j := 0; j < width; j++ {
			color := getCellColor(i, j)
			chessPiece := boardState[i][j]
			chessBoardBuilder.WriteString(color + chessPiece + chessConstants.AnsiResetColorCommand)
		}

		err = tryAddPlayerNamesByLineAndHeight(i, height, &chessBoardBuilder, player1.GetName(), player2.GetName())

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		chessBoardBuilder.WriteString("\n")
	}

	chessBoardBuilder.WriteString(" " + chessConstants.Alphabet[:width])
	chessBoard = chessBoardBuilder.String()

	fmt.Println("Сгенерированное шахматное поле:")
	fmt.Println(chessBoard)
}

func tryAddPlayerNamesByLineAndHeight(line, height int, chessBoardBuilder *strings.Builder, firstName, secondName string) error {
	if &chessBoardBuilder == nil {
		return errors.New("не удалось получить string Builder при добавлении имени игрока")
	}

	if line == 0 {
		chessBoardBuilder.WriteString(chessConstants.Indentation + "Игрок 2: " + secondName)

		return nil
	}

	if line == height-1 {
		chessBoardBuilder.WriteString(chessConstants.Indentation + "Игрок 1: " + firstName)

		return nil
	}

	return nil
}

func clearConsole() {
	fmt.Print("\033[2J\033[H\033[3J")
}

func handleMoveInput() (string, string, error) {
	var oldPosition, newPosition string

	fmt.Println("Введите начальную позицию:")
	_, err := fmt.Scan(&oldPosition)

	if err != nil {
		return "", "", errors.New("не удалось прочитать начальную позицию")
	}

	fmt.Println("Введите новую позицию:")
	_, err = fmt.Scan(&newPosition)

	if err != nil {
		return "", "", errors.New("не удалось прочитать конечную позицию")
	}

	return strings.ToUpper(oldPosition), strings.ToUpper(newPosition), nil
}

func executeMove(game *chessModels.Game, board *chessModels.Board, player chessModels.Player, oldPosition, newPosition string) error {
	state := board.GetBoardState()

	oldRow, oldColumn, err := parsePosition(oldPosition, len(state), len(state[0]))

	if err != nil {
		return err
	}

	newRow, newColumn, err := parsePosition(newPosition, len(state), len(state[0]))
	if err != nil {
		return err
	}

	if oldRow == newRow && oldColumn == newColumn {
		return errors.New("начальная и конечная позиции совпадают")
	}

	chessPiece := state[oldRow][oldColumn]

	state[newRow][newColumn] = chessPiece
	state[oldRow][oldColumn] = chessConstants.Empty

	board.UpdateBoard(state)
	game.UpdateBoard(state)

	game.AddMove(
		player,
		oldPosition,
		newPosition,
		chessPiece,
	)

	return nil
}

func parsePosition(position string, height, width int) (int, int, error) {
	position = strings.ToUpper(position)

	if len(position) < 2 {
		return 0, 0, errors.New("неверный формат координат")
	}

	column := int(position[0] - 'A')

	row, err := strconv.Atoi(position[1:])
	if err != nil {
		return 0, 0, errors.New("неверный номер строки")
	}

	row = height - row

	if column < 0 || column >= width || row < 0 || row >= height {
		return 0, 0, errors.New("координаты находятся за пределами доски")
	}

	return row, column, nil
}

func generateRandomPositions(width, height int) (string, string) {
	oldColumn := 'A' + rune(rand.Intn(width))
	oldRow := rand.Intn(height) + 1

	newColumn := 'A' + rune(rand.Intn(width))
	newRow := rand.Intn(height) + 1

	for oldColumn == newColumn && oldRow == newRow {
		newColumn = 'A' + rune(rand.Intn(width))
		newRow = rand.Intn(height) + 1
	}

	oldPosition := fmt.Sprintf("%c%d", oldColumn, oldRow)
	newPosition := fmt.Sprintf("%c%d", newColumn, newRow)

	return oldPosition, newPosition
}
