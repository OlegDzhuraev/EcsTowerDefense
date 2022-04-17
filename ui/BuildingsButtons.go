package ui

import (
	. "RtsGame/settings"
	rl "github.com/gen2brain/raylib-go/raylib"
	"strconv"
)

type BuildButtonsDrawSystem struct {
	ScreenPos rl.Vector2
}

func (system *BuildButtonsDrawSystem) Init() {
	system.ScreenPos = rl.Vector2{X: ScreenSize.X / 2, Y: ScreenSize.Y - 60}
}

func (system *BuildButtonsDrawSystem) Update() {
	offset := 110
	totalOffset := offset * len(BuildingMakers) / 2

	for i, bm := range BuildingMakers {
		color := rl.White

		if PlayerResources[bm.PriceResource] < bm.Price {
			color = rl.Red
		}

		posX := int32(system.ScreenPos.X + float32(offset*i-totalOffset))
		text := "[" + strconv.Itoa(i+1) + "] " + bm.Name
		priceText := "Price: " + strconv.Itoa(int(bm.Price))
		rl.DrawText(text, posX, int32(system.ScreenPos.Y), 16, color)
		rl.DrawText(priceText, posX, int32(system.ScreenPos.Y+20), 16, color)
	}
}
