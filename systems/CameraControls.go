package systems

import (
	. "TowerDefenseTalosEcs/engine"
	"TowerDefenseTalosEcs/settings"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type CameraControls struct {
	cameraSpeed float32
}

func (system *CameraControls) Init() {
	system.cameraSpeed = 10

	settings.MainCamera = rl.Camera3D{
		Up:         rl.Vector3{Y: 1},
		Position:   rl.Vector3{Y: 10, Z: 10},
		Fovy:       45,
		Projection: rl.CameraPerspective,
	}
}

func (system *CameraControls) Update() {
	dir := rl.Vector3{X: HorizontalAxis.GetValue(), Z: -VerticalAxis.GetValue()}
	camPos := settings.MainCamera.Position
	camTarget := settings.MainCamera.Target
	dirScaled := rl.Vector3Scale(dir, system.cameraSpeed*rl.GetFrameTime())
	settings.MainCamera.Position = rl.Vector3Add(camPos, dirScaled)
	settings.MainCamera.Target = rl.Vector3Add(camTarget, dirScaled)

	mouseWheel := rl.GetMouseWheelMove()
	yPos := settings.MainCamera.Position.Y

	settings.MainCamera.Position.Y = rl.Clamp(yPos-float32(mouseWheel), 5, 100)
}
