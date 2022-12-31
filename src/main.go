package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/HidemaruOwO/OGP-Generate-API/src/lib"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	lib.Intialize()

	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{lib.WebDomain(), lib.ApiDomain()},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	publicImageDir := filepath.Join("public", "image")

	// Static Router
	e.Static("/public/image/", publicImageDir)

	// Router
	e.GET("/", rootHandler)
	e.POST("/generate", generateHandler)

	e.Logger.Fatal(e.Start(":3090"))
}

func rootHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Running OGP-Generate-API.")
}

type generateJson struct {
	Text string `json:"text"`
	Font string `json:"font"`
}

type responseGenerateJson struct {
	URL string `json:"url"`
}

func generateHandler(c echo.Context) error {
	post := new(generateJson)
	if err := c.Bind(post); err != nil {
		lib.Error(err)
		return c.String(http.StatusInternalServerError, "JSONデータのバインドに失敗しました。")
	}

	font := "nemui"

	charImage, err := lib.ImageCreate(post.Text, &lib.ImageCreateOptions{Font: &font})
	if err != nil {
		lib.Error(err)
		return c.String(http.StatusInternalServerError, err.Error())
	}

	syntheticImage := lib.ImageSynthetic(charImage)
	if charImage == nil {
		errMessage := "画像の合成にてエラーが発生しました。"
		lib.Error(fmt.Errorf(errMessage))
		return c.String(http.StatusInternalServerError, errMessage)
	}

	randomString := lib.RandomString(10)

	imagePath := filepath.Join("public", "image", "synthetic"+randomString+".png")

	lib.CreateImageFile(imagePath, syntheticImage)

	responseStruct := responseGenerateJson{
		URL: lib.ApiDomain() + "/" + imagePath,
	}

	responseJson, err := json.Marshal(responseStruct)

	if err != nil {
		errMessage := "JSONの変換にてエラーが発生したため処理を終了します。"
		lib.Error(fmt.Errorf(errMessage))
		return c.JSON(http.StatusInternalServerError, errMessage)
	}

	return c.JSON(http.StatusOK, string(responseJson))
}
