package components

type MineBuilding struct {
	MiningResource string
	RequireField   bool
	Field          *ResourceField
	AddPerSecond   int32
	Timer          float32
}
