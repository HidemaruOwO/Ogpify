package lib

import (
	"bytes"
	"fmt"
	"image"
	"image/png"
	"math/rand"
	"os"
	"path/filepath"

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

func Intialize() {
	publicDir := filepath.Join("public")

	if _, err := os.Stat(publicDir); os.IsNotExist(err) {
		os.Mkdir(publicDir, 0755)
	}

	publicImageDir := filepath.Join("public", "image")
	if _, err := os.Stat(publicImageDir); os.IsNotExist(err) {
		os.MkdirAll(publicImageDir, 0755)
	} else {
		RemoveFiles(publicImageDir)
	}

	fmt.Printf("Intialized!!\n")
}

func RemoveFiles(dir string) error {
	d, err := os.Open(dir)
	if err != nil {
		return err
	}
	defer d.Close()
	names, err := d.Readdirnames(-1)
	if err != nil {
		return err
	}
	for _, name := range names {
		err = os.RemoveAll(filepath.Join(dir, name))
		if err != nil {
			return err
		}
	}
	return nil
}

func RandomString(num int) string {
	letter := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	str := make([]rune, num)
	for i := range str {
		str[i] = letter[rand.Intn(len(letter))]
	}

	return string(str)
}

func ApiDomain() string {
	return "https://ogp-api.v-sli.me"
}

func WebDomain() string {
	return "https://v-sli.me"

}
