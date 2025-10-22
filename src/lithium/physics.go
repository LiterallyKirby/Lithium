package lithium



// PhysicsSettings holds physics configuration
type PhysicsSettings struct {
	RaycastSteps     int
	MaxVelocity      float64
	Friction         float64
	CollisionPadding float64
}

var DefaultPhysics = PhysicsSettings{
	RaycastSteps:     10,
	MaxVelocity:      20,
	Friction:         0.9,
	CollisionPadding: 0.1,
}

func checkCollisions() {
	// Raycast-based collision detection
	for _, a := range world.Entities {
		if !a.Visible || !a.Solid {
			continue
		}

		// Check if entity is moving
		if a.Vel.X != 0 || a.Vel.Y != 0 {
			// Perform raycast collision detection
			raycastCollision(a)
		}

		// Check overlaps for collision callbacks
		for _, b := range world.Entities {
			if a == b || !b.Visible || !b.Solid {
				continue
			}
			if aabbOverlap(a, b) {
				if a.OnCollide != nil {
					a.OnCollide(b)
				}
			}
		}
	}
}

func raycastCollision(e *Entity) {
	if e.Vel.X == 0 && e.Vel.Y == 0 {
		return
	}

	// Cast rays from entity corners to detect collisions before they happen
	corners := []Vec2{
		{e.Pos.X, e.Pos.Y},                             // top-left
		{e.Pos.X + e.Scale.X, e.Pos.Y},                 // top-right
		{e.Pos.X, e.Pos.Y + e.Scale.Y},                 // bottom-left
		{e.Pos.X + e.Scale.X, e.Pos.Y + e.Scale.Y},     // bottom-right
		{e.Pos.X + e.Scale.X/2, e.Pos.Y + e.Scale.Y/2}, // center
	}

	for _, corner := range corners {
		targetX := corner.X + e.Vel.X
		targetY := corner.Y + e.Vel.Y

		// Cast ray in direction of movement
		if hit, entity := RaycastIgnore(corner, Vec2{targetX, targetY}, DefaultPhysics.RaycastSteps, e); hit {
			resolveRaycastCollision(e, entity)
		}
	}
}

func resolveRaycastCollision(a, b *Entity) {
	if !b.Solid {
		return
	}

	// Get bounding boxes
	ax1, ay1, ax2, ay2 := GetBounds(a)
	bx1, by1, bx2, by2 := GetBounds(b)

	// Calculate overlap
	overlapX := min(ax2, bx2) - max(ax1, bx1)
	overlapY := min(ay2, by2) - max(ay1, by1)

	// Resolve collision based on smallest overlap
	if overlapX < overlapY {
		// Horizontal collision
		if a.Pos.X < b.Pos.X {
			a.Pos.X -= overlapX + DefaultPhysics.CollisionPadding
		} else {
			a.Pos.X += overlapX + DefaultPhysics.CollisionPadding
		}
		a.Vel.X = 0
	} else {
		// Vertical collision
		if a.Pos.Y < b.Pos.Y {
			// Landing on top
			a.Pos.Y -= overlapY + DefaultPhysics.CollisionPadding
			a.Vel.Y = 0
		} else {
			// Hit from below
			a.Pos.Y += overlapY + DefaultPhysics.CollisionPadding
			if a.Vel.Y < 0 {
				a.Vel.Y = 0
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

func ApplyForce(e *Entity, fx, fy float64) {
	e.Vel.X += fx
	e.Vel.Y += fy
}

func ApplyImpulse(e *Entity, direction Vec2, force float64) {
	normalized := direction.Normalize()
	e.Vel.X += normalized.X * force
	e.Vel.Y += normalized.Y * force
}

func SetGravity(g float64) {
	world.mu.Lock()
	defer world.mu.Unlock()
	world.Gravity = g
}

func GetGravity() float64 {
	world.mu.RLock()
	defer world.mu.RUnlock()
	return world.Gravity
}

// IsGrounded checks if entity is standing on something
func IsGrounded(e *Entity) bool {
	// Cast ray slightly below the entity
	start := Vec2{e.Pos.X + e.Scale.X/2, e.Pos.Y + e.Scale.Y}
	end := Vec2{start.X, start.Y + 2}
	hit, other := RaycastIgnore(start, end, 5, e)
	return hit && other != nil && other.Solid
}

// ClampVelocity limits the velocity to a maximum value
func ClampVelocity(e *Entity, maxVel float64) {
	vel := Vec2{e.Vel.X, e.Vel.Y}
	if vel.Length() > maxVel {
		vel = vel.Normalize().Scale(maxVel)
		e.Vel = vel
	}
}
