package ui

import (
	"TowerDefenseTalosEcs/settings"
	rl "github.com/gen2brain/raylib-go/raylib"
	"strconv"
)

type GameRulesDrawSystem struct {
	centerX  int32
	centerY  int32
	fontSize int32
}

func (system *GameRulesDrawSystem) Init() {
	system.centerX = int32(settings.ScreenSize.X / 2)
	system.centerY = int32(settings.ScreenSize.Y / 2)
	system.fontSize = 50
}

func (system *GameRulesDrawSystem) Update() {
	isAnyState := settings.GameOver || settings.GameWin

	if isAnyState {
		posX := system.centerX - 128
		posY := system.centerY - 64

		rl.DrawRectangle(0, 0, int32(settings.ScreenSize.X), int32(settings.ScreenSize.Y), rl.Color{A: 192})

		text := "Game over"
		color := rl.Red

		if settings.GameWin {
			text = "Game win"
			color = rl.Green
		}

		rl.DrawText(text, posX, posY, system.fontSize, color)
	}

	rl.DrawText("Score: "+strconv.Itoa(int(settings.Score)), system.centerX-64, 20, 24, rl.White)
}
