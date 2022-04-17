package components

type Attack struct {
	Damage     float32
	ReloadTime float32
	Distance   float32

	ReloadTimeLeft float32
	IsLoaded       bool
}
