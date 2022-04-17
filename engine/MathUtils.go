package engine

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func GetDirection3d(from rl.Vector3, to rl.Vector3) rl.Vector3 {
	return rl.Vector3Normalize(rl.Vector3{X: to.X - from.X, Y: to.Y - from.Y, Z: to.Z - from.Z})
}

func LerpFloat(a, b, val float32) float32      { return a + val*(b-a) }
func InverseLerpFloat(a, b, v float32) float32 { return (v - a) / (b - a) }
func LerpU8(a, b uint8, val float32) uint8     { return uint8(LerpFloat(float32(a), float32(b), val)) }
