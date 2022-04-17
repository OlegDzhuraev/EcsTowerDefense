package engine

import rl "github.com/gen2brain/raylib-go/raylib"

type Timer struct {
	Time     float32
	timeLeft float32
}

func (t *Timer) DoTick() {
	t.DoTickUntilReady()

	if t.IsReady() {
		t.Reset()
	}
}

func (t *Timer) DoTickUntilReady() {
	if t.timeLeft > 0 {
		t.timeLeft -= rl.GetFrameTime()
	}
}

func (t Timer) IsReady() bool { return t.timeLeft <= 0 }
func (t *Timer) Reset()       { t.timeLeft = t.Time }
