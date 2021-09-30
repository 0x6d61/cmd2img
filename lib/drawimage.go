package lib

import (
	"bufio"
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io/ioutil"
	"os"

	_ "cmd2img/lib/statik"

	"github.com/golang/freetype/truetype"
	"github.com/rakyll/statik/fs"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

func readFontFile() font.Face {
	statikFs, err := fs.New()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	file, err := statikFs.Open("/RictyDiminished-Bold.ttf")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	font, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()
	ft, err := truetype.Parse(font)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	option := truetype.Options{
		Size:              20,
		DPI:               0,
		Hinting:           0,
		GlyphCacheEntries: 0,
		SubPixelsX:        0,
		SubPixelsY:        0,
	}

	return truetype.NewFace(ft, &option)
}

func DrawImage(text string, pngFileName string) {

	row, column := maxLine(text)
	imageWidth := 15 * column
	imageHeight := 20 * (row + 1)

	textTopMargin := 20

	bgColor := color.RGBA{0, 0, 0, 0xff}
	img := image.NewRGBA(image.Rect(0, 0, imageWidth, imageHeight))
	draw.Draw(img, img.Bounds(), &image.Uniform{bgColor}, image.ZP, draw.Src)

	face := readFontFile()
	draw := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(color.RGBA{0, 230, 64, 1}),
		Face: face,
		Dot:  fixed.Point26_6{},
	}

	buf := bytes.NewBufferString(text)
	scanner := bufio.NewScanner(buf)
	countLine := 0
	for scanner.Scan() {
		countLine++
		line := scanner.Text()
		draw.Dot.X = fixed.I(20)
		draw.Dot.Y = fixed.I(textTopMargin * countLine)
		draw.DrawString(line)
	}

	buf = &bytes.Buffer{}
	err := png.Encode(buf, img)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	file, err := os.Create(pngFileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()

	file.Write(buf.Bytes())

}

func maxLine(text string) (int, int) {
	buf := bytes.NewBufferString(text)
	scanner := bufio.NewScanner(buf)
	maxColumn := 0
	maxRow := 0
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > maxColumn {
			maxColumn = len(line)
		}
		maxRow++
	}

	return maxRow, maxColumn
}
