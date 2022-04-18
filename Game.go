package main

import (
	. "TowerDefenseTalosEcs/engine"
	. "TowerDefenseTalosEcs/engine/render"
	"TowerDefenseTalosEcs/settings"
	. "TowerDefenseTalosEcs/systems"
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
	mainLayer.
		Add(&ChildrenSystem{}).
		Add(&TimedDestroySystem{}).
		Add(&CurveScaleSystem{}).
		Add(&ModelRenderSystem{})
}

func addGameSystems() {
	mainLayer.
		Add(&CameraControlsSystem{}).
		Add(&BuildInputSystem{}).
		Add(&MinesInitSystem{}).
		Add(&EnemySpawnSystem{}).
		Add(&TargetSearchSystem{}).
		Add(&RemoveTargetSystem{}).
		Add(&EnemySpawnSystem{}).
		Add(&TargetSearchSystem{}).
		Add(&RemoveTargetSystem{}).
		Add(&AttackSystem{}).
		Add(&MoveSystem{}).
		Add(&DeathSystem{}).
		Add(&MineSystem{}).
		Add(&EnemyMoveSystem{}).
		Add(&BulletFxSystem{}).
		Add(&BuildEffectSystem{}).
		Add(&ShootEffectSystem{}).
		Add(&GameRulesSystem{}).
		Add(&GenerateResourcesSystem{}).
		Add(&HivesSystem{})

	uiRenderLayer.
		Add(&ui.ResourcesDrawSystem{}).
		Add(&ui.HealthbarsDrawSystem{}).
		Add(&ui.BuildButtonsDrawSystem{}).
		Add(&ui.GameRulesDrawSystem{}).
		Add(&ui.DebugUiSystem{})

	postProcessLayer.
		Add(&PostProcessSystem{})
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
		rl.BeginTextureMode(RenderTexture)
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
