package lithium


// Raycast performs a simple stepped raycast from start to end and
// returns the first entity hit (naive but fine for demos).
func Raycast(start, end Vec2, steps int) (bool, *Entity) {
if steps <= 0 {
steps = 100
}
dx := (end.X - start.X) / float64(steps)
dy := (end.Y - start.Y) / float64(steps)
x := start.X
y := start.Y
for i := 0; i <= steps; i++ {
for _, e := range world.Entities {
if pointInsideEntity(x, y, e) {
return true, e
}
}
x += dx
y += dy
}
return false, nil
}


func pointInsideEntity(x, y float64, e *Entity) bool {
return x >= e.Pos.X && x <= e.Pos.X+e.Scale.X && y >= e.Pos.Y && y <= e.Pos.Y+e.Scale.Y
}
