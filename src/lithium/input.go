package lithium


import (
ebiten "github.com/hajimehoshi/ebiten/v2"
)


// Key constants exported for students
var (
KeyLeft = ebiten.KeyArrowLeft
KeyRight = ebiten.KeyArrowRight
KeyUp = ebiten.KeyArrowUp
KeyDown = ebiten.KeyArrowDown
KeySpace = ebiten.KeySpace
KeyA = ebiten.KeyA
KeyW = ebiten.KeyW
KeyS = ebiten.KeyS
KeyD = ebiten.KeyD
)


var prevKeys = map[ebiten.Key]bool{}


func updateInputState() {
// populate prevKeys for JustPressed checks
for k := range prevKeys {
prevKeys[k] = ebiten.IsKeyPressed(k)
}
}


func IsPressed(k ebiten.Key) bool {
return ebiten.IsKeyPressed(k)
}


func JustPressed(k ebiten.Key) bool {
// very simple implementation: true if currently pressed and previously not recorded as pressed
cur := ebiten.IsKeyPressed(k)
prev := prevKeys[k]
if cur && !prev {
prevKeys[k] = true
return true
}
if !cur {
prevKeys[k] = false
}
return false
}
