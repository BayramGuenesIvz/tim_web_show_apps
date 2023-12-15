package readSettings

import (
	"flag"
	"os"
	"strings"

	"github.com/BayramGuenesIvz/tim_web_show_apps/internal/data"
)

func LoadExternalSettings() {

	loadFromOSFlags()
	// Überschreiben OS Flags durch ENV Param
	loadFromOSEnv()
	// Überschreiben Parameter, falls durch args mitgegeben
	loadFromOSArgs()
	//println("ThisPort:" + ThisPort)
	//	println("ServiceDB, PortDB:" + ServiceDB + ":" + PortDB)
	return
}

func loadFromOSFlags() {
	confLocationPointer := flag.String("confLocation", "", "config_addr.json path")

	servAdrPointer := flag.String("serviceAddr ", "127.0.0.1", "service Addr")
	portPointer := flag.String("port", "18000", "application port")

	flag.Parse()
	data.ThisConfLocation = *confLocationPointer
	data.ThisServerAddr = *servAdrPointer
	data.ThisPort = *portPointer
	return
}

func loadFromOSEnv() {

	confLocation := data.ThisConfLocation
	serviceAddr := data.ThisServerAddr
	port := data.ThisPort
	environList := os.Environ()
	leng := len(environList)
	for i := 0; i < leng; i++ {
		if i > 0 {
			osenvparamval := environList[i]
			splittedString := strings.Split(osenvparamval, "=")
			paramname := splittedString[0]
			paramval := splittedString[1]

			if paramname == "configfile" ||
				paramname == "confLocation" || paramname == "ConfLocation" ||
				paramname == "conflocation" || paramname == "Conflocation" || paramname == "CONFLOCATION" {
				confLocation = paramval
			}
			if paramname == "SVCTimShowApps" || paramname == "serviceAddr" || paramname == "SERVICEADDR" {
				serviceAddr = paramval
			}

			if paramname == "port" || paramname == "PORT" || paramname == "Port" ||
				paramname == "PortTimShowApps" || paramname == "PORTTIMSHOWAPPS" {
				port = paramval
			}

		}
	}
	data.ThisConfLocation = confLocation
	data.ThisServerAddr = serviceAddr
	data.ThisPort = port

}

func loadFromOSArgs() {
	confLocation := data.ThisConfLocation
	serviceAddr := data.ThisServerAddr
	port := data.ThisPort

	leng := len(os.Args)
	for i := 0; i < leng; i++ {
		if i > 0 {
			osparam := os.Args[i]
			splittedString := strings.Split(osparam, "=")
			var namevalues []string
			namevalues = append(namevalues, splittedString...)
			leng := len(namevalues)
			if leng > 1 {
				paramname := splittedString[0]
				paramval := splittedString[1]

				if paramname == "configfile" ||
					paramname == "confLocation" || paramname == "ConfLocation" ||
					paramname == "conflocation" || paramname == "Conflocation" || paramname == "CONFLOCATION" {
					confLocation = paramval
				}
				if paramname == "SVCTimShowApps" || paramname == "serviceAddr" || paramname == "SERVICEADDR" {
					serviceAddr = paramval
				}

				if paramname == "port" || paramname == "PORT" || paramname == "Port" ||
					paramname == "PortTimShowApps" || paramname == "PORTTIMSHOWAPPS" {
					port = paramval
				}

			}

		}
	}
	data.ThisConfLocation = confLocation
	data.ThisServerAddr = serviceAddr
	data.ThisPort = port

}
