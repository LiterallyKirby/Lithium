package main

import (
	"fmt"
	lithium "Lithium/src/lithium"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	// Set window size
	lithium.SetWindowSize(800, 600)

	// Set background color
	lithium.SetBackgroundColor(lithium.Black)

	// Student-friendly example: a player square and platforms
	player := lithium.AddEntity("Player", lithium.Rect, lithium.Vec2{X: 100, Y: 100}, true)
	lithium.Scale(player, 32, 32)
	lithium.SetColor(player, lithium.Cyan)

	// Ground platform
	ground := lithium.AddEntity("Ground", lithium.Rect, lithium.Vec2{X: 0, Y: 550}, false)
	lithium.Scale(ground, 800, 50)
	lithium.SetColor(ground, lithium.Green)

	// Floating platforms
	platform1 := lithium.AddEntity("Platform1", lithium.Rect, lithium.Vec2{X: 200, Y: 400}, false)
	lithium.Scale(platform1, 150, 20)
	lithium.SetColor(platform1, lithium.Yellow)

	platform2 := lithium.AddEntity("Platform2", lithium.Rect, lithium.Vec2{X: 450, Y: 300}, false)
	lithium.Scale(platform2, 150, 20)
	lithium.SetColor(platform2, lithium.Orange)

	// Moving platform
	movingPlatform := lithium.AddEntity("MovingPlatform", lithium.Rect, lithium.Vec2{X: 600, Y: 450}, false)
	lithium.Scale(movingPlatform, 120, 20)
	lithium.SetColor(movingPlatform, lithium.Magenta)
	movingDir := 1.0

	// Collectible coin
	coin := lithium.AddEntity("Coin", lithium.Circle, lithium.Vec2{X: 300, Y: 200}, false)
	lithium.Scale(coin, 20, 20)
	lithium.SetColor(coin, lithium.Yellow)
	coin.Solid = false

	score := 0
	coinCollected := false

	// Enemy
	enemy := lithium.AddEntity("Enemy", lithium.Rect, lithium.Vec2{X: 500, Y: 500}, false)
	lithium.Scale(enemy, 30, 30)
	lithium.SetColor(enemy, lithium.Red)
	enemy.Solid = false
	enemy.Vel.X = 2

	// Collision handlers
	player.OnCollide = func(other *lithium.Entity) {
		if other.Name == "Coin" && !coinCollected {
			score += 10
			coinCollected = true
			lithium.RemoveEntity("Coin")
		}
		if other.Name == "Enemy" {
			// Reset player position
			lithium.Teleport(player, 100, 100)
			lithium.SetVelocity(player, 0, 0)
			score = 0
		}
	}

	lithium.OnUpdate(func() {
		// Player movement
		if lithium.IsPressed(lithium.KeyRight) || lithium.IsPressed(lithium.KeyD) {
			lithium.Move(player, 3, 0)
		}
		if lithium.IsPressed(lithium.KeyLeft) || lithium.IsPressed(lithium.KeyA) {
			lithium.Move(player, -3, 0)
		}
		
		// Jump only when grounded
		if (lithium.JustPressed(lithium.KeySpace) || lithium.JustPressed(lithium.KeyW)) && lithium.IsGrounded(player) {
			lithium.ApplyForce(player, 0, -12)
		}

		// Moving platform logic
		movingPlatform.Pos.X += movingDir
		if movingPlatform.Pos.X > 700 || movingPlatform.Pos.X < 500 {
			movingDir = -movingDir
		}

		// Enemy AI - simple patrol
		if enemy.Pos.X > 700 {
			enemy.Vel.X = -2
		} else if enemy.Pos.X < 400 {
			enemy.Vel.X = 2
		}
		enemy.Pos.X += enemy.Vel.X

		// Rotate coin for visual effect
		if !coinCollected {
			c := lithium.GetEntity("Coin")
			if c != nil {
				lithium.Rotate(c, 0.05)
			}
		}

		// Keep player in bounds
		if player.Pos.Y > 600 {
			lithium.Teleport(player, 100, 100)
			lithium.SetVelocity(player, 0, 0)
		}

		// Mouse interaction - click to spawn a platform
		if lithium.JustClickedMouse(lithium.MouseButtonLeft) {
			mousePos := lithium.GetMousePositionF()
			platformName := fmt.Sprintf("CustomPlatform_%d", lithium.GetFrameCount())
			customPlatform := lithium.AddEntity(platformName, lithium.Rect, mousePos, false)
			lithium.Scale(customPlatform, 80, 15)
			lithium.SetColor(customPlatform, lithium.Blue)
		}
	})

	// Custom drawing
	lithium.OnDraw(func(screen *ebiten.Image) {
		// Draw score
		lithium.DrawText(screen, fmt.Sprintf("Score: %d", score), 10, 10)
		
		// Draw controls
		lithium.DrawText(screen, "Arrow Keys/WASD: Move", 10, 30)
		lithium.DrawText(screen, "Space/W: Jump", 10, 50)
		lithium.DrawText(screen, "Left Click: Place Platform", 10, 70)
		lithium.DrawText(screen, "F: Debug Info", 10, 90)
		
		// Draw debug info
		if lithium.IsPressed(lithium.KeyF) {
			lithium.DrawDebugInfo(screen)
			
			// Draw raycast visualization
			center := lithium.GetCenter(player)
			mousePos := lithium.GetMousePositionF()
			hit, hitEntity := lithium.Raycast(center, mousePos, 50)
			
			if hit && hitEntity != nil {
				lithium.DrawLine(screen, center.X, center.Y, mousePos.X, mousePos.Y, lithium.Red)
				lithium.DrawText(screen, fmt.Sprintf("Hit: %s", hitEntity.Name), 10, 130)
			} else {
				lithium.DrawLine(screen, center.X, center.Y, mousePos.X, mousePos.Y, lithium.Green)
			}
		}
	})

	lithium.Run()
}
