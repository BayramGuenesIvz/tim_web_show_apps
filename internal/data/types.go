package data

type ApplConfItemStruct struct {
	Addr    string
	Port    string
	IntPort string
	Notes   string
}
type ApplConfStruct struct {
	DB             ApplConfItemStruct
	Adminer        ApplConfItemStruct
	NumRangeServer ApplConfItemStruct
	LogServer      ApplConfItemStruct
	Overview       ApplConfItemStruct
	MockProducer   ApplConfItemStruct
	Receiver       ApplConfItemStruct
	Repo           ApplConfItemStruct
	DBOrgLogger    ApplConfItemStruct
	DBOrgReceiver  ApplConfItemStruct
	DBOrgRepo      ApplConfItemStruct
	DBOrga         ApplConfItemStruct
	SysPortainer   ApplConfItemStruct
}

var ApplConf ApplConfStruct

var (
	ThisServerAddr   string
	ThisPort         string
	ThisConfLocation string
)
