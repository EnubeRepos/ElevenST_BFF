package api

import (
	"encoding/json"
	"io"

	apiError "github.com/EnubeRepos/ElevenST_BFF/errors"
)

type PixGenerate struct {
	LegalName   string  `json:"legalName"`
	Amount      float64 `json:"amount"`
	Description string  `json:"description"`
	Pixkey      string  `json:"Pixkey"`
}

func (pg *PixGenerate) Bind(data io.Reader) apiError.RestErr {
	jsonData, err := io.ReadAll(data)
	if err != nil {
		return apiError.NewInternalServerError("Error trying to read request body", err)
	}
	// dataByte, err := json.Marshal(jsonData)
	// if err != nil {
	// 	return apiError.NewInternalServerError("Error trying to marshal request body", err)

	// }
	err = json.Unmarshal(jsonData, pg)
	if err != nil {
		return apiError.NewInternalServerError("Error trying to unmarshal request body", err)
	}
	return nil
}

func (pg *PixGenerate) ValidateBody() apiError.RestErr {
	if pg.LegalName == "" {
		return apiError.NewBadRequestError("field legalName must be set")
	}
	if pg.Amount == 0.0 {
		return apiError.NewBadRequestError("field amount must be set")
	}
	if pg.Pixkey == "" {
		return apiError.NewBadRequestError("field pixKey must be set")
	}
	return nil
}
