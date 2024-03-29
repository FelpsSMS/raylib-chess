package main

import rl "github.com/gen2brain/raylib-go/raylib"

type PieceType int

const (
	KING PieceType = iota
	QUEEN
	BISHOP
	ROOK
	PAWN
	KNIGHT
)

type Piece struct {
	id          int
	pieceType   PieceType
	pos         rl.Vector2
	originalPos rl.Vector2
	isBlack     bool
}

func (p *Piece) CheckForPlayerMove() {
	bufferZone := float32(15.0)
	mousePos := rl.GetMousePosition()

	if rl.IsMouseButtonDown(rl.MouseButtonLeft) {
		if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
			if rl.CheckCollisionPointCircle(mousePos, p.pos, CIRCLE_RADIUS) ||
				rl.CheckCollisionPointCircle(mousePos, rl.Vector2{X: p.pos.X, Y: p.pos.Y}, CIRCLE_RADIUS+bufferZone) {
				currentPiece = p
				dragOffset = rl.NewVector2(mousePos.X-p.pos.X, mousePos.Y-p.pos.Y)
			}
		}

		if currentPiece == p {
			p.pos.X = mousePos.X - dragOffset.X
			p.pos.Y = mousePos.Y - dragOffset.Y
		}
	} else {
		if currentPiece == p {
		outer:
			for _, tile := range tilesToBeHighlighted {
				if rl.CheckCollisionPointRec(rl.GetMousePosition(), tile.box) {
					for _, piece := range pieces {
						if tile.center == piece.pos {

							if piece.isBlack {
								currentPiece = nil
								p.pos = p.originalPos
								break outer

							} else {
								index := FindElementIndex(pieces, piece)

								if index != -1 {
									pieces = RemoveFromSlice(pieces, index)
								}
							}
						}
					}
					p.originalPos = tile.center
					break
				}
			}

			currentPiece = nil
			p.pos = p.originalPos
		}
	}
}
