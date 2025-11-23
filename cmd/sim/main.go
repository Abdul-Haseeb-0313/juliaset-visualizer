package main

import(
	"github.com/Abdul-Haseeb-0313/juliaset-visualizer/pkg/game"
	"github.com/Abdul-Haseeb-0313/juliaset-visualizer/pkg/window"
	"github.com/hajimehoshi/ebiten/v2"
	
	"log"
)



func main(){
	
	ebiten.SetWindowSize(window.WIDTH, window.HEIGHT)
	ebiten.SetWindowTitle("Juliaset visualizer")

	g := game.NewGame()

	if err := ebiten.RunGame(g); err != nil {
        log.Fatal(err)
    }
}
