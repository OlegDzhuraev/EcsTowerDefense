package components

import ecs "github.com/OlegDzhuraev/talosecs"

type TargetSearcher struct {
	Target      ecs.Entity
	MaxDistance float32
}
