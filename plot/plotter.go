// Package plot implemants the Fractal Growth algorithm described in The Computational Beauty of Nature by Gary William Flake.
package plot

import (
	"image"
	"math"

	"github.com/llgcode/draw2d"
	"github.com/llgcode/draw2d/draw2dimg"
)

type Plotter struct {
	InitialAngle   float64
	AngleDelta     float64
	Step           float64
	StepDelta      float64
	LineWidth      float64
	LineWidthDelta float64
}

func (p *Plotter) Plot(s string, gc draw2d.GraphicContext, x, y float64) error {
	t := NewTurtle(gc)
	t.MoveTo(x, y)
	t.RotateTo(p.InitialAngle)
	p.do(t, s)
	return nil
}

func (p *Plotter) do(t *Turtle, s string) {
	depth := 0.0
	step := p.Step
	t.SetLineWidth(p.LineWidth)
	for _, c := range s {
		switch c {
		case 'F':
			t.DrawForward(step)
		case 'G':
			t.GoForward(step)
		case '+':
			t.Rotate(p.AngleDelta)
		case '-':
			t.Rotate(-p.AngleDelta)
		case '[':
			t.SaveState()
		case ']':
			t.RestoreState()
		case '|':
			t.DrawForward(step)
		case '<':
			depth++
			step = p.Step * math.Pow(p.StepDelta, depth)
			t.SetLineWidth(p.LineWidth * math.Pow(p.LineWidthDelta, depth))
		case '>':
			if depth > 0 {
				depth--
			}
			step = p.Step * math.Pow(p.StepDelta, depth)
			t.SetLineWidth(p.LineWidth * math.Pow(p.LineWidthDelta, depth))
		}
	}
}

func (p *Plotter) Measure(s string) (float64, float64, float64, float64) {
	img := image.NewRGBA(image.Rect(0, 0, 1, 1))
	gc := draw2dimg.NewGraphicContext(img)
	t := NewTurtle(gc)
	t.MoveTo(0, 0)
	t.RotateTo(p.InitialAngle)
	depth := 0.0
	step := p.Step
	for _, c := range s {
		switch c {
		case 'F':
			t.GoForward(step)
		case 'G':
			t.GoForward(step)
		case '+':
			t.Rotate(p.AngleDelta)
		case '-':
			t.Rotate(-p.AngleDelta)
		case '[':
			t.SaveState()
		case ']':
			t.RestoreState()
		case '|':
			t.GoForward(step)
		case '<':
			depth++
			step = p.Step * math.Pow(p.StepDelta, depth)
		case '>':
			if depth > 0 {
				depth--
			}
			step = p.Step * math.Pow(p.StepDelta, depth)
		}
	}
	return t.Bounds()
}
