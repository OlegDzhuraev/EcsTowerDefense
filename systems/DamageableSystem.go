package systems

import (
	. "TowerDefenseTalosEcs/components"
	"TowerDefenseTalosEcs/settings"
	. "github.com/OlegDzhuraev/talosecs"
)

type DeathSystem struct {
}

func (system *DeathSystem) Update() {
	damageables := FilterWith[*Damageable]()

	for _, damageable := range damageables {
		if damageable.Health <= 0 {
			e := GetEntity(damageable)

			if sc, ok := GetComponent[*ScoreForDestroy](e); ok {
				settings.Score += uint(sc.Score)
			}

			KillEntity(e)
		}
	}
}
