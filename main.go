package main

import (
	"log"
	"os"
	"reflect"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var SCREEN_WIDTH = 1024
var SCREEN_HEIGHT = 800
var GAME_TITLE = "GO CHESS"
var logger = log.New(os.Stdout, "LOG: ", log.Ldate|log.Ltime|log.Lshortfile)
var tiles []*Tile
var boardMarkers []BoardMarker
var pieces []*Piece
var CIRCLE_RADIUS = float32(30)
var currentPiece *Piece
var isDragging = false
var dragOffset rl.Vector2
var selectedTile *Tile

func FindElementIndex[T any](slice []T, element T) int {
	for index, elementInSlice := range slice {
		if reflect.DeepEqual(elementInSlice, element) {
			return index
		}
	}

	return -1
}

func RemoveFromSlice[T any](slice []T, index int) []T {
	return append(slice[:index], slice[index+1:]...)
}

func main() {
	rl.InitWindow(int32(SCREEN_WIDTH), int32(SCREEN_HEIGHT), GAME_TITLE)
	defer rl.CloseWindow()

	//Disable esc key for closing the game
	rl.SetExitKey(0)

	rl.SetTargetFPS(60)

	CreateBoard()
	setupDebugPieces()

	for !rl.WindowShouldClose() {

		for _, piece := range pieces {
			piece.CheckForPlayerMove()
		}

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		drawTiles()
		drawDebugPieces()

		rl.EndDrawing()
	}
}

func drawTiles() {
	BOARD_MARKER_FONT_SIZE := int32(20)

	for _, tile := range tiles {
		tile.CheckForPiecePlacement()
		color := rl.White

		if tile.isBlack {
			color = rl.Black
		}

		rl.DrawRectangleRec(tile.box, color)

		if selectedTile != nil && currentPiece != nil {
			logger.Printf("SELECTED ROW %v | COL %v | ID %v", selectedTile.row, selectedTile.col, selectedTile)
			logger.Printf("TILE ROW %v | COL %v", tile.row, tile.col)

			if tile.row == selectedTile.row && tile.col == selectedTile.col {
				rl.DrawRectangleLines(int32(tile.box.X), int32(tile.box.Y), int32(tile.box.Width), int32(tile.box.Height), rl.Red)
			}
		}

	}

	for _, marker := range boardMarkers {
		rl.DrawText(marker.text, marker.X, marker.Y, BOARD_MARKER_FONT_SIZE, rl.DarkBrown)
	}
}

func setupDebugPieces() {
	id := 0
	var pieceType PieceType

	for _, tile := range tiles {
		isEmpty := true

		if tile.row == 2 || tile.row == 7 {
			pieceType = PAWN
			isEmpty = false
		}

		if (tile.row == 8 && tile.col == 1) ||
			(tile.row == 8 && tile.col == 8) ||
			(tile.row == 1 && tile.col == 1) ||
			(tile.row == 1 && tile.col == 8) {
			pieceType = ROOK
			isEmpty = false
		}

		if (tile.row == 8 && tile.col == 2) ||
			(tile.row == 8 && tile.col == 7) ||
			(tile.row == 1 && tile.col == 2) ||
			(tile.row == 1 && tile.col == 7) {
			pieceType = KNIGHT
			isEmpty = false
		}

		if (tile.row == 8 && tile.col == 3) ||
			(tile.row == 8 && tile.col == 6) ||
			(tile.row == 1 && tile.col == 3) ||
			(tile.row == 1 && tile.col == 6) {
			pieceType = BISHOP
			isEmpty = false
		}

		if (tile.row == 8 && tile.col == 4) || (tile.row == 1 && tile.col == 4) {
			pieceType = QUEEN
			isEmpty = false
		}

		if (tile.row == 8 && tile.col == 5) || (tile.row == 1 && tile.col == 5) {
			pieceType = KING
			isEmpty = false
		}

		if !isEmpty {
			piece := Piece{pieceType: pieceType, pos: tile.center, originalPos: tile.center, id: id}
			pieces = append(pieces, &piece)

			id++
		}
	}
}

func drawDebugPieces() {

	for _, piece := range pieces {
		rl.DrawCircle(int32(piece.pos.X), int32(piece.pos.Y), CIRCLE_RADIUS, rl.Brown)
		var text string

		switch piece.pieceType {

		case PAWN:
			text = "P"

		case ROOK:
			text = "R"

		case KNIGHT:
			text = "K"

		case BISHOP:
			text = "B"

		case KING:
			text = "KG"

		case QUEEN:
			text = "Q"

		default:
			text = "P"
		}

		rl.DrawText(text, int32(piece.pos.X-CIRCLE_RADIUS/4), int32(piece.pos.Y-CIRCLE_RADIUS/4), 16, rl.Red)
	}
}
