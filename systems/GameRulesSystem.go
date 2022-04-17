package systems

import (
	. "RtsGame/components"
	. "RtsGame/entities"
	. "RtsGame/settings"
	. "RtsGame/tags"
	ecs "github.com/OlegDzhuraev/talosecs"
)

type GameRulesSystem struct {
}

func (system *GameRulesSystem) Init() {
	NewHeadquartersBuilding()

	BuildingMakers = []BuildData{
		{
			Name:          "Mine",
			PriceResource: "iron",
			Price:         20,
			Action:        func() ecs.Entity { return NewMineBuilding() },
		},
		{
			Name:          "Turret",
			PriceResource: "iron",
			Price:         40,
			Action:        func() ecs.Entity { return NewTurretBuilding() },
		},
	}
}

func (system *GameRulesSystem) Update() {
	if GameOver || GameWin {
		return
	}

	headquarters := ecs.FilterWith[*HeadquartersTag]()

	isAlive := false
	for range headquarters {
		isAlive = true
		break
	}

	if !isAlive {
		GameOver = true

		ownedByPlayer := ecs.FilterWith[*PlayerOwnedTag]()
		for _, c := range ownedByPlayer {
			ecs.KillEntity(ecs.GetEntity(c))
		}
	}

	hives := ecs.FilterWith[*Hive]()

	isAnyHiveAlive := false
	for range hives {
		isAnyHiveAlive = true
	}

	if !isAnyHiveAlive {
		GameWin = true
	}
}
