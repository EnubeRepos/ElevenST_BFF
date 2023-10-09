package model

type ClinicReferenceList struct {
	List []ClinicReference `json:"list"`
}

type ClinicReference struct {
	Name string `json:"clinicName"`
	URL  string `json:"url"`
}

type ResponseContact struct {
	Total int                   `json:"total"`
	List  []ResponseContactList `json:"list"`
}

type ResponseContactList struct {
	ClinicName string  `json:"clinicName"`
	URL        string  `json:"url"`
	Entity     string  `json:"entity"`
	Content    Contact `json:"response"`
}
