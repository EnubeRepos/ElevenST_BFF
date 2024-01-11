package config

import (
	"os"
	"time"

	"github.com/EnubeRepos/ElevenST_BFF/model"
	"github.com/sirupsen/logrus"
)

type Config struct {
	Log           *logrus.Logger
	Authorization string
	APIAuth       string
	Clinics       []model.ClinicReference
}

func New() (*Config, error) {
	var clinicRef model.ClinicReferenceList
	var cfg Config

	// header := map[string]string{
	// 	"Method": "GET",
	// }

	// var ClinicReferenceList = os.Getenv("CLINIC_REFERENCE")
	var AuthorizationAPI = os.Getenv("API_AUTH")
	// _, err := web.MakeRequest(ClinicReferenceList, header, nil, &clinicRef)
	// if err != nil {
	// 	return &cfg, err
	// }

	l := logrus.New()
	l.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat:   time.RFC3339Nano,
		DisableHTMLEscape: true,
		PrettyPrint:       true,
	})

	cfg.Authorization = os.Getenv("AUTH")
	cfg.Clinics = clinicRef.List
	cfg.APIAuth = AuthorizationAPI
	cfg.Log = l

	return &cfg, nil
}
