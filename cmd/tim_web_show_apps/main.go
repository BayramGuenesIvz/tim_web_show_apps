package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/BayramGuenesIvz/tim_web_show_apps/internal/data"
	config "github.com/BayramGuenesIvz/tim_web_show_apps/internal/readConfig"
	settings "github.com/BayramGuenesIvz/tim_web_show_apps/internal/readSettings"
	"github.com/BayramGuenesIvz/tim_web_show_apps/pkg/monitor"
)

var ()

func main() {
	println()
	routerGinGonic()
}

func routerGinGonic() {
	settings.LoadExternalSettings()
	_, err := config.GetApplConf()
	if err != nil {
		return
	}
	println("hello 20240111 2")
	router := gin.New()

	router.GET("/", amAliveHandler)

	// -----------------------------------------------------------------//
	router.LoadHTMLGlob("web/templates/*")
	//router.Static("/static", "./public")
	//router.Static("/static", "web/public")
	router.Static("/static", "web/static")

	router.GET("/index", monitor.GotoMonitor)
	//router.POST("/performOp", monitor.ServicesPerformOp)
	// -----------------------------------------------------------------//

	log.Print("Starting Server on Port", ":"+data.ThisPort) //setImageServiceRoutes(router)

	http.ListenAndServe(":"+data.ThisPort, router)
}
func amAliveHandler(c *gin.Context) {
	c.JSON(http.StatusOK, "tim_web_show_apps (micro-)service is alive")
}
