package main

import (
	"image/color"
	"log"
	"strconv"

	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/text"
	"golang.org/x/image/font"
)

var (
	tileImage       *ebiten.Image
	mplusSmallFont  font.Face
	mplusNormalFont font.Face
	mplusBigFont    font.Face
)

func init() {
	tileImage, _ = ebiten.NewImage(tileSize, tileSize, ebiten.FilterDefault)
	tileImage.Fill(color.White)

	tt, err := truetype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}

	const dpi = 72
	mplusSmallFont = truetype.NewFace(tt, &truetype.Options{
		Size:    24,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	mplusNormalFont = truetype.NewFace(tt, &truetype.Options{
		Size:    18,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	mplusBigFont = truetype.NewFace(tt, &truetype.Options{
		Size:    48,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
}

// Tile is structure for Tile
type Tile struct {
	x        int
	y        int
	isMine   bool
	numMines int
	isOpen   bool
	isFlag   bool
}

// NewTile is constructor for Tile
func NewTile(x, y int, isMine bool) *Tile {
	return &Tile{
		x:      x,
		y:      y,
		isMine: isMine,
	}
}

// GetIsMine is getter for isMine
func (t *Tile) GetIsMine() bool {
	return t.isMine
}

// SetNumMines is setter for numMines
func (t *Tile) SetNumMines(n int) {
	t.numMines = n
}

// GetNumMines is getter for numMines
func (t *Tile) GetNumMines() int {
	return t.numMines
}

// GetIsOpen is getter for isOpen
func (t *Tile) GetIsOpen() bool {
	return t.isOpen
}

// SetIsOpen is setter for isOpen
func (t *Tile) SetIsOpen(isOpen bool) {
	t.isOpen = isOpen
}

// GetIsFlag getter for isFlag
func (t *Tile) GetIsFlag() bool {
	return t.isFlag
}

// SetIsFlag is setter for isFlag
func (t *Tile) SetIsFlag(isFlag bool) {
	t.isFlag = isFlag
}

// Update updates a tile
func (t *Tile) Update() error {
	return nil
}

// Draw draws a tile
func (t *Tile) Draw(boardImage *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	x := t.x*tileSize + (t.x+1)*tileMargin
	y := t.y*tileSize + (t.y+1)*tileMargin
	op.GeoM.Translate(float64(x), float64(y))
	r, g, b, a := colorToScale(tileBackgroundColor(t.isOpen))
	op.ColorM.Scale(r, g, b, a)
	boardImage.DrawImage(tileImage, op)

	drawString := func(s string) {
		f := mplusNormalFont
		bound, _ := font.BoundString(f, s)
		w := (bound.Max.X - bound.Min.X).Ceil()
		h := (bound.Max.Y - bound.Min.Y).Ceil()
		x = x + (tileSize-w)/2
		y = y + (tileSize-h)/2 + h
		text.Draw(boardImage, s, f, x, y, tileColor())
	}

	if t.isOpen {
		var s string
		if !t.isMine {
			if t.numMines != 0 {
				s = strconv.Itoa(t.numMines)
			}
		} else {
			s = "b"
		}
		drawString(s)
	} else if t.isFlag {
		drawString("?")
	}
}
