package math

import(
	"image/color"
	"math/cmplx"
	"math"
	"github.com/Abdul-Haseeb-0313/juliaset-visualizer/pkg/pallete"
)

const Xmin = -2.0
const Xmax =  2.0
const Ymin = -2.0
const Ymax =  2.0
const maxIter = 100

func JuliaColorAt(x, y, cx, cy float64) color.Color{
	z := complex(x, y);
	c := complex(cx, cy);

	iter := 0

	for ; iter < maxIter; iter++{
		z = z*z + c
		if cmplx.Abs(z) > 2{
			break
		}
	}
	
	if iter == maxIter{
		return color.Black
	}
	
	nu := float64(iter) + 1 - math.Log2(math.Log2(cmplx.Abs(z)))
	t := nu / float64(maxIter);

	return pallete.ColorFromPalette(t)
}



