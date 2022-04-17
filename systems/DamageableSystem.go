package systems

import (
	. "TowerDefenseTalosEcs/components"
	"TowerDefenseTalosEcs/settings"
	ecs "github.com/OlegDzhuraev/talosecs"
)

type DeathSystem struct {
}

func (system *DeathSystem) Update() {
	damageables := ecs.FilterWith[*Damageable]()

	for _, damageable := range damageables {
		if damageable.Health <= 0 {
			e := ecs.GetEntity(damageable)

			if sc, ok := ecs.GetComponent[*ScoreForDestroy](e); ok {
				settings.Score += uint(sc.Score)
			}

			ecs.KillEntity(e)
		}
	}
}
