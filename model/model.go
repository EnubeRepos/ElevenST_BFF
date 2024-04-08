package model

import (
	"encoding/json"

	apiError "github.com/EnubeRepos/ElevenST_BFF/errors"
)

type ContactList struct {
	Total int       `json:"total"`
	List  []Contact `json:"list"`
}

func (cl *ContactList) Bind(data []byte) apiError.RestErr {
	err := json.Unmarshal(data, cl)
	if err != nil {
		return apiError.NewInternalServerError("Error trying to unmarshal request body", err)
	}
	return nil
}

func (cl *ContactList) Unbind() ([]byte, apiError.RestErr) {
	data, err := json.Marshal(cl)
	if err != nil {
		return nil, apiError.NewInternalServerError("Error trying to unmarshal request body", err)
	}
	return data, nil
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

func (c *Contact) Bind(data []byte) apiError.RestErr {
	err := json.Unmarshal(data, c)
	if err != nil {
		return apiError.NewInternalServerError("Error trying to unmarshal request body", err)
	}
	return nil
}

func (c *Contact) Unbind() ([]byte, apiError.RestErr) {
	data, err := json.Marshal(c)
	if err != nil {
		return nil, apiError.NewInternalServerError("Error trying to unmarshal request body", err)
	}
	return data, nil
}
