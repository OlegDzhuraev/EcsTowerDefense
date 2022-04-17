package main

import (
	"TowerDefenseTalosEcs/engine"
	"TowerDefenseTalosEcs/engine/render"
	"TowerDefenseTalosEcs/settings"
	"TowerDefenseTalosEcs/systems"
	"TowerDefenseTalosEcs/ui"
	ecs "github.com/OlegDzhuraev/talosecs"
	rl "github.com/gen2brain/raylib-go/raylib"
	"math/rand"
	"time"
)

var mainLayer *ecs.Layer
var uiRenderLayer *ecs.Layer
var postProcessLayer *ecs.Layer

func makeLayers() {
	mainLayer = ecs.NewLayer()
	uiRenderLayer = ecs.NewLayer()
	postProcessLayer = ecs.NewLayer()
}

func addEngineSystems() {
	mainLayer.Add(&engine.ChildrenSystem{})
	mainLayer.Add(&engine.TimedDestroySystem{})
	mainLayer.Add(&engine.CurveScaleSystem{})

	mainLayer.Add(&render.ModelRenderSystem{})
}

func addGameSystems() {
	mainLayer.Add(&systems.CameraControls{})
	mainLayer.Add(&systems.BuildInputSystem{})
	mainLayer.Add(&systems.MinesInitSystem{})

	mainLayer.Add(&systems.EnemySpawnSystem{})
	mainLayer.Add(&systems.TargetSearchSystem{})
	mainLayer.Add(&systems.RemoveTargetSystem{})

	mainLayer.Add(&systems.AttackSystem{})
	mainLayer.Add(&systems.MoveSystem{})
	mainLayer.Add(&systems.DeathSystem{})
	mainLayer.Add(&systems.MineSystem{})
	mainLayer.Add(&systems.EnemyMoveSystem{})
	mainLayer.Add(&systems.BulletFxSystem{})

	mainLayer.Add(&systems.BuildEffectSystem{})
	mainLayer.Add(&systems.ShootEffectSystem{})
	mainLayer.Add(&systems.GameRulesSystem{})
	mainLayer.Add(&systems.GenerateResourcesSystem{})
	mainLayer.Add(&systems.HivesSystem{})

	uiRenderLayer.Add(&ui.ResourcesDrawSystem{})
	uiRenderLayer.Add(&ui.HealthbarsDrawSystem{})
	uiRenderLayer.Add(&ui.BuildButtonsDrawSystem{})
	uiRenderLayer.Add(&ui.GameRulesDrawSystem{})
	uiRenderLayer.Add(&ui.DebugUiSystem{})

	postProcessLayer.Add(&systems.PostProcessSystem{})
}

func gameInit() {
	rl.InitWindow(int32(settings.ScreenSize.X), int32(settings.ScreenSize.Y), "Tower defense Talos ECS")
	rl.SetTargetFPS(2000)
	rand.Seed(time.Now().UnixNano())

	ecs.AddLayer(mainLayer)

	uiRenderLayer.Init()
	postProcessLayer.Init()

	ecs.Init()
}

func main() {
	makeLayers()
	addEngineSystems()
	addGameSystems()
	gameInit()

	for !rl.WindowShouldClose() {
		rl.BeginTextureMode(systems.RenderTexture)
		rl.ClearBackground(settings.ClearColor)
		rl.BeginMode3D(settings.MainCamera)
		ecs.Update()
		rl.EndMode3D()
		rl.EndTextureMode()

		rl.BeginDrawing()
		postProcessLayer.Update()
		uiRenderLayer.Update()
		rl.EndDrawing()
	}

	rl.CloseWindow()
}
