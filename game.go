package main

import (
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten"
)

const (
	// Easy mode
	Easy int = iota
	// Normal mode
	Normal
	//Hard mode
	Hard
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Game is structure for Game
type Game struct {
	board      *Board
	boardImage *ebiten.Image
}

// NewGame is constructor for Game
func NewGame(difficulty int) *Game {
	return &Game{
		board: NewBoard(difficulty),
	}
}

// Size returns size of Board as pixel
func (g *Game) Size() (int, int) {
	return g.board.Size()
}

// Update updates Game
func (g *Game) Update() error {
	if err := g.board.Update(); err != nil {
		return err
	}
	return nil
}

// Draw draws Game
func (g *Game) Draw(screen *ebiten.Image) {
	if g.boardImage == nil {
		w, h := g.board.Size()
		g.boardImage, _ = ebiten.NewImage(w, h, ebiten.FilterDefault)
	}
	screen.Fill(backgroundColor)
	g.board.Draw(g.boardImage)
	op := &ebiten.DrawImageOptions{}
	sw, sh := screen.Size()
	bw, bh := g.boardImage.Size()
	x := (sw - bw) / 2
	y := (sh - bh) / 2
	op.GeoM.Translate(float64(x), float64(y))
	screen.DrawImage(g.boardImage, op)
}
