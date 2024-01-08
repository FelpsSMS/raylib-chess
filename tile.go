package main

import (
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Tile struct {
	isBlack bool
	box     rl.Rectangle
	center  rl.Vector2
	row     int
	col     int
}

type BoardMarker struct {
	text string
	X    int32
	Y    int32
}

var boardRows = map[int]string{
	1: "a",
	2: "b",
	3: "c",
	4: "d",
	5: "e",
	6: "f",
	7: "g",
	8: "h",
}

func (t *Tile) CheckForPiecePlacement() {

	if currentPiece != nil && rl.CheckCollisionPointRec(rl.GetMousePosition(), t.box) {
		selectedTile = t

	} else {
		selectedTile = nil
	}
}

func CreateBoard() {
	black := true
	boardOffset := float32(30)
	tileWidth := float32((SCREEN_WIDTH - int(boardOffset)) / 8)
	tileHeight := float32((SCREEN_HEIGHT - int(boardOffset)) / 8)

	for i := 1; i < 9; i++ {
		black = !black

		for j := 1; j < 9; j++ {
			rect := rl.NewRectangle(tileWidth*float32(j-1)+boardOffset, tileHeight*float32(i-1), tileWidth, tileHeight)
			centerPoint := rl.Vector2{X: rect.X + rect.Width/2, Y: rect.Y + rect.Height/2}

			tile := Tile{isBlack: black, box: rect, center: centerPoint, col: j, row: i}
			black = !black

			tiles = append(tiles, &tile)

			boardMarkers = append(boardMarkers, BoardMarker{X: int32(rect.X) + int32(tileWidth)/2, Y: int32(SCREEN_HEIGHT) - int32(boardOffset), text: boardRows[j]})
		}

		boardMarkers = append(boardMarkers, BoardMarker{X: int32(boardOffset) / 2, Y: int32(tiles[len(tiles)-1].box.Y) + int32(tileHeight)/2, text: strconv.Itoa(i)})
	}
}
