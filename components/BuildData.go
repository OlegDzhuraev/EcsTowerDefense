package components

import ecs "github.com/OlegDzhuraev/talosecs"

type BuildData struct {
	Name          string
	PriceResource string
	Price         int32
	Action        func() ecs.Entity
}
