package systems

import (
	. "TowerDefenseTalosEcs/engine"
	. "TowerDefenseTalosEcs/entities"
	"TowerDefenseTalosEcs/signals"
	. "TowerDefenseTalosEcs/tags"
	ecs "github.com/OlegDzhuraev/talosecs"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type EnemySpawnSystem struct {
	timer              Timer
	actualHive         int32
	minRespawnTime     float32
	timeReducePerSpawn float32
}

func (system *EnemySpawnSystem) Init() {
	system.timer = Timer{Time: 5}
	system.minRespawnTime = 1
	system.timeReducePerSpawn = 0.1
}

func (system *EnemySpawnSystem) Update() {
	system.timer.DoTickUntilReady()

	if system.timer.IsReady() {
		hives, transforms := ecs.FilterWith2[*HiveTag, *Transform]()

		var hivesAmount int32 = 0
		var newActualHive int32 = 0

		for i := range hives {
			if int32(i) == system.actualHive {
				enemyEnt := NewEnemy()
				tr := transforms[i]

				if enemyTr, ok := ecs.GetComponent[*Transform](enemyEnt); ok {
					enemyTr.Position = tr.Position
					ecs.TryAddSignal(&signals.SpawnFxSignal{Position: tr.Position})
				}

				newActualHive = system.actualHive + 1
			}

			hivesAmount++
		}

		if newActualHive >= hivesAmount {
			newActualHive = 0
		}

		system.actualHive = newActualHive
		system.timer.Time = rl.Clamp(system.timer.Time-system.timeReducePerSpawn, system.minRespawnTime, 999)
		system.timer.Reset()
	}
}
