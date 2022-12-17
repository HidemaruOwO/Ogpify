package lib

import (
	"bytes"
	"fmt"
	"image"
	"image/png"
	"os"

	"github.com/fatih/color"
)

func Error(err error) {
	errorMessage := color.RedString(err.Error())
	fmt.Fprintf(os.Stderr, errorMessage+"\n")
}

func ErrorExit(err error) {
	Error(err)
	os.Exit(1)
}

func CreateImageFile(path string, image image.Image) {
	buf := &bytes.Buffer{}
	err := png.Encode(buf, image)

	if err != nil {
		Error(err)
		return
	}

	file, err := os.Create(path)
	if err != nil {
		Error(err)
		return
	}
	defer file.Close()

	file.Write(buf.Bytes())
}
