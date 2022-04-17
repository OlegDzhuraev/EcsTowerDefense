package main

import (
	"TowerDefenseTalosEcs/entities"
	"TowerDefenseTalosEcs/systems"
	"github.com/OlegDzhuraev/talosecs"
	"testing"
)

func BenchmarkSample(b *testing.B) {
	for i := 0; i < b.N; i++ {
		talosecs.AddSystem(&systems.TargetSearchSystem{})
		talosecs.AddSystem(&systems.MoveSystem{})
		talosecs.AddSystem(&systems.AttackSystem{})
		talosecs.AddSystem(&systems.EnemyMoveSystem{})
		entities.NewEnemy()
		talosecs.Init()
		talosecs.Update()
	}
}
