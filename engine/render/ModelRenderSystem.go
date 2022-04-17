package render

import (
	"RtsGame/engine"
	"github.com/OlegDzhuraev/talosecs"
)

type ModelRenderSystem struct {
}

func (mrs *ModelRenderSystem) Update() {
	models, transformsModels := talosecs.FilterWith2[*ModelRenderer, *engine.Transform]()

	for i, model := range models {
		model.Render(transformsModels[i])
	}
}
