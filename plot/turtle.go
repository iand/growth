package plot

import (
	"math"

	"github.com/llgcode/draw2d"
)

func NewTurtle(gc draw2d.GraphicContext) *Turtle {
	return &Turtle{
		state: &state{},
		gc:    gc,
	}
}

type Turtle struct {
	state          *state
	gc             draw2d.GraphicContext
	x0, x1, y0, y1 float64
}

type state struct {
	x, y float64
	h    float64 // heading in degrees
	prev *state
}

func (t *Turtle) MoveTo(x, y float64) {
	t.state.x = x
	t.state.y = y
	t.updateBounds()
}

func (t *Turtle) RotateTo(h float64) {
	t.state.h = h
}
func (t *Turtle) SetLineWidth(w float64) {
	t.gc.SetLineWidth(w)
}

func (t *Turtle) DrawForward(d float64) {
	t.gc.BeginPath()
	t.gc.MoveTo(t.state.x, t.state.y)
	t.GoForward(d)
	t.gc.LineTo(t.state.x, t.state.y)
	t.gc.Close()
	t.gc.Stroke()
}

func (t *Turtle) GoForward(d float64) {
	s, c := math.Sincos(t.state.h * math.Pi / 180.0)
	t.state.x += d * s
	t.state.y -= d * c
	t.updateBounds()
}

func (t *Turtle) Rotate(a float64) {
	t.state.h += a
	for t.state.h >= 360 {
		t.state.h -= 360
	}
	for t.state.h < 0 {
		t.state.h += 360
	}
}

func (t *Turtle) SaveState() {
	st := &state{
		x:    t.state.x,
		y:    t.state.y,
		h:    t.state.h,
		prev: t.state,
	}
	t.state = st
}

func (t *Turtle) RestoreState() {
	if t.state.prev != nil {
		t.state = t.state.prev
	}
}

func (t *Turtle) Bounds() (float64, float64, float64, float64) {
	return t.x0, t.x1, t.y0, t.y1
}

func (t *Turtle) updateBounds() {
	if t.state.x < t.x0 {
		t.x0 = t.state.x
	}
	if t.state.x > t.x1 {
		t.x1 = t.state.x
	}

	if t.state.y < t.y0 {
		t.y0 = t.state.y
	}
	if t.state.y > t.y1 {
		t.y1 = t.state.y
	}

}
