package model

type ContactList struct {
	Total int       `json:"total"`
	List  []Contact `json:"list"`
}

type Contact struct {
	ID                string `json:"id,omitempty"`
	Name              string `json:"name,omitempty"`
	FirstName         string `json:"firstName,omitempty"`
	LastName          string `json:"lastName,omitempty"`
	EmailAddress      string `json:"emailAddress,omitempty"`
	PhoneNumber       string `json:"phoneNumber,omitempty"`
	AddressStreet     string `json:"addressStreet,omitempty"`
	AddressCity       string `json:"addressCity,omitempty"`
	AddressState      string `json:"addressState,omitempty"`
	AddressCountry    string `json:"addressCountry,omitempty"`
	AddressPostalCode string `json:"addressPostalCode,omitempty"`
	DocumentNumber    string `json:"documentNumber,omitempty"`
	DocumentNumberRG  string `json:"documentNumberRG,omitempty"`
	BirthDate         string `json:"birthDate,omitempty"`
	CardNumber        string `json:"cardNumber,omitempty"`
	AccountID         string `json:"accountId,omitempty"`
	AccountName       string `json:"accountName,omitempty"`
}
