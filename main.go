package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"
)

var (
	game *Game
)

func update(screen *ebiten.Image) error {

	if err := game.Update(); err != nil {
		return err
	}
	if ebiten.IsDrawingSkipped() {
		return nil
	}
	game.Draw(screen)
	return nil
}

func main() {
	game = NewGame(Hard)
	w, h := game.Size()
	if err := ebiten.Run(update, w, h, 1, "Minesweeper"); err != nil {
		log.Fatal(err)
	}
}
