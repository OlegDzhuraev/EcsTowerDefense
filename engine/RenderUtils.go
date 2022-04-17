package engine

import rl "github.com/gen2brain/raylib-go/raylib"

// todo maybe some same models pool and mesh instancing?

func MakeSphereModel() rl.Model {
	mesh := rl.GenMeshSphere(1, 12, 12)
	model := rl.LoadModelFromMesh(mesh)

	return model
}

func MakeCubeModel() rl.Model {
	mesh := rl.GenMeshCube(1, 1, 1)
	model := rl.LoadModelFromMesh(mesh)

	return model
}

func Lerp(a, b rl.Color, val float32) rl.Color {
	return rl.Color{
		R: LerpU8(a.R, b.R, val),
		G: LerpU8(a.G, b.G, val),
		B: LerpU8(a.B, b.B, val),
		A: LerpU8(a.A, b.A, val)}
}
