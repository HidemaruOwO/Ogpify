package lib

import (
	"bytes"
	"fmt"
	"image"
	"image/png"
	"math/rand"
	"os"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
)

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

func Domain2Url(domain string) string {
	return "https://" + domain
}

// Turn array and index
func SplitText(text string, splitlen int) ([]string, int) {
	array := []string{}
	runes := []rune(text)
	for i := 0; i < len(runes); i += splitlen {
		if i+splitlen < len(runes) {
			array = append(array, string(runes[i:(i+splitlen)]))
		} else {
			array = append(array, string(runes[i:]))
		}
	}
	return array, len(array)
}

// Turn array and index
func SplitCharText(text string, char string) ([]string, int) {
	array := strings.Split(text, char)

	return array, len(array)
}

func Version() string {
	return "v1.0-beta2"
}

func Infof(format string, a ...any) {
	logType := color.HiCyanString("INFO")
	fmt.Printf("[%s] %s", logType, fmt.Sprintf(format, a...))
}

func Error(err error) {
	logType := color.RedString("ERROR")
	errorMessage := color.RedString(err.Error())
	fmt.Fprintf(os.Stderr, "[%s] %s", logType, errorMessage)
}

func ErrorExit(err error) {
	Error(err)
	os.Exit(1)
}

func Warnf(format string, a ...any) {
	logType := color.HiYellowString("WARN")
	fmt.Printf("[%s] %s", logType, fmt.Sprintf(format, a...))
}

func Criticalf(format string, a ...any) {
	logType := color.HiMagentaString("CRITICAL")
	fmt.Printf("[%s] %s", logType, fmt.Sprintf(format, a...))
}

func Debugf(isDebug bool, format string, a ...any) {
	if isDebug {
		logType := color.HiGreenString("DEBUG")
		fmt.Printf("[%s] %s", logType, fmt.Sprintf(format, a...))
	}
}
