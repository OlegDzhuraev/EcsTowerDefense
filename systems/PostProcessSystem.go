package systems

import (
	"TowerDefenseTalosEcs/settings"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var RenderTexture rl.RenderTexture2D

type PostProcessSystem struct {
	shader   rl.Shader
	amountId int32
}

func (system *PostProcessSystem) Init() {
	system.shader = rl.LoadShader("0", "resources/postproc.fs")

	system.amountId = rl.GetShaderLocation(system.shader, "amount")
	renderWidthId := rl.GetShaderLocation(system.shader, "renderWidth")
	renderHeightId := rl.GetShaderLocation(system.shader, "renderHeight")

	rl.SetShaderValue(system.shader, renderWidthId, []float32{settings.ScreenSize.X}, rl.ShaderUniformFloat)
	rl.SetShaderValue(system.shader, renderHeightId, []float32{settings.ScreenSize.Y}, rl.ShaderUniformFloat)

	RenderTexture = rl.LoadRenderTexture(int32(settings.ScreenSize.X), int32(settings.ScreenSize.Y))
}

func (system *PostProcessSystem) Update() {
	rl.SetShaderValue(system.shader, system.amountId, []float32{rl.GetTime()}, rl.ShaderUniformFloat)
	rl.BeginShaderMode(system.shader)

	// NOTE: Render texture must be y-flipped due to default OpenGL coordinates (left-bottom)
	rl.DrawTextureRec(
		RenderTexture.Texture,
		rl.Rectangle{Width: float32(RenderTexture.Texture.Width), Height: float32(-RenderTexture.Texture.Height)},
		rl.Vector2{},
		rl.White)

	rl.EndShaderMode()
}
