package models

type Patient struct {
	Id	int
	UserID string
	Name string
	EName string
	Sex string `orm:"size(8)"`
	Age int8
	AgeType string `orm:"size(8)"`
	Birth string `orm:"size(8)"`
	Addr string
	Tel string `orm:"size(16)"`
	EMail string `orm:"size(64)"`
}


type Examine struct {
	Id	int
	PatientId int
	ChkDate string `orm:"size(8)"`
	ChkTime string `orm:"size(6)"`
	ChkType string `orm:"size(16)"`
	ChkPart1 string `orm:"size(32)"`
	ChkPart2 string `orm:"size(32)"`
	ChkPart3 string `orm:"size(32)"`
	ChkMode string `orm:"size(16)"`
	ShowInfo string `orm:"type(text)"`
	Conclusion string `orm:"type(text)"`
	Suggestion string `orm:"type(text)"`
	RepDate string `orm:"size(8)"`
	RepDoctor string
	
}