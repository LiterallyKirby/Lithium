package main


import (
lithium "github.com/LiterallyKirby/Lithium"
)


func main() {
// Student-friendly example: a player square and a ground platform.
player := lithium.AddEntity("Player", lithium.Rect, lithium.Vec2{X: 100, Y: 100}, true)
lithium.Scale(player, 32, 32)
lithium.SetColor(player, lithium.Cyan)


ground := lithium.AddEntity("Ground", lithium.Rect, lithium.Vec2{X: 0, Y: 420}, false)
lithium.Scale(ground, 640, 40)
lithium.SetColor(ground, lithium.Green)


lithium.OnUpdate(func() {
if lithium.IsPressed(lithium.KeyRight) {
lithium.Move(player, 2, 0)
}
if lithium.IsPressed(lithium.KeyLeft) {
lithium.Move(player, -2, 0)
}
if lithium.JustPressed(lithium.KeySpace) {
lithium.ApplyForce(player, 0, -8)
}
})


lithium.Run()
}
