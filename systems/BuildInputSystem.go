package systems

import (
	. "TowerDefenseTalosEcs/engine"
	. "TowerDefenseTalosEcs/oneframes"
	. "TowerDefenseTalosEcs/settings"
	. "TowerDefenseTalosEcs/signals"
	. "TowerDefenseTalosEcs/tags"
	ecs "github.com/OlegDzhuraev/talosecs"
	rl "github.com/gen2brain/raylib-go/raylib"
	"math"
)

type BuildInputSystem struct {
	buildEnabled bool
	selectedId   int

	worldCursorPos rl.Vector3

	isMoneyEnough  bool
	canBuildByDist bool
}

func (system *BuildInputSystem) Update() {
	for i := 1; i < 10; i++ {
		realId := i - 1
		if rl.IsKeyPressed(NumKeys[i]) && realId < len(BuildingMakers) {
			system.selectedId = realId
			building := BuildingMakers[system.selectedId]

			if PlayerResources[building.PriceResource] >= building.Price {
				system.buildEnabled = true
			}
		}
	}

	if system.buildEnabled && rl.IsMouseButtonPressed(rl.MouseRightButton) {
		system.buildEnabled = false
	}

	if system.buildEnabled {
		g0 := rl.Vector3{X: -100.0, Z: -100.0}
		g1 := rl.Vector3{X: -500.0, Z: 100.0}
		g2 := rl.Vector3{X: 100.0, Z: 100.0}
		g3 := rl.Vector3{X: 100.0, Z: -100.0}
		ray := rl.GetMouseRay(rl.GetMousePosition(), MainCamera)

		groundHitInfo := rl.GetRayCollisionQuad(ray, g0, g1, g2, g3)
		system.worldCursorPos = groundHitInfo.Point

		building := BuildingMakers[system.selectedId]
		system.isMoneyEnough = PlayerResources[building.PriceResource] >= building.Price

		if system.isMoneyEnough {
			playersTransforms, _ := ecs.FilterWith2[*Transform, *PlayerOwnedTag]()

			var minDist float32 = math.MaxFloat32

			for _, tr := range playersTransforms {
				dist := rl.Vector2Distance(rl.Vector2{X: system.worldCursorPos.X, Y: system.worldCursorPos.Z}, rl.Vector2{X: tr.Position.X, Y: tr.Position.Z})

				if dist < minDist {
					minDist = dist
				}
			}

			system.canBuildByDist = minDist < 7.5

			if rl.IsMouseButtonDown(rl.MouseLeftButton) && system.canBuildByDist {
				ent := BuildingMakers[system.selectedId].Action()

				ent.OneFrame(&MineInitOneFrame{Position: system.worldCursorPos})

				if tr, ok := ecs.GetComponent[*Transform](ent); ok { // todo refactor, move to a separated system
					tr.Position = system.worldCursorPos
				}

				ecs.TryAddSignal(&SpawnFxSignal{Position: system.worldCursorPos})

				PlayerResources[building.PriceResource] -= building.Price

				system.buildEnabled = false
			}
		}
	}

	system.drawBuildMode()
}

func (system BuildInputSystem) drawBuildMode() {
	if system.buildEnabled {
		color := rl.Yellow

		if !system.canBuildByDist {
			color = rl.Red
		}

		rl.DrawSphere(system.worldCursorPos, 0.75, color)
	}
}
