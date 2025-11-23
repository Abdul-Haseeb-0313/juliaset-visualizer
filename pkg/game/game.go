package game

import(

	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/Abdul-Haseeb-0313/juliaset-visualizer/pkg/window"
	"github.com/Abdul-Haseeb-0313/juliaset-visualizer/pkg/math"
    "github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
}

func convToCod(x, y float64) (float64, float64){
	return x * ((math.Xmax - math.Xmin)/window.WIDTH) + math.Xmin, y * ((math.Ymax - math.Ymin)/window.HEIGHT) + math.Ymin;
} 

var cursorX = 0;
var cursorY = 0;



func (g *Game) Update() error {
	

	return nil
}


func (g *Game) Draw(screen *ebiten.Image) {
	
	x, y := ebiten.CursorPosition()

	if(! (x < 0 || x > window.WIDTH || y < 0 || y > window.HEIGHT) ){
		cursorX, cursorY = x, y
	}
	
	cx, cy := convToCod(float64(cursorX), float64(cursorY));

	
	
	fmt.Println(cx, cy)
	
	for i := 0; i <= window.WIDTH; i++{
		for j := 0; j <= window.HEIGHT; j++{
			px, py := convToCod(float64(i), float64(j));
			col := math.JuliaColorAt(px, py, cx, cy);
			
			screen.Set(i, j, col)
			
		}
	}
	msg := fmt.Sprintf("c = %f + i%f", cx, cy);
	ebitenutil.DebugPrint(screen, msg)
	
}


func (g *Game) Layout(outW, outH int) (int, int){
	return window.WIDTH, window.HEIGHT
}


func NewGame() *Game{
	return &Game{}
}



