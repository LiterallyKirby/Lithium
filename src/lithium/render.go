package lithium


import (
"github.com/hajimehoshi/ebiten/v2/ebitenutil"
"image/color"
"github.com/hajimehoshi/ebiten/v2"
)


// basic colors available to students
var (
Red = color.RGBA{R: 220, G: 50, B: 50, A: 255}
Green = color.RGBA{R: 50, G: 200, B: 50, A: 255}
Blue = color.RGBA{R: 50, G: 100, B: 220, A: 255}
White = color.RGBA{255, 255, 255, 255}
Black = color.RGBA{0, 0, 0, 255}
Yellow = color.RGBA{240, 220, 70, 255}
Cyan = color.RGBA{80, 200, 200, 255}
)


func drawRect(screen *ebiten.Image, e *Entity) {
ebitenutil.DrawRect(screen, e.Pos.X, e.Pos.Y, e.Scale.X, e.Scale.Y, e.Color)
}


// drawText is a small helper for debugging or scoreboard (keeps it simple)
func DrawText(screen *ebiten.Image, s string, x, y float64) {
ebitenutil.DebugPrintAt(screen, s, int(x), int(y))
}
