package lithium


func checkCollisions() {
// naive O(n^2) collision checking, fine for small scenes
for _, a := range world.Entities {
for _, b := range world.Entities {
if a == b {
continue
}
if a.Visible && b.Visible && aabbOverlap(a, b) {
// simple resolution: push A outside B on Y axis
resolveAABB(a, b)
if a.OnCollide != nil {
a.OnCollide(b)
}
}
}
}
}


func aabbOverlap(a, b *Entity) bool {
ax1 := a.Pos.X
ay1 := a.Pos.Y
ax2 := a.Pos.X + a.Scale.X
ay2 := a.Pos.Y + a.Scale.Y


bx1 := b.Pos.X
by1 := b.Pos.Y
bx2 := b.Pos.X + b.Scale.X
by2 := b.Pos.Y + b.Scale.Y


return ax1 < bx2 && ax2 > bx1 && ay1 < by2 && ay2 > by1
}


func resolveAABB(a, b *Entity) {
// naive Y-axis resolution for platformer-style behavior
// determine previous overlap depth and push up
ax1 := a.Pos.X
ay1 := a.Pos.Y
ax2 := a.Pos.X + a.Scale.X
ay2 := a.Pos.Y + a.Scale.Y


bx1 := b.Pos.X
by1 := b.Pos.Y
bx2 := b.Pos.X + b.Scale.X
by2 := b.Pos.Y + b.Scale.Y


overlapX := min(ax2, bx2) - max(ax1, bx1)
overlapY := min(ay2, by2) - max(ay1, by1)


if overlapX < overlapY {
// horizontal push
if a.Pos.X < b.Pos.X {
a.Pos.X -= overlapX
} else {
a.Pos.X += overlapX
}
a.Vel.X = 0
} else {
// vertical push
if a.Pos.Y < b.Pos.Y {
// a is above b -> land on it
a.Pos.Y -= overlapY
a.Vel.Y = 0
} else {
// a is below b -> push down
a.Pos.Y += overlapY
a.Vel.Y = 0
}
}
}


func ApplyForce(e *Entity, fx, fy float64) {
e.Vel.X += fx
e.Vel.Y += fy
}
