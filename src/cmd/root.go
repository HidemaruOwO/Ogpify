package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/HidemaruOwO/OGP-Generate-API/src/lib"
	"github.com/fatih/color"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/cobra"
)

type CmdOptions struct {
	optApiDomain      string
	optWebDomain      string
	optPort           string
	optDebug          bool
	optReverseProxyed bool
	optHttps          bool
}

var o = &CmdOptions{}

var RootCmd = &cobra.Command{
	Run: func(cmd *cobra.Command, args []string) {
		if o.optApiDomain != "" && o.optWebDomain != "" {
			server()
		} else if o.optWebDomain != "" {
			fmt.Printf("Need --api-domain flag\n")
		} else if o.optApiDomain != "" {
			fmt.Printf("Need --page-domain flag\n")
		} else {
			fmt.Printf("Ogpify %s\n", lib.Version())
			fmt.Printf("✨ Thank you for installing Ogpify!!\n")
			fmt.Printf("Please run the help command for usage. \nRun:\n")
			color.New(color.Bold).Printf("\t" + color.BlueString("$ ") + "ogpify --help\n")
		}
	},
}

func init() {
	cobra.OnInitialize()
	RootCmd.AddCommand()
	RootCmd.Flags().StringVarP(&o.optPort, "port", "", "", "Port (Default: 3090)")
	RootCmd.Flags().StringVarP(&o.optApiDomain, "api-domain", "a", "", "API Domain option (Example: api.ogc.v-sli.me)")
	RootCmd.Flags().StringVarP(&o.optWebDomain, "page-domain", "p", "", "Domain of the site used for the Post (Example: ogc.v-sli.me)")
	RootCmd.Flags().BoolVarP(&o.optDebug, "debug", "d", false, "Enable this flag causes logging in debug mode")
	RootCmd.Flags().BoolVarP(&o.optReverseProxyed, "isReverseProxyed", "", false, "Enabling this flag removes the port of return. If you are using a reverse proxy, please enable this")
	RootCmd.Flags().BoolVarP(&o.optHttps, "isHttps", "", false, "If you enable this flag, the url in the return value will start with https; if your API server is SSL-enabled, please enable this flag")
}

func server() {
	lib.Intialize()

	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{lib.Domain2Url(o.optWebDomain, o.optReverseProxyed, o.optHttps, o.optPort), lib.Domain2Url(o.optApiDomain, o.optReverseProxyed, o.optHttps, o.optPort)},
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

	if o.optPort != "" {
		port, err := strconv.Atoi(o.optPort)
		if err != nil {
			lib.ErrorExit(err)
		}

		e.Logger.Fatal(e.Start(":" + strconv.Itoa(port)))
	} else {
		e.Logger.Fatal(e.Start(":3090"))
	}
}

func rootHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Running OGP-Generate-API.")
}

type generateJson struct {
	Text      string `json:"text"`
	Font      string `json:"font"`
	SingleUrl bool   `json:bool`
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

	charImage, err := lib.ImageCreate(post.Text, o.optDebug, &lib.ImageCreateOptions{Font: &font})
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
		URL: lib.Domain2Url(o.optApiDomain, o.optReverseProxyed, o.optHttps, o.optPort) + "/" + imagePath,
	}

	responseJson, err := json.Marshal(responseStruct)

	if err != nil {
		errMessage := "An error occurred during conversion to JSON"
		lib.Error(fmt.Errorf(errMessage))
		return c.JSON(http.StatusInternalServerError, errMessage)
	}

	// ここのコメントを考えてよこパイロットくん
	if post.SingleUrl {
		return c.String(http.StatusOK, lib.Domain2Url(o.optApiDomain, o.optReverseProxyed, o.optHttps, o.optPort)+"/"+imagePath)
	} else {
		return c.JSON(http.StatusOK, string(responseJson))
	}
}
