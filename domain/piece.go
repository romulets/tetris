package domain

type PieceColor uint32

const (
	RED    PieceColor = 0
	BLUE   PieceColor = 1
	GREEN  PieceColor = 2
	YELLOW PieceColor = 3
	PURPLE PieceColor = 4
	WHITE  PieceColor = 5
	ORANGE PieceColor = 6

	// T Visual hack to ease dev
	T bool = true
	// F Visual hack to ease dev
	F bool = false
)

type Piece struct {
	shape [][]bool
	color PieceColor
}

func buildSquareShape() Piece {
	return Piece{[][]bool{
		{T, T},
		{T, T},
	}, RED}
}

func buildSShape() Piece {
	return Piece{[][]bool{
		{F, T, F},
		{T, T, F},
		{T, F, F},
	}, BLUE}
}

func buildZShape() Piece {
	return Piece{[][]bool{
		{T, F, F},
		{T, T, F},
		{F, T, F},
	}, GREEN}
}

func buildTShape() Piece {
	return Piece{[][]bool{
		{T, T, T},
		{F, T, F},
		{F, F, F},
	}, YELLOW}
}

func buildLShape() Piece {
	return Piece{[][]bool{
		{T, F, F},
		{T, F, F},
		{T, T, F},
	}, PURPLE}
}

func buildMirroredLShape() Piece {
	return Piece{[][]bool{
		{F, F, T},
		{F, F, T},
		{F, T, T},
	}, WHITE}
}

func buildLineShape() Piece {
	return Piece{[][]bool{
		{T, T, T, T},
		{F, F, F, F},
		{F, F, F, F},
		{F, F, F, F},
	}, ORANGE}
}

func transposePieceAnticlockwise(originalPiece Piece) Piece {
	originalShape := originalPiece.shape
	newShape := make([][]bool, len(originalShape))

	for i := range originalShape {
		newShape[i] = make([]bool, len(originalShape[i]))
		copy(newShape[i], originalShape[i])
	}

	rowSize := len(originalShape)
	for i := 0; i < rowSize/2; i++ {
		colSize := len(originalShape[i])
		for j := 0; j < colSize-i-1; j++ {
			temp := originalShape[i][j]
			newShape[i][j] = originalShape[j][colSize-1-i]
			newShape[j][colSize-1-i] = originalShape[colSize-1-i][colSize-1-j]
			newShape[colSize-1-i][colSize-1-j] = originalShape[colSize-1-j][i]
			newShape[colSize-1-j][i] = temp
		}
	}

	return Piece{newShape, originalPiece.color}
}

func transposePieceClockwise(originalPiece Piece) Piece {
	originalShape := originalPiece.shape
	newShape := make([][]bool, len(originalShape))

	for i := range originalShape {
		newShape[i] = make([]bool, len(originalShape[i]))
		copy(newShape[i], originalShape[i])
	}

	rowSize := len(originalShape)
	for i := 0; i < rowSize/2; i++ {
		colSize := len(originalShape[i])
		for j := 0; j < colSize-i-1; j++ {
			temp := originalShape[i][j]
			newShape[i][j] = originalShape[colSize-1-j][i]
			newShape[colSize-1-j][i] = originalShape[colSize-1-i][colSize-1-j]
			newShape[colSize-1-i][colSize-1-j] = originalShape[j][colSize-1-i]
			newShape[j][colSize-1-i] = temp
		}
	}

	return Piece{newShape, originalPiece.color}
}
