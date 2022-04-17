package components

type Damageable struct {
	MaxHealth float32
	Health    float32
}

func (d Damageable) IsAlive() bool { return d.Health > 0 }
