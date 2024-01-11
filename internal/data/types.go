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
	Loadchannel    ApplConfItemStruct
	Repo           ApplConfItemStruct
	RepoChanges    ApplConfItemStruct
	RuiRepo        ApplConfItemStruct
	ELS            ApplConfItemStruct
	ELSGui         ApplConfItemStruct
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
