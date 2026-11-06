package constants

const (
	MinHeight = 4 //ограничение, чтобы уместить все ряды шахмат
	MaxHeight = 26
	MinWidth  = 1
	MaxWidth  = 26 //ограничене по количество букв в латинице
	BaseWidth = 8

	WhiteKing   = "♔"
	WhiteQueen  = "♕"
	WhiteRook   = "♖"
	WhiteBishop = "♗"
	WhiteKnight = "♘"
	WhitePawn   = "♙"

	BlackKing   = "♚"
	BlackQueen  = "♛"
	BlackRook   = "♜"
	BlackBishop = "♝"
	BlackKnight = "♞"
	BlackPawn   = "♟"

	Empty = "⛶"

	AnsiWhiteColorCommand = "\033[47m"
	AnsiBlackColorCommand = "\033[40m"
	AnsiResetColorCommand = "\033[0m"

	Indentation = "   "

	Alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
)
