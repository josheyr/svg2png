package tests

import (
	svg2png "github.com/josheyr/svg2png/pkg"
	"os"
	"testing"
)

const (
	testSvg = "<?xml version=\"1.0\" encoding=\"UTF-8\" standalone=\"no\" ?>\n<!DOCTYPE svg PUBLIC \"-//W3C//DTD SVG 1.1//EN\" \"http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd\">\n<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" version=\"1.1\" width=\"500\" height=\"500\" viewBox=\"0 0 500 500\" xml:space=\"preserve\">\n<desc>Created with Fabric.js 5.2.4</desc>\n<defs>\n</defs>\n<g transform=\"matrix(6.54 0 0 6.54 253.46 242.97)\"  >\n<circle style=\"stroke: none; stroke-width: 1; stroke-dasharray: none; stroke-linecap: butt; stroke-dashoffset: 0; stroke-linejoin: miter; stroke-miterlimit: 4; fill: rgb(255,85,85); fill-rule: nonzero; opacity: 1;\"  cx=\"0\" cy=\"0\" r=\"30\" />\n</g>\n</svg>"
)

func TestSvgToPng(t *testing.T) {
	pngitem, err := svg2png.SvgToPng(testSvg, 500, 500)
	if err != nil {
		t.Error(err)
	}

	// show png
	file, err := os.Create("test.png")
	if err != nil {
		t.Error(err)
	}

	_, err = file.Write(pngitem)
	if err != nil {
		t.Error(err)
	}

	err = file.Close()
	if err != nil {
		t.Error(err)
	}

	// print file path
	t.Log("file path: " + file.Name())
}
