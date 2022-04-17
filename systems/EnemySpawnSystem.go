package systems

import (
	"TowerDefenseTalosEcs/components"
	. "TowerDefenseTalosEcs/engine"
	. "TowerDefenseTalosEcs/entities"
	"TowerDefenseTalosEcs/signals"
	ecs "github.com/OlegDzhuraev/talosecs"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type EnemySpawnSystem struct {
	Timer              Timer
	ActualHive         int32
	minRespawnTime     float32
	timeReducePerSpawn float32
}

func (system *EnemySpawnSystem) Init() {
	system.Timer = Timer{Time: 5}
	system.minRespawnTime = 1
	system.timeReducePerSpawn = 0.1
}

func (system *EnemySpawnSystem) Update() {
	system.Timer.DoTickUntilReady()

	if system.Timer.IsReady() {
		hives, transforms := ecs.FilterWith2[*components.Hive, *Transform]()

		var hivesAmount int32 = 0
		var newActualHive int32 = 0

		for i := range hives {
			if int32(i) == system.ActualHive {
				enemyEnt := NewEnemy()
				tr := transforms[i]

				if enemyTr, ok := ecs.GetComponent[*Transform](enemyEnt); ok {
					enemyTr.Position = tr.Position
					ecs.TryAddSignal(&signals.SpawnFxSignal{Position: tr.Position})
				}

				newActualHive = system.ActualHive + 1
			}

			hivesAmount++
		}

		if newActualHive >= hivesAmount {
			newActualHive = 0
		}

		system.ActualHive = newActualHive
		system.Timer.Time = rl.Clamp(system.Timer.Time-system.timeReducePerSpawn, system.minRespawnTime, 999)
		system.Timer.Reset()
	}
}
