package monitor

/* ----------------------------------- */
import "github.com/BayramGuenesIvz/tim_web_show_apps/internal/data"

// ViewPageDataInfo ...
type ViewPageDataInfo struct {
	AddrRef data.ApplConfStruct
	//StockLoading []LoadingStruct
}

var viewpagedata ViewPageDataInfo

/* ----------------------------------- */

/* ============Konstanten ============================================= */

/* ============Types ============================================= */

type nameValStruct struct {
	name string
	val  string
}
