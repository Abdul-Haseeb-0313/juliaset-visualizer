package pallete

import(
	"image/color"
)

var p = []color.RGBA{
    {68, 1, 84, 255},    // dark purple
    {59, 82, 139, 255},  // blue
    {33, 145, 140, 255}, // teal
    {94, 201, 98, 255},  // green-yellow
    {253, 231, 37, 255}, // bright yellow
}

func lerp(a, b uint8, t float64) uint8 {
    return uint8(float64(a) + (float64(b)-float64(a))*t)
}

func ColorFromPalette(t float64) color.RGBA {
    if t <= 0 {
        return p[0]
    }
    if t >= 1 {
        return p[len(p)-1]
    }

    scaled := t * float64(len(p)-1)
    i := int(scaled)
    frac := scaled - float64(i)

    c1 := p[i]
    c2 := p[i+1]

    return color.RGBA{
        lerp(c1.R, c2.R, frac),
        lerp(c1.G, c2.G, frac),
        lerp(c1.B, c2.B, frac),
        255,
    }
}
