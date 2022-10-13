package domain

import (
	"reflect"
	"testing"
)

func shapesEqual(t *testing.T, expected [][]bool, actual [][]bool) {
	if !reflect.DeepEqual(expected, actual) {
		t.Error("Expected shape ", expected, "Got shape", actual)
	}
}

func TestBuildSquareShape(t *testing.T) {
	expectedShape := [][]bool{
		{T, T},
		{T, T},
	}

	piece := buildSquareShape()
	shapesEqual(t, expectedShape, piece.shape)
}

func TestRotateSquareShapeClockwise(t *testing.T) {
	piece := buildSquareShape()

	transposedPiece1 := transposePieceClockwise(piece)
	expectedTransposedShape1 := [][]bool{
		{T, T},
		{T, T},
	}

	transposedPiece2 := transposePieceClockwise(transposedPiece1)
	expectedTransposedShape2 := [][]bool{
		{T, T},
		{T, T},
	}

	transposedPiece3 := transposePieceClockwise(transposedPiece2)
	expectedTransposedShape3 := [][]bool{
		{T, T},
		{T, T},
	}

	transposedPiece4 := transposePieceClockwise(transposedPiece3)
	expectedTransposedShape4 := [][]bool{
		{T, T},
		{T, T},
	}

	shapesEqual(t, expectedTransposedShape1, transposedPiece1.shape)
	shapesEqual(t, expectedTransposedShape2, transposedPiece2.shape)
	shapesEqual(t, expectedTransposedShape3, transposedPiece3.shape)
	shapesEqual(t, expectedTransposedShape4, transposedPiece4.shape)
}

func TestRotateSquareShapeAnticlockwise(t *testing.T) {
	piece := buildSquareShape()

	transposedPiece1 := transposePieceAnticlockwise(piece)
	expectedTransposedShape1 := [][]bool{
		{T, T},
		{T, T},
	}

	transposedPiece2 := transposePieceAnticlockwise(transposedPiece1)
	expectedTransposedShape2 := [][]bool{
		{T, T},
		{T, T},
	}

	transposedPiece3 := transposePieceAnticlockwise(transposedPiece2)
	expectedTransposedShape3 := [][]bool{
		{T, T},
		{T, T},
	}

	transposedPiece4 := transposePieceAnticlockwise(transposedPiece3)
	expectedTransposedShape4 := [][]bool{
		{T, T},
		{T, T},
	}

	shapesEqual(t, expectedTransposedShape1, transposedPiece1.shape)
	shapesEqual(t, expectedTransposedShape2, transposedPiece2.shape)
	shapesEqual(t, expectedTransposedShape3, transposedPiece3.shape)
	shapesEqual(t, expectedTransposedShape4, transposedPiece4.shape)
}

func TestBuildSShape(t *testing.T) {
	expectedShape := [][]bool{
		{F, T, F},
		{T, T, F},
		{T, F, F},
	}

	piece := buildSShape()
	shapesEqual(t, expectedShape, piece.shape)
}

func TestRotateSShapeClockwise(t *testing.T) {
	piece := buildSShape()

	transposedPiece1 := transposePieceClockwise(piece)
	expectedTransposedShape1 := [][]bool{
		{T, T, F},
		{F, T, T},
		{F, F, F},
	}

	transposedPiece2 := transposePieceClockwise(transposedPiece1)
	expectedTransposedShape2 := [][]bool{
		{F, F, T},
		{F, T, T},
		{F, T, F},
	}

	transposedPiece3 := transposePieceClockwise(transposedPiece2)
	expectedTransposedShape3 := [][]bool{
		{F, F, F},
		{T, T, F},
		{F, T, T},
	}

	transposedPiece4 := transposePieceClockwise(transposedPiece3)
	expectedTransposedShape4 := [][]bool{
		{F, T, F},
		{T, T, F},
		{T, F, F},
	}

	shapesEqual(t, expectedTransposedShape1, transposedPiece1.shape)
	shapesEqual(t, expectedTransposedShape2, transposedPiece2.shape)
	shapesEqual(t, expectedTransposedShape3, transposedPiece3.shape)
	shapesEqual(t, expectedTransposedShape4, transposedPiece4.shape)
}

func TestRotateSShapeAnticlockwise(t *testing.T) {
	piece := buildSShape()

	transposedPiece1 := transposePieceAnticlockwise(piece)
	expectedTransposedShape1 := [][]bool{
		{F, F, F},
		{T, T, F},
		{F, T, T},
	}

	transposedPiece2 := transposePieceAnticlockwise(transposedPiece1)
	expectedTransposedShape2 := [][]bool{
		{F, F, T},
		{F, T, T},
		{F, T, F},
	}

	transposedPiece3 := transposePieceAnticlockwise(transposedPiece2)
	expectedTransposedShape3 := [][]bool{
		{T, T, F},
		{F, T, T},
		{F, F, F},
	}

	transposedPiece4 := transposePieceAnticlockwise(transposedPiece3)
	expectedTransposedShape4 := [][]bool{
		{F, T, F},
		{T, T, F},
		{T, F, F},
	}

	shapesEqual(t, expectedTransposedShape1, transposedPiece1.shape)
	shapesEqual(t, expectedTransposedShape2, transposedPiece2.shape)
	shapesEqual(t, expectedTransposedShape3, transposedPiece3.shape)
	shapesEqual(t, expectedTransposedShape4, transposedPiece4.shape)
}

func TestBuildZShape(t *testing.T) {
	expectedShape := [][]bool{
		{T, F, F},
		{T, T, F},
		{F, T, F},
	}

	piece := buildZShape()
	shapesEqual(t, expectedShape, piece.shape)
}

func TestRotateZShapeClockwise(t *testing.T) {
	piece := buildZShape()

	transposedPiece1 := transposePieceClockwise(piece)
	expectedTransposedShape1 := [][]bool{
		{F, T, T},
		{T, T, F},
		{F, F, F},
	}

	transposedPiece2 := transposePieceClockwise(transposedPiece1)
	expectedTransposedShape2 := [][]bool{
		{F, T, F},
		{F, T, T},
		{F, F, T},
	}

	transposedPiece3 := transposePieceClockwise(transposedPiece2)
	expectedTransposedShape3 := [][]bool{
		{F, F, F},
		{F, T, T},
		{T, T, F},
	}

	transposedPiece4 := transposePieceClockwise(transposedPiece3)
	expectedTransposedShape4 := [][]bool{
		{T, F, F},
		{T, T, F},
		{F, T, F},
	}

	shapesEqual(t, expectedTransposedShape1, transposedPiece1.shape)
	shapesEqual(t, expectedTransposedShape2, transposedPiece2.shape)
	shapesEqual(t, expectedTransposedShape3, transposedPiece3.shape)
	shapesEqual(t, expectedTransposedShape4, transposedPiece4.shape)
}

func TestRotateZShapeAnticlockwise(t *testing.T) {
	piece := buildZShape()

	transposedPiece1 := transposePieceAnticlockwise(piece)
	expectedTransposedShape1 := [][]bool{
		{F, F, F},
		{F, T, T},
		{T, T, F},
	}

	transposedPiece2 := transposePieceAnticlockwise(transposedPiece1)
	expectedTransposedShape2 := [][]bool{
		{F, T, F},
		{F, T, T},
		{F, F, T},
	}

	transposedPiece3 := transposePieceAnticlockwise(transposedPiece2)
	expectedTransposedShape3 := [][]bool{
		{F, T, T},
		{T, T, F},
		{F, F, F},
	}

	transposedPiece4 := transposePieceAnticlockwise(transposedPiece3)
	expectedTransposedShape4 := [][]bool{
		{T, F, F},
		{T, T, F},
		{F, T, F},
	}

	shapesEqual(t, expectedTransposedShape1, transposedPiece1.shape)
	shapesEqual(t, expectedTransposedShape2, transposedPiece2.shape)
	shapesEqual(t, expectedTransposedShape3, transposedPiece3.shape)
	shapesEqual(t, expectedTransposedShape4, transposedPiece4.shape)
}

func TestBuildTShape(t *testing.T) {
	expectedShape := [][]bool{
		{T, T, T},
		{F, T, F},
		{F, F, F},
	}

	piece := buildTShape()
	shapesEqual(t, expectedShape, piece.shape)
}

func TestRotateTShapeClockwise(t *testing.T) {
	piece := buildTShape()

	transposedPiece1 := transposePieceClockwise(piece)
	expectedTransposedShape1 := [][]bool{
		{F, F, T},
		{F, T, T},
		{F, F, T},
	}

	transposedPiece2 := transposePieceClockwise(transposedPiece1)
	expectedTransposedShape2 := [][]bool{
		{F, F, F},
		{F, T, F},
		{T, T, T},
	}

	transposedPiece3 := transposePieceClockwise(transposedPiece2)
	expectedTransposedShape3 := [][]bool{
		{T, F, F},
		{T, T, F},
		{T, F, F},
	}

	transposedPiece4 := transposePieceClockwise(transposedPiece3)
	expectedTransposedShape4 := [][]bool{
		{T, T, T},
		{F, T, F},
		{F, F, F},
	}

	shapesEqual(t, expectedTransposedShape1, transposedPiece1.shape)
	shapesEqual(t, expectedTransposedShape2, transposedPiece2.shape)
	shapesEqual(t, expectedTransposedShape3, transposedPiece3.shape)
	shapesEqual(t, expectedTransposedShape4, transposedPiece4.shape)
}

func TestRotateTShapeAnticlockwise(t *testing.T) {
	piece := buildTShape()

	transposedPiece1 := transposePieceAnticlockwise(piece)
	expectedTransposedShape1 := [][]bool{
		{T, F, F},
		{T, T, F},
		{T, F, F},
	}

	transposedPiece2 := transposePieceAnticlockwise(transposedPiece1)
	expectedTransposedShape2 := [][]bool{
		{F, F, F},
		{F, T, F},
		{T, T, T},
	}

	transposedPiece3 := transposePieceAnticlockwise(transposedPiece2)
	expectedTransposedShape3 := [][]bool{
		{F, F, T},
		{F, T, T},
		{F, F, T},
	}

	transposedPiece4 := transposePieceAnticlockwise(transposedPiece3)
	expectedTransposedShape4 := [][]bool{
		{T, T, T},
		{F, T, F},
		{F, F, F},
	}

	shapesEqual(t, expectedTransposedShape1, transposedPiece1.shape)
	shapesEqual(t, expectedTransposedShape2, transposedPiece2.shape)
	shapesEqual(t, expectedTransposedShape3, transposedPiece3.shape)
	shapesEqual(t, expectedTransposedShape4, transposedPiece4.shape)
}

func TestBuildLShape(t *testing.T) {
	expectedShape := [][]bool{
		{T, F, F},
		{T, F, F},
		{T, T, F},
	}

	piece := buildLShape()
	shapesEqual(t, expectedShape, piece.shape)
}

func TestRotateLShapeClockwise(t *testing.T) {
	piece := buildLShape()

	transposedPiece1 := transposePieceClockwise(piece)
	expectedTransposedShape1 := [][]bool{
		{T, T, T},
		{T, F, F},
		{F, F, F},
	}

	transposedPiece2 := transposePieceClockwise(transposedPiece1)
	expectedTransposedShape2 := [][]bool{
		{F, T, T},
		{F, F, T},
		{F, F, T},
	}

	transposedPiece3 := transposePieceClockwise(transposedPiece2)
	expectedTransposedShape3 := [][]bool{
		{F, F, F},
		{F, F, T},
		{T, T, T},
	}

	transposedPiece4 := transposePieceClockwise(transposedPiece3)
	expectedTransposedShape4 := [][]bool{
		{T, F, F},
		{T, F, F},
		{T, T, F},
	}

	shapesEqual(t, expectedTransposedShape1, transposedPiece1.shape)
	shapesEqual(t, expectedTransposedShape2, transposedPiece2.shape)
	shapesEqual(t, expectedTransposedShape3, transposedPiece3.shape)
	shapesEqual(t, expectedTransposedShape4, transposedPiece4.shape)
}

func TestRotateLShapeAnticlockwise(t *testing.T) {
	piece := buildLShape()

	transposedPiece1 := transposePieceAnticlockwise(piece)
	expectedTransposedShape1 := [][]bool{
		{F, F, F},
		{F, F, T},
		{T, T, T},
	}

	transposedPiece2 := transposePieceAnticlockwise(transposedPiece1)
	expectedTransposedShape2 := [][]bool{
		{F, T, T},
		{F, F, T},
		{F, F, T},
	}

	transposedPiece3 := transposePieceAnticlockwise(transposedPiece2)
	expectedTransposedShape3 := [][]bool{
		{T, T, T},
		{T, F, F},
		{F, F, F},
	}

	transposedPiece4 := transposePieceAnticlockwise(transposedPiece3)
	expectedTransposedShape4 := [][]bool{
		{T, F, F},
		{T, F, F},
		{T, T, F},
	}

	shapesEqual(t, expectedTransposedShape1, transposedPiece1.shape)
	shapesEqual(t, expectedTransposedShape2, transposedPiece2.shape)
	shapesEqual(t, expectedTransposedShape3, transposedPiece3.shape)
	shapesEqual(t, expectedTransposedShape4, transposedPiece4.shape)
}

func TestBuildMirroredLShape(t *testing.T) {
	expectedShape := [][]bool{
		{F, F, T},
		{F, F, T},
		{F, T, T},
	}

	piece := buildMirroredLShape()
	shapesEqual(t, expectedShape, piece.shape)
}

func TestRotateMirroredLShapeClockwise(t *testing.T) {
	piece := buildMirroredLShape()

	transposedPiece1 := transposePieceClockwise(piece)
	expectedTransposedShape1 := [][]bool{
		{F, F, F},
		{T, F, F},
		{T, T, T},
	}

	transposedPiece2 := transposePieceClockwise(transposedPiece1)
	expectedTransposedShape2 := [][]bool{
		{T, T, F},
		{T, F, F},
		{T, F, F},
	}

	transposedPiece3 := transposePieceClockwise(transposedPiece2)
	expectedTransposedShape3 := [][]bool{
		{T, T, T},
		{F, F, T},
		{F, F, F},
	}

	transposedPiece4 := transposePieceClockwise(transposedPiece3)
	expectedTransposedShape4 := [][]bool{
		{F, F, T},
		{F, F, T},
		{F, T, T},
	}

	shapesEqual(t, expectedTransposedShape1, transposedPiece1.shape)
	shapesEqual(t, expectedTransposedShape2, transposedPiece2.shape)
	shapesEqual(t, expectedTransposedShape3, transposedPiece3.shape)
	shapesEqual(t, expectedTransposedShape4, transposedPiece4.shape)
}

func TestRotateMirroredLShapeAnticlockwise(t *testing.T) {
	piece := buildMirroredLShape()

	transposedPiece1 := transposePieceAnticlockwise(piece)
	expectedTransposedShape1 := [][]bool{
		{T, T, T},
		{F, F, T},
		{F, F, F},
	}

	transposedPiece2 := transposePieceAnticlockwise(transposedPiece1)
	expectedTransposedShape2 := [][]bool{
		{T, T, F},
		{T, F, F},
		{T, F, F},
	}

	transposedPiece3 := transposePieceAnticlockwise(transposedPiece2)
	expectedTransposedShape3 := [][]bool{
		{F, F, F},
		{T, F, F},
		{T, T, T},
	}

	transposedPiece4 := transposePieceAnticlockwise(transposedPiece3)
	expectedTransposedShape4 := [][]bool{
		{F, F, T},
		{F, F, T},
		{F, T, T},
	}

	shapesEqual(t, expectedTransposedShape1, transposedPiece1.shape)
	shapesEqual(t, expectedTransposedShape2, transposedPiece2.shape)
	shapesEqual(t, expectedTransposedShape3, transposedPiece3.shape)
	shapesEqual(t, expectedTransposedShape4, transposedPiece4.shape)
}

func TestBuildLineShape(t *testing.T) {
	expectedShape := [][]bool{
		{T, T, T, T},
		{F, F, F, F},
		{F, F, F, F},
		{F, F, F, F},
	}

	piece := buildLineShape()
	shapesEqual(t, expectedShape, piece.shape)
}

func TestRotateLineShapeClockwise(t *testing.T) {
	piece := buildLineShape()

	transposedPiece1 := transposePieceClockwise(piece)
	expectedTransposedShape1 := [][]bool{
		{F, F, F, T},
		{F, F, F, T},
		{F, F, F, T},
		{F, F, F, T},
	}

	transposedPiece2 := transposePieceClockwise(transposedPiece1)
	expectedTransposedShape2 := [][]bool{
		{F, F, F, F},
		{F, F, F, F},
		{F, F, F, F},
		{T, T, T, T},
	}

	transposedPiece3 := transposePieceClockwise(transposedPiece2)
	expectedTransposedShape3 := [][]bool{
		{T, F, F, F},
		{T, F, F, F},
		{T, F, F, F},
		{T, F, F, F},
	}

	transposedPiece4 := transposePieceClockwise(transposedPiece3)
	expectedTransposedShape4 := [][]bool{
		{T, T, T, T},
		{F, F, F, F},
		{F, F, F, F},
		{F, F, F, F},
	}

	shapesEqual(t, expectedTransposedShape1, transposedPiece1.shape)
	shapesEqual(t, expectedTransposedShape2, transposedPiece2.shape)
	shapesEqual(t, expectedTransposedShape3, transposedPiece3.shape)
	shapesEqual(t, expectedTransposedShape4, transposedPiece4.shape)
}

func TestRotateLineShapeAnticlockwise(t *testing.T) {
	piece := buildLineShape()

	transposedPiece1 := transposePieceAnticlockwise(piece)
	expectedTransposedShape1 := [][]bool{
		{T, F, F, F},
		{T, F, F, F},
		{T, F, F, F},
		{T, F, F, F},
	}

	transposedPiece2 := transposePieceAnticlockwise(transposedPiece1)
	expectedTransposedShape2 := [][]bool{
		{F, F, F, F},
		{F, F, F, F},
		{F, F, F, F},
		{T, T, T, T},
	}

	transposedPiece3 := transposePieceAnticlockwise(transposedPiece2)
	expectedTransposedShape3 := [][]bool{
		{F, F, F, T},
		{F, F, F, T},
		{F, F, F, T},
		{F, F, F, T},
	}

	transposedPiece4 := transposePieceAnticlockwise(transposedPiece3)
	expectedTransposedShape4 := [][]bool{
		{T, T, T, T},
		{F, F, F, F},
		{F, F, F, F},
		{F, F, F, F},
	}

	shapesEqual(t, expectedTransposedShape1, transposedPiece1.shape)
	shapesEqual(t, expectedTransposedShape2, transposedPiece2.shape)
	shapesEqual(t, expectedTransposedShape3, transposedPiece3.shape)
	shapesEqual(t, expectedTransposedShape4, transposedPiece4.shape)
}
