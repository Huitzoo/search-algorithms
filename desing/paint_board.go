package desing

import (
	"image"
	"image/color"
	"image/png"
	"os"
	"search/utils"
)

func PaintBoard(
	size []int,
	goalID int,
	startID int,
	exploredIDs map[int]struct{},
	blockedCellIDs map[int]struct{},
) {

	width := size[0] * utils.SizeSquareForCell
	height := size[1] * utils.SizeSquareForCell

	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	lastMaxY := 0
	maxY := 0

	for y := 0; y < size[1]; y++ {
		maxY += utils.SizeSquareForCell
		lastMaxX := 0
		maxX := 0
		for x := 0; x < size[0]; x++ {
			colorCell := determinateColor(
				[]int{x, y},
				size[0],
				exploredIDs,
				blockedCellIDs,
				goalID,
				startID,
			)
			maxX += utils.SizeSquareForCell

			for yPaint := lastMaxY; yPaint < maxY; yPaint++ {
				for xPaint := lastMaxX; xPaint < maxX; xPaint++ {

					if xPaint < maxX-utils.LinesBoardWide && yPaint < maxY-utils.LinesBoardWide {
						img.Set(xPaint, yPaint, colorCell)
					} else {
						img.Set(xPaint, yPaint, color.Black)
					}
				}
			}
			lastMaxX = maxX
		}
		lastMaxY = maxY
	}

	f, _ := os.Create("image2.png")
	png.Encode(f, img)
}

func determinateColor(
	point []int,
	columns int,
	exploredIDs map[int]struct{},
	blockedCellIDs map[int]struct{},
	goalID int,
	startID int,
) color.Color {
	id := utils.CalculateIDStateByCoords(point, columns)
	if _, exists := blockedCellIDs[id]; exists {
		return color.RGBA{48, 48, 48, 0xff}
	} else if _, exists := exploredIDs[id]; exists {
		if id == startID {
			return color.RGBA{50, 205, 50, 0xff}
		}
		return color.RGBA{253, 218, 13, 0xff}
	} else if goalID == id {
		return color.RGBA{220, 20, 60, 0xff}
	}
	return color.White
}

func PaintCell() {

}
