package ui

import (
	. "TowerDefenseTalosEcs/components"
	. "TowerDefenseTalosEcs/engine"
	"TowerDefenseTalosEcs/settings"
	"TowerDefenseTalosEcs/tags"
	. "github.com/OlegDzhuraev/talosecs"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type HealthbarsDrawSystem struct {
	width    int32
	height   int32
	upOffset rl.Vector3
}

func (system *HealthbarsDrawSystem) Init() {
	system.width = 64
	system.height = 10
	system.upOffset = rl.Vector3{Y: 1.75}
}

func (system *HealthbarsDrawSystem) Update() {
	damageables, transforms := FilterWith2[*Damageable, *Transform]()

	for i, damageable := range damageables {
		tr := transforms[i]

		color := rl.Red
		if _, ok := GetComponent[*tags.PlayerOwnedTag](GetEntity(damageable)); ok {
			color = rl.Green
		}

		percents := damageable.Health / damageable.MaxHealth

		if percents >= 1 {
			continue
		}

		screenPos := rl.GetWorldToScreen(rl.Vector3Add(tr.Position, system.upOffset), settings.MainCamera)

		drawHealthbar(system, screenPos, percents, color)
	}
}

func drawHealthbar(healthbars *HealthbarsDrawSystem, screenPos rl.Vector2, percents float32, color rl.Color) {
	var posX = int32(screenPos.X) - healthbars.width/2
	var posY = int32(screenPos.Y)
	var offset int32 = 2
	var ofsettedWidth = healthbars.width - offset*2
	var ofsettedHeight = healthbars.height - offset*2

	rl.DrawRectangle(posX, posY, healthbars.width, healthbars.height, rl.White)
	rl.DrawRectangle(posX+offset, posY+offset, ofsettedWidth, ofsettedHeight, rl.Black)
	rl.DrawRectangle(posX+offset, posY+offset, int32(float32(ofsettedWidth)*percents), ofsettedHeight, color)
}
