package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	s int32 = 20
	x int32 = 10
	y int32 = 20
)

func main() {
	rl.InitWindow(800, 450, "raylib [core] example - basic window")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	vale := rl.LoadTexture("imgs/val.png")

	for !rl.WindowShouldClose() {

		inputs()

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		rl.DrawRectangle(x, y, 50, 50, rl.Green)

		rl.DrawTexture(vale, x, y, rl.White)

		rl.EndDrawing()
	}
}

func inputs() {
	if rl.IsKeyPressed(rl.KeyD) {
		x += s
	}
	if rl.IsKeyPressed(rl.KeyA) {
		x -= s
	}
	if rl.IsKeyPressed(rl.KeyW) {
		y -= s
	}
	if rl.IsKeyPressed(rl.KeyS) {
		y += s
	}
}
