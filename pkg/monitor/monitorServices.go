package monitor

import (
	"net/http"

	//"tim_presse/tim_test_ms_util_gen_numrange/bootstrap"

	"github.com/BayramGuenesIvz/tim_web_show_apps/internal/readConfig"
	"github.com/gin-gonic/gin"
)

func GotoMonitor(c *gin.Context) {

	viewpagedata = GetAppsAddrRef()
	//println("viewpagedata.AddrRef.DB.Port" + viewpagedata.AddrRef.DB.Port)
	//t.Execute(w, viewpagedata)
	c.HTML(http.StatusOK, "01_mainview.html", gin.H{"AppsAddrInfo": viewpagedata.AddrRef})

}

func GetAppsAddrRef() ViewPageDataInfo {
	lViewpagedata := ViewPageDataInfo{}
	lViewpagedata.AddrRef, _ = readConfig.GetApplConf()
	return lViewpagedata
}
