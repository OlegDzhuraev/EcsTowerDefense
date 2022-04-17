package systems

import (
	. "TowerDefenseTalosEcs/components"
	. "TowerDefenseTalosEcs/engine"
	. "TowerDefenseTalosEcs/tags"
	ecs "github.com/OlegDzhuraev/talosecs"
	rl "github.com/gen2brain/raylib-go/raylib"
	"math"
)

type TargetSearchSystem struct {
}

func (system *TargetSearchSystem) Update() {
	searchers, searchersTeams, searchersTransforms := ecs.FilterWith3[*TargetSearcher, *Team, *Transform]()
	possibleTargets, possibleTargetsTeams, possibleTargetsTransforms := ecs.FilterWith3[*TargetTag, *Team, *Transform]()

	for i, s := range searchers {
		searcherTr := searchersTransforms[i]

		if ecs.IsAlive(s.Target) {
			continue
		}

		searcherTeam := searchersTeams[i].Id

		var trMap = map[*Transform]ecs.Entity{}

		for i2, t := range possibleTargets {
			tTeam := possibleTargetsTeams[i2].Id
			tTr := possibleTargetsTransforms[i2]

			targetEnt := ecs.GetEntity(t)

			if ecs.IsAlive(targetEnt) && searcherTeam != tTeam {
				trMap[tTr] = targetEnt
			}
		}

		s.Target = GetNearest(searcherTr.Position, trMap)
	}
}

func GetNearest(fromPos rl.Vector3, targets map[*Transform]ecs.Entity) ecs.Entity {
	var minDist float32 = math.MaxFloat32
	var nearestEnt ecs.Entity

	for tr, ent := range targets {
		dist := rl.Vector3Distance(fromPos, tr.Position)
		if dist < minDist {
			minDist = dist
			nearestEnt = ent
		}
	}

	return nearestEnt
}
