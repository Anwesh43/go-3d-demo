package main

import (
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	var w int32 = 800
	var h int32 = 450
	rl.InitWindow(w, h, "raylib core example")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)
	var x int32 = 0
	var y int32 = h / 2
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		rl.DrawText("First Text", w/2, 0, 20, rl.LightGray)
		rl.DrawRectangle(x, y, int32(50), int32(50), rl.Green)
		rl.EndDrawing()

		x += (int32(w) / 10)
		if x > int32(w) {
			x = 0
		}
		time.Sleep(100 * time.Millisecond)
	}
}
