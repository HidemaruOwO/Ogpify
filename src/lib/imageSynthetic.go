package lib

import (
	"image"
	"os"

	"image/draw"
)

func ImageSynthetic(charImage image.Image) image.Image {
	originImagePath := "./images/ogp_back.png"

	originFile, err := os.Open(originImagePath)
	if err != nil {
		Error(err)
		return nil
	}
	originImage, _, err := image.Decode(originFile)

	charStartPoint := image.Point{0, 330}

	charRectangle := image.Rectangle{charStartPoint, charStartPoint.Add(charImage.Bounds().Size())}
	originRectangle := image.Rectangle{image.Point{0, 0}, originImage.Bounds().Size()}

	img := image.NewRGBA(originRectangle)
	draw.Draw(img, originRectangle, originImage, image.Point{0, 0}, draw.Src)
	draw.Draw(img, charRectangle, charImage, image.Point{0, 0}, draw.Over)

	return img
}

func createdImageSynthetic(one image.Image, two image.Image, fontSize float64) image.Image {
	charStartPoint := image.Point{0, int(fontSize)}

	charRectangle := image.Rectangle{charStartPoint, charStartPoint.Add(one.Bounds().Size())}
	originRectangle := image.Rectangle{image.Point{0, 0}, one.Bounds().Size()}

	img := image.NewRGBA(originRectangle)
	draw.Draw(img, originRectangle, one, image.Point{0, 0}, draw.Src)
	draw.Draw(img, charRectangle, two, image.Point{0, 0}, draw.Over)

	return img
}
