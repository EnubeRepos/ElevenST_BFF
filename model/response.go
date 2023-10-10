package model

type ClinicReferenceList struct {
	List []ClinicReference `json:"list"`
}

type ClinicReference struct {
	Name string `json:"clinicName"`
	URL  string `json:"url"`
}

type ResponseContactList struct {
	Total int               `json:"total"`
	List  []ResponseContact `json:"list"`
}

type ResponseContact struct {
	ClinicName string  `json:"clinicName"`
	URL        string  `json:"url"`
	Entity     string  `json:"entity"`
	Content    Contact `json:"response"`
}
