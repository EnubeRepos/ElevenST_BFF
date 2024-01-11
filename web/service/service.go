package service

import (
	"fmt"

	apiError "github.com/EnubeRepos/ElevenST_BFF/errors"
	"github.com/EnubeRepos/ElevenST_BFF/model/api"
	pixRepo "github.com/EnubeRepos/ElevenST_BFF/repositories/pix"
)

func GeneratePix(body *api.PixGenerate) (*api.PixGenerateResponse, apiError.RestErr) {
	body.LegalName, _ = pixRepo.NormalizeInput(body.LegalName)
	body.Description = "Pagamento de Consulta"
	paste, err := pixRepo.GeneratePaste(
		body.Amount,
		body.LegalName,
		body.Description,
		body.Pixkey,
	)
	if err != nil {
		return nil, apiError.NewBadRequestError(fmt.Sprintf("Error trying to generate Copy and Paste pix. Details: %s", err.Error()))
	}
	res := new(api.PixGenerateResponse)
	res.CopyPaste = paste
	return res, nil
}
