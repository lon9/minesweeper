package main

import (
	"fmt"
	"math/rand"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

// Board is board structure
type Board struct {
	xSize    int
	ySize    int
	mines    int
	tiles    [][]*Tile
	flagMode bool
}

// NewBoard is constructor
func NewBoard(difficalty int) *Board {
	var (
		xSize int
		ySize int
		mines int
	)
	if difficalty == Easy {
		xSize = 9
		ySize = 9
		mines = 10
	} else if difficalty == Normal {
		xSize = 16
		ySize = 16
		mines = 40
	} else if difficalty == Hard {
		xSize = 30
		ySize = 16
		mines = 99
	}
	tiles := make([][]*Tile, ySize)
	selected := make(map[int][]int)
	for i := 0; i < mines; i++ {
		var (
			xCandidate int
			yCandidate int
		)
		for {
			xCandidate = rand.Intn(xSize)
			yCandidate = rand.Intn(ySize)
			if ys, ok := selected[xCandidate]; ok {
				var isUsed bool
				for _, y := range ys {
					if y == yCandidate {
						// Already used
						isUsed = true
						break
					}
				}
				if !isUsed {
					selected[xCandidate] = append(selected[xCandidate], yCandidate)
					break
				}
			} else {
				selected[xCandidate] = []int{yCandidate}
				break
			}
		}
	}

	for y := 0; y < ySize; y++ {
		tiles[y] = make([]*Tile, xSize)
	L1:
		for x := 0; x < xSize; x++ {
			if ys, ok := selected[x]; ok {
				for _, yCand := range ys {
					if yCand == y {
						tiles[y][x] = NewTile(x, y, true)
						continue L1
					}
				}
			}
			tiles[y][x] = NewTile(x, y, false)
		}
	}

	for y := 0; y < ySize; y++ {
		for x := 0; x < xSize; x++ {
			var cnt int
			for y1 := y - 1; y1 < y+2; y1++ {
				for x1 := x - 1; x1 < x+2; x1++ {
					if y1 < 0 || y1 > ySize-1 || x1 < 0 || x1 > xSize-1 {
						continue
					}
					if tiles[y1][x1].GetIsMine() {
						cnt++
					}
				}
			}
			tiles[y][x].SetNumMines(cnt)
		}
	}

	return &Board{
		xSize: xSize,
		ySize: ySize,
		mines: mines,
		tiles: tiles,
	}
}

func (b *Board) showBoard() {
	for y := 0; y < b.ySize; y++ {
		for x := 0; x < b.xSize; x++ {
			if b.tiles[y][x].GetIsMine() {
				fmt.Print("b")
			} else {
				fmt.Print(b.tiles[y][x].GetNumMines())
			}
		}
		fmt.Println()
	}
}

// Size returns a size of Board as pixel
func (b *Board) Size() (int, int) {
	x := b.xSize*tileSize + (b.xSize+1)*tileMargin
	y := b.ySize*tileSize + (b.ySize+1)*tileMargin
	return x, y
}

// Update updares Board
func (b *Board) Update() error {
	for y := 0; y < b.ySize; y++ {
		for x := 0; x < b.xSize; x++ {
			if err := b.tiles[y][x].Update(); err != nil {
				return err
			}
		}
	}

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		posX, posY := ebiten.CursorPosition()
		x := posX / (tileSize + tileMargin)
		y := posY / (tileSize + tileMargin)
		if b.flagMode {
			b.tiles[y][x].SetIsFlag(!b.tiles[y][x].GetIsFlag())
		} else {
			if !b.tiles[y][x].GetIsFlag() {
				if b.tiles[y][x].GetIsMine() {
					fmt.Println("Game over")
					b.tiles[y][x].SetIsOpen(true)
				} else {
					if b.tiles[y][x].GetNumMines() == 0 {
						b.open(x, y)
					} else {
						b.tiles[y][x].SetIsOpen(true)
					}
				}
			}
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		b.flagMode = !b.flagMode
	}
	return nil
}

func (b *Board) open(xPos, yPos int) {
	b.tiles[yPos][xPos].SetIsOpen(true)
	for y := yPos - 1; y < yPos+2; y++ {
		for x := xPos - 1; x < xPos+2; x++ {
			if y < 0 || y > b.ySize-1 || x < 0 || x > b.xSize-1 {
				continue
			}
			if b.tiles[y][x].GetNumMines() == 0 && !b.tiles[y][x].GetIsOpen() {
				b.open(x, y)
			} else {
				if b.tiles[y][x].GetNumMines() != 0 && !b.tiles[y][x].GetIsMine() {
					b.tiles[y][x].SetIsOpen(true)
				}
			}
		}
	}
}

// Draw draws Board
func (b *Board) Draw(boardImage *ebiten.Image) {
	boardImage.Fill(frameColor)
	for y := 0; y < b.ySize; y++ {
		for x := 0; x < b.xSize; x++ {
			op := &ebiten.DrawImageOptions{}
			w := x*tileSize + (x+1)*tileMargin
			h := y*tileSize + (y+1)*tileMargin
			op.GeoM.Translate(float64(w), float64(h))
			r, g, b, a := colorToScale(tileBackgroundColor(false))
			op.ColorM.Scale(r, g, b, a)
			boardImage.DrawImage(tileImage, op)
		}
	}
	for y := 0; y < b.ySize; y++ {
		for x := 0; x < b.xSize; x++ {
			b.tiles[y][x].Draw(boardImage)
		}
	}
}
