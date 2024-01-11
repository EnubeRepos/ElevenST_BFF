package crm

import (
	"fmt"
	"io"
	"net/http"

	apiError "github.com/EnubeRepos/ElevenST_BFF/errors"
)

const (
	METHOD_GET   = "GET"
	METHOD_POST  = "POST"
	METHOD_DEL   = "DELETE"
	METHOD_PATCH = "PATCH"
)

type CRM interface {
	GET(URL, Entity, Resource string) ([]byte, apiError.RestErr)
	GETByID(URL, Entity, ID string) ([]byte, apiError.RestErr)
	PATCH(URL, Entity, ID string) ([]byte, apiError.RestErr)
	DELETE(URL, Entity, ID string) ([]byte, apiError.RestErr)
	POST(URL, Entity string) ([]byte, apiError.RestErr)
}

type CRMconfig struct {
	Auth string
}

func New(Authorization string) *CRMconfig {
	return &CRMconfig{
		Auth: fmt.Sprintf("Basic %s", Authorization),
	}
}

func (c *CRMconfig) GET(URL, Entity, Resource string) ([]byte, apiError.RestErr) {
	url := fmt.Sprintf("%s/%s?%s", URL, Entity, Resource)

	client := &http.Client{}
	req, err := http.NewRequest(METHOD_GET, url, nil)

	if err != nil {
		fmt.Println(err)
		return nil, apiError.NewInternalServerError("Error trying to create NewRequest", err)
	}

	req.Header.Add("Authorization", c.Auth)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, apiError.NewInternalServerError("Error trying to do request", err)
	}

	if res.StatusCode == 404 {
		return nil, apiError.NewNotFoundError(fmt.Sprintf("Not Found | URL: %s, Entity: %s, Resource: %s", URL, Entity, Resource))
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil, apiError.NewInternalServerError("Error trying read body", err)
	}

	return body, nil
}
