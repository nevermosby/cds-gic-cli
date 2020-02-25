package model

type Vdc struct {
	VdcName  string `header:"Vdc Name"`
	VdcId    string `header:"Vdc ID"`
	RegionId string `header:"Region Id"`
}
