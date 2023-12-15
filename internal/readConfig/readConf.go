package readConfig

import (
	"encoding/json"
	"errors"
	"log" //"strconv"
	"os"

	"github.com/BayramGuenesIvz/tim_web_show_apps/internal/data"
)

func GetApplConf() (eConf data.ApplConfStruct, err error) {

	data.ApplConf = data.ApplConfStruct{}
	if len(data.ThisConfLocation) == 0 {
		errString :=
			"Achtung !!!  Service konnte nicht gestartet werden." +
				" Bitte Pfad der Konfigurationsdatei configshowapps.json  über Umgebungsparameter" +
				" 'confLocation' bzw als Startargument(osparam) 'confLocation=<path>' angeben."
		println(errString)
		return data.ApplConf, errors.New(errString)
	}

	filePathAndName := data.ThisConfLocation
	file, _ := os.Open(filePathAndName)
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&data.ApplConf)

	if err != nil {
		println("Decoding configshowapps.json"+" FAILED", err.Error())
		log.Fatal(err.Error() + ". Möglicherweise fehlt die configshowapps.json Datei im Pfad " + data.ThisConfLocation)
		return data.ApplConf, err
	}
	eConf = data.ApplConf

	return eConf, nil

}
