//go:build !test

package main

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	windowWidth  = 800
	windowHeight = 600

	// Texture radius for the baked circle sprite.
	circleTextureRadius = 256

	ballChaseSpeed  float32 = 15.0
	smearChaseSpeed float32 = ballChaseSpeed * 0.45

	radToDeg = float32(180.0 / math.Pi)
)

// makeBallTexture creates a white filled circle texture.
// Generated on CPU via ImageDrawCircle to avoid render texture coordinate flips.
func makeBallTexture() rl.Texture2D {
	diameter := circleTextureRadius * 2
	img := rl.GenImageColor(diameter, diameter, rl.Blank)
	rl.ImageDrawCircle(img, int32(circleTextureRadius), int32(circleTextureRadius), int32(circleTextureRadius), rl.White)
	tex := rl.LoadTextureFromImage(img)
	rl.UnloadImage(img)
	return tex
}

func main() {
	rl.InitWindow(windowWidth, windowHeight, "blob")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	ballTexture := makeBallTexture()
	defer rl.UnloadTexture(ballTexture)

	targetX := float32(windowWidth / 2)
	targetY := float32(windowHeight / 2)

	ballX, ballY := targetX, targetY
	smearX, smearY := targetX, targetY

	for !rl.WindowShouldClose() {
		// Left click or drag to steer target
		if rl.IsMouseButtonDown(rl.MouseButtonLeft) {
			mouse := rl.GetMousePosition()
			targetX = mouse.X
			targetY = mouse.Y
		}

		dt := rl.GetFrameTime()

		// Head chases cursor
		ballX = moveToward(ballX, targetX, dt, ballChaseSpeed)
		ballY = moveToward(ballY, targetY, dt, ballChaseSpeed)

		// Tail chases head with lag, creating the smear
		smearX = moveToward(smearX, ballX, dt, smearChaseSpeed)
		smearY = moveToward(smearY, ballY, dt, smearChaseSpeed)

		rl.BeginDrawing()
		rl.ClearBackground(rl.Color{R: 0x18, G: 0x18, B: 0x18, A: 0xFF})

		texW := float32(ballTexture.Width)
		texH := float32(ballTexture.Height)
		dispW := texW * 0.25
		dispH := texH * 0.25

		// Stretched dome tail connecting ball to smear
		tailLen := distance(ballX, ballY, smearX, smearY)
		if tailLen > 1 {
			src := rl.Rectangle{
				X:      0,
				Y:      0,
				Width:  texW,
				Height: texH * 0.5,
			}
			dst := rl.Rectangle{
				X:      ballX,
				Y:      ballY,
				Width:  dispW,
				Height: tailLen,
			}
			origin := rl.Vector2{X: dispW * 0.5, Y: tailLen}
			angle := float32(math.Atan2(float64(smearY-ballY), float64(smearX-ballX)))*radToDeg + 90.0

			rl.DrawTexturePro(ballTexture, src, dst, origin, angle, rl.Green)
		}

		// Ball head
		src := rl.Rectangle{X: 0, Y: 0, Width: texW, Height: texH}
		dst := rl.Rectangle{
			X:      ballX - dispW*0.5,
			Y:      ballY - dispH*0.5,
			Width:  dispW,
			Height: dispH,
		}
		rl.DrawTexturePro(ballTexture, src, dst, rl.Vector2{}, 0, rl.Green)

		rl.EndDrawing()
	}
}
