package systems

import (
	. "TowerDefenseTalosEcs/components"
	. "TowerDefenseTalosEcs/engine"
	. "TowerDefenseTalosEcs/entities"
	. "TowerDefenseTalosEcs/oneframes"
	ecs "github.com/OlegDzhuraev/talosecs"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type AttackSystem struct {
}

func (system *AttackSystem) Update() {
	transforms, attacks, targetSearchers := ecs.FilterWith3[*Transform, *Attack, *TargetSearcher]()

	for i, attack := range attacks {
		if !attack.IsLoaded {
			attack.ReloadTimeLeft = attack.ReloadTime
			attack.IsLoaded = true
			continue
		}

		targetSearcher := targetSearchers[i]
		target := targetSearcher.Target
		transform := transforms[i]

		if attack.ReloadTimeLeft > 0 {
			attack.ReloadTimeLeft -= rl.GetFrameTime()
		} else if ecs.IsAlive(target) {
			targetPos := rl.Vector3{}
			if targetTransform, ok := ecs.GetComponent[*Transform](target); ok { // todo to attack separated system
				dist := rl.Vector3Distance(transform.Position, targetTransform.Position)

				targetPos = targetTransform.Position
				if dist > attack.Distance {
					continue
				}
			}

			if damageable, ok := ecs.GetComponent[*Damageable](target); ok {
				damageable.Health = rl.Clamp(damageable.Health-attack.Damage, 0, damageable.MaxHealth)
				attack.ReloadTimeLeft = attack.ReloadTime

				NewBulletEffect(transform.Position, targetPos)
				ecs.GetEntity(attack).OneFrame(&ShootOneFrame{Position: transform.Position})
			}
		}
	}
}
