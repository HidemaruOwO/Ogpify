package main

import (
	"fmt"

	"github.com/HidemaruOwO/OGP-Generate-API/src/lib"
)

func main() {
	font := "nemui"

	charImage := lib.ImageCreate("unko", &lib.ImageCreateOptions{Font: &font})
	if charImage == nil {
		lib.Error(fmt.Errorf("エラーが発生したため処理を終了します。"))
		return
	}

	syntheticImage := lib.ImageSynthetic(charImage)
	if charImage == nil {
		lib.Error(fmt.Errorf("エラーが発生したため処理を終了します。"))
		return
	}

	lib.CreateImageFile("./synthetic.png", syntheticImage)
}
