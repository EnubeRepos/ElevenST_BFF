package web

import (
	"fmt"

	uri "net/url"

	"github.com/EnubeRepos/ElevenST_BFF/model"
	"github.com/sirupsen/logrus"
)

const (
	API_PATH = "/api/v1/"
)

type Web struct {
	Authorization string
	Log           *logrus.Logger
}

func New(l *logrus.Logger, auth string) *Web {
	return &Web{
		Authorization: auth,
		Log:           l,
	}
}

func (w Web) GetContactByDocumentNumber(url, docNumber string) (model.ContactList, error) {
	var list model.ContactList
	var total int
	offset := 0
	maxSize := 50

	header := map[string]string{
		"Method":        "GET",
		"Authorization": w.Authorization,
	}

	for {
		var temp model.ContactList

		filter := fmt.Sprintf(`?maxSize=%d&offset=%d&orderBy=name&order=asc&where[0][type]=equals&where[0][attribute]=documentNumber&where[0][value]=%s`, maxSize, offset, uri.QueryEscape(docNumber))

		_, err := MakeRequest(url+API_PATH+"Contact"+filter, header, nil, temp)
		if err != nil {
			w.Log.Error("Error making GET request")
			return list, err
		}

		total = temp.Total
		offset += maxSize

		list.Total = total
		list.List = append(list.List, temp.List...)

		if offset >= total {
			break
		}
	}

	return list, nil
}
