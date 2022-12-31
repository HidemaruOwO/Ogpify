package lib

import (
	"fmt"
	"image"
	"image/color"
	"io/ioutil"
	"path/filepath"
	"unicode/utf8"

	"github.com/golang/freetype/truetype"
	imageFont "golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

type ImageCreateOptions struct {
	Font *string
}

func ImageCreate(text string, options *ImageCreateOptions) image.Image {
	//Optionのハンドル
	textFont := "Koruri-Bold.ttf"
	if options.Font != nil {
		textFont = *options.Font
	}
	fmt.Printf("text: %s, textfont: %s\n", text, textFont)

	fontPath := filepath.Join("fonts/Koruri-Bold.ttf")

	fontBinary, err := ioutil.ReadFile(fontPath)
	if err != nil {
		Error(err)
		return nil
	}

	font, err := truetype.Parse(fontBinary)
	if err != nil {
		Error(err)
		return nil
	}

	fontSize := fontSize(text)

	textOptions := truetype.Options{
		Size:              fontSize,
		DPI:               0,
		Hinting:           0,
		GlyphCacheEntries: 0,
		SubPixelsX:        0,
		SubPixelsY:        0,
	}

	imageWidth := 1280
	imageHeight := 510

	img := image.NewRGBA(image.Rect(0, 0, imageWidth, imageHeight))

	fontFace := truetype.NewFace(font, &textOptions)

	fontDrawer := &imageFont.Drawer{
		Dst:  img,
		Src:  image.NewUniform(color.RGBA{52, 52, 52, 255}),
		Face: fontFace,
		Dot:  fixed.Point26_6{},
	}

	fontDrawer.Dot.X = (fixed.I(imageWidth) - fontDrawer.MeasureString(text)) / 2
	fontDrawer.Dot.Y = fixed.I((int(fontSize)))

	fontDrawer.DrawString(text)

	return img
}

func fontSize(text string) float64 {
	textLength := utf8.RuneCountInString(text)
	var fontSize float64

	if textLength <= 6 {
		fontSize = 150
	} else if textLength <= 10 {
		fontSize = 100
	} else if textLength <= 15 {
		fontSize = 70
	}

	return fontSize
}
