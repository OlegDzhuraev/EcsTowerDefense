package main

import (
	"RtsGame/engine"
	"RtsGame/engine/render"
	"RtsGame/settings"
	"RtsGame/systems"
	"RtsGame/ui"
	ecs "github.com/OlegDzhuraev/talosecs"
	rl "github.com/gen2brain/raylib-go/raylib"
	"math/rand"
	"time"
)

var mainLayer *ecs.Layer
var uiRenderLayer *ecs.Layer
var postProcessLayer *ecs.Layer

func addEngineSystems() {
	mainLayer.Add(&engine.ChildrenSystem{})
	mainLayer.Add(&engine.TimedDestroySystem{})
	mainLayer.Add(&engine.CurveScaleSystem{})

	mainLayer.Add(&render.ModelRenderSystem{})
}

func mainInit() {
	rl.InitWindow(int32(settings.ScreenSize.X), int32(settings.ScreenSize.Y), "They are debiches")
	rl.SetTargetFPS(2000)
	rand.Seed(time.Now().UnixNano())
}

func main() {
	mainLayer = ecs.NewLayer()
	uiRenderLayer = ecs.NewLayer()
	postProcessLayer = ecs.NewLayer()

	addEngineSystems()

	mainLayer.Add(&systems.CameraControls{CameraSpeed: 10})
	mainLayer.Add(&systems.BuildInput{})
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

	postProcessLayer.Add(&systems.PostProcessSystem{})

	mainInit()

	ecs.AddLayer(mainLayer)

	uiRenderLayer.Init()
	postProcessLayer.Init()
	ecs.Init()

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
		rl.DrawFPS(10, 128)
		rl.EndDrawing()
	}

	rl.CloseWindow()
}
