package systems

import (
	. "RtsGame/components"
	. "RtsGame/engine"
	. "RtsGame/engine/render"
	. "RtsGame/oneframes"
	ecs "github.com/OlegDzhuraev/talosecs"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type MinesInitSystem struct {
	maxDistToField float32
}

func (system *MinesInitSystem) Init() {
	system.maxDistToField = 1.25
}

func (system *MinesInitSystem) Update() {
	inits, mines := ecs.FilterWith2[*MineInitOneFrame, *MineBuilding]()
	noMinesRequireInit := true

	for range inits {
		noMinesRequireInit = false
		break
	}

	if noMinesRequireInit {
		return
	}

	fields, fieldsTransforms := ecs.FilterWith2[*ResourceField, *Transform]()

	for i, mine := range mines {
		if !mine.RequireField {
			continue
		}

		pos := inits[i].Position

		for k, field := range fields {
			tr := fieldsTransforms[k]

			if rl.Vector3Distance(pos, tr.Position) < system.maxDistToField {
				mine.Field = field
				break
			}
		}

		if mine.Field == nil {
			e := ecs.GetEntity(mine)
			ecs.DelComponent[*MineBuilding](e) // mark mine as wrong placed, it will stay on ground but stop work

			if mr, ok := ecs.GetComponent[*ModelRenderer](e); ok {
				mr.SetColor(rl.Gray)
			}
		}
	}
}
