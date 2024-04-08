package service

import (
	"fmt"
	"net/url"
	"sync"

	"github.com/EnubeRepos/ElevenST_BFF/config"
	apiError "github.com/EnubeRepos/ElevenST_BFF/errors"
	"github.com/EnubeRepos/ElevenST_BFF/model"
	"github.com/EnubeRepos/ElevenST_BFF/model/api"
	"github.com/EnubeRepos/ElevenST_BFF/repositories/crm"
	pixRepo "github.com/EnubeRepos/ElevenST_BFF/repositories/pix"
)

func GeneratePix(body *api.PixGenerate) (*api.PixGenerateResponse, apiError.RestErr) {
	body.LegalName, _ = pixRepo.NormalizeInput(body.LegalName)
	body.Description = "Pagamento de Consulta"
	paste, err := pixRepo.GeneratePixCopyPast(
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

func GetContactsInAllPlatforms(cpf string, cfg *config.Config) (interface{}, apiError.RestErr) {
	crmlib := crm.New(cfg.Authorization)

	var contactList model.ResponseContactList
	var rsp []model.ResponseContact

	var wg sync.WaitGroup
	wg.Add(len(cfg.Clinics.List))
	for _, clinic := range cfg.Clinics.List {
		go func(clinic model.ClinicReference) {
			defer wg.Done()
			contactlist := new(model.ContactList)
			filter := fmt.Sprintf(`maxSize=1&offset=0&orderBy=name&order=asc&where[0][type]=equals&where[0][attribute]=documentNumber&where[0][value]=%s`, url.QueryEscape(cpf))
			res, err := crmlib.GET(clinic.URL+"/api/v1", "Contact", filter)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(string(res))
			err = contactlist.Bind(res)
			if err != nil {
				fmt.Println(err)
			}
			if contactlist.Total == 0 {
				return
			}
			rsp = append(rsp, model.ResponseContact{
				ClinicName: clinic.Name,
				URL:        clinic.URL,
				Entity:     "Contact",
				Content:    contactlist.List[0],
				ID:         contactlist.List[0].ID,
			})
		}(clinic)
	}

	wg.Wait()

	contactList.Total = len(rsp)
	contactList.List = rsp

	return contactList, nil
}
