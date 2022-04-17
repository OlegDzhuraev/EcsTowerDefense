package engine

import (
	"sort"
)

type FloatCurve struct {
	Keys map[float32]float32
}

func (fc *FloatCurve) Evaluate(time float32) float32 {
	var keys []float32

	for key := range fc.Keys {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

	for i, key := range keys {
		if key > time {
			if i == 0 {
				return fc.Keys[key]
			} else {
				prevKey := keys[i-1]
				lerpFor := InverseLerpFloat(prevKey, key, time)

				return LerpFloat(fc.Keys[prevKey], fc.Keys[key], lerpFor)
			}
		} else if i == len(keys)-1 {
			return fc.Keys[key]
		}
	}

	return 0
}

func (fc *FloatCurve) GetTimeLength() float32 {
	var maxTime float32 = 0.

	for key := range fc.Keys {
		if key > maxTime {
			maxTime = key
		}
	}

	return maxTime
}
