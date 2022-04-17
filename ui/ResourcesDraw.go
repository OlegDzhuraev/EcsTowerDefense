package ui

import (
	. "TowerDefenseTalosEcs/settings"
	rl "github.com/gen2brain/raylib-go/raylib"
	"strconv"
)

type ResourcesDrawSystem struct {
}

func (system *ResourcesDrawSystem) Update() {
	iterator := 0
	for name, val := range PlayerResources {
		rl.DrawText(name+": "+strconv.Itoa(int(val)), 20, 20+int32(iterator*32), 22, rl.White)
		iterator++
	}
}
