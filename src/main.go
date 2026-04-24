package main

import (
	"log"

	"github.com/bhaeussermann/ebitengine-astronaut/src/game"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
  ebiten.SetWindowTitle("Hello")
  ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
  ebiten.SetWindowSize(800, 600)
  ebiten.SetWindowSizeLimits(300, 300, -1, -1)
  error := ebiten.RunGame(game.NewGame())
  if error != nil {
    log.Fatal(error)
  }
}
