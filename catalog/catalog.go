package catalog

import (
	"github.com/iand/growth/plot"
)

type Entry struct {
	Depth int
	plot.Grammar
	plot.Plotter
}

var Entries = map[string]Entry{

	"big-h": {
		Depth: 10,
		Plotter: plot.Plotter{
			InitialAngle:   0,
			AngleDelta:     90,
			Step:           200,
			StepDelta:      0.65,
			LineWidth:      9,
			LineWidthDelta: 0.65,
		},
		Grammar: plot.Grammar{
			Seed: "[F]--F",
			Rules: map[rune]string{
				'F': "|[+F][-F]",
			},
		},
	},

	"bent-big-h": {
		Depth: 10,
		Plotter: plot.Plotter{
			InitialAngle:   0,
			AngleDelta:     80,
			Step:           200,
			StepDelta:      0.65,
			LineWidth:      9,
			LineWidthDelta: 0.65,
		},
		Grammar: plot.Grammar{
			Seed: "[F]--F",
			Rules: map[rune]string{
				'F': "|[+F][-F]",
			},
		},
	},

	"two-ys": {
		Depth: 10,
		Plotter: plot.Plotter{
			InitialAngle:   0,
			AngleDelta:     45,
			Step:           200,
			StepDelta:      0.65,
			LineWidth:      9,
			LineWidthDelta: 0.65,
		},
		Grammar: plot.Grammar{
			Seed: "[F]----F",
			Rules: map[rune]string{
				'F': "|[+F][-F]",
			},
		},
	},

	"twig": {
		Depth: 10,
		Plotter: plot.Plotter{
			InitialAngle:   0,
			AngleDelta:     20,
			Step:           200,
			StepDelta:      0.65,
			LineWidth:      9,
			LineWidthDelta: 0.65,
		},
		Grammar: plot.Grammar{
			Seed: "F",
			Rules: map[rune]string{
				'F': "|[-F][+F]",
			},
		},
	},

	"weed-1": {
		Depth: 6,
		Plotter: plot.Plotter{
			InitialAngle:   0,
			AngleDelta:     25,
			Step:           200,
			StepDelta:      0.45,
			LineWidth:      9,
			LineWidthDelta: 0.45,
		},
		Grammar: plot.Grammar{
			Seed: "F",
			Rules: map[rune]string{
				'F': "F[-F]F[+F]F",
			},
		},
	},

	"weed-2": {
		Depth: 6,
		Plotter: plot.Plotter{
			InitialAngle:   0,
			AngleDelta:     25,
			Step:           200,
			StepDelta:      0.45,
			LineWidth:      9,
			LineWidthDelta: 0.45,
		},
		Grammar: plot.Grammar{
			Seed: "F",
			Rules: map[rune]string{
				'F': "|[-F]|[+F]F",
			},
		},
	},

	"weed-3": {
		Depth: 4,
		Plotter: plot.Plotter{
			InitialAngle:   0,
			AngleDelta:     20,
			Step:           200,
			StepDelta:      0.45,
			LineWidth:      9,
			LineWidthDelta: 0.45,
		},
		Grammar: plot.Grammar{
			Seed: "F",
			Rules: map[rune]string{
				'F': "|[-F]|[+F][-F]F",
			},
		},
	},

	"bush-1": {
		Depth: 5,
		Plotter: plot.Plotter{
			InitialAngle:   0,
			AngleDelta:     25,
			Step:           200,
			StepDelta:      0.45,
			LineWidth:      9,
			LineWidthDelta: 0.45,
		},
		Grammar: plot.Grammar{
			Seed: "F",
			Rules: map[rune]string{
				'F': "FF+[+F-F-F]-[-F+F+F]",
			},
		},
	},

	"bush-2": {
		Depth: 8,
		Plotter: plot.Plotter{
			InitialAngle:   0,
			AngleDelta:     20,
			Step:           200,
			StepDelta:      0.45,
			LineWidth:      9,
			LineWidthDelta: 0.45,
		},
		Grammar: plot.Grammar{
			Seed: "F",
			Rules: map[rune]string{
				'F': "|[+F]|[-F]+F",
			},
		},
	},

	"tree-1": {
		Depth: 6,
		Plotter: plot.Plotter{
			InitialAngle:   0,
			AngleDelta:     20,
			Step:           200,
			StepDelta:      0.50,
			LineWidth:      9,
			LineWidthDelta: 0.5,
		},
		Grammar: plot.Grammar{
			Seed: "F",
			Rules: map[rune]string{
				'F': "|[---F][+++F]|[--F][++F]|F",
			},
		},
	},

	"tree-2": {
		Depth: 5,
		Plotter: plot.Plotter{
			InitialAngle:   0,
			AngleDelta:     8,
			Step:           200,
			StepDelta:      0.5,
			LineWidth:      9,
			LineWidthDelta: 0.5,
		},
		Grammar: plot.Grammar{
			Seed: "F",
			Rules: map[rune]string{
				'F': "|[+++++F][-------F]-|[++++F][------F]-|[+++F][-----F]-|F",
			},
		},
	},

	"tree-3": {
		Depth: 9,
		Plotter: plot.Plotter{
			InitialAngle:   0,
			AngleDelta:     20,
			Step:           200,
			StepDelta:      0.65,
			LineWidth:      9,
			LineWidthDelta: 0.65,
		},
		Grammar: plot.Grammar{
			Seed: "F",
			Rules: map[rune]string{
				'F': "|[--F][+F]-F",
			},
		},
	},

	"carpet": {
		Depth: 5,
		Plotter: plot.Plotter{
			InitialAngle:   0,
			AngleDelta:     90,
			Step:           200,
			StepDelta:      0.5,
			LineWidth:      9,
			LineWidthDelta: 0.5,
		},
		Grammar: plot.Grammar{
			Seed: "F-F-F-F",
			Rules: map[rune]string{
				'F': "F[F]-F+F[--F]+F-F",
			},
		},
	},

	"sierpinksi-square": {
		Depth: 5,
		Plotter: plot.Plotter{
			InitialAngle:   0,
			AngleDelta:     90,
			Step:           200,
			StepDelta:      0.5,
			LineWidth:      2,
			LineWidthDelta: 0.8,
		},
		Grammar: plot.Grammar{
			Seed: "F-F-F-F",
			Rules: map[rune]string{
				'F': "FF[-F-F-F]F",
			},
		},
	},

	"rug": {
		Depth: 5,
		Plotter: plot.Plotter{
			InitialAngle:   0,
			AngleDelta:     90,
			Step:           200,
			StepDelta:      0.5,
			LineWidth:      2,
			LineWidthDelta: 0.8,
		},
		Grammar: plot.Grammar{
			Seed: "F-F-F-F",
			Rules: map[rune]string{
				'F': "F[-F-F]FF",
			},
		},
	},

	"koch-island": {
		Depth: 5,
		Plotter: plot.Plotter{
			InitialAngle:   0,
			AngleDelta:     60,
			Step:           200,
			StepDelta:      0.5,
			LineWidth:      2,
			LineWidthDelta: 0.8,
		},
		Grammar: plot.Grammar{
			Seed: "F++F++F",
			Rules: map[rune]string{
				'F': "F-F++F-F",
			},
		},
	},

	"quadric-koch-island": {
		Depth: 4,
		Plotter: plot.Plotter{
			InitialAngle:   0,
			AngleDelta:     90,
			Step:           200,
			StepDelta:      0.35,
			LineWidth:      1,
			LineWidthDelta: 0.9,
		},
		Grammar: plot.Grammar{
			Seed: "F-F-F-F",
			Rules: map[rune]string{
				'F': "F-F+F+FF-F-F+F",
			},
		},
	},

	"square-spikes": {
		Depth: 5,
		Plotter: plot.Plotter{
			InitialAngle:   0,
			AngleDelta:     5,
			Step:           200,
			StepDelta:      0.5,
			LineWidth:      1,
			LineWidthDelta: 0.9,
		},
		Grammar: plot.Grammar{
			Seed: "F------------------F------------------F------------------F",
			Rules: map[rune]string{
				'F': "F-----------------F++++++++++++++++++++++++++++++++++F-----------------F",
			},
		},
	},

	"sierpinksi-gasket": {
		Depth: 8,
		Plotter: plot.Plotter{
			InitialAngle:   -30,
			AngleDelta:     60,
			Step:           200,
			StepDelta:      0.65,
			LineWidth:      1,
			LineWidthDelta: 0.9,
		},
		Grammar: plot.Grammar{
			Seed: "F--F--F",
			Rules: map[rune]string{
				'F': "F--F--F--GG",
				'G': "GG",
			},
		},
	},

	"sierpinksi-maze": {
		Depth: 8,
		Plotter: plot.Plotter{
			InitialAngle:   -30,
			AngleDelta:     60,
			Step:           200,
			StepDelta:      0.65,
			LineWidth:      1,
			LineWidthDelta: 0.9,
		},
		Grammar: plot.Grammar{
			Seed: "F",
			Rules: map[rune]string{
				'F': "[GF][+G---F][G+G+F]",
				'G': "GG",
			},
		},
	},

	"sierpinksi-arrowhead": {
		Depth: 8,
		Plotter: plot.Plotter{
			InitialAngle:   -30,
			AngleDelta:     60,
			Step:           200,
			StepDelta:      0.65,
			LineWidth:      1,
			LineWidthDelta: 0.9,
		},
		Grammar: plot.Grammar{
			Seed: "F",
			Rules: map[rune]string{
				'F': "[-G+++F][-G+F][GG--F]",
				'G': "GG",
			},
		},
	},

	"penrose-snowflake": {
		Depth: 4,
		Plotter: plot.Plotter{
			InitialAngle:   0,
			AngleDelta:     18,
			Step:           200,
			StepDelta:      0.5,
			LineWidth:      1,
			LineWidthDelta: 0.9,
		},
		Grammar: plot.Grammar{
			Seed: "F----F----F----F----F",
			Rules: map[rune]string{
				'F': "F----F----F----------F++F----F",
			},
		},
	},

	"penrose-tile": {
		Depth: 6,
		Plotter: plot.Plotter{
			InitialAngle:   0,
			AngleDelta:     36,
			Step:           200,
			StepDelta:      0.65,
			LineWidth:      1,
			LineWidthDelta: 0.9,
		},
		Grammar: plot.Grammar{
			Seed: "[X]++[X]++[X]++[X]++[X]",
			Rules: map[rune]string{
				'W': "YF++ZF----XF[-YF----WF]++",
				'X': "+YF--ZF[---WF--XF]+",
				'Y': "-WF++XF[+++YF++ZF]-",
				'Z': "--YF++++WF[+ZF++++XF]--XF",
				'F': "",
			},
		},
	},

	"dragon-curve": {
		Depth: 11,
		Plotter: plot.Plotter{
			InitialAngle:   90,
			AngleDelta:     45,
			Step:           200,
			StepDelta:      0.75,
			LineWidth:      1,
			LineWidthDelta: 1,
		},
		Grammar: plot.Grammar{
			Seed: "F",
			Rules: map[rune]string{
				'F': "[+F][+G--G----F]",
				'G': "-G++G-",
			},
		},
	},
}
