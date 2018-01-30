package main

import (
	"errors"
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"math"
	"os"

	"github.com/iand/growth/catalog"
	"github.com/llgcode/draw2d/draw2dimg"
)

const (
	// complexityLimit is the maximum size of the expanded ruleset
	complexityLimit = 250000

	// maxWidth is the maximum width of the output image
	maxWidth = 2048

	// maxHeight is the maximum height of the output image
	maxHeight = 2048
)

var (
	margin         = flag.Int("m", 10, "Number of pixels to leave as a margin around the plotted image.")
	outputFilename = flag.String("o", "growth.png", "Name of file image will be written to.")
	entryName      = flag.String("e", "big-h", "Name of catalog entry to plot.")
	depth          = flag.Int("d", 0, "Overide the number of grammar expansion iterations.")
	dumpGrammer    = flag.Bool("g", false, "Dump the grammar used for the plot to stdout.")
)

func main() {
	flag.Parse()
	if err := Main(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}

func Main() error {
	entry, ok := catalog.Entries[*entryName]
	if !ok {
		return fmt.Errorf("unknown catalog entry: %s", *entryName)
	}

	if *depth == 0 {
		*depth = entry.Depth
	}

	s, err := entry.Grammar.Iterate(*depth, complexityLimit)
	if err != nil {
		return err
	}

	if *dumpGrammer {
		fmt.Println(s)
	}

	x0, x1, y0, y1 := entry.Plotter.Measure(s)
	w := int(math.Ceil(x1 - x0 + entry.Plotter.LineWidth))
	h := int(math.Ceil(y1 - y0 + entry.Plotter.LineWidth))
	if w <= 0 {
		return errors.New("ruleset resulted in a zero width image")
	}
	if h <= 0 {
		return errors.New("ruleset resulted in a zero height image")
	}
	if w > maxWidth {
		return fmt.Errorf("ruleset resulted in an image that was too wide (%d pixels)", w)
	}
	if h > maxHeight {
		return fmt.Errorf("ruleset resulted in an image that was too tall (%d pixels)", w)
	}

	img := image.NewRGBA(image.Rect(0, 0, w+2*(*margin), h+2*(*margin)))
	draw.Draw(img, img.Bounds(), &image.Uniform{color.RGBA{255, 255, 255, 255}}, image.ZP, draw.Src)
	gc := draw2dimg.NewGraphicContext(img)

	if err := entry.Plotter.Plot(s, gc, -x0+float64(*margin), -y0+float64(*margin)); err != nil {
		return err
	}

	f, err := os.Create(*outputFilename)
	if err != nil {
		return err
	}
	defer f.Close()

	return png.Encode(f, img)
}
