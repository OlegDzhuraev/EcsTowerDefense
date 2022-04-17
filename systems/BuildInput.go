package systems

import (
	. "TowerDefenseTalosEcs/engine"
	"TowerDefenseTalosEcs/oneframes"
	. "TowerDefenseTalosEcs/settings"
	"TowerDefenseTalosEcs/signals"
	"TowerDefenseTalosEcs/tags"
	ecs "github.com/OlegDzhuraev/talosecs"
	rl "github.com/gen2brain/raylib-go/raylib"
	"math"
)

type BuildInput struct {
	BuildEnabled   bool
	isMoneyEnough  bool
	canBuildByDist bool
	WorldCursorPos rl.Vector3

	SelectedId int
}

func (system *BuildInput) Update() {
	for i := 1; i < 10; i++ {
		realId := i - 1
		if rl.IsKeyPressed(NumKeys[i]) && realId < len(BuildingMakers) {
			system.SelectedId = realId
			building := BuildingMakers[system.SelectedId]

			if PlayerResources[building.PriceResource] >= building.Price {
				system.BuildEnabled = true
			}
		}
	}

	if system.BuildEnabled && rl.IsMouseButtonPressed(rl.MouseRightButton) {
		system.BuildEnabled = false
	}

	if system.BuildEnabled {
		g0 := rl.Vector3{X: -100.0, Z: -100.0}
		g1 := rl.Vector3{X: -500.0, Z: 100.0}
		g2 := rl.Vector3{X: 100.0, Z: 100.0}
		g3 := rl.Vector3{X: 100.0, Z: -100.0}
		ray := rl.GetMouseRay(rl.GetMousePosition(), MainCamera)

		groundHitInfo := rl.GetRayCollisionQuad(ray, g0, g1, g2, g3)
		system.WorldCursorPos = groundHitInfo.Point

		building := BuildingMakers[system.SelectedId]
		system.isMoneyEnough = PlayerResources[building.PriceResource] >= building.Price

		if system.isMoneyEnough {
			playersTransforms, _ := ecs.FilterWith2[*Transform, *tags.PlayerOwnedTag]()

			var minDist float32 = math.MaxFloat32

			for _, tr := range playersTransforms {
				dist := rl.Vector2Distance(rl.Vector2{X: system.WorldCursorPos.X, Y: system.WorldCursorPos.Z}, rl.Vector2{X: tr.Position.X, Y: tr.Position.Z})

				if dist < minDist {
					minDist = dist
				}
			}

			system.canBuildByDist = minDist < 7.5

			if rl.IsMouseButtonDown(rl.MouseLeftButton) && system.canBuildByDist {
				ent := BuildingMakers[system.SelectedId].Action()

				// todo сделать систему, которая по SpawnFxSignal будет проставлять позицию?
				// можно этот сигнал вешать на сам объект и фильтровать по нему. Тогда не нужен GetComponent

				ent.OneFrame(&oneframes.MineInitOneFrame{Position: system.WorldCursorPos})

				if tr, ok := ecs.GetComponent[*Transform](ent); ok {
					tr.Position = system.WorldCursorPos
				}

				ecs.TryAddSignal(&signals.SpawnFxSignal{Position: system.WorldCursorPos})

				PlayerResources[building.PriceResource] -= building.Price

				system.BuildEnabled = false
			}
		}
	}

	system.drawBuildMode()
}

func (system BuildInput) drawBuildMode() {
	if system.BuildEnabled {
		color := rl.Yellow

		if !system.canBuildByDist {
			color = rl.Red
		}

		rl.DrawSphere(system.WorldCursorPos, 0.75, color)
	}
}
