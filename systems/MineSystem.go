package systems

import (
	. "TowerDefenseTalosEcs/components"
	. "TowerDefenseTalosEcs/engine"
	. "TowerDefenseTalosEcs/engine/render"
	. "TowerDefenseTalosEcs/entities"
	"TowerDefenseTalosEcs/settings"
	ecs "github.com/OlegDzhuraev/talosecs"
	rl "github.com/gen2brain/raylib-go/raylib"
	"math"
)

type MineSystem struct {
	maxDistToField float32
}

func (system *MineSystem) Init() {
	system.maxDistToField = 1.25
}

func (system *MineSystem) Update() {
	mines := ecs.FilterWith[*MineBuilding]()
	dTime := rl.GetFrameTime()

	for _, mine := range mines {
		mine.Timer -= dTime

		if mine.Timer <= 0 {
			addValue := mine.AddPerSecond
			mine.Timer = 1

			if mine.RequireField {
				addValue = int32(math.Min(float64(mine.Field.Amount), float64(mine.AddPerSecond)))
				mine.Field.Amount -= addValue

				if fieldTr, ok := ecs.GetComponent[*Transform](ecs.GetEntity(mine.Field)); ok {
					scale := GetResourceFieldScale(mine.Field.Amount)
					fieldTr.Scale = rl.Vector3{X: scale, Y: fieldTr.Scale.Y, Z: scale}
				}

				if mine.Field.Amount <= 0 {
					e := ecs.GetEntity(mine)
					if mr, ok := ecs.GetComponent[*ModelRenderer](e); ok {
						mr.SetColor(rl.Gray)
					}
					ecs.DelComponent[*MineBuilding](e)
				}
			}

			settings.PlayerResources[mine.MiningResource] += addValue
		}
	}
}
